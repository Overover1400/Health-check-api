package v1

import (
	"github.com/overover1400/api-health-check/adaptor/store"
	"github.com/overover1400/api-health-check/interactor/health_check_api"
)

func HealthCheckApi(store store.MySqlStore) error {

	if err := healthCheckApi.NewCronHealthCheck(store).CronHealthCheck(); err != nil {
		return err
	}
	return nil
}
