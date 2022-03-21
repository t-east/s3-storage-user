package controllers

import (
	"bytes"
	"io"
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
	Param *entities.Param
}

func LoadContentController(param *entities.Param) *ContentController {
	return &ContentController{
		Param: param,
	}
}

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("content")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	content := &entities.ContentIn{
		Content:     buf.Bytes(),
		ContentName: r.FormValue("name"),
		PrivKey:     []byte(r.FormValue("name")),
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	crypt := crypt.NewUserCrypt(cc.Param)
	outputPort := presenters.NewContentOutputPort(w)
	inputPort := interactor.NewContentInputPort(outputPort, crypt)
	inputPort.Upload(content)
}

func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {
	crypt := crypt.NewUserCrypt(cc.Param)
	outputPort := presenters.NewContentOutputPort(w)
	inputPort := interactor.NewContentInputPort(outputPort, crypt)
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
