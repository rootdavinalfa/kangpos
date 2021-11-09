/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package common

import (
	"encoding/json"
	"fmt"
	"kangpos/lang"
	response2 "kangpos/response"
	"net/http"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, r *http.Request, status int, payload interface{}, keyLang string) {

	langID := r.URL.Query().Get("lang")
	if len(langID) == 0 {
		langID = "en-US"
	}

	message := lang.GetLangValue(lang.LangParams{
		LangID: langID,
		Key:    keyLang,
	})

	var errs = true
	if FindfirstDigit(status) != 4 || FindfirstDigit(status) != 5 {
		errs = false
	}
	payload = response2.Response{
		Error:   errs,
		Message: message,
		Data:    payload,
	}
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write([]byte(response))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}
