package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"user/src/domains/entities"
	"user/src/usecases/port"
)

type User struct {
	w http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) port.UserOutputPort {
	return &User{
		w: w,
	}
}

func (u *User) Render(user *entities.User, statusCode int) {
	fmt.Println(user)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(u.w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.w.WriteHeader(statusCode)
	u.w.Header().Set("Content-Type", "application/json")
	u.w.Write(res)
}

func (u *User) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	http.Error(u.w, err.Error(), http.StatusInternalServerError)
}
