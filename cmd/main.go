package main

import (
	"fmt"
	"net/http"

	brandDel "github.com/ell1jah/db_cp/internal/brand/delivery"
	brandRepo "github.com/ell1jah/db_cp/internal/brand/repo"
	brandServ "github.com/ell1jah/db_cp/internal/brand/service"
	itemDel "github.com/ell1jah/db_cp/internal/item/delivery"
	itemRepo "github.com/ell1jah/db_cp/internal/item/repo"
	itemServ "github.com/ell1jah/db_cp/internal/item/service"
	userDel "github.com/ell1jah/db_cp/internal/user/delivery"
	userRepo "github.com/ell1jah/db_cp/internal/user/repo"
	userServ "github.com/ell1jah/db_cp/internal/user/service"
	"github.com/ell1jah/db_cp/pkg/context"
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

	sessionManager := session.JWTSessionsManager{}
	contextManager := context.ContextManager{}

	authManager := middleware.AuthManager{
		SessionManager: sessionManager,
		Logger:         logger,
		ContextManager: contextManager,
	}

	userHandler := userDel.UserHandler{
		Logger:   logger,
		Sessions: sessionManager,
		UserService: userServ.UserService{
			UserRepo: &userRepo.PgUserRepo{
				Logger: logger,
				DB:     db,
			},
			Logger: logger,
		},
	}

	brandHandler := brandDel.BrandHandler{
		Logger: logger,
		BrandService: brandServ.BrandService{
			BrandRepo: &brandRepo.PgBrandRepo{
				Logger: logger,
				DB:     db,
			},
			Logger: logger,
		},
	}

	itemHandler := itemDel.ItemHandler{
		Logger: logger,
		ItemService: itemServ.ItemService{
			ItemRepo: &itemRepo.PgItemRepo{
				Logger: logger,
				DB:     db,
			},
			Logger: logger,
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	r.HandleFunc("/brand/{BRAND_ID:[0-9]+}", http.HandlerFunc(brandHandler.Get)).Methods("GET")
	r.Handle("/brand", authManager.Auth(http.HandlerFunc(brandHandler.Create), "admin")).Methods("PUT")
	r.Handle("/brand/{BRAND_ID:[0-9]+}", authManager.Auth(http.HandlerFunc(brandHandler.Update), "admin")).Methods("POST")
	r.Handle("/brand/{BRAND_ID:[0-9]+}", authManager.Auth(http.HandlerFunc(brandHandler.Delete), "admin")).Methods("DELETE")

	r.HandleFunc("/item/{ITEM_ID:[0-9]+}", http.HandlerFunc(itemHandler.Get)).Methods("GET")
	r.Handle("/item", authManager.Auth(http.HandlerFunc(itemHandler.Create), "admin")).Methods("PUT")
	r.Handle("/item/{ITEM_ID:[0-9]+}", authManager.Auth(http.HandlerFunc(itemHandler.Update), "admin")).Methods("POST")
	r.Handle("/item/{ITEM_ID:[0-9]+}", authManager.Auth(http.HandlerFunc(itemHandler.Delete), "admin")).Methods("DELETE")
	r.HandleFunc("/items}", http.HandlerFunc(itemHandler.Get)).Methods("GET")

	//TODO добавление, удаление, изменение новых пользователей
	//TODO GET /items

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
