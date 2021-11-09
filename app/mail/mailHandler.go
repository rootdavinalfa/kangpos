/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package mail

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
	"kangpos/common"
	"net/http"
)

func getAllMail(w http.ResponseWriter, r *http.Request) {
	var mails []Mail
	DB.Preload("Attachments").Preload("MailTarget").Find(&mails)
	common.RespondJSON(w, r, http.StatusOK, mails, "SUCCESS_FETCHED")
}

func getMail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mailID := params["id"]
	var mail Mail
	DB.Preload("Attachments").Preload("MailTarget").Find(&mail, mailID)
	common.RespondJSON(w, r, http.StatusOK, mail, "SUCCESS_FETCHED")
}

func newMail(w http.ResponseWriter, r *http.Request) {
	var mail Mail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {

	}

	DB.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&mail)
	common.RespondJSON(w, r, http.StatusCreated, mail, "SUCCESS_FETCHED")
}
