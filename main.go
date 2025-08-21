package main

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pxgo/go-fm/middlewares"
	"github.com/pxgo/go-fm/modules"
	"github.com/pxgo/go-fm/routes"
	"io/fs"
	"net/http"
	"os"
)

//go:embed web/*
var publicFiles embed.FS

func main() {

	modules.InitReader()

	e := echo.New()

	e.HideBanner = true
	e.HTTPErrorHandler = middlewares.CustomHTTPErrorHandler
	e.Use(middlewares.LoggerIn)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	var fileSystem http.FileSystem
	if modules.Config.Debug {
		fileSystem = http.FS(os.DirFS("web"))
	} else {
		fsys, err := fs.Sub(publicFiles, "web")
		if err != nil {
			modules.Logger.Error(err)
			panic(err)
		}

		fileSystem = http.FS(fsys)
	}

	assetHandler := http.FileServer(fileSystem)

	routes.InitRoutes(e)

	e.GET("/*", echo.WrapHandler(assetHandler))

	err := e.Start(fmt.Sprintf("%s:%d", modules.Config.Host, modules.Config.Port))
	if err != nil {
		modules.Logger.Error(err)
	}
}
