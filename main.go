package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"literm/goFile/service"
	"log"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := &App{}
	log.Println("初始化CollectionService")
	collectionService := &service.CollectionService{}

	err := app.startup()
	if err != nil {
		println("Error:", err.Error())
	}
	// Create application with options
	err = wails.Run(&options.App{
		Title:  "literm",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 245, G: 245, B: 245, A: 100},
		Bind: []interface{}{
			app,
			collectionService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
