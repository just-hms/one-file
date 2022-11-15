package main

import (
	"one-file/pkg/controllers"
	"one-file/pkg/models"
)

func main() {
	models.Build()
	controllers.HandleRequests()
}
