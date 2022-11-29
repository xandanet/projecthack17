package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"podcast/src/domains/segments"
	"podcast/src/services"
	"podcast/src/utils"
)

type segmentControllerInterface interface {
	List(ctx *gin.Context)
	Search(ctx *gin.Context)
	GetContent(ctx *gin.Context)
}

type segmentController struct{}

var SegmentController segmentControllerInterface = &segmentController{}

func (c *segmentController) List(ctx *gin.Context) {
	result, err := services.SegmentService.List()

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *segmentController) Search(ctx *gin.Context) {
	var input segments.SegmentSearchInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.SegmentService.Search(&input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}

func (c *segmentController) GetContent(ctx *gin.Context) {
	var input segments.SubtitleContentInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	result, err := services.SegmentService.GetContent(&input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: result,
		Code: http.StatusOK,
	})
}
