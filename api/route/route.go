package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/hiroki-Fukumoto/matching-app/api/controller"
	"github.com/hiroki-Fukumoto/matching-app/api/controller/service"

	_ "github.com/hiroki-Fukumoto/matching-app/api/docs"

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

	return route
}
