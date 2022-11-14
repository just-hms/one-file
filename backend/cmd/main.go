package main

import (
	"one-file/controllers"
	"one-file/models"
)

func main() {
	models.Build()
	controllers.HandleRequests()
}
