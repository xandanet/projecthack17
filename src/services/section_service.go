package services

import (
	"podcast/src/domains/sections"
	"podcast/src/utils"
)

type SectionServiceI interface {
	List(req *sections.SectionListInput) ([]sections.SectionDTO, utils.RestErrorI)
}

type sectionService struct{}

var SectionService SectionServiceI = &sectionService{}

func (s *sectionService) List(req *sections.SectionListInput) ([]sections.SectionDTO, utils.RestErrorI) {
	result, err := sections.SectionDao.List(req.PodcastID)
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	return result, nil
}
