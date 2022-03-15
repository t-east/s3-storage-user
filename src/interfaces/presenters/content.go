package presenters

import (
	"encoding/json"
	"net/http"

	"user/src/domains/entities"
	"user/src/usecases/port"
)

type Content struct {
	w http.ResponseWriter
}

func NewContentOutputPort(w http.ResponseWriter) port.ContentOutputPort {
	return &Content{
		w: w,
	}
}

func (c *Content) Render(content *entities.Content, statusCode int) {
	res, err := json.Marshal(content)
	if err != nil {
		http.Error(c.w, err.Error(), http.StatusInternalServerError)
		return
	}
	c.w.WriteHeader(statusCode)
	c.w.Header().Set("Content-Type", "application/json")
	c.w.Write(res)
}

func (c *Content) RenderError(err error) {
	c.w.WriteHeader(http.StatusInternalServerError)
	http.Error(c.w, err.Error(), http.StatusInternalServerError)
}
