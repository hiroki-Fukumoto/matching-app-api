package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/hiroki-Fukumoto/matching-app-api/api/config"
	"github.com/hiroki-Fukumoto/matching-app-api/api/controller"
	"github.com/hiroki-Fukumoto/matching-app-api/api/middleware"
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/service"

	_ "github.com/hiroki-Fukumoto/matching-app-api/api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:8081",
		},
		AllowMethods: []string{
			"*",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	appApiGroup := route.Group("/api")
	appApiV1 := appApiGroup.Group("/v1")

	gHealthCheck := appApiV1.Group("health-check")
	{
		c := controller.NewHealthCheckController()

		gHealthCheck.GET("", c.HealthCheck)
	}

	gInitial := appApiV1.Group("initial")
	{
		s := service.NewInitialService()
		c := controller.NewInitialController(s)

		gInitial.GET("", c.Initial)
	}

	gAuth := appApiV1.Group("")
	{
		ur := repository.NewUserRepository(config.Connect())
		s := service.NewAuthService(ur)
		c := controller.NewAuthController(s)

		gAuth.POST("login", c.Login)
	}

	gPrefecture := appApiV1.Group("prefectures")
	{
		s := service.NewPrefectureService()
		c := controller.NewPrefectureController(s)

		gPrefecture.GET("", c.FindAll)
	}

	gUser := appApiV1.Group("users")
	{
		db := config.Connect()
		r := repository.NewUserRepository(db)
		fr := repository.NewFavoriteRepository(db)
		s := service.NewUserService(r, fr)
		c := controller.NewUserController(s)

		gUser.POST("", c.Create)

		gUser.Use(middleware.CheckApiToken())
		gUser.Use(middleware.DBTransactionMiddleware(db))
		gUser.GET("info/me", c.Me)
		gUser.GET("pickup/today", c.PickupToday)
		gUser.POST("all", c.FindAll)
		gUser.GET(":id", c.FindByID)
		gUser.PATCH("", c.Update)
	}

	gSendLike := appApiV1.Group("likes")
	{
		db := config.Connect()
		r := repository.NewFavoriteRepository(db)
		s := service.NewFavoriteService(r)
		c := controller.NewFavoriteController(s)

		gSendLike.Use(middleware.CheckApiToken())
		gSendLike.POST("", c.SendLike)
		gSendLike.DELETE(":receiverID/cancel", c.CancelLike)
		gSendLike.GET("/send", c.FindSendLikes)
		gSendLike.GET("/receive", c.FindReceiveLikes)
	}

	gMessage := appApiV1.Group("messages")
	{
		db := config.Connect()
		r := repository.NewMessageRepository(db)
		s := service.NewMessageService(r)
		c := controller.NewMessageController(s)

		gMessage.Use(middleware.CheckApiToken())
		gMessage.GET("", c.FindReceiveMessages)
		gMessage.POST("", c.SendMessage)
		gMessage.PUT(":id/read", c.ReadMessage)
	}

	gHobby := appApiV1.Group("hobbies")
	{
		db := config.Connect()
		r := repository.NewHobbyRepository(db)
		s := service.NewHobbyService(r)
		c := controller.NewHobbyController(s)

		gHobby.Use(middleware.CheckApiToken())
		gHobby.GET("", c.FindAll)
	}

	return route
}
