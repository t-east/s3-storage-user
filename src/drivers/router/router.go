package router

import (
	"net/http"

	"user/src/interfaces/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer(cc controllers.ContentController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	api := e.Group("/api")
	api.POST("/content/meta", cc.MetaGen)
	api.GET("/content", cc.GetLog)
	api.POST("/key", cc.SetKey)
	api.POST("/log", cc.GetLog)

	return e
}

// func GetParam() (*entities.Param, error) {
// 	url := "http://localhost:4001/api/param"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	param := &entities.Param{}
// 	err = json.NewDecoder(resp.Body).Decode(&param)
// 	if err != nil {
// 		log.Print(err)
// 		return nil, err
// 	}
// 	return param, err
// }
