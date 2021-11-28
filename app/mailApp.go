/*
 * Copyright (c) 2021-2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package app

import (
	"github.com/gorilla/mux"
	"kangpos/common"
	"kangpos/service/mail"
)

func SetRoutes(route *mux.Router) {
	common.Get(route, "/mail/list", mail.GetAllMail)
	common.Get(route, "/mail/get/{id}", mail.GetMail)
	common.Post(route, "/mail", mail.NewMail)
	common.Post(route, "/mail/send/{id}", mail.SendNow)
	common.Delete(route, "/mail/{id}", mail.Delete)
}
