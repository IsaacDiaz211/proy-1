package main

import Config "proy-1/app/Config"

func main() {
	router := Config.SetupRoutes()
	router.Run(":8080")
}
