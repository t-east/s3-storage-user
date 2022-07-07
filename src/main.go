package main

import (
	"log"
	"user/src/drivers/ethereum"
	router "user/src/drivers/router"
	"user/src/interfaces/controllers"
)

func realMain() {
	param, err := ethereum.GetParam()
	if err != nil {
		log.Print(err)
	}
	cc := *controllers.LoadContentController(param)

	e := router.NewServer(cc)
	e.Logger.Fatal(e.Start(":4000"))
}

func main() {
	realMain()
}
