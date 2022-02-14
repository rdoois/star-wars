package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rdoois/star-wars/controller"
)

var (
	planetController controller.PlanetController = controller.NewPlanetController()
)

func init() {
	godotenv.Load()
}

func main() {
	router := gin.Default()

	router.POST("/planets", planetController.Create)
	router.GET("/planets", planetController.List)
	router.GET("/planets/search", planetController.Search)
	router.GET("/planets/:id", planetController.Show)
	router.DELETE("/planets/:id", planetController.Delete)

	router.Run(":" + os.Getenv("PORT"))
}
