package routers

import (
	"OctavianoRyan25/GinGo/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controller.CreateCar)
	router.PUT("/cars/:carID", controller.UpdateCar)
	router.GET("/cars/:carID", controller.GetCar)
	router.GET("/cars/all", controller.GetAllCar)
	router.DELETE("/cars/:carID", controller.DeleteCar)
	return router
}