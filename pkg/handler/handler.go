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
		cars := api.Group("/cars")
		{
			cars.POST("/", h.createCar)
			cars.GET("/", h.getCarsList)
			cars.GET("/:id", h.getCarById)
			cars.PUT("/:id", h.updateCarById)
			cars.DELETE("/:id", h.deleteById)
			cars.GET("/getall", h.getAllCarsAndClients)
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
