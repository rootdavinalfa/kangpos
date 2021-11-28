/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package cron

import (
	cron2 "github.com/robfig/cron/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	config2 "kangpos/config"
	"log"
	"time"
)

var DB *gorm.DB

func Initialize() {
	db, err := gorm.Open(sqlite.Open(config2.GetConfig().DB.Dsn))
	if err != nil {
		log.Fatal("Could not connect database")
	}
	DB = db
	timeZone, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron2.New(cron2.WithLocation(timeZone), cron2.WithSeconds(), cron2.WithParser(cron2.NewParser(
		cron2.SecondOptional|cron2.Minute|cron2.Hour|cron2.Dom|cron2.Month|cron2.Dow|cron2.Descriptor,
	)))

	//_, err = scheduler.AddFunc("* * * * * *", mail.MailGun)
	if err != nil {
		log.Println(err.Error())
	}
	go scheduler.Run()

}
