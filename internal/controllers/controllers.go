package controllers

import (
	"net/http"

	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/models"
	"github.com/gin-gonic/gin"
)

func ShowSites(c *gin.Context) {
	var site []models.Site
	database.DB.Find(&site)
	c.JSON(200, site)
}

func CreateSite(c *gin.Context) {
	var site models.Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error: ": err.Error()})
		return
	}
	database.DB.Create(&site)
	c.JSON(http.StatusOK, site)
}
