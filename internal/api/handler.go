package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Api) AuthHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "auth.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/styleauth.css",
	})
}
func (s *Api) AuthHandler2(c echo.Context) error {

	return c.Render(http.StatusOK, "auth2.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/styleauth2.css",
	})
}
func (s *Api) AuthHandler3(c echo.Context) error {

	return c.Render(http.StatusOK, "auth3.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/styleauth3.css",
	})
}
func (s *Api) BdHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "bd.html", map[string]interface{}{
		"name":    "AUTH",
		"CSSPath": "/static/css/style_bd.css",
	})

}
