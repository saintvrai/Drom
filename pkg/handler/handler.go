package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/saintvrai/Drom/docs"
	"github.com/saintvrai/Drom/pkg/service"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createCar)
			lists.GET("/", h.getCarsList)
			lists.GET("/:id", h.getCarById)
			lists.PUT("/:id", h.updateCarById)
			lists.DELETE("/:id", h.deleteById)
		}
		clients := api.Group("/clients")
		{
			clients.POST("/", h.createClient)
			clients.GET("/", h.getClients)
			clients.GET("/:id", h.getClientById)
			clients.PUT("/:id", h.updateClientById)
			clients.DELETE("/:id", h.deleteClientById)
		}
	}
	return router
}
