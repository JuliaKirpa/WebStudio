package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in")
	}
	order := router.Group("/order")
	{
		client := order.Group("/client")
		{
			client.POST("/")
			client.GET("/")
			client.GET("/:id")
			client.PUT("/:id")
			client.DELETE("/:id")
		}

		servise := router.Group("/service")
		{
			servise.POST("/")
			servise.GET("/")
			servise.DELETE("/:id")
		}
	}
	return router
}
