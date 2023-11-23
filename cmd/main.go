package main

import (
	"Skyline/internal/database"
	"Skyline/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("Hello World ")
	appSetting, err := utils.LoadAppConfig("./configs")
	if err != nil {
		panic("AppSetting cannot be loaded")
	}

	err = database.InitDB(appSetting.DbConnection)
	if err != nil {
		panic("Cannot connected to the database!")
	}
}
