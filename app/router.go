package app

import (
	_ "database/sql"
	Controller "stockup-go/controllers"
	ControllerPage "stockup-go/controllers/page"
	"stockup-go/database/migrations"
	"stockup-go/database/seeders"
	"stockup-go/helper"
	// _ "github.com/go-sql-driver/mysql"
	// "gitlab.com/fireexitco/cmu-aqi-monitor/backend/config"
	// "stockup-go/config"
	// "gitlab.com/fireexitco/cmu-aqi-monitor/backend/controller/government"
	// L "gitlab.com/fireexitco/cmu-aqi-monitor/backend/controller/login"
	// "gitlab.com/fireexitco/cmu-aqi-monitor/backend/models"
)

func Url() {
	helper.Log("Url()")
	// router.GET("/ping", func(c *gin.Context) {

	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error":   "not_found",
	// 		"message": "aqi data not found for village 1",
	// 	})

	// })

	router.POST("/api/login-by-device-id", Controller.LoginByDeviceID)
	router.POST("/api/login", Controller.Login)
	router.POST("/api/logout", Controller.Logout)
	router.POST("/api/save-register", Controller.SaveRegister)

	router.GET("/api/image/:folder/:id", Controller.GetImage)
	router.GET("/api/get-company-by-is-register", Controller.GetCompanyByIsRegister)

	router.POST("/api/get-data-dashboard", ControllerPage.GetDataDashboard)

	router.GET("/run-migrations", migrations.RunMigrations)

	router.GET("/run-seeders", seeders.RunSeeders)

	// router.GET("/text", Text)
	// var rrouter *gin.Engine
	// rrouter = gin.Default()
	// ccorsConfig := cors.DefaultConfig()
	// ccorsConfig.AllowAllOrigins = true
	// ccorsConfig.AddAllowMethods("OPTION")
	// ccorsConfig.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"}

	// rrouter.Use(cors.New(ccorsConfig))
	// rrr.
	// router.Static("/storage/stock.jpg", "./storage/stock.jpg")
	// router.Static("/txt", yourHandler)

	// router.POST("/change-password", TokenAuthMiddleware(), L.ChangePassword)
	// router.GET("/forget-password", L.ForgetPassword)
	// router.POST("/reset-password", TokenAuthMiddleware(), L.ResetPassword)

	// router.GET("/me", TokenAuthMiddleware(), L.GetMe)
	// router.GET("/get-user", TokenAuthMiddleware(), L.GetUser)
	// router.GET("/get-list-user", TokenAuthMiddleware(), L.GetListUser)
	// router.POST("/create-user", TokenAuthMiddleware(), L.CreateUser)
	// router.POST("/edit-user/:user_id", TokenAuthMiddleware(), L.EditUser)
	// router.POST("/delete-user/:user_id", TokenAuthMiddleware(), L.DeleteUser)

	// router.POST("/add-log", TokenAuthMiddleware(), L.AddLog)
	// router.GET("/get-list-log", TokenAuthMiddleware(), L.GetListLog)

	// router.GET("/get-list-government", TokenAuthMiddleware(), government.GetListGovernment)
	// router.GET("/get-government", TokenAuthMiddleware(), government.GetGovernment)
	// router.POST("/create-government", TokenAuthMiddleware(), government.CreateGovernment)
	// router.POST("/edit-government/:government_id", TokenAuthMiddleware(), government.EditGovernment)
	// router.POST("/delete-government/:government_id", TokenAuthMiddleware(), government.DeleteGovernment)
	// router.GET("/get-list-subdistrict", TokenAuthMiddleware(), government.GetSubdistrict)
	// router.GET("/get-list-village", TokenAuthMiddleware(), government.GetVillage)
	// router.GET("/get-all-village", government.GetAllVillage)

	// router.GET("/get-aqi", Aql.GetAqi)
	// router.GET("/get-list-aqi", TokenAuthMiddleware(), Aql.GetListAqi)
	// router.GET("/test-data", Aql.TestData)

}

// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.Request.Header["Authorization"]
// 		if token == nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
// 			return
// 		}
// 		var user models.Users
// 		config.GetDB().Select("users.id,users.`name`,users.surname,users.email,users.phone").
// 			Joins("left join user_tokens on user_tokens.user_id = users.id").
// 			Where("token = ?", token).
// 			First(&user)
// 		if user.Id == 0 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization, please login again"})
// 			return
// 		}
// 		c.Set("userID", user.Id)
// 		c.Next()
// 	}
// }
