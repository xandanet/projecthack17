package services

import (
	"golang.org/x/exp/rand"
	"podcast/src/domains/plays"
	"podcast/src/utils"
	"time"
)

type PlayServiceI interface {
	Create(req *plays.PlayCreateInput) utils.RestErrorI
	Seed() utils.RestErrorI
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
	for i := int64(1); i < 4; i++ {
		startDate := time.Date(2022, 10, 6, 20, 31, 0, 0, time.UTC)
		for j := int64(0); j < 4000000; j++ {
			numberOfPlays := rand.Int63n(15)
			for k := int64(0); k < numberOfPlays; k++ {
				_ = plays.PlayDao.Create(&plays.PlayCreateInput{
					PodcastID: i,
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
