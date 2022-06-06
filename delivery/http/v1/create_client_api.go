package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/overover1400/api-health-check/adaptor/store"
	"github.com/overover1400/api-health-check/dto"
	"github.com/overover1400/api-health-check/interactor/client_api"
	"net/http"
)

func CreateClientApi(store store.MySqlStore) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := dto.CreateClientApiRequest{}

		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		//TODO write validator

		resp, err := clientApi.NewClientApi(store).CreateClientApi(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}
