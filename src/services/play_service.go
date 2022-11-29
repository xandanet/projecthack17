package services

import (
	"golang.org/x/exp/rand"
	"podcast/src/domains/plays"
	"podcast/src/domains/podcasts"
	"podcast/src/resterror"
	"podcast/src/utils"
	"time"
)

type PlayServiceI interface {
	Create(req *plays.PlayCreateInput) utils.RestErrorI
	Seed() utils.RestErrorI
	Statistics(req *plays.PlayStatisticsInput) ([]plays.PlayStatisticsOutput, utils.RestErrorI)
	PerDay(req *plays.PlayStatisticsPerDayInput) ([]plays.PlayStatisticsPerDayOutput, utils.RestErrorI)
	SegmentPopularity(req *plays.PlayStatisticsPerDayInput) ([]plays.PlaySegmentPopularityOutput, utils.RestErrorI)
}

type playService struct{}

var PlayService PlayServiceI = &playService{}

func (s *playService) Create(req *plays.PlayCreateInput) utils.RestErrorI {
	req.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := plays.PlayDao.Create(req)
	if err != nil {
		return utils.NewInternalServerError(err.Error())
	}

	return nil
}

func (s *playService) Seed() utils.RestErrorI {
	durations, err := podcasts.PodcastDao.GetMaxDuration()
	if err != nil {
		return resterror.NewBadRequestError(err.Error())
	}

	for _, duration := range durations {
		startDate := time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC)
		for j := int64(0); j < duration.Duration; j++ {
			numberOfPlays := rand.Int63n(15)
			for k := int64(0); k < numberOfPlays; k++ {
				_ = plays.PlayDao.Create(&plays.PlayCreateInput{
					PodcastID: duration.ID,
					Position:  j,
					CreatedAt: startDate.Format("2006-01-02 15:04:05"),
				})

				startDate = startDate.Add(time.Duration(rand.Int63n(20)) * time.Second)
			}
			//Jump ahead
			j += rand.Int63n(500)
		}
	}

	return nil
}

func (s *playService) Statistics(req *plays.PlayStatisticsInput) ([]plays.PlayStatisticsOutput, utils.RestErrorI) {
	result, err := plays.PlayDao.Statistics(req)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (s *playService) PerDay(req *plays.PlayStatisticsPerDayInput) ([]plays.PlayStatisticsPerDayOutput, utils.RestErrorI) {
	result, err := plays.PlayDao.PerDay(req)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	return result, nil
}

func (s *playService) SegmentPopularity(req *plays.PlayStatisticsPerDayInput) ([]plays.PlaySegmentPopularityOutput, utils.RestErrorI) {
	result, err := plays.PlayDao.SegmentPopularity(req)
	if err != nil {
		return nil, utils.NewInternalServerError(err.Error())
	}
	return result, nil
}
