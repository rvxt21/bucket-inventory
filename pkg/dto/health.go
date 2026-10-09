package dto

type ReadyResponse struct {
	Status  string `json:"status"`
	Service string `json:"service,omitempty"`
}
