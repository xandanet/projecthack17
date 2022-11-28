package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthControllerInterface interface {
	Ping(ctx *gin.Context)
}

type healthController struct{}

var HealthController healthControllerInterface = &healthController{}

// Ping godoc
// @Summary Health Check
// @Description Returns a 200 response if the system is running
// @Produce json
// @Success 200
// @Failure 500
// @Router / [get]
func (c *healthController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "")
}
