package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/overover1400/api-health-check/adaptor/store"
	"github.com/overover1400/api-health-check/dto"
	"github.com/overover1400/api-health-check/interactor/client_api"
	"net/http"
	"strconv"
)

func DeleteClientApi(store store.MySqlStore) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientApiID := c.Param("id")
		clientApiIDToInt, err := strconv.Atoi(clientApiID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		req := dto.DeleteClientApiRequest{ID: uint(clientApiIDToInt)}
		if err = clientApi.NewClientApi(store).DeleteClientApi(c.Request().Context(), req); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
