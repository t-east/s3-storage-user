package main

import (
	"log"
	"user/src/core"
	router "user/src/drivers/router"
	"user/src/interfaces/controllers"

	"github.com/joho/godotenv"
)

func realMain() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}
	// param, err := ethereum.GetParam()
	param, _, err := core.CreateParamMock()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(param)
	cc := *controllers.LoadContentController(param)

	e := router.NewServer(cc)
	e.Logger.Fatal(e.Start(":4000"))
}

func main() {
	realMain()
}
