package searches

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/utils/helpers"
	"podcast/src/zlog"
)

type SearchDaoI interface {
	GetByID(id int64) (*SearchDTO, error)
	Find(text string) (*SearchDTO, error)
	CreateOrUpdate(text string, sentiment string) (int64, error)
	List() ([]SearchDTO, error)

	TopSearches() ([]TopSearchesOutput, error)
	ListLocations() ([]SearchLocationsOutput, error)
}
type searchDao struct{}

var SearchDao SearchDaoI = &searchDao{}

func (s searchDao) GetByID(id int64) (*SearchDTO, error) {
	var search SearchDTO

	err := mysql.Client.Get(&search, queryByID, id)
	if err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>GetByID: %s", err))
		}
		return nil, err
	}

	return &search, nil
}

func (s searchDao) Find(text string) (*SearchDTO, error) {
	var search SearchDTO

	err := mysql.Client.Get(&search, text)
	if err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>Find: %s", err))
		}
		return nil, err
	}

	return &search, nil
}

func (s searchDao) CreateOrUpdate(text string, sentiment string) (int64, error) {
	var search SearchDTO
	err := mysql.Client.Get(&search, queryFind, text)

	if err == sql.ErrNoRows {

		searchInput := SearchInput{
			Text:      text,
			Sentiment: sentiment,
		}

		qMap, err := helpers.ConvertStructToMap(searchInput, "db")
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return 0, err
		}

		row, err := mysql.Client.NamedExec(queryCreate, qMap)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return 0, err
		}

		id, err := row.LastInsertId()
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("CreateDao=>Create: %s", err))
			return 0, err
		}
		return id, nil
	} else {
		_, err := mysql.Client.Exec(queryUpdateCount, text)
		if err != nil {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>Update: %s", err))
			return search.ID, nil
		}
	}

	return search.ID, nil
}

func (s searchDao) List() ([]SearchDTO, error) {
	var results []SearchDTO

	// Get the records
	if err := mysql.Client.Select(&results, queryList); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>List=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}

func (s searchDao) TopSearches() ([]TopSearchesOutput, error) {
	var results []TopSearchesOutput

	// Get the records
	if err := mysql.Client.Select(&results, queryTopSegmentFromSearch); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>TopSearches=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}

func (s searchDao) ListLocations() ([]SearchLocationsOutput, error) {
	var results []SearchLocationsOutput
	// Get the records
	if err := mysql.Client.Select(&results, queryGetSearchLocations); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SearchDao=>SearchLocations=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}
