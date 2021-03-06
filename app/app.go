/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"kangpos/config"
	"kangpos/service/mail"
	"log"
	"net/http"
)

type App struct {
}

var Router *mux.Router

func (a *App) Initialize(config *config.Config) {

	db, err := gorm.Open(sqlite.Open(config.DB.Dsn))
	if err != nil {
		log.Fatal("Could not connect database")
	}
	Router = mux.NewRouter()
	SetRoutes(Router)
	mail.DBMigrate(db)
	//cron.Initialize()
}

func (a *App) Run(host string) {
	fmt.Println("Listened on port " + host)
	log.Fatal(http.ListenAndServe(host, Router))
}
