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
	List(podcastID int64) ([]SegmentDTO, error)
	ListTextOnly() ([]string, error)
	SearchByText(search string) (*SegmentDTO, error)
	SearchByNaturalSearch(search string) ([]SegmentDTO, error)
	GetSearchLogByID(id int64) (*SearchSegmentOutput, error)
	CreateSearchLog(input *SearchSegmentInput) (*SearchSegmentOutput, error)
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

func (d *segmentDao) List(podcastID int64) ([]SegmentDTO, error) {
	var results []SegmentDTO

	// Get the records
	if err := mysql.Client.Select(&results, queryList, podcastID); err != nil {
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
			zlog.Logger.Error(fmt.Sprintf("SegmentDao=>SearchByText=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return &result, nil
}

func (d *segmentDao) SearchByNaturalSearch(search string) ([]SegmentDTO, error) {
	var result []SegmentDTO

	// Get the records
	if err := mysql.Client.Select(&result, querySearchByNaturalSearchText, search, search, search, search); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SegmentDao=>SearchByNaturalSearch=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return result, nil
}

func (d *segmentDao) GetSearchLogByID(id int64) (*SearchSegmentOutput, error) {
	var searchSubtitle SearchSegmentOutput

	err := mysql.Client.Get(&searchSubtitle, queryGetSearchBySubtitleByID, id)
	if err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>GetSearchLogByID: %s", err))
		}
		return nil, err
	}

	return &searchSubtitle, nil
}

func (d *segmentDao) CreateSearchLog(input *SearchSegmentInput) (*SearchSegmentOutput, error) {
	var searchSegment SearchSegmentOutput
	if err := mysql.Client.Get(&searchSegment, querySearchBySubtitleIdSearchId, input.SegmentId, input.SearchId); err == sql.ErrNoRows {

		qMap, err := helpers.ConvertStructToMap(input, "db")
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("SegmentDao=>Create: %s", err))
			return nil, err
		}

		row, err := mysql.Client.NamedExec(queryCreateSearchSubtitle, qMap)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("SegmentDao=>Create: %s", err))
			return nil, err
		}

		id, err := row.LastInsertId()
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("SegmentDao=>Create: %s", err))
			return nil, err
		}

		err = mysql.Client.Get(&searchSegment, queryGetSearchBySubtitleByID, id)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return nil, err
		}

		return &searchSegment, nil
	} else {

		_, err := mysql.Client.Exec(queryUpdateCount, input.SegmentId, input.SearchId)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("SegmentDao=>UpdateCount: %s", err))
			return nil, err
		}
	}

	ipLocation := helpers.GetLocationFromIp(input.IpAddress)

	searchInput := SearchLogInput{
		SearchId:  input.SearchId,
		IpAddress: input.IpAddress,
		Region:    ipLocation.Postal.Code,
		City:      ipLocation.City.Names["en"],
		Country:   ipLocation.Country.Names["en"],
	}
	fmt.Println(searchInput)
	qMap, err := helpers.ConvertStructToMap(searchInput, "db")
	if err != nil {
		zlog.Logger.Error(fmt.Sprintf("SegmentDao=>Create: %s", err))
		return nil, err
	}

	_, err = mysql.Client.NamedExec(queryCreateLog, qMap)
	if err != nil {
		zlog.Logger.Error(fmt.Sprintf("SegmentDao=>CreateSearchLog: %s", err))
		return nil, err
	}

	return &searchSegment, nil
}
