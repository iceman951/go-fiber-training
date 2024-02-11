package health

type HealthResponse struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	ENV       string `json:"env"`
	Timestamp string `json:"timestamp"`
}
