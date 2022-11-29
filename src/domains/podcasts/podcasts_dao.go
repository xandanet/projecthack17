package podcasts

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type PodcastDaoI interface {
	GetMaxDuration() ([]PodcastDuration, error)
	List() ([]PodcastDTO, error)
	Single(id int64) (*PodcastDTO, error)
	Interventions(id int64) ([]PodcastInterventionsOutput, error)
	Sentiment(id int64) ([]PodcastSentimentOutput, error)
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

func (d *podcastDao) List() ([]PodcastDTO, error) {
	var results []PodcastDTO

	// Get the records
	if err := mysql.Client.Select(&results, queryList); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("PodcastDao=>List=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}

func (d *podcastDao) Single(id int64) (*PodcastDTO, error) {
	var result PodcastDTO

	// Get the records
	if err := mysql.Client.Get(&result, querySingle, id); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("PodcastDao=>Single=>Get: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return &result, nil
}

func (d *podcastDao) Interventions(id int64) ([]PodcastInterventionsOutput, error) {
	var results []PodcastInterventionsOutput

	// Get the records
	if err := mysql.Client.Select(&results, queryInterventions, id); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("PodcastDao=>Interventions=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}

func (d *podcastDao) Sentiment(id int64) ([]PodcastSentimentOutput, error) {
	var results []PodcastSentimentOutput

	// Get the records
	if err := mysql.Client.Select(&results, querySentiment, id); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("PodcastDao=>Sentiment=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}
