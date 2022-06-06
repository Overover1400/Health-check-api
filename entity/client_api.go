package entity

type ClientApi struct {
	ID              uint   `json:"id"`
	HealthCheckTime uint   `json:"health_check_time"`
	ToggleClientApi bool   `json:"toggle_client_api"`
	ListenerUrl     string `json:"listener_url"`
	Url             string `json:"url"`
	HttpMethod      string `json:"http_method"`
	Header          string `json:"header"`
	Body            string `json:"body"`
	IntervalDate    uint64 `json:"interval_date"`
}
