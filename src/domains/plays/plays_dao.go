package plays

import (
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type PlayDaoI interface {
	Create(req *PlayCreateInput) error
	Statistics(req *PlayStatisticsInput) ([]PlayStatisticsOutput, error)
}

type playDao struct{}

var PlayDao PlayDaoI = &playDao{}

func (d *playDao) Create(req *PlayCreateInput) error {
	if _, err := mysql.Client.Exec(queryCreate, req.PodcastID, req.Position, req.CreatedAt); err != nil {
		zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>List=>Select: %s", err))
		return err
	}

	return nil
}

func (d *playDao) Statistics(req *PlayStatisticsInput) ([]PlayStatisticsOutput, error) {
	var result []PlayStatisticsOutput

	if err := mysql.Client.Select(&result, queryStatistics, req.PodcastID, req.StartDate, req.EndDate); err != nil {
		zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>Statistics=>Select: %s", err))
		return nil, err
	}

	return result, nil
}
