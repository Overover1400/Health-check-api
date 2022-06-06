package store

import (
	"context"
	"github.com/overover1400/api-health-check/entity"
)

func (m MySqlStore) CreateClientApi(ctx context.Context, clientApi entity.ClientApi) (entity.ClientApi, error) {
	client := mapFromClientApi(clientApi)
	if err := m.db.WithContext(ctx).Create(&client).Error; err != nil {
		return entity.ClientApi{}, err
	}
	return mapToClientApi(client), nil
}

func (m MySqlStore) ToggleClientApi(ctx context.Context, ID uint) error {
	client := ClientApi{ID: ID}

	if err := m.db.WithContext(ctx).Save(&client).Error; err != nil {
		return err
	}
	return nil
}

func (m MySqlStore) FindClientApis(ctx context.Context) ([]entity.ClientApi, error) {

	var clientApi []ClientApi
	if err := m.db.WithContext(ctx).Find(&clientApi).Error; err != nil {
		return nil, err
	}

	clientApiEntity := make([]entity.ClientApi, len(clientApi))

	for i := range clientApi {
		clientApiEntity[i] = mapToClientApi(clientApi[i])
	}

	return clientApiEntity, nil
}

func (m MySqlStore) DeleteClientApi(ctx context.Context, ID uint) error {
	var client ClientApi

	if err := m.db.WithContext(ctx).Where("id = ?", ID).First(&client).Error; err != nil {
		return err
	}
	return nil
}
