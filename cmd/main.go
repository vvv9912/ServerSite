package main

import (
	"ServerSite/internal/api"
	"ServerSite/internal/authorization"
	"ServerSite/internal/config"
	"ServerSite/internal/server"
	"ServerSite/internal/service"
	"ServerSite/internal/storage"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

// todo config
const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "supersecretkey"

func main() {

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDSN)
	defer db.Close()

	if err != nil {
		fmt.Println(db)
		return
	}

	sProducts := storage.NewProductsPostgresStorage(db)

	auth := authorization.NewAutorization(TOKEN_EXP, SECRET_KEY)

	DbUser := storage.NewUsersPostgresStorage()

	Service := service.NewService(auth, sProducts, DbUser)
	Api := api.NewApi(Service)
	s := server.NewServer(Api)

	ctx := context.Background()

	serverAddr := config.Get().ServerAddress
	err = s.ServerStart(ctx, serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
}
