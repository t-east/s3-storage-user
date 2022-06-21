package controllers

import (
	"encoding/json"
	"net/http"
	"user/src/domains/entities"
	"user/src/interfaces/crypt"
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
	Param *entities.Param
}

func LoadContentController(param *entities.Param) *ContentController {
	return &ContentController{
		Param: param,
	}
}

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	contentIn := &entities.ContentIn{}
	err := json.NewDecoder(r.Body).Decode(&contentIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	crypt := crypt.NewContentCrypt(cc.Param)
	inputPort := interactor.NewContentInputPort(crypt)
	newContent, err := inputPort.Upload(contentIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(newContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(201)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {
	crypt := crypt.NewContentCrypt(cc.Param)
	inputPort := interactor.NewContentInputPort(crypt)
	inputPort.GetKey()
}

func (cc *ContentController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cc.Post(w, r)
	case "GET":
		cc.Get(w, r)
	default:
		http.NotFound(w, r)
	}
}
