package server

import (
	"ServerSite/internal/api"
	"ServerSite/internal/mw"
	"ServerSite/internal/service"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"html/template"
	"io"
	"path"
)

type Server struct {
	echo    *echo.Echo
	service *service.Service
}
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func NewServer(api *api.Api) *Server {
	s := &Server{echo: echo.New()}
	s.echo.Static("/static", "static")

	templates := make(map[string]*template.Template)

	//templates["auth.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "auth.html"), path.Join("static", "templates", "base.html")))
	//templates["auth2.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "auth2.html"), path.Join("static", "templates", "base.html")))
	templates["auth3.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "auth3.html"), path.Join("static", "templates", "base.html")))
	templates["bd.html"] = template.Must(template.ParseFiles(path.Join("static", "templates", "bd.html"), path.Join("static", "templates", "base.html")))

	s.echo.Renderer = &TemplateRegistry{templates: templates}

	mw := mw.NewMW(api.Service)

	s.echo.GET("/auth", api.AuthHandler3)
	s.echo.GET("/bd", api.BdHandler, mw.CheckAuthorization)

	s.echo.POST("/api/post-auth", api.PostAuth)
	s.echo.GET("/api/get-data-db", api.GetDataDb, mw.CheckAuthorization)
	s.echo.GET("/api/get-file-bd", api.GetDownloadDB, mw.CheckAuthorization)

	s.echo.POST("/api/save-change-bd/change", api.PostChangeBD, mw.CheckAuthorization)
	s.echo.POST("/api/save-change-bd/add", api.PostNewAddBD, mw.CheckAuthorization)
	return s
}

func (s *Server) ServerStart(ctx context.Context, addr string) error {
	err := s.echo.Start(addr)
	defer s.echo.Close()
	return err
}
