package controllers

import (
	"encoding/json"
	"net/http"
	"user/src/domains/entities"
	"user/src/usecases/port"
)

type ContentController struct {
	// -> crypt.NewContentCrypt
	CryptFactory func(p *entities.Param) port.ContentCrypt
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
	outputPort := cc.OutputFactory(w)
	crypt := cc.CryptFactory(content.Param)
	inputPort := cc.InputFactory(outputPort, crypt)
	inputPort.Upload(content)
}
