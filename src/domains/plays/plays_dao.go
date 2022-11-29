package plays

import (
	"fmt"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type PlayDaoI interface {
	Create(req *PlayCreateInput) error
}

type playDao struct{}

var PlayDao PlayDaoI = &playDao{}

func (d *playDao) Create(req *PlayCreateInput) error {
	if _, err := mysql.Client.Exec(queryCreate, req.PodcastID, req.Position, req.CreatedAt); err != nil {
		zlog.Logger.Error(fmt.Sprintf("SubtitleDao=>List=>Select: %s", err))
		return nil
	}

	return nil
}
