package store

import (
	"github.com/overover1400/api-health-check/entity"
)

type ClientApi struct {
	ID              uint   `json:"id" gorm:"primary_key"`
	HealthCheckTime uint   `json:"health_check_time"`
	ToggleClientApi bool   `json:"toggle_client_api" gorm:"default:true"`
	ListenerUrl     string `json:"listener_url"`
	//TODO to avoid replicated URLs, URL column should be set to UNIQUE
	Url          string `json:"url"`
	HttpMethod   string `json:"http_method"`
	Header       string `json:"header"`
	Body         string `json:"body"`
	IntervalDate uint64 `json:"interval_date"`
}

func mapFromClientApi(clientApi entity.ClientApi) ClientApi {
	return ClientApi{
		ID:              clientApi.ID,
		ToggleClientApi: clientApi.ToggleClientApi,
		ListenerUrl:     clientApi.ListenerUrl,
		Url:             clientApi.Url,
		HttpMethod:      clientApi.HttpMethod,
		Header:          clientApi.Header,
		Body:            clientApi.Body,
		IntervalDate:    clientApi.IntervalDate,
	}
}

func mapToClientApi(clientApi ClientApi) entity.ClientApi {
	return entity.ClientApi{
		ID:              clientApi.ID,
		ToggleClientApi: clientApi.ToggleClientApi,
		ListenerUrl:     clientApi.ListenerUrl,
		Url:             clientApi.Url,
		HttpMethod:      clientApi.HttpMethod,
		Header:          clientApi.Header,
		Body:            clientApi.Body,
		IntervalDate:    clientApi.IntervalDate,
	}
}
