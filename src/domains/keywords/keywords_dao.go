package keywords

import (
	"database/sql"
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type KeywordDaoI interface {
	Create(word string) error
	CreateRelationshipSubtitle(keywordID int64, subtitleID int64) error
	FindID(word string) (int64, error)
}

type keywordDao struct{}

var KeywordDao KeywordDaoI = &keywordDao{}

func (d *keywordDao) Create(word string) error {
	//Create the record
	if _, err := mysql.Client.Exec(queryCreate, word); err != nil {
		zlog.Logger.Error(fmt.Sprintf("KeywordDao=>Create=>Select: %s", err))
		return err
	}

	return nil
}

func (d *keywordDao) CreateRelationshipSubtitle(keywordID int64, subtitleID int64) error {
	//Create the record
	if _, err := mysql.Client.Exec(queryCreateRelationshipSubtitle, keywordID, subtitleID); err != nil {
		zlog.Logger.Error(fmt.Sprintf("KeywordDao=>CreateRelationshipSubtitle=>Select: %s", err))
		return err
	}

	return nil
}

func (d *keywordDao) FindID(word string) (int64, error) {
	result := int64(0)

	if err := mysql.Client.Get(&result, queryFindID, word); err != nil {
		if err != sql.ErrNoRows {
			zlog.Logger.Error(fmt.Sprintf("KeywordDao=>FindID=>Select: %s", err))
		}
		return result, err
	}

	return result, nil
}
