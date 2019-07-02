package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "create.html", map[string]interface{}{
		"name": "HOME",
		"msg":  "Framework Echo Workedsss!sss",
	})
}
