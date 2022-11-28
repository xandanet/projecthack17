package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"podcast/src/domains/subtitles"
	"podcast/src/services"
	"podcast/src/utils"
)

type subtitleControllerInterface interface {
	List(ctx *gin.Context)
	Search(ctx *gin.Context)
	Topics(ctx *gin.Context)
}

type subtitleController struct{}

var SubtitleController subtitleControllerInterface = &subtitleController{}

func (c *subtitleController) List(ctx *gin.Context) {
	result, err := services.SubtitleService.List()

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *subtitleController) Search(ctx *gin.Context) {
	var input subtitles.SubtitleSearchInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.SubtitleService.Search(&input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *subtitleController) Topics(ctx *gin.Context) {
	err := services.SubtitleService.ParseAllPodcasts()

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: "PODCASTS_PARSED",
		Code: http.StatusOK,
	})
}
