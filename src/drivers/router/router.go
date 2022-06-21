package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"user/src/core"
	"user/src/domains/entities"

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

func GetParam() (*entities.Param, error) {
	url := "http://localhost:4001/api/param"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	param := &entities.Param{}
	err = json.NewDecoder(resp.Body).Decode(&param)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return param, err
}

func ServerHandlerPublic(w http.ResponseWriter, r *http.Request) {
	// param, err := GetParam()
	param, _, err := core.CreateParamMock()
	if err != nil {
		http.Error(w, fmt.Sprintf("connect: connection refused: %v", err.Error()), http.StatusNotFound)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	allowOptionsMiddleware(w, r)
	var head string
	_, r.URL.Path = core.ShiftPath(r.URL.Path)
	head, r.URL.Path = core.ShiftPath(r.URL.Path)
	switch head {
	case "content":
		uc := controllers.LoadContentController(param)
		uc.Dispatch(w, r)
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
