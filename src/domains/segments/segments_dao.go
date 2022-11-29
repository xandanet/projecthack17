package segments

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type SegmentDaoI interface {
	List() ([]SegmentDTO, error)
	ListTextOnly() ([]string, error)
	SearchByText(search string) (*SegmentDTO, error)
	SearchByNaturalSearch(search string) ([]SegmentDTO, error)
}

type segmentDao struct{}

var SegmentDao SegmentDaoI = &segmentDao{}

func (d *segmentDao) List() ([]SegmentDTO, error) {
	var results []SegmentDTO

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

func (d *segmentDao) ListTextOnly() ([]string, error) {
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

func (d *segmentDao) SearchByText(search string) (*SegmentDTO, error) {
	var result SegmentDTO

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

func (d *segmentDao) SearchByNaturalSearch(search string) ([]SegmentDTO, error) {
	var result []SegmentDTO

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
