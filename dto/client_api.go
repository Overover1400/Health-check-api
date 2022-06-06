package dto

import "github.com/overover1400/api-health-check/entity"

type CreateClientApiRequest struct {
	ID              uint   `json:"id"`
	ToggleClientApi bool   `json:"toggle_client_Api"`
	ListenerUrl     string `json:"monitor_url"`
	Url             string `json:"req_url"`
	HttpMethod      string `json:"req_http_method"`
	Header          string `json:"req_header"`
	Body            string `json:"req_body"`
	IntervalDate    uint64 `json:"interval_date"`
}

type CreateClientApiResponse struct {
	ClientApi entity.ClientApi
}

type ToggleClientApiRequest struct {
	ID uint `json:"id"`
}

//type ToggleClientApiResponse struct{}

//type FindClientApiRequest struct{}

type FindClientApiResponse struct {
	ClientApi []entity.ClientApi `json:"client_api"`
}

type DeleteClientApiRequest struct {
	ID uint `json:"id"`
}

//type DeleteClientApiResponse struct{}
