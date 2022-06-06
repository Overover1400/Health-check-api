package clientApi

// Application business logic

import (
	"context"
	"github.com/overover1400/api-health-check/contract"
	"github.com/overover1400/api-health-check/dto"
	"github.com/overover1400/api-health-check/entity"
	"time"
)

// by store abstraction we can make our interactor or business logic layer independent of database layer
type ClientApiInteractor struct {
	store contract.ClientApiStore
}

func NewClientApi(store contract.ClientApiStore) ClientApiInteractor {
	return ClientApiInteractor{store: store}
}

func (i ClientApiInteractor) CreateClientApi(ctx context.Context, req dto.CreateClientApiRequest) (dto.CreateClientApiResponse, error) {
	clientApi, err := i.store.CreateClientApi(ctx, entity.ClientApi{
		ToggleClientApi: req.ToggleClientApi,
		ListenerUrl:     req.ListenerUrl,
		Url:             req.Url,
		HttpMethod:      req.HttpMethod,
		Header:          req.Header,
		Body:            req.Body,
		IntervalDate:    uint64(time.Now().Unix()),
	})
	if err != nil {
		return dto.CreateClientApiResponse{}, err
	}

	return dto.CreateClientApiResponse{ClientApi: clientApi}, nil
}

func (i ClientApiInteractor) ToggleClientApi(ctx context.Context, req dto.ToggleClientApiRequest) error {

	if err := i.store.ToggleClientApi(ctx, req.ID); err != nil {
		return err
	}

	return nil
}

func (i ClientApiInteractor) FindClientApis(ctx context.Context) (dto.FindClientApiResponse, error) {

	findClientApi, err := i.store.FindClientApis(ctx)

	if err != nil {
		return dto.FindClientApiResponse{}, err
	}

	return dto.FindClientApiResponse{ClientApi: findClientApi}, nil
}

func (i ClientApiInteractor) DeleteClientApi(ctx context.Context, req dto.DeleteClientApiRequest) error {

	if err := i.store.DeleteClientApi(ctx, req.ID); err != nil {
		return err
	}

	return nil
}
