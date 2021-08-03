package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/virusdefender/site-comment/api"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	e := echo.New()
	e.IPExtractor = func(request *http.Request) string {
		xff := request.Header.Get(echo.HeaderXForwardedFor)
		if xff == "" {
			return ""
		}
		s := strings.Split(xff, ",")
		if len(s) == 0 {
			return ""
		}
		return strings.TrimSpace(s[len(s)-1])
	}
	e.Use(middleware.Recover())
	e.POST("/api/comment", api.CreateComment)
	e.GET("/api/comment", api.ListComment)
	e.DELETE("/api/comment", api.DeleteComment)

	e.GET("/api/backup", api.BackupDB)

	e.HTTPErrorHandler = func(err error, context echo.Context) {
		code := http.StatusOK
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		context.JSON(code, map[string]interface{}{
			"err": err.Error(),
		})
	}
	e.Logger.Fatal(e.Start(":" + port))
}
