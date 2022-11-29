package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"podcast/src/services"
	"podcast/src/utils"
)

type SearchControllerInterface interface {
	ListLocations(ctx *gin.Context)
}

type searchController struct{}

var SearchController SearchControllerInterface = &searchController{}

func (c *searchController) ListLocations(ctx *gin.Context) {
	results, err := services.SearchService.ListLocations()

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: results,
		Code: http.StatusOK,
	})
}
