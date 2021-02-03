package app

import (
	"stockup-go/config"
	"stockup-go/helper"
	"stockup-go/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	helper.Log("init()")
	router = gin.Default()
	// router.Use(InitServerHeader())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("OPTION")
	corsConfig.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}

	router.Use(cors.New(corsConfig))
}

func StartApp() {
	helper.Log("StartApp()")
	// init database
	config.DbConn()

	Url()

	configs := util.Configs()
	apiPort := configs.Find("API_PORT")
	if apiPort == "not found" {
		apiPort = "8080"
	}

	if configs.Find("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Run(":" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
