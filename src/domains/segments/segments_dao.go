package segments

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/utils/helpers"
	"podcast/src/zlog"
)

type SegmentDaoI interface {
	GetByID(id int64) (*SegmentDTO, error)
	List() ([]SegmentDTO, error)
	ListTextOnly() ([]string, error)
	SearchByText(search string) (*SegmentDTO, error)
	SearchByNaturalSearch(search string) ([]SegmentDTO, error)
	GetSearchLogByID(id int64) (*SearchSubtitleOutput, error)
	CreateSearchLog(input *SearchSubtitleInput) (*SearchSubtitleOutput, error)
}

type segmentDao struct{}

var SegmentDao SegmentDaoI = &segmentDao{}

func (d *segmentDao) GetByID(id int64) (*SegmentDTO, error) {
	var segment SegmentDTO

	err := mysql.Client.Get(&segment, queryGetByID, id)
	if err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>GetByID: %s", err))
		}
		return nil, err
	}

	return &segment, nil
}

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
	search = ">" + search
	if err := mysql.Client.Select(&result, querySearchByNaturalSearchText, search, search); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>SearchByNaturalSearch=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return result, nil
}

func (d *segmentDao) GetSearchLogByID(id int64) (*SearchSubtitleOutput, error) {
	var searchSubtitle SearchSubtitleOutput

	err := mysql.Client.Get(&searchSubtitle, queryGetSearchBySubtitleByID, id)
	if err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>GetSearchLogByID: %s", err))
		}
		return nil, err
	}

	return &searchSubtitle, nil
}

func (d *segmentDao) CreateSearchLog(input *SearchSubtitleInput) (*SearchSubtitleOutput, error) {

	var searchSubtitle *SearchSubtitleOutput
	err := mysql.Client.Get(&searchSubtitle, querySearchBySubtitleIdSearchId, input.SubtitleId, input.SearchId)

	if err == sql.ErrNoRows {

		qMap, err := helpers.ConvertStructToMap(input, "db")
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return nil, err
		}

		row, err := mysql.Client.NamedExec(queryCreateSearchSubtitle, qMap)
		fmt.Println(row)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return nil, err
		}

		id, err := row.LastInsertId()
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return nil, err
		}

		err = mysql.Client.Get(&searchSubtitle, queryGetSearchBySubtitleByID, id)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return nil, err
		}
		return searchSubtitle, nil

	} else {
		_, err := mysql.Client.Exec(queryUpdateCount, input.SubtitleId, input.SearchId)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>Update: %s", err))
			return nil, err
		}
	}

	return searchSubtitle, nil
}
