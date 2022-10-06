package controllers

import (
	"net/http"
	"user/src/domains/entities"
	"user/src/interfaces/contracts"
	"user/src/interfaces/crypt"
	"user/src/usecases/interactor"
	"user/src/usecases/port"

	"github.com/labstack/echo/v4"
)

type ContentController struct {
	// -> crypt.NewContentCrypt
	CryptFactory func() port.CryptPort
	// -> interactor.NewContentCreateMetaDataputPort
	InputFactory func(
		cr port.CryptPort,
		cc port.ContractPort,
		// co port.ContentContracts,
	) port.ContentInputPort
	Param *entities.Param
}

func LoadContentController(param *entities.Param) *ContentController {
	return &ContentController{
		Param: param,
	}
}

func (cc *ContentController) MetaGen(c echo.Context) error {
	req := &entities.ContentCreateMetaData{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	content := &entities.ContentCreateMetaData{
		Content: req.Content,
		PrivKey: req.PrivKey,
		Address: req.Address,
	}
	crypt := crypt.NewContentCrypt(cc.Param)
	contract := contracts.NewContentContracts()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
	)
	receipt, err := inputPort.MetaGen(content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, receipt)
}

type SetKeyReq struct {
	PubKey     []byte `json:"pub_key"`
	EthPrivKey string `json:"eth_priv_key"`
}

func (cc *ContentController) SetKey(c echo.Context) error {
	req := &SetKeyReq{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	crypt := crypt.NewContentCrypt(cc.Param)
	contract := contracts.NewContentContracts()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
	)
	err := inputPort.SetKey(req.PubKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (cc *ContentController) GetLog(c echo.Context) error {
	crypt := crypt.NewContentCrypt(cc.Param)
	contract := contracts.NewContentContracts()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
	)
	content, err := inputPort.ListLog()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, content)
}
