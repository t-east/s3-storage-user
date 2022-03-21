package router

import (
	"fmt"
	"log"
	"net/http"

	"user/src/core"

	"user/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
)

func allowOptionsMiddleware(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return nil
	}
	return nil
}

func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
	param, _, _ := mocks.CreateParamMock()
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	allowOptionsMiddleware(w, r)
	var head string
	_, r.URL.Path = core.ShiftPath(r.URL.Path)
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
