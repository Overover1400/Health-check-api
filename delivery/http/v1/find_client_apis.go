package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/overover1400/api-health-check/adaptor/store"
	clientApi "github.com/overover1400/api-health-check/interactor/client_api"
	"net/http"
)

func FindClientApis(store store.MySqlStore) echo.HandlerFunc {
	return func(c echo.Context) error {

		resp, err := clientApi.NewClientApi(store).FindClientApis(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, resp)
	}
}
