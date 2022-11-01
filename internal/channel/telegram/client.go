package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	sendMessageAction = "sendMessage"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

type client struct {
	host   string
	apiKey string
}

func (c *client) init(host, apiKey string) *client {
	c.host = host
	c.apiKey = apiKey
	return c
}

func (c *client) SendMessage(chatId string, message string, format string) error {
	params := url.Values{}
	params.Add("chat_id", chatId)
	params.Add("text", message)
	params.Add("parse_mode", format)

	_, err := c.do(sendMessageAction, params)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (c *client) makeRequest(action string, params url.Values) (*http.Request, error) {
	apiURL := c.makeURL(action)
	apiURL.RawQuery = params.Encode()

	return http.NewRequest(http.MethodGet, apiURL.String(), nil)
}

func (c *client) do(action string, params url.Values) ([]byte, error) {
	request, err := c.makeRequest(action, params)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http response body: %w", err)
	}

	return body, nil
}

func (c *client) makeURL(action string) *url.URL {
	botApiKeyPath := fmt.Sprintf("bot%s", c.apiKey)

	return &url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(botApiKeyPath, action),
	}
}
