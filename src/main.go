package main

import (
	"aero-router/src/models"
	"aero-router/src/utils"
	"fmt"
)

func main() {
	fmt.Println("Aero Router started successfully")
	models.Init()
	utils.Test()
}
