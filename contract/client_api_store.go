package contract

import (
	"context"
	"github.com/overover1400/api-health-check/entity"
)

type ClientApiStore interface {
	CreateClientApi(ctx context.Context, req entity.ClientApi) (entity.ClientApi, error)
	ToggleClientApi(ctx context.Context, ID uint) error
	FindClientApis(ctx context.Context) ([]entity.ClientApi, error)
	DeleteClientApi(ctx context.Context, ID uint) error
}

type HealthCheckApiStore interface {
	FindClientApiToHealthCheck() error
}
