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
	"log"
	"net/http"
)

func GetAllMail(w http.ResponseWriter, r *http.Request) {
	var mails []Mail
	DB.Preload("Attachments").Preload("MailTarget").Find(&mails)
	common.RespondJSON(w, r, http.StatusOK, mails, "SUCCESS_FETCHED")
}

func GetMail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mailID := params["id"]
	var mail Mail
	DB.Preload("Attachments").Preload("MailTarget").Find(&mail, mailID)
	common.RespondJSON(w, r, http.StatusOK, mail, "SUCCESS_FETCHED")
}

func NewMail(w http.ResponseWriter, r *http.Request) {
	var mail Mail
	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		log.Println("Error occurred! " + err.Error())
	}

	DB.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&mail)
	common.RespondJSON(w, r, http.StatusCreated, mail, "SUCCESS_FETCHED")
}

func SendNow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mailID := params["id"]
	var mail Mail
	DB.Preload("Attachments").Preload("MailTarget").Find(&mail, mailID)
	mailer := common.Mailer{}
	mailer.InitializeMail(&common.ClientConfig{
		Host:     "smtp.gmail.com",
		Username: "",
		Password: "",
		Port:     "587",
	})

	body := mailer.BuildEmailData(&common.MailTemplate{
		From:    mailer.Configuration.Username,
		To:      mail.To,
		Subject: mail.Subject,
	}, mail.Body)

	err := mailer.SendMail(mail.To, body)
	if err != nil {
		common.RespondJSON(w, r, http.StatusInternalServerError, mail, "ERROR_SEND_MAIL")
		return
	}

	common.RespondJSON(w, r, http.StatusOK, mail, "SUCCESS_SEND_MAIL")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mailID := params["id"]
	DB.Delete(&Mail{}, mailID)
	common.RespondJSON(w, r, http.StatusOK, mailID, "SUCCESS_DELETED")
}

func MailGun() {

}
