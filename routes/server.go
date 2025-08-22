package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pxgo/go-fm/modules"
	"github.com/pxgo/go-fm/tools"
	"net/http"
)

func GetServerInfo(ctx echo.Context) error {
	musicInfo := modules.MusicReader.GetMusicInfo()
	err := ctx.JSON(http.StatusOK, tools.Response.GetResponseBody(struct {
		Name    string              `json:"name"`
		Version string              `json:"version"`
		Time    int64               `json:"time"`
		FMInfo  *modules.IMusicInfo `json:"FMInfo"`
	}{
		Name:    modules.Config.Name,
		Version: modules.Config.Version,
		Time:    modules.Config.Time,
		FMInfo:  musicInfo,
	}))
	if err != nil {
		modules.Logger.Error(err)
		return err
	}
	return nil
}

// PWA Manifest结构
type PWAManifest struct {
	Name            string        `json:"name"`
	ShortName       string        `json:"short_name"`
	Description     string        `json:"description"`
	StartURL        string        `json:"start_url"`
	Display         string        `json:"display"`
	BackgroundColor string        `json:"background_color"`
	ThemeColor      string        `json:"theme_color"`
	Orientation     string        `json:"orientation"`
	Categories      []string      `json:"categories"`
	Icons           []PWAIcon     `json:"icons"`
}

type PWAIcon struct {
	Src     string `json:"src"`
	Sizes   string `json:"sizes"`
	Type    string `json:"type"`
	Purpose string `json:"purpose"`
}

func GetManifest(ctx echo.Context) error {
	// 获取当前应用名称
	appName := modules.Config.Name
	
	manifest := PWAManifest{
		Name:            appName,
		ShortName:       appName,
		Description:     "Progressive Web App for " + appName + " Audio Streaming",
		StartURL:        "/",
		Display:         "standalone",
		BackgroundColor: "#1e293b",
		ThemeColor:      "#1e293b",
		Orientation:     "portrait-primary",
		Categories:      []string{"music", "entertainment"},
		Icons: []PWAIcon{
			{
				Src:     "/cover.jpg",
				Sizes:   "192x192",
				Type:    "image/jpeg",
				Purpose: "any maskable",
			},
			{
				Src:     "/cover.jpg",
				Sizes:   "512x512",
				Type:    "image/jpeg",
				Purpose: "any maskable",
			},
		},
	}

	ctx.Response().Header().Set("Content-Type", "application/manifest+json")
	
	// 直接返回JSON响应
	return ctx.JSON(http.StatusOK, manifest)
}
