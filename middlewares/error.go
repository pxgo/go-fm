package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/pxgo/go-fm/modules"
	"github.com/pxgo/go-fm/settings"
	"github.com/pxgo/go-fm/tools"
	"net/http"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var status int
	var resType string
	var message string

	if iErr, ok := err.(tools.IError); ok {
		status = int(iErr.Status)
		resType = string(iErr.Type)
		message = string(iErr.Type)
	} else {
		status = http.StatusInternalServerError
		resType = string(settings.ResponseTypes.ServerInternalError)
		message = err.Error()
	}
	modules.Logger.Error(err)
	err = c.JSON(status, settings.IResponseBody{
		Code:    0,
		Type:    resType,
		Message: message,
	})
	if err != nil {
		modules.Logger.Error(err)
	}
}
