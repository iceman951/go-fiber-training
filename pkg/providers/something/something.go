package something

import (
	"encoding/json"
	"calcal/pkg/utils/client"

	"github.com/gofiber/fiber/v2"
)

type SomethingProviderInterface interface {
	GetSomeData(id string) (map[string]interface{}, error)
}
type somethingProvider struct{}

func NewSomethingProvider(url string) SomethingProviderInterface {
	return &somethingProvider{}
}

func (s *somethingProvider) GetSomeData(id string) (map[string]interface{}, error) {
	headers := client.HttpHeaders{
		"key": "value",
	}

	// create new client for making a request
	cli := client.New()
	code, resp, err := cli.Get("https://www.example.com", headers)
	if err != nil {
		return nil, err
	}

	// some example to handle the http response
	if code != fiber.StatusOK {
		return nil, nil
	}

	// example map response
	var data map[string]interface{}
	errUnmarshal := json.Unmarshal(resp, &data)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}
	return data, nil
}
