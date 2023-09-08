package main

import (
	"fmt"
	"net/http"

	"github.com/ell1jah/db_cp/internal/user/delivery"
	"github.com/ell1jah/db_cp/internal/user/repo"
	"github.com/ell1jah/db_cp/internal/user/service"
	"github.com/ell1jah/db_cp/pkg/middleware"
	"github.com/ell1jah/db_cp/pkg/session"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const port = ":8080"

func main() {
	zapLogger := zap.Must(zap.NewDevelopment())
	logger := zapLogger.Sugar()

	params := "user=postgres dbname=clothshop password=postgres host=localhost port=5432 sslmode=disable"
	db, err := sqlx.Connect("postgres", params)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	userHandler := delivery.UserHandler{
		Logger:   logger,
		Sessions: session.JWTSessionsManager{},
		UserService: service.UserService{
			UserRepo: &repo.PgUserRepo{
				Logger: logger,
				DB:     db,
			},
			Logger: logger,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	//TODO authmidddleware для других путей
	//TODO добавление, удаление, изменение новых пользователей админом

	mux := middleware.AccessLog(logger, r)
	mux = middleware.Panic(logger, mux)

	logger.Infow("starting server",
		"type", "START",
		"port", port,
	)

	logger.Errorln(http.ListenAndServe(port, mux))

	err = zapLogger.Sync()
	if err != nil {
		fmt.Println(err)
	}
}
