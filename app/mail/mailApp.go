/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package mail

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"kangpos/common"
)

var DB *gorm.DB

func SetRoutes(route *mux.Router, db *gorm.DB) {
	DB = db
	common.Get(route, "/mail/list", getAllMail)
	common.Get(route, "/mail/by-id/{id}", getMail)
	common.Post(route, "/mail", newMail)
}
