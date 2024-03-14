package api

import (
	"ServerSite/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Api) PostAuth(c echo.Context) error {

	auth := new(model.Auth)
	if err := c.Bind(auth); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	jwt, err := s.Service.GetTokenAuthorization(auth.Login, auth.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "Invalid credentials"})
	}

	cookie := new(http.Cookie)
	cookie.Name = "JWT"
	cookie.Value = jwt

	c.SetCookie(cookie)

	return c.JSON(
		http.StatusOK, map[string]interface{}{
			"success": true,
			"jwt":     jwt,
		})
}
