package main

import (
	"context"
	"embed"
	"log"
	"os"
	"path/filepath"
	rt "runtime"

	"Pebble/internal/database"
	"Pebble/internal/handlers"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Initialize database
	homeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(homeDir, "砾石", "砾石.db")
	if err := database.InitDB(dbPath); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Create handlers
	projectHandler := handlers.NewProjectHandler()
	issueHandler := handlers.NewIssueHandler()
	documentHandler := handlers.NewDocumentHandler()
	tagHandler := handlers.NewTagHandler()
	searchHandler := handlers.NewSearchHandler()
	settingsHandler := handlers.NewSettingsHandler()
	activityHandler := handlers.NewActivityHandler()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "砾石",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         rt.GOOS != "darwin",
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		DragAndDrop: DragAndDropOptions(),
		OnDomReady: func(ctx context.Context) {
			runtime.OnFileDrop(ctx, func(x, y int, paths []string) {
			})
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.WARNING,
		OnStartup:        app.startup,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
			projectHandler,
			issueHandler,
			documentHandler,
			tagHandler,
			searchHandler,
			settingsHandler,
			activityHandler,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
			WebviewUserDataPath: "",
			BackdropType:        windows.Acrylic,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "砾石",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func DragAndDropOptions() *options.DragAndDrop {
	if rt.GOOS == "windows" {
		return &options.DragAndDrop{
			EnableFileDrop:     true,
			DisableWebViewDrop: true,
		}
	} else {
		return &options.DragAndDrop{
			EnableFileDrop: true,
		}
	}
}