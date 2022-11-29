package keywords

import (
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type KeywordDaoI interface {
	List() ([]KeywordDTO, error)
}

type keywordDao struct{}

var KeywordDao KeywordDaoI = &keywordDao{}

func (d *keywordDao) List() ([]KeywordDTO, error) {
	var result []KeywordDTO

	if err := mysql.Client.Select(&result, queryList); err != nil {
		zlog.Logger.Error(fmt.Sprintf("KeywordDao=>List=>Select: %s", err))
		return nil, err
	}

	return result, nil
}
