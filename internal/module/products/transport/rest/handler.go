package rest

import (
	"github.com/labstack/echo"
	"service/internal/module/products/usecase"
	"service/pkg/echoserver/response"
)

type rest struct {
	usecase usecase.Usecase
}

func NewHandler(echo *echo.Group, usecase usecase.Usecase) {
	transport := rest{usecase}

	routes := echo.Group("/products")
	routes.GET("", transport.list)
}

func (r *rest) list(c echo.Context) error {
	val, err := r.usecase.Check(c.Request().Context(), c.QueryParam("prefix"))
	if err != nil {
		return response.Error(c, err)
	}
	return response.Render(c, val)
}
