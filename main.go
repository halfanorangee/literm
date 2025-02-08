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
	shellService := &service.ShellService{}

	err := app.startup()
	if err != nil {
		println("Error:", err.Error())
	}
	// Create application with options
	err = wails.Run(&options.App{
		Title:            "literm",
		WindowStartState: options.Maximised,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 245, G: 245, B: 245, A: 255},
		Bind: []interface{}{
			collectionService,
			shellService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
