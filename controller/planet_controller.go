package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PlanetController interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Show(c *gin.Context)
	Search(c *gin.Context)
	Delete(c *gin.Context)
}

type planetController struct {
}

func NewPlanetController() PlanetController {
	return &planetController{}
}

func (controller *planetController) Create(c *gin.Context) {
	log.Println("Creating a planet")
}

func (controller *planetController) List(c *gin.Context) {
	log.Println("Listing all planets")
}

func (controller *planetController) Show(c *gin.Context) {
	log.Println("Showing a planet by id")
}

func (controller *planetController) Search(c *gin.Context) {
	log.Println("Searching a planet")
}

func (controller *planetController) Delete(c *gin.Context) {
	log.Println("Deleting a planet")
}
