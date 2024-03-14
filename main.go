package main

import (
	_ "github.com/abylaymoldabek/finSystem/docs"
	"github.com/abylaymoldabek/finSystem/internal/app"
)

const configsDir = "config.json"

// @title Tag Products API
// @version 1.0
// @description This is a sample API for managing products.
// @host localhost:8088

func main() {

	app.Run(configsDir)
}
