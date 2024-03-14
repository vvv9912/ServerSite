package service

import (
	"ServerSite/internal/authorization"
	"ServerSite/internal/model"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

type PostStorager interface {
	AddProduct(ctx context.Context, product model.Products) error
	ChangeProductByArticle(ctx context.Context, product model.Products) error
	SelectAllProducts(ctx context.Context) ([]model.Products, error)
}

type DbUser interface {
	CheckUserCredentials(username, password string) (bool, int, string, error)
	CheckId(id int) (bool, string, error)
}

type Service struct {
	Authorization *authorization.Autorization
	Storage       PostStorager
	DbUser
}

func NewService(authorization *authorization.Autorization, storage PostStorager, DbUser DbUser) *Service {
	return &Service{Authorization: authorization, Storage: storage, DbUser: DbUser}
}

func (s *Service) GetTokenAuthorization(login, password string) (string, error) {

	passwordHash := Sha256Hash(password)

	ok, id, role, err := s.DbUser.CheckUserCredentials(login, passwordHash)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("bad Autorization")
	}

	jwt, err := s.Authorization.BuildJWTString(id, role)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return jwt, nil
}
func Sha256Hash(input string) string {
	// Создаем новый хеш SHA-256
	hasher := sha256.New()

	// Преобразуем строку в байты и передаем хеш-функции
	hasher.Write([]byte(input))

	// Получаем хеш в виде среза байтов
	hashBytes := hasher.Sum(nil)

	// Преобразуем срез байтов в строку в шестнадцатеричном формате
	hashedString := hex.EncodeToString(hashBytes)

	return hashedString
}
