package services

import (
	"podcast/src/domains/podcasts"
	"podcast/src/resterror"
	"podcast/src/utils"
)

type PodcastServiceI interface {
	List() ([]podcasts.PodcastDTO, utils.RestErrorI)
	Interventions(id int64) ([]podcasts.PodcastInterventionsOutput, resterror.RestErrorI)
	Sentiment(id int64) ([]podcasts.PodcastSentimentOutput, resterror.RestErrorI)
	Bookmark(input podcasts.BookmarkInput) resterror.RestErrorI
	GetBookmark(input podcasts.BookmarkSearchInput) ([]podcasts.GetBookmarkSearchOutput, resterror.RestErrorI)
}

type podcastService struct{}

var PodcastService PodcastServiceI = &podcastService{}

func (s *podcastService) List() ([]podcasts.PodcastDTO, utils.RestErrorI) {
	result, err := podcasts.PodcastDao.List()
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	return result, nil
}

func (s *podcastService) Interventions(id int64) ([]podcasts.PodcastInterventionsOutput, resterror.RestErrorI) {
	result, err := podcasts.PodcastDao.Interventions(id)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	return result, nil
}

func (s *podcastService) Sentiment(id int64) ([]podcasts.PodcastSentimentOutput, resterror.RestErrorI) {
	sentiments, err := podcasts.PodcastDao.Sentiment(id)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	//Remove duplicates
	var result []podcasts.PodcastSentimentOutput
	currentSentiment := ""
	for _, sentiment := range sentiments {
		if sentiment.Sentiment != currentSentiment {
			currentSentiment = sentiment.Sentiment
			result = append(result, sentiment)
		}
	}

	return result, nil
}

func (s *podcastService) Bookmark(input podcasts.BookmarkInput) resterror.RestErrorI {
	err := podcasts.PodcastDao.CreateBookmark(input)
	if err != nil {
		return utils.NewInternalServerError("ERROR_CREATING_BOOKMARK")
	}
	return nil
}

func (s *podcastService) GetBookmark(input podcasts.BookmarkSearchInput) ([]podcasts.GetBookmarkSearchOutput, resterror.RestErrorI) {
	bookmarks, err := podcasts.PodcastDao.GetBookmark(input.PodId)

	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}
	return bookmarks, nil
}
