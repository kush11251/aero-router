package controllers

import (
	"aero-router/src/models"
	"fmt"
)

func GetRoute() {
	fmt.Println("Route retrieved successfully")
	models.Init()
}
