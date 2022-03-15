package router

import (
	// "database/sql"
	// "fmt"
	"log"
	"net/http"

	// "os"

	// blank import for MySQL driver

	"user/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
)

// Serve はserverを起動させます．
func Serve(addr string) {
	// コントローラの準備
	_ = controllers.LoadContentController()

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
