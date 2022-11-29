package services

import (
	"podcast/src/domains/searches"
	"podcast/src/utils"
)

type SearchServiceI interface {
	ListLocations() ([]searches.SearchLocationsOutput, utils.RestErrorI)
}

type searchService struct{}

var SearchService SearchServiceI = &searchService{}

func (s *searchService) ListLocations() ([]searches.SearchLocationsOutput, utils.RestErrorI) {
	result, err := searches.SearchDao.ListLocations()
	if err != nil {
		return nil, utils.NewInternalServerError(utils.ErrorGetList)
	}

	return result, nil
}
