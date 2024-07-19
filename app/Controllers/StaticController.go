package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"title": "Home"})
}

func GetCatalogue(c *gin.Context) {
	c.JSON(200, gin.H{"title": "Catalogo"})
}

func GetAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.html", gin.H{"title": "Sobre nosotros"})
}

func GetCommer(c *gin.Context) {
	c.HTML(http.StatusOK, "commercialization.html", gin.H{"title": "Comercialización"})
}

func GetTerms(c *gin.Context) {
	c.HTML(http.StatusOK, "terms.html", gin.H{"title": "Términos de uso"})
}
