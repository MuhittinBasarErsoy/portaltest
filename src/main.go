package main

import (
	"app"
	"config"
	"os"
)

func main() {
	config := config.GetConfig()
	port := os.Getenv("PORT")
	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + port)
}
