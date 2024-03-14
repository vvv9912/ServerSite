package authorization

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Autorization struct {
	TokenExp  time.Duration
	SecretKey string
}

func NewAutorization(tokenExp time.Duration, secretKey string) *Autorization {
	return &Autorization{TokenExp: tokenExp, SecretKey: secretKey}
}

// Claims — структура утверждений, которая включает стандартные утверждения
// и одно пользовательское — UserID
type Claims struct {
	jwt.RegisteredClaims
	UserID int
	Role   string
}

// BuildJWTString создаёт токен и возвращает его в виде строки.
func (a *Autorization) BuildJWTString(userId int, role string) (string, error) {
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.TokenExp)),
		},
		// собственное утверждение
		UserID: userId,
		Role:   role,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}
func (a *Autorization) GetUserId(tokenString string) (int, string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			//доп проверка заголовка
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(a.SecretKey), nil
		})
	if err != nil {
		return -1, "", nil
	}

	if !token.Valid {
		//		fmt.Println("Token is not valid")
		return -1, "", nil
	}

	//fmt.Println("Token os valid")
	return claims.UserID, claims.Role, nil
}
