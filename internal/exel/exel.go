package exel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"strings"
)

type Excel struct {
	nameFile string
	f        *excelize.File
}

func NewExcel(nameFile string) *Excel {
	return &Excel{nameFile: nameFile}
}
func (e *Excel) CreateFile(f *excelize.File) {
	e.f = f
}
func (e *Excel) CloseFile() {
	if err := e.f.SaveAs(e.nameFile); err != nil {
		fmt.Println(err)
	}
}

func (e *Excel) WriteStamp(sheet string) {
	//sheet := "Catalog"
	index, err := e.f.NewSheet(sheet)
	if err != nil {
		log.Println(err)
		return
	}

	e.f.SetActiveSheet(index)

	err = e.f.DeleteSheet("Sheet1")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, ArticleCell(1), "Article")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, CatalogCell(1), "Catalog")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, NameCell(1), "Name")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, DescriptionCell(1), "Description")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, PriceCell(1), "Price")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, LengthCell(1), "Length")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, WidthCell(1), "Width")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, HeightCell(1), "Height ")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, WeightCell(1), "Weight")
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, PhotoCell(1), "Photo")
	if err != nil {
		log.Println(err)
		return
	}

}
func (e *Excel) Write(numberCell int, p ProductsPars, sheet string) {

	err := e.f.SetCellValue(sheet, ArticleCell(numberCell), p.Article)
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, CatalogCell(numberCell), p.Catalog)
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, NameCell(numberCell), p.Name)
	if err != nil {
		log.Println(err)
		return
	}
	err = e.f.SetCellValue(sheet, DescriptionCell(numberCell), p.Description)
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, PriceCell(numberCell), fmt.Sprintf("%f", p.Price))
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, LengthCell(numberCell), fmt.Sprintf("%d", p.Length))
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, WidthCell(numberCell), fmt.Sprintf("%d", p.Width))
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, HeightCell(numberCell), fmt.Sprintf("%d", p.Height))
	if err != nil {
		log.Println(err)
		return
	}

	err = e.f.SetCellValue(sheet, WeightCell(numberCell), fmt.Sprintf("%d", p.Weight))
	if err != nil {
		log.Println(err)
		return
	}

	for k := range p.PhotoUrl {
		if k != (len(p.PhotoUrl) - 1) {
			//Добавить путь
			err = e.f.SetCellValue(sheet, PhotoCell(numberCell), fmt.Sprintf("%s,", p.PhotoUrl[k]))
			if err != nil {
				log.Println(err)
				return
			}

		} else {
			err = e.f.SetCellValue(sheet, PhotoCell(numberCell), p.PhotoUrl[k])
			if err != nil {
				log.Println(err)
				return
			}
		}

	}

}

const (
	articleCell     = "A"
	catalogCell     = "B"
	nameCell        = "C"
	descriptionCell = "D"
	priceCell       = "E"
	lengthCell      = "F"
	widthCell       = "G"
	heightCell      = "H"
	weightCell      = "I"
	photoCell       = "J"
)

func ArticleCell(number int) string {
	return fmt.Sprintf("%s%d", articleCell, number)
}
func CatalogCell(number int) string {
	return fmt.Sprintf("%s%d", catalogCell, number)
}
func NameCell(number int) string {
	return fmt.Sprintf("%s%d", nameCell, number)
}
func DescriptionCell(number int) string {
	return fmt.Sprintf("%s%d", descriptionCell, number)
}
func PriceCell(number int) string {
	return fmt.Sprintf("%s%d", priceCell, number)
}
func LengthCell(number int) string {
	return fmt.Sprintf("%s%d", lengthCell, number)
}
func WidthCell(number int) string {
	return fmt.Sprintf("%s%d", widthCell, number)
}
func HeightCell(number int) string {
	return fmt.Sprintf("%s%d", heightCell, number)
}
func WeightCell(number int) string {
	return fmt.Sprintf("%s%d", weightCell, number)
}
func PhotoCell(number int) string {
	return fmt.Sprintf("%s%d", photoCell, number)
}

func (e *Excel) Read() *[]ProductsPars {
	f, err := excelize.OpenFile(e.nameFile)

	if err != nil {
		log.Println(err)
		return nil
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Лист1")

	products := make([]ProductsPars, len(rows)-1)

	if err != nil {
		log.Println(err)
		return nil
	}

	for i := 1; i < len(rows); i++ {
		if len(rows[i]) > 10 || len(rows) < 9 {
			log.Println(err)
			break
		}

		products[i-1].Article, err = strconv.Atoi(rows[i][0])
		if err != nil {
			log.Println("strconv  Atoi stroka %v, err: %v", i, err)
			return nil
		}

		products[i-1].Catalog = rows[i][1]
		products[i-1].Name = rows[i][2]
		products[i-1].Description = rows[i][3]
		products[i-1].Price, err = strconv.ParseFloat(rows[i][4], 64)
		if err != nil {
			log.Println("ParseFloat stroka %v, err: %v", i, err)
			return nil
		}
		products[i-1].Length, err = strconv.Atoi(rows[i][5])
		if err != nil {
			log.Println("strconv  Atoi stroka %v, err: %v", i, err)
			return nil
		}
		products[i-1].Width, err = strconv.Atoi(rows[i][6])
		if err != nil {
			log.Println("strconv  Atoi stroka %v, err: %v", i, err)
			return nil
		}
		products[i-1].Height, err = strconv.Atoi(rows[i][7])
		if err != nil {
			log.Println("strconv  Atoi stroka %v, err: %v", i, err)
			return nil
		}
		products[i-1].Weight, err = strconv.Atoi(rows[i][8])
		if err != nil {
			log.Println("strconv  Atoi stroka %v, err: %v", i, err)
			return nil
		}
		if len(rows[i]) > 9 {
			a := rows[i][9]
			b := strings.Split(a, ",")
			products[i-1].PhotoUrl = b
		}
	}

	return &products
}

type ProductsPars struct {
	Article     int      `json:"article,omitempty"`
	Catalog     string   `json:"catalog,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	PhotoUrl    []string `json:"photo_url,omitempty"`
	Price       float64  `json:"price,omitempty"`
	Length      int      `json:"length"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Weight      int      `json:"weight"`
}
