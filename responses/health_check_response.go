package responses

type HealthCheckResponse struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}
