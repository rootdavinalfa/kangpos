/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dsn string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dsn: "kangpos.sqlite",
		},
	}
}
