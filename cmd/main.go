package main

import (
	"Skyline/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("Hello World ")
	appSetting, err := utils.LoadAppConfig("./internal/configs")
	if err != nil {
		panic("AppSetting cannot be loaded")
	}

	err = utils.InitDB(appSetting.DbConnection)
	if err != nil {
		panic("Cannot connected to the database!")
	}
}
