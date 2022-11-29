package services

import (
	"fmt"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
	"podcast/src/domains/searches"
	"podcast/src/domains/segments"
	"podcast/src/utils"
	"sort"
)

var stopWords = []string{"a", "about", "above", "above", "across", "after", "afterwards", "again", "against", "all", "almost", "alone", "along", "already", "also", "although", "always", "am", "among", "amongst", "amoungst", "amount", "an", "and", "another", "any", "anyhow", "anyone", "anything", "anyway", "anywhere", "are", "around", "as", "at", "back", "be", "became", "because", "become", "becomes", "becoming", "been", "before", "beforehand", "behind", "being", "below", "beside", "besides", "between", "beyond", "bill", "both", "bottom", "but", "by", "call", "can", "cannot", "cant", "co", "con", "could", "couldnt", "cry", "de", "describe", "detail", "do", "done", "down", "due", "during", "each", "eg", "eight", "either", "eleven", "else", "elsewhere", "empty", "enough", "etc", "even", "ever", "every", "everyone", "everything", "everywhere", "except", "few", "fifteen", "fify", "fill", "find", "fire", "first", "five", "for", "former", "formerly", "forty", "found", "four", "from", "front", "full", "further", "get", "give", "go", "had", "has", "hasnt", "have", "he", "hence", "her", "here", "hereafter", "hereby", "herein", "hereupon", "hers", "herself", "him", "himself", "his", "how", "however", "hundred", "ie", "if", "in", "inc", "indeed", "interest", "into", "is", "it", "its", "itself", "keep", "last", "latter", "latterly", "least", "less", "ltd", "made", "many", "may", "me", "meanwhile", "might", "mill", "mine", "more", "moreover", "most", "mostly", "move", "much", "must", "my", "myself", "name", "namely", "neither", "never", "nevertheless", "next", "nine", "no", "nobody", "none", "noone", "nor", "not", "nothing", "now", "nowhere", "of", "off", "often", "on", "once", "one", "only", "onto", "or", "other", "others", "otherwise", "our", "ours", "ourselves", "out", "over", "own", "part", "per", "perhaps", "please", "put", "rather", "re", "same", "see", "seem", "seemed", "seeming", "seems", "serious", "several", "she", "should", "show", "side", "since", "sincere", "six", "sixty", "so", "some", "somehow", "someone", "something", "sometime", "sometimes", "somewhere", "still", "such", "system", "take", "ten", "than", "that", "the", "their", "them", "themselves", "then", "thence", "there", "thereafter", "thereby", "therefore", "therein", "thereupon", "these", "they", "thickv", "thin", "third", "this", "those", "though", "three", "through", "throughout", "thru", "thus", "to", "together", "too", "top", "toward", "towards", "twelve", "twenty", "two", "un", "under", "until", "up", "upon", "us", "very", "via", "was", "we", "well", "were", "what", "whatever", "when", "whence", "whenever", "where", "whereafter", "whereas", "whereby", "wherein", "whereupon", "wherever", "whether", "which", "while", "whither", "who", "whoever", "whole", "whom", "whose", "why", "will", "with", "within", "without", "would", "yet", "you", "your", "yours", "yourself", "yourselves"}

type SegmentServiceI interface {
	List() ([]segments.SegmentDTO, utils.RestErrorI)
	Search(input *segments.SegmentSearchInput) (*segments.SearchSubtitleDTO, utils.RestErrorI)
	GetContent(input *segments.SearchSubtitleInput) (*segments.SearchSubtitleOutput, utils.RestErrorI)
}

type segmentService struct{}

var SegmentService SegmentServiceI = &segmentService{}

func (s *segmentService) List() ([]segments.SegmentDTO, utils.RestErrorI) {
	result, err := segments.SegmentDao.List()
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	return result, nil
}

func (s *segmentService) Search(input *segments.SegmentSearchInput) (*segments.SearchSubtitleDTO, utils.RestErrorI) {
	//Do a natural search first
	var naturalResult []segments.SegmentDTO
	naturalResult, err := segments.SegmentDao.SearchByNaturalSearch(input.Text)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetSearch)
	}

	if len(naturalResult) > 0 {
		searchId, err := searches.SearchDao.CreateOrUpdate(input.Text)

		if err == nil {
			return &segments.SearchSubtitleDTO{
				SearchID:    searchId,
				SubtitleDTO: naturalResult,
			}, nil
		}
		return nil, utils.NewInternalServerError(utils.ErrorSaveSearch)
	}

	textOnly, err := segments.SegmentDao.ListTextOnly()
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	var results []segments.SegmentDTO

	searchResults := s.searchInText(input.Text, textOnly, 0.9)

	//Find the entries that match the results
	for _, searchResult := range searchResults {
		detail, err := segments.SegmentDao.SearchByText(searchResult.Text)
		if err != nil {
			return nil, utils.NewInternalServerError(utils.ErrorGetSearch)
		}
		detail.Similarity = searchResult.Similarity
		results = append(results, *detail)
	}

	//Order results by similarity
	sort.Slice(results, func(i, j int) bool {
		return results[i].Similarity > results[j].Similarity
	})

	searchId, err := searches.SearchDao.CreateOrUpdate(input.Text)
	return &segments.SearchSubtitleDTO{
		SearchID:    searchId,
		SubtitleDTO: results,
	}, nil
}

func (s *segmentService) searchInText(query string, testCorpus []string, minSimilarity float64) []segments.TextSearchAnalysis {
	vectoriser := nlp.NewCountVectoriser(stopWords...)
	transformer := nlp.NewTfidfTransformer()

	// set k (the number of dimensions following truncation) to 4
	reducer := nlp.NewTruncatedSVD(4)

	lsiPipeline := nlp.NewPipeline(vectoriser, transformer, reducer)

	// Transform the corpus into an LSI fitting the model to the documents in the process
	lsi, err := lsiPipeline.FitTransform(testCorpus...)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return nil
	}

	// run the query through the same pipeline that was fitted to the corpus and
	// to project it into the same dimensional space
	queryVector, err := lsiPipeline.Transform(query)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return nil
	}

	// iterate over document feature vectors (columns) in the LSI matrix and compare
	// with the query vector for similarity.  Similarity is determined by the difference
	// between the angles of the vectors known as the cosine similarity
	var matched []segments.TextSearchAnalysis
	_, docs := lsi.Dims()
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		if similarity >= minSimilarity {
			matched = append(matched, segments.TextSearchAnalysis{
				Text:       testCorpus[i],
				Similarity: similarity,
			})
		}
	}

	return matched
}

func (s *segmentService) GetContent(input *segments.SearchSubtitleInput) (*segments.SearchSubtitleOutput, utils.RestErrorI) {
	_, err := segments.SegmentDao.GetByID(input.SubtitleId)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	_, err = searches.SearchDao.GetByID(input.SearchId)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	result, err := segments.SegmentDao.CreateSearchLog(input)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}
	return result, nil
}
