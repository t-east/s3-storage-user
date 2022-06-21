package controllers

import (
	"net/http"
	"user/src/domains/entities"
	"user/src/interfaces/crypt"
	"user/src/usecases/interactor"
	"user/src/usecases/port"

	"github.com/labstack/echo/v4"
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

func (cc *ContentController) Post(c echo.Context) error {
	req := &entities.ContentIn{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	content := &entities.ContentIn{
		Content: req.Content,
		PrivKey: req.PrivKey,
		Address: req.Address,
	}
	crypt := crypt.NewContentCrypt(cc.Param)
	inputPort := interactor.NewContentInputPort(
		crypt,
	)
	receipt, err := inputPort.MetaGen(content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, receipt)
}
