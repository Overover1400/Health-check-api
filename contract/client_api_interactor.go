package contract

import (
	"context"
	"github.com/overover1400/api-health-check/dto"
)

type ClientApiInteractor interface {
	CreateClientApi(ctx context.Context, req dto.CreateClientApiRequest) (dto.CreateClientApiResponse, error)
	ToggleClientApi(ctx context.Context, req dto.ToggleClientApiRequest) error
	FindClientApis(ctx context.Context) (dto.FindClientApiResponse, error)
	DeleteClientApi(ctx context.Context, req dto.DeleteClientApiRequest) error
}
