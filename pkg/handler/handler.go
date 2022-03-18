package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
	}
	order := router.Group("/order")
	{
		order.POST("/", h.createOrder)
		order.GET("/", h.getAllOrders)
		order.GET("/:id", h.getOrderById)
		order.PUT("/:id", h.updateOrder)
		order.DELETE("/:id", h.deleteOrder)

		client := order.Group("/client")
		{
			client.POST("/", h.createClient)
			client.GET("/", h.getClient)
			client.GET("/:id", h.getClientById)
			client.PUT("/:id", h.updateClient)
			client.DELETE("/:id", h.deleteClient)
		}

		servise := order.Group("/service")
		{
			servise.POST("/", h.addService)
			servise.DELETE("/:id", h.deleteService)
		}
	}
	return router
}
