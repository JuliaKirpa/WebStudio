package handler

import (
	"ClientBaseWEBStudio/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
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
			client.GET("/", h.getAllClients)
			client.GET("/:id", h.getClientById)
			client.PUT("/:id", h.updateClient)
			client.DELETE("/:id", h.deleteClient)
		}

		product := order.Group("/service")
		{
			product.POST("/", h.addProduct)
			product.DELETE("/:id", h.deleteProduct)
		}
	}
	return router
}
