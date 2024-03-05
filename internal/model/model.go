package model

type Auth struct {
	Login    string `json:"login_login" form:"login_login" query:"login_login"`
	Password string `json:"password_password" form:"password_password" query:"password_password"`
}
type Products struct {
	Article     int      `json:"article,omitempty"`
	Catalog     string   `json:"catalog,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	PhotoUrl    [][]byte `json:"photo_url"`
	Price       float64  `json:"price,omitempty"`
	Length      int      `json:"length"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Weight      int      `json:"weight"`
}
