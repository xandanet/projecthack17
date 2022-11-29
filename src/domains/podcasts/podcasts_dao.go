package podcasts

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type PodcastDaoI interface {
	GetMaxDuration() ([]PodcastDuration, error)
}

type podcastDao struct{}

var PodcastDao PodcastDaoI = &podcastDao{}

func (d *podcastDao) GetMaxDuration() ([]PodcastDuration, error) {
	var durations []PodcastDuration

	err := mysql.Client.Select(&durations, queryMaxDuration)
	if err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("PodcastDao=>GetMaxDuration: %s", err))
		}
		return nil, err
	}

	return durations, nil
}
