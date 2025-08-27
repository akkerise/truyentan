package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	app "github.com/truyentan/backend/internal/app"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		srv := app.NewServer()
		if err := srv.Run("127.0.0.1:" + port); err != nil {
			log.Fatal(err)
		}
	}()

	bindings := NewBindings()

	err := wails.Run(&options.App{
		Title:       "truyen-reader",
		AssetServer: &assetserver.Options{Assets: assets},
		OnStartup:   bindings.startup,
		Bind:        []interface{}{bindings},
	})
	if err != nil {
		log.Fatal(err)
	}
}
