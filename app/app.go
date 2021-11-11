package app

import (
	"log"

	"TubesBigData/config"
	"TubesBigData/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	router *fiber.App
)

func StartApplication() {
	serverPort := config.Config("PORT")
	// Try connecting to the database
	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	router = fiber.New()
	router.Use(cors.New())
	mapURLs()
	_ = router.Listen(":" + serverPort)
}
