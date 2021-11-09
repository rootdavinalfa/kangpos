/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package mail

import (
	"fmt"
	"gorm.io/gorm"
)

type Mail struct {
	gorm.Model
	Subject     string
	Body        string
	MailTarget  []Target     `gorm:"foreignKey:MailID;references:ID"`
	Attachments []Attachment `gorm:"foreignKey:MailID;references:ID"`
}

type Attachment struct {
	MailID   uint
	FileName string
	Data     string
}

type Target struct {
	MailID uint
	Target string
	Kind   string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&Mail{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = db.AutoMigrate(&Attachment{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = db.AutoMigrate(&Target{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return db
}
