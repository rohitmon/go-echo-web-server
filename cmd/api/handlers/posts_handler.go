package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rohitmon/go-echo-web-server/cmd/api/service"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		return c.String(http.StatusBadGateway, "Unable to process the request")
	}

	res := make(map[string]any)
	res["status"] = http.StatusOK
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadGateway, "Unable to process the request")
	}

	data, err := service.Get(idx)

	if err != nil {
		return c.String(http.StatusBadGateway, "Unable to process the request")
	}

	res := make(map[string]any)
	res["status"] = http.StatusOK
	res["data"] = data

	return c.JSON(http.StatusOK, res)

}
