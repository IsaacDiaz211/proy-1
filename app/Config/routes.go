package config

import (
	controllers "proy-1/app/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Static("/assets", "./assets")
	/*router.Static("/assets/css", "./assets/css")
	router.Static("/assets/js", "./assets/js")
	router.Static("/assets/img", "./assets/img")*/
	router.LoadHTMLGlob("app/Views/*")

	router.GET("/", controllers.GetHome)
	router.GET("/catalogue", controllers.GetCatalogue)
	router.GET("/info", controllers.GetInfo)
	router.GET("/about-us", controllers.GetAbout)
	router.GET("/commercialization", controllers.GetCommer)
	router.GET("/terms", controllers.GetTerms)

	return router
}
