/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package main

import (
	app2 "kangpos/app"
	config2 "kangpos/config"
)

func main() {
	config := config2.GetConfig()
	app := app2.App{}
	app.Initialize(config)
	app.Run(":3000")
}
