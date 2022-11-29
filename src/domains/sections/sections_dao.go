package sections

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type SectionDaoI interface {
	List(podcastID int64) ([]SectionDTO, error)
}

type sectionDao struct{}

var SectionDao SectionDaoI = &sectionDao{}

func (d *sectionDao) List(podcastID int64) ([]SectionDTO, error) {
	var results []SectionDTO

	// Get the records
	if err := mysql.Client.Select(&results, queryList, podcastID); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("SectionDao=>List=>Select: %s", err))
			return nil, err
		}
		return nil, nil
	}

	return results, nil
}
