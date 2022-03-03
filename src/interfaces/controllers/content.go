package controllers

import (
	"encoding/json"
	"net/http"
	"pairing_test/src/user/domains/entities"
	"pairing_test/src/user/interfaces/contracts"
	"pairing_test/src/user/usecases/port"

	"gorm.io/gorm"
)

type ContentController struct {
	// -> gateway.NewContentRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	// -> contracts.NewContentContracts
	ContractFactory func() port.ContentContract
	// -> crypt.NewContentCrypt
	CryptFactory func(p contracts.Param) port.ContentCrypt
	// -> presenter.NewContentOutputPort
	OutputFactory func(w http.ResponseWriter) port.ContentOutputPort
	// -> interactor.NewContentInputPort
	InputFactory func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		cr port.ContentCrypt,
		// co port.ContentContracts,
	) port.ContentInputPort
	Param contracts.Param
	Conn  *gorm.DB
}

func LoadContentController(db *gorm.DB, param contracts.Param) *ContentController {
	return &ContentController{Conn: db, Param: param}
}

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	content := &entities.ContentInput{}
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	outputPort := cc.OutputFactory(w)
	repository := cc.RepoFactory(cc.Conn)
	crypt := cc.CryptFactory(cc.Param)
	inputPort := cc.InputFactory(outputPort, repository, crypt)
	inputPort.Upload(content)
}

func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {
	//  TODO: idの取得をちゃんとしたやつにする
	id := "1"

	outputPort := cc.OutputFactory(w)
	repository := cc.RepoFactory(cc.Conn)
	crypt := cc.CryptFactory(cc.Param)
	inputPort := cc.InputFactory(outputPort, repository, crypt)
	inputPort.FindByID(id)
}
