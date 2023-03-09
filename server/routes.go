package server

import (
	"bunzz-fizz-buzz/controllers"
	"bunzz-fizz-buzz/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files

	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes for healthcheck of api server
	healthcheck := router.Group("health")
	{
		health := new(controllers.HealthController)
		ping := new(controllers.PingController)
		healthcheck.GET("/health", health.Status)
		healthcheck.GET("/ping", ping.Ping)
	}

	// Routes for fizzbuzz api server
	fizzbuzz := router.Group("fizzbuzz")
	{
		fizz := new(controllers.FizzBuzzController)
		fizzbuzz.GET("/messages", fizz.GetMessageByCount)
		fizzbuzz.POST("/fizzbuzz", fizz.FizzBuzz)
	}

	// Routes for swagger
	swagger := router.Group("swagger")
	{
		// programatically set swagger info
		docs.SwaggerInfo.Title = "FizzBuzz Service"
		docs.SwaggerInfo.Description = "This is a fizz buzz backend written in Go."
		docs.SwaggerInfo.Version = "1.0"
		// docs.SwaggerInfo.Host = "cloudfactory.swagger.io"
		// docs.SwaggerInfo.BasePath = "/v1"

		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		swagger.POST("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	return router

}
