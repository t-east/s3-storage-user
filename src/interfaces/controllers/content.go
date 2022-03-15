package controllers

import (
	"encoding/json"
	"net/http"
	"user/src/domains/entities"
	"user/src/interfaces/crypt"
	"user/src/interfaces/presenters"
	"user/src/usecases/interactor"
	"user/src/usecases/port"
)

type ContentController struct {
	// -> crypt.NewContentCrypt
	CryptFactory func() port.ContentCrypt
	// -> presenter.NewContentOutputPort
	OutputFactory func(w http.ResponseWriter) port.ContentOutputPort
	// -> interactor.NewContentInputPort
	InputFactory func(
		o port.ContentOutputPort,
		cr port.ContentCrypt,
		// co port.ContentContracts,
	) port.ContentInputPort
}

func LoadContentController() *ContentController {
	return &ContentController{}
}

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	content := &entities.ContentInput{}
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	crypt := crypt.NewUserCrypt()
	outputPort := presenters.NewContentOutputPort(w)
	inputPort := interactor.NewContentInputPort(outputPort, crypt)
	inputPort.Upload(content)
}
