package routes

import (
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/sites", controllers.ShowSites)
	r.POST("/add_site", controllers.CreateSite)
	r.Run()
}
