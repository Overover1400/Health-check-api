package healthCheckApi

import (
	"github.com/overover1400/api-health-check/contract"
)

type HealthCheckApiInteractor struct {
	store contract.HealthCheckApiStore
}

func NewCronHealthCheck(store contract.HealthCheckApiStore) HealthCheckApiInteractor {
	return HealthCheckApiInteractor{store: store}
}

func (c HealthCheckApiInteractor) CronHealthCheck() error {
	if err := c.store.FindClientApiToHealthCheck(); err != nil {
		return err
	}
	return nil
}
