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

	gUser := appApiV1.Group("users")
	{
		db := config.Connect()
		r := repository.NewUserRepository(db)
		s := service.NewUserService(r)
		c := controller.NewUserController(s)

		gUser.POST("", c.Create)

		gUser.Use(middleware.CheckApiToken())
		gUser.GET("me", c.Me)
	}

	return route
}
