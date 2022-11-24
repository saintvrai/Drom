package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/saintvrai/Drom/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createCar)
			lists.GET("/", h.getCarsList)
			lists.GET("/:id", h.getCarById)
			lists.PUT("/:id", h.updateCarById)
			lists.DELETE("/:id", h.deleteById)
		}
	}
	return router
}
