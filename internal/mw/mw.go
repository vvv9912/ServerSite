package mw

import (
	"ServerSite/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MW struct {
	*service.Service
}

func NewMW(s *service.Service) *MW {
	return &MW{s}
}

func (m *MW) CheckAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if len(c.Cookies()) != 0 {
			token, err := c.Cookie("jwt")

			id, role, err := m.Authorization.GetUserId(token.Value)

			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/auth")
			}

			ok, roleDb, err := m.DbUser.CheckId(id)
			if err != nil || !ok {
				return c.Redirect(http.StatusTemporaryRedirect, "/auth")
			}

			if role != roleDb {
				return c.Redirect(http.StatusTemporaryRedirect, "/auth")
			}

			c.Set("id", id)
			c.Set("role", role)
			return next(c)
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/auth")
	}
}
