package api

import (
	"ServerSite/internal/exel"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mholt/archiver"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
)

func (s *Api) GetDownloadDB(c echo.Context) error {
	products, err := s.Storage.SelectAllProducts(context.TODO())
	if err != nil {
		fmt.Println(err)
		c.Response().WriteHeader(http.StatusBadRequest)
		return err
	}

	pathBD := "testBD"
	_, err = os.Stat(pathBD)
	if !os.IsNotExist(err) {
		log.Println("!os.IsNotExist1")
		err = os.RemoveAll(pathBD)
		if err != nil {
			log.Println("!os.IsNotExist(err)", err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}
	}

	err = os.Mkdir(pathBD, os.ModePerm)
	if err != nil {
		log.Println("os.Mkdir", err)
		log.Println("pathBD:", pathBD)
	}

	ex2 := exel.NewExcel(path.Join(pathBD, "test.xlsx"))

	f := excelize.NewFile()
	ex2.CreateFile(f)
	ex2.WriteStamp("Catalog")

	for i := range products {
		var photos []string
		pathCatalog := path.Join(pathBD, products[i].Catalog)
		if len(products[i].PhotoUrl) != 0 {
			if _, err := os.Stat(pathCatalog); os.IsNotExist(err) {
				os.Mkdir(pathCatalog, os.ModePerm)
			} else {
				log.Println("Папка уже существует:", pathCatalog)
			}
			if err != nil {
				c.Response().WriteHeader(http.StatusBadRequest)
				return err
			}
		}

		for k := range products[i].PhotoUrl {
			fileName := products[i].Catalog + "_" + products[i].Name + "_" + generateRandomName()
			err := ioutil.WriteFile(path.Join(pathCatalog, fileName), products[i].PhotoUrl[k], 0644)
			if err != nil {
				log.Println("ioutil.WriteFile", err)
				c.Response().WriteHeader(http.StatusBadRequest)
				return err
			}
			photos = append(photos, "/"+products[i].Catalog+"/"+fileName)
		}

		ex2.Write(i+2, exel.ProductsPars{
			Article:     products[i].Article,
			Catalog:     products[i].Catalog,
			Name:        products[i].Name,
			Description: products[i].Description,
			PhotoUrl:    photos,
			Price:       products[i].Price,
			Length:      products[i].Length,
			Width:       products[i].Width,
			Height:      products[i].Height,
			Weight:      products[i].Weight,
		}, "Catalog")
	}

	ex2.CloseFile()

	fileZip := "bd.zip"
	_, err = os.Stat(fileZip)
	if !os.IsNotExist(err) {
		log.Println("!os.IsNotExist2")
		err = os.RemoveAll(fileZip)
		if err != nil {
			log.Println("!os.IsNotExist(err)2", err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}
	}
	err = archiver.Archive([]string{pathBD}, fileZip)

	if err != nil {
		log.Println("archiver.Archive([]string{pathBD}, fileZip)", err)
		c.Response().WriteHeader(http.StatusBadRequest)
		return err
	}

	_, err = os.Stat(pathBD)
	if !os.IsNotExist(err) {
		log.Println("!os.IsNotExist3")
		err = os.RemoveAll(pathBD)
		if err != nil {
			log.Println("!os.IsNotExist(err)3", err)
			c.Response().WriteHeader(http.StatusBadRequest)
			return err
		}
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/zip")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename="+fileZip)
	c.Response().WriteHeader(http.StatusOK)

	return c.File(fileZip)

}
func generateRandomName() string {
	// Генерируем случайное имя для файла
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(randomBytes) + ".jpg"
}
