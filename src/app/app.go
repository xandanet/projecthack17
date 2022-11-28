package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"podcast/src/clients/mysql"
	"podcast/src/zlog"
)

type App struct {
	DB     *sqlx.DB
	Router *gin.Engine
}

type Credentials struct {
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
	DBName2 string
	DBTls   string
	AppPort string
}

func StartApplication(c *Credentials) {
	db := mysql.ConnectToDatabase(c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, c.DBTls)

	application := &App{
		DB:     db,
		Router: gin.Default(),
	}

	application.SetupRoutes()

	zlog.Logger.Info(fmt.Sprintf("Starting application on port: %s", c.AppPort))
	if err := application.Router.Run(fmt.Sprintf(":%s", c.AppPort)); err != nil {
		zlog.Logger.Panic(fmt.Sprintf("failed to start application: %s", err))
	}
}
