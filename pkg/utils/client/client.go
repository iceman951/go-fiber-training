package client

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ClientInterface interface {
	Get(url string, headers HttpHeaders) (int, []byte, error)
	Post(url string, payload interface{}, headers HttpHeaders) (int, []byte, error)
	Patch(url string, payload interface{}, headers HttpHeaders) (int, []byte, error)
	Put(url string, payload interface{}, headers HttpHeaders) (int, []byte, error)
	Delete(url string, payload interface{}, headers HttpHeaders) (int, []byte, error)
	Use(method, url string, payload interface{}, headers HttpHeaders) (int, []byte, error) // -- for dynamic method option
}

type clientProperties struct {
	agent  *fiber.Agent
	config ClientConfig
}

func New(config ...ClientConfig) ClientInterface {
	agent, cfg := setupAgent(config...)
	return &clientProperties{
		agent:  agent,
		config: cfg,
	}
}

// getConfig: utils for getting current config
func (cli *clientProperties) getConfig() ClientConfig {
	return cli.config
}

// getResponse: utils for calling api
func (cli *clientProperties) getResponse(url string, headers HttpHeaders, payload interface{}) (int, []byte, error) {
	if cli.config.SkipSSLVerify {
		cli.agent.InsecureSkipVerify()
	}

	req := cli.agent.Request()
	req.SetRequestURI(url)

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// payload is not nil
	if payload != nil {
		cli.agent.JSON(payload)
	}

	if err := cli.agent.Parse(); err != nil {
		// case: unable to call
		return fiber.StatusInternalServerError, nil, err
	}

	code, resp, errs := cli.agent.Bytes()
	if len(errs) != 0 {
		return code, resp, errs[0]
	}
	return code, resp, nil
}

// Get: Send request with GET
func (cli *clientProperties) Get(url string, headers HttpHeaders) (int, []byte, error) {
	cli.agent = fiber.Get(url)

	return cli.getResponse(url, headers, nil)
}

// Post: Send request with POST
func (cli *clientProperties) Post(url string, payload interface{}, headers HttpHeaders) (int, []byte, error) {
	cli.agent = fiber.Post(url)

	return cli.getResponse(url, headers, payload)
}

// Patch: Send request with POST
func (cli *clientProperties) Patch(url string, payload interface{}, headers HttpHeaders) (int, []byte, error) {
	cli.agent = fiber.Patch(url)

	return cli.getResponse(url, headers, payload)
}

// Put: Send request with POST
func (cli *clientProperties) Put(url string, payload interface{}, headers HttpHeaders) (int, []byte, error) {
	cli.agent = fiber.Put(url)

	return cli.getResponse(url, headers, payload)
}

// Delete: Send request with POST
func (cli *clientProperties) Delete(url string, payload interface{}, headers HttpHeaders) (int, []byte, error) {
	cli.agent = fiber.Delete(url)

	return cli.getResponse(url, headers, payload)
}

// Use: Send request with customizable method
func (cli *clientProperties) Use(method, url string, payload interface{}, headers HttpHeaders) (int, []byte, error) {
	methodUpper := strings.ToUpper(method)
	switch methodUpper {
	case fiber.MethodPost:
		return cli.Post(url, payload, headers)
	case fiber.MethodPatch:
		return cli.Patch(url, payload, headers)
	case fiber.MethodPut:
		return cli.Put(url, payload, headers)
	case fiber.MethodDelete:
		return cli.Delete(url, payload, headers)
	default:
		return cli.Get(url, headers)
	}
}

// Helper
func setupAgent(config ...ClientConfig) (*fiber.Agent, ClientConfig) {
	agent := fiber.AcquireAgent()
	agent.ContentType(fiber.MIMEApplicationJSONCharsetUTF8)

	var cfg ClientConfig

	// Return default config if nothing provided
	if len(config) < 1 {
		cfg = DefaultClientConfig
	} else {
		cfg = config[0]
	}

	// set timeout when specific, else used default
	// Ex. timeout := time.Duration(10) * time.Second
	if cfg.Timeout != 0 {
		agent.Timeout(cfg.Timeout)
	}

	return agent, cfg
}
