package router

import (
	"fmt"
	"log"
	"net/http"

	"user/src/core"

	"user/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = core.ShiftPath(r.URL.Path)
	switch head {
	case "content":
		uc := controllers.LoadContentController()
		uc.Post(w, r)
	default:
		http.Error(w, fmt.Sprintf("method not allowed request. req: %v", r.URL), http.StatusNotFound)
	}
}

// Serve はserverを起動させます．
func Serve() {
	sm := http.NewServeMux()
	sm.Handle("/api/", http.HandlerFunc(ServerHandlerPublic))
	err := http.ListenAndServe(":4000", sm)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
