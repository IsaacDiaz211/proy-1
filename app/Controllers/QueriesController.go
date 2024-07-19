package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "queries.html", gin.H{"title": "Información de contacto"})
}
