package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"podcast/src/domains/podcasts"
	"podcast/src/domains/sections"
	"podcast/src/services"
	"podcast/src/utils"
	"podcast/src/utils/helpers"
)

type podcastControllerInterface interface {
	List(ctx *gin.Context)
	Single(ctx *gin.Context)
	Subtitles(ctx *gin.Context)
	Interventions(ctx *gin.Context)
	Sentiment(ctx *gin.Context)
	BookMark(ctx *gin.Context)
	GetBookMark(ctx *gin.Context)
}

type podcastController struct{}

var PodcastController podcastControllerInterface = &podcastController{}

func (c *podcastController) List(ctx *gin.Context) {
	result, err := services.PodcastService.List()

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *podcastController) Single(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id := helpers.ConvertStringToInt64(idStr)
	
	result, err := services.PodcastService.Single(id)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *podcastController) Subtitles(ctx *gin.Context) {
	var input sections.SectionListInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.SectionService.List(&input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *podcastController) Interventions(ctx *gin.Context) {
	var input podcasts.PodcastInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.PodcastService.Interventions(input.ID)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *podcastController) Sentiment(ctx *gin.Context) {
	var input podcasts.PodcastInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.PodcastService.Sentiment(input.ID)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *podcastController) BookMark(ctx *gin.Context) {
	var input podcasts.BookmarkInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	err := services.PodcastService.Bookmark(input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: "BOOKMARK SAVED",
		Code: http.StatusOK,
	})
}

func (c *podcastController) GetBookMark(ctx *gin.Context) {
	var input podcasts.BookmarkSearchInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.PodcastService.GetBookmark(input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}
