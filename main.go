package main

import (
	"add-service/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":3035")
}
