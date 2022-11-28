package subtitles

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type SubtitleDaoI interface {
	List() ([]SubtitleDTO, error)
	ListTextOnly() ([]string, error)
	SearchByText(search string) (*SubtitleDTO, error)
	SearchByNaturalSearch(search string) ([]SubtitleDTO, error)
}

type subtitleDao struct{}

var SubtitleDao SubtitleDaoI = &subtitleDao{}

func (d *subtitleDao) List() ([]SubtitleDTO, error) {
	var results []SubtitleDTO

	// Get the records
	if err := mysql.Client.Select(&results, queryList); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>List=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}

func (d *subtitleDao) ListTextOnly() ([]string, error) {
	var results []string

	// Get the records
	if err := mysql.Client.Select(&results, queryListTextOnly); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>ListTextOnly=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}

func (d *subtitleDao) SearchByText(search string) (*SubtitleDTO, error) {
	var result SubtitleDTO

	// Get the records
	if err := mysql.Client.Get(&result, querySearchByText, search); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>SearchByText=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return &result, nil
}

func (d *subtitleDao) SearchByNaturalSearch(search string) ([]SubtitleDTO, error) {
	var result []SubtitleDTO

	// Get the records
	if err := mysql.Client.Select(&result, querySearchByNaturalSearchText, search, search); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>SearchByNaturalSearch=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return result, nil
}
