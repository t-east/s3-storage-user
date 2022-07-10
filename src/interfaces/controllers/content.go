package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"user/src/domains/entities"
	"user/src/usecases/port"

	"github.com/labstack/echo/v4"
)

type ContentController struct {
	// -> crypt.NewContentCrypt
	CryptFactory func() port.ContentCrypt
	// -> interactor.NewContentInputPort
	InputFactory func(
		cr port.ContentCrypt,
		cc port.ContentContract,
		// co port.ContentContracts,
	) port.ContentInputPort
	Param *entities.Param
}

func LoadContentController(param *entities.Param) *ContentController {
	return &ContentController{
		Param: param,
	}
}

type MetaRes struct {
	Receipt entities.Content `json:"receipt"`
	PrivKey []byte           `json:"priv_key"`
	PubKey  []byte           `json:"pub_key"`
}

func (cc *ContentController) MetaGen(c echo.Context) error {
	req := &entities.ContentIn{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// crypt := crypt.NewContentCrypt(cc.Param)
	// contract := contracts.NewContentContracts()
	// inputPort := interactor.NewContentInputPort(
	// 	crypt,
	// 	contract,
	// )
	// receipt, privKey, pubKey, err := inputPort.MetaGen(content)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	raw, err := ioutil.ReadFile("./content.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc entities.Content

	json.Unmarshal(raw, &fc)
	return c.JSON(http.StatusCreated, MetaRes{
		Receipt: entities.Content{
			Content:  req.Content,
			MetaData: fc.MetaData,
		},
		PrivKey: req.PrivKey,
	})
}

type SetKeyReq struct {
	PubKey []byte `json:"pub_key"`
}

func (cc *ContentController) SetKey(c echo.Context) error {
	// req := &SetKeyReq{}
	// if err := c.Bind(req); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	// crypt := crypt.NewContentCrypt(cc.Param)
	// contract := contracts.NewContentContracts()
	// inputPort := interactor.NewContentInputPort(
	// 	crypt,
	// 	contract,
	// )
	// err := inputPort.SetKey(req.PubKey)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	return c.NoContent(http.StatusCreated)
}

func (cc *ContentController) GetLog(c echo.Context) error {
	// crypt := crypt.NewContentCrypt(cc.Param)
	// contract := contracts.NewContentContracts()
	// inputPort := interactor.NewContentInputPort(
	// 	crypt,
	// 	contract,
	// )
	// content, err := inputPort.GetLog()
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	// }
	raw, err := ioutil.ReadFile("./log.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fc entities.Log

	json.Unmarshal(raw, &fc)
	logs := []entities.Log{
		fc,
		fc,
	}
	return c.JSON(http.StatusOK, logs)
}
