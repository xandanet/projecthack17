package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"podcast/src/domains/plays"
	"podcast/src/services"
	"podcast/src/utils"
)

type playControllerInterface interface {
	Create(ctx *gin.Context)
	Seed(ctx *gin.Context)
}

type playController struct{}

var PlayController playControllerInterface = &playController{}

func (c *playController) Create(ctx *gin.Context) {
	var input plays.PlayCreateInput

	if ok := utils.GinShouldPassAll(ctx,
		utils.GinShouldBind(&input),
		utils.GinShouldValidate(&input),
	); !ok {
		return
	}

	err := services.PlayService.Create(&input)

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: "CREATED",
		Code: http.StatusOK,
	})
}

func (c *playController) Seed(ctx *gin.Context) {
	err := services.PlayService.Seed()

	if err != nil {
		ctx.JSON(err.Code(), err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoErrorData{
		Data: "SEEDED",
		Code: http.StatusOK,
	})
}
