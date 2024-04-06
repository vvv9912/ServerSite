package api

import "ServerSite/internal/service"

type Api struct {
	*service.Service
}

func NewApi(service *service.Service) *Api {
	return &Api{Service: service}
}

type ProductsGet struct {
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
type Data struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
