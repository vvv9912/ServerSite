package api

import (
	"ServerSite/internal/model"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (s *Api) PostNewAddBD(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(body)
		return err
	}

	//todo перенос в слой сервиса
	var a []ProductsGet
	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Println(err)
		return err
	}

	photosBytes := make([][]byte, 0)
	for i, product := range a {
		log.Println("POST:save:", product.Article, product.Name, "len:photo", len(product.PhotoUrl))
		for _, base64Img := range product.PhotoUrl {
			prefix := "data:image/jpeg;base64,"
			encodedString := strings.TrimPrefix(base64Img, prefix)
			data, _ := base64.StdEncoding.DecodeString(encodedString)
			photosBytes = append(photosBytes, data)
		}

		err = s.Storage.AddProduct(context.TODO(), model.Products{
			Article:     a[i].Article,
			Catalog:     a[i].Catalog,
			Name:        a[i].Name,
			Description: a[i].Description,
			PhotoUrl:    photosBytes,
			Price:       a[i].Price,
			Length:      a[i].Length,
			Width:       a[i].Width,
			Height:      a[i].Height,
			Weight:      a[i].Weight,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
		photosBytes = make([][]byte, 0)
	}

	return c.JSONBlob(
		http.StatusOK,
		[]byte(""),
	)
}

func (s *Api) PostChangeBD(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(body)
		return err
	}

	//todo перенос в слой сервиса
	var a []ProductsGet
	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Println(err)
		return err
	}
	photosBytes := make([][]byte, 0)
	for i, product := range a {
		log.Println("POST:changed:", product.Article, product.Name, "len:photo", len(product.PhotoUrl))
		for _, base64Img := range product.PhotoUrl {
			prefix := "data:image/jpeg;base64,"
			encodedString := strings.TrimPrefix(base64Img, prefix)
			data, _ := base64.StdEncoding.DecodeString(encodedString)
			photosBytes = append(photosBytes, data)
			// Теперь `data` содержит байтовый массив изображения
		}

		err = s.Storage.ChangeProductByArticle(context.TODO(), model.Products{
			Article:     a[i].Article,
			Catalog:     a[i].Catalog,
			Name:        a[i].Name,
			Description: a[i].Description,
			PhotoUrl:    photosBytes,
			Price:       a[i].Price,
			Length:      a[i].Length,
			Width:       a[i].Width,
			Height:      a[i].Height,
			Weight:      a[i].Weight,
		})
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return c.JSONBlob(
		http.StatusOK,
		[]byte(""),
	)

}
