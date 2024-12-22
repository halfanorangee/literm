package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"literm/goFile/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
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
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Bind: []interface{}{
			app,
			collectionService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
