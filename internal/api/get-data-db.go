package api

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (s *Api) GetDataDb(c echo.Context) error {

	products, err := s.Storage.SelectAllProducts(context.TODO())
	if err != nil {
		log.Println(err)
		return err
	}

	msg, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSONBlob(http.StatusOK, msg)
}
