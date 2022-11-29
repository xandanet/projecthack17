package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"podcast/src/controllers"
	"podcast/src/zlog"
	"strings"
	"time"
)

func (app *App) SetupRoutes() {
	app.Router.ForwardedByClientIP = true
	app.Router.RemoteIPHeaders = []string{"X-Forwarded-For"}
	err := app.Router.SetTrustedProxies(nil)
	if err != nil {
		zlog.Logger.Error(err.Error())
	}

	// Adding CORS
	app.Router.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Origin",
			"Cache-Control",
			"X-Requested-With",
		},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "X-Requested-With", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    false,
	}))

	app.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
			"code":  http.StatusNotFound,
		})
	})

	app.Router.GET("", controllers.HealthController.Ping)

	v1Routes := app.Router.Group("/v1")

	//Podcasts
	podcastRoutes := v1Routes.Group("podcasts")
	{
		podcastRoutes.GET("", controllers.PodcastController.List)
		podcastRoutes.GET("/subtitles", controllers.PodcastController.Subtitles)
		podcastRoutes.GET("/interventions", controllers.PodcastController.Interventions)
		podcastRoutes.GET("/sentiment", controllers.PodcastController.Sentiment)
		podcastRoutes.GET("/text", controllers.SegmentController.List)
		podcastRoutes.GET("/search", controllers.SegmentController.Search)
		podcastRoutes.GET("/search/statistics", controllers.SegmentController.Statistics)
		podcastRoutes.POST("/content", controllers.SegmentController.GetContent)
		podcastRoutes.POST("/search/content", controllers.SegmentController.GetContent)
		podcastRoutes.GET("/search-generator", controllers.SegmentController.SearchGenerator)
		podcastRoutes.POST("/bookmark", controllers.PodcastController.BookMark)
		podcastRoutes.GET("/bookmark", controllers.PodcastController.GetBookMark)
	}

	//Plays
	playsRoutes := v1Routes.Group("plays")
	{
		playsRoutes.POST("", controllers.PlayController.Create)
		playsRoutes.POST("/seed", controllers.PlayController.Seed)
		playsRoutes.GET("/statistics", controllers.PlayController.Statistics)
		playsRoutes.GET("/per-day", controllers.PlayController.PerDay)
		playsRoutes.GET("/segment-popularity", controllers.PlayController.SegmentPopularity)
	}
}
