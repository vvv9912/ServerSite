package api

import (
	"ServerSite/internal/logger"
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

	for i, product := range a {
		photosBytes := make([][]byte, 0)
		logger.Log.CustomInfo("Post save", map[string]interface{}{
			"article":   product.Article,
			"catalog":   product.Catalog,
			"name":      product.Name,
			"lenPhoto":  len(product.PhotoUrl),
			"sizePhoto": len(product.PhotoUrl[i]),
		})

		for _, base64Img := range product.PhotoUrl {
			prefix := "data:image/jpeg;base64,"
			encodedString := strings.TrimPrefix(base64Img, prefix)
			//todo
			if len(encodedString) > 0 && encodedString[0] == ' ' {
				encodedString = strings.TrimSpace(encodedString)
			}
			data, err := base64.StdEncoding.DecodeString(encodedString)
			if err != nil {
				logger.Log.AddOriginalError(err).CustomError("Err base64 encoding", map[string]interface{}{
					"encodedString": encodedString,
				})
				return err
			}
			photosBytes = append(photosBytes, data)
		}
		for k := range photosBytes {
			logger.Log.CustomInfo("SizePhoto", map[string]interface{}{
				"sizePhoto": photosBytes[k],
			})
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

		logger.Log.CustomInfo("Post changed", map[string]interface{}{
			"article":   product.Article,
			"catalog":   product.Catalog,
			"name":      product.Name,
			"lenPhoto":  len(product.PhotoUrl),
			"sizePhoto": len(product.PhotoUrl[i]),
		})

		for _, base64Img := range product.PhotoUrl {
			prefix := "data:image/jpeg;base64,"
			encodedString := strings.TrimPrefix(base64Img, prefix)
			//todo
			if len(encodedString) > 0 && encodedString[0] == ' ' {
				encodedString = strings.TrimSpace(encodedString)
			}

			data, err := base64.StdEncoding.DecodeString(encodedString)
			if err != nil {
				logger.Log.AddOriginalError(err).CustomError("Err base64 encoding", map[string]interface{}{
					"encodedString": encodedString,
					"base64Img":     base64Img})
				return err
			}
			photosBytes = append(photosBytes, data)
		}
		for k := range photosBytes {
			logger.Log.CustomInfo("SizePhoto", map[string]interface{}{
				"sizePhoto": photosBytes[k],
			})
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
