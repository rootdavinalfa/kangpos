/*
 * Copyright (c) 2021-2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package lang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type LangParams struct {
	LangID string
	Key    string
}

type LangModel struct {
	LangID      string     `json:"lang_id"`
	Description string     `json:"description"`
	Data        []LangData `json:"data"`
}

type LangData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func readJson(langID string) *LangModel {
	path := "./resources/lang/" + langID + ".json"
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	lang := LangModel{}
	err = json.Unmarshal(jsonFile, &lang)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &lang
}

func GetLangValue(param LangParams) string {

	var lang = readJson(param.LangID)
	if lang != nil {
		log.Println("Failed to parse language!")
	}

	if lang.LangID != param.LangID {
		log.Println("Returned language not match")
	}
	for _, value := range lang.Data {
		if value.Key == param.Key {
			return value.Value
		}
	}
	return "n/a"
}
