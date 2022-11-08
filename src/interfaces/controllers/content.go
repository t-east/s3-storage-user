package controllers

import (
	"net/http"
	"user/src/domains/entities"
	"user/src/interfaces/contracts"
	"user/src/interfaces/crypt"
	"user/src/interfaces/random"
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
	req := &MetaDataReq{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	content := ContentAPISchemaToEntity(req)
	crypt := crypt.NewContentCrypt(cc.Param)
	contract := contracts.NewContentContracts()
	random := random.NewRandomID()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
		random,
	)
	metaData, err := inputPort.MetaGen(content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, metaData)
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
	random := random.NewRandomID()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
		random,
	)
	err := inputPort.SetKey(req.PubKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

type InitIndexLogRes struct {
	ID string `json:"id"`
}

func (cc *ContentController) InitIndexLog(c echo.Context) error {
	crypt := crypt.NewContentCrypt(cc.Param)
	contract := contracts.NewContentContracts()
	random := random.NewRandomID()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
		random,
	)
	id, err := inputPort.InitIndexLog()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, InitIndexLogRes{ID: id})
}

func (cc *ContentController) ListLog(c echo.Context) error {
	crypt := crypt.NewContentCrypt(cc.Param)
	contract := contracts.NewContentContracts()
	random := random.NewRandomID()
	inputPort := interactor.NewContentInputPort(
		crypt,
		contract,
		random,
	)
	logList, err := inputPort.ListLog()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, logList)
}
