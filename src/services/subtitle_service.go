package services

import (
	"fmt"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
	"math"
	"podcast/src/domains/keywords"
	"podcast/src/domains/subtitles"
	"podcast/src/resterror"
	"podcast/src/utils"
	"podcast/src/utils/helpers"
	"sort"
)

var stopWords = []string{"a", "about", "above", "above", "across", "after", "afterwards", "again", "against", "all", "almost", "alone", "along", "already", "also", "although", "always", "am", "among", "amongst", "amoungst", "amount", "an", "and", "another", "any", "anyhow", "anyone", "anything", "anyway", "anywhere", "are", "around", "as", "at", "back", "be", "became", "because", "become", "becomes", "becoming", "been", "before", "beforehand", "behind", "being", "below", "beside", "besides", "between", "beyond", "bill", "both", "bottom", "but", "by", "call", "can", "cannot", "cant", "co", "con", "could", "couldnt", "cry", "de", "describe", "detail", "do", "done", "down", "due", "during", "each", "eg", "eight", "either", "eleven", "else", "elsewhere", "empty", "enough", "etc", "even", "ever", "every", "everyone", "everything", "everywhere", "except", "few", "fifteen", "fify", "fill", "find", "fire", "first", "five", "for", "former", "formerly", "forty", "found", "four", "from", "front", "full", "further", "get", "give", "go", "had", "has", "hasnt", "have", "he", "hence", "her", "here", "hereafter", "hereby", "herein", "hereupon", "hers", "herself", "him", "himself", "his", "how", "however", "hundred", "ie", "if", "in", "inc", "indeed", "interest", "into", "is", "it", "its", "itself", "keep", "last", "latter", "latterly", "least", "less", "ltd", "made", "many", "may", "me", "meanwhile", "might", "mill", "mine", "more", "moreover", "most", "mostly", "move", "much", "must", "my", "myself", "name", "namely", "neither", "never", "nevertheless", "next", "nine", "no", "nobody", "none", "noone", "nor", "not", "nothing", "now", "nowhere", "of", "off", "often", "on", "once", "one", "only", "onto", "or", "other", "others", "otherwise", "our", "ours", "ourselves", "out", "over", "own", "part", "per", "perhaps", "please", "put", "rather", "re", "same", "see", "seem", "seemed", "seeming", "seems", "serious", "several", "she", "should", "show", "side", "since", "sincere", "six", "sixty", "so", "some", "somehow", "someone", "something", "sometime", "sometimes", "somewhere", "still", "such", "system", "take", "ten", "than", "that", "the", "their", "them", "themselves", "then", "thence", "there", "thereafter", "thereby", "therefore", "therein", "thereupon", "these", "they", "thickv", "thin", "third", "this", "those", "though", "three", "through", "throughout", "thru", "thus", "to", "together", "too", "top", "toward", "towards", "twelve", "twenty", "two", "un", "under", "until", "up", "upon", "us", "very", "via", "was", "we", "well", "were", "what", "whatever", "when", "whence", "whenever", "where", "whereafter", "whereas", "whereby", "wherein", "whereupon", "wherever", "whether", "which", "while", "whither", "who", "whoever", "whole", "whom", "whose", "why", "will", "with", "within", "without", "would", "yet", "you", "your", "yours", "yourself", "yourselves"}

var ignoreWords = []string{"i", "me", "my", "myself", "we", "our", "ours", "ourselves", "you", "your", "yours", "yourself", "yourselves", "he", "him", "his", "himself", "she", "her", "hers", "herself", "it", "its", "itself", "they", "them", "their", "theirs", "themselves", "what", "which", "who", "whom", "this", "that", "these", "those", "am", "is", "are", "was", "were", "be", "been", "being", "have", "has", "had", "having", "do", "does", "did", "doing", "a", "an", "the", "and", "but", "if", "or", "because", "as", "until", "while", "of", "at", "by", "for", "with", "about", "against", "between", "into", "through", "during", "before", "after", "above", "below", "to", "from", "up", "down", "in", "out", "on", "off", "over", "under", "again", "further", "then", "once", "here", "there", "when", "where", "why", "how", "all", "any", "both", "each", "few", "more", "most", "other", "some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too", "very", "s", "t", "can", "will", "just", "don", "should", "now", "yeah", "known", "like", "let", "going", "guys", "want", "said", "actually"}

type SubtitleServiceI interface {
	List() ([]subtitles.SubtitleDTO, utils.RestErrorI)
	Search(input *subtitles.SubtitleSearchInput) ([]subtitles.SubtitleDTO, utils.RestErrorI)
	ParseAllPodcasts() utils.RestErrorI
}

type subtitleService struct{}

var SubtitleService SubtitleServiceI = &subtitleService{}

func (s *subtitleService) List() ([]subtitles.SubtitleDTO, utils.RestErrorI) {
	result, err := subtitles.SubtitleDao.List()
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	return result, nil
}

func (s *subtitleService) Search(input *subtitles.SubtitleSearchInput) ([]subtitles.SubtitleDTO, utils.RestErrorI) {
	//Do a natural search first
	naturalResult, err := subtitles.SubtitleDao.SearchByNaturalSearch(input.Text)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetSearch)
	}

	if len(naturalResult) > 0 {
		return naturalResult, nil
	}

	textOnly, err := subtitles.SubtitleDao.ListAllText()
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	var results []subtitles.SubtitleDTO

	searchResults := s.searchInText(input.Text, textOnly, 0.9)

	//Find the entries that match the results
	for _, searchResult := range searchResults {
		detail, err := subtitles.SubtitleDao.SearchByText(searchResult.Text)
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

	return results, nil
}

func (s *subtitleService) searchInText(query string, testCorpus []string, minSimilarity float64) []subtitles.TextSearchAnalysis {
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
	var matched []subtitles.TextSearchAnalysis
	_, docs := lsi.Dims()
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		if similarity >= minSimilarity {
			matched = append(matched, subtitles.TextSearchAnalysis{
				Text:       testCorpus[i],
				Similarity: similarity,
			})
		}
	}

	return matched
}

func (s *subtitleService) Topics_(podcastID int64) (*subtitles.ParagraphAnalysis, utils.RestErrorI) {
	textOnly, err := subtitles.SubtitleDao.ListTextOnly(podcastID)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	var result subtitles.ParagraphAnalysis

	// Create a pipeline with a count vectoriser and LDA transformer for 2 topics
	vectoriser := nlp.NewCountVectoriser(stopWords...)
	lda := nlp.NewLatentDirichletAllocation(3)
	pipeline := nlp.NewPipeline(vectoriser, lda)

	docsOverTopics, err := pipeline.FitTransform(textOnly...)
	if err != nil {
		fmt.Printf("Failed to model topics for documents because %v", err)
		return nil, resterror.NewBadRequestError(err.Error())
	}

	// Examine Document over topic probability distribution
	dr, dc := docsOverTopics.Dims()
	for doc := 0; doc < dc; doc++ {
		paragraphResult := subtitles.ParagraphInformation{ParagraphID: int64(doc), Topic: 0, Score: 0}
		for topic := 0; topic < dr; topic++ {
			if docsOverTopics.At(topic, doc) > paragraphResult.Score {
				paragraphResult.Topic = int64(topic)
				paragraphResult.Score = docsOverTopics.At(topic, doc)
			}
		}
		result.Paragraphs = append(result.Paragraphs, paragraphResult)
	}

	// Examine Topic over word probability distribution
	topicsOverWords := lda.Components()
	tr, tc := topicsOverWords.Dims()

	vocab := make([]string, len(vectoriser.Vocabulary))
	for k, v := range vectoriser.Vocabulary {
		vocab[v] = k
	}
	for topic := 0; topic < tr; topic++ {
		var topicWords []subtitles.TopicWords
		for word := 0; word < tc; word++ {
			if len(vocab[word]) > 1 && !helpers.StringInMap(vocab[word], stopWords) && !helpers.StringInMap(vocab[word], ignoreWords) {
				topicWords = append(topicWords, subtitles.TopicWords{
					Text:  vocab[word],
					Score: topicsOverWords.At(topic, word),
				})
			}
		}

		//Order results by similarity
		sort.Slice(topicWords, func(i, j int) bool {
			return topicWords[i].Score > topicWords[j].Score
		})

		//Only save the top 5
		result.Topics = append(result.Topics, subtitles.Topic{
			ID:    int64(topic),
			Words: topicWords[0:5],
		})
	}

	return &result, nil
}

// ParseAllPodcasts This version analyses each paragraph individually
func (s *subtitleService) ParseAllPodcasts() utils.RestErrorI {
	minWordLength := 3

	for podcastID := int64(1); podcastID < 4; podcastID++ {
		textOnly, err := subtitles.SubtitleDao.ListByPodcast(podcastID)
		if err != nil {
			return utils.NewInternalServerError(utils.ErrorGetList)
		}

		var globalResult []subtitles.ParagraphAnalysis

		for _, textParagraph := range textOnly {
			var result subtitles.ParagraphAnalysis

			// Create a pipeline with a count vectoriser and LDA transformer for 2 topics
			vectoriser := nlp.NewCountVectoriser(textParagraph.Content)
			lda := nlp.NewLatentDirichletAllocation(1)
			pipeline := nlp.NewPipeline(vectoriser, lda)

			docsOverTopics, err := pipeline.FitTransform(textParagraph.Content)
			if err != nil {
				fmt.Printf("Failed to model topics for documents because %v", err)
				return resterror.NewBadRequestError(err.Error())
			}

			// Examine Document over topic probability distribution
			dr, dc := docsOverTopics.Dims()
			for doc := 0; doc < dc; doc++ {
				paragraphResult := subtitles.ParagraphInformation{ParagraphID: textParagraph.ID, Topic: 0, Score: 0}
				for topic := 0; topic < dr; topic++ {
					if docsOverTopics.At(topic, doc) > paragraphResult.Score {
						paragraphResult.Topic = int64(topic)
						paragraphResult.Score = docsOverTopics.At(topic, doc)
					}
				}
				result.Paragraphs = append(result.Paragraphs, paragraphResult)
			}

			if len(vectoriser.Vocabulary) == 0 {
				continue
			}

			// Examine Topic over word probability distribution
			topicsOverWords := lda.Components()
			tr, tc := topicsOverWords.Dims()

			vocab := make([]string, len(vectoriser.Vocabulary))
			for k, v := range vectoriser.Vocabulary {
				vocab[v] = k
			}
			for topic := 0; topic < tr; topic++ {
				var topicWords []subtitles.TopicWords
				for word := 0; word < tc; word++ {
					if len(vocab[word]) > minWordLength && !helpers.StringInMap(vocab[word], stopWords) && !helpers.StringInMap(vocab[word], ignoreWords) {
						topicWords = append(topicWords, subtitles.TopicWords{
							Text:  vocab[word],
							Score: topicsOverWords.At(topic, word),
						})
					}
				}

				//Order results by similarity
				sort.Slice(topicWords, func(i, j int) bool {
					return topicWords[i].Score > topicWords[j].Score
				})

				//Only save the top 5
				result.Topics = append(result.Topics, subtitles.Topic{
					ID:    int64(topic),
					Words: topicWords[0:int(math.Min(10, float64(len(topicWords))))],
				})

				//Add the words to the database
				for _, word := range topicWords[0:int(math.Min(5, float64(len(topicWords))))] {
					keywordID, err := keywords.KeywordDao.FindID(word.Text)
					if err != nil {
						//Create word
						err := keywords.KeywordDao.Create(word.Text)
						if err != nil {
							return resterror.NewBadRequestError(err.Error())
						}
						//Get ID
						keywordID, err = keywords.KeywordDao.FindID(word.Text)
						if err != nil {
							return resterror.NewBadRequestError(err.Error())
						}
					}
					//Create relationship
					err = keywords.KeywordDao.CreateRelationshipSubtitle(keywordID, textParagraph.ID)
					if err != nil {
						return resterror.NewBadRequestError(err.Error())
					}
				}

			}

			globalResult = append(globalResult, result)
		}
	}

	return nil
}
