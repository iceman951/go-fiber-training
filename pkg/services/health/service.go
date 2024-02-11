package health

import "github.com/spf13/viper"

type HealthInterface interface {
	Get() HealthInfo
}

type HealthInfo struct {
	Name    string
	Version string
	ENV     string
	Status  string
}

func New() HealthInterface {
	return &HealthInfo{
		Name:    viper.GetString("ProjectName"),
		Version: viper.GetString("Version"),
		ENV:     viper.GetString("ENV"),
		Status:  "I'm OK.",
	}
}

func (h *HealthInfo) Get() HealthInfo {
	return HealthInfo{
		Name:    h.Name,
		Version: h.Version,
		ENV:     h.ENV,
		Status:  h.Status,
	}
}
