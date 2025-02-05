package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	base_url     = "https://pokeapi.co/api/v2/"
	limit_param  = "limit"
	offset_param = "offset"
)

type Cache interface {
	// set, with value being a structure
	Set(endpoint string, value any) error
	// get, with return values being first an unmarshalled structure, then bool whether somethnig was found, and then error if something went wrong
	Get(endpoint string, value any) (bool, error)
}

type noopCache struct{}

func (*noopCache) Set(endpoint string, value any) error {
	return nil
}

func (*noopCache) Get(endpoint string, value any) (bool, error) {
	return false, nil
}

type Client struct {
	cache   Cache
	client  *http.Client
	baseUrl string
}

func NewClient(cache Cache, client *http.Client) *Client {
	if cache == nil {
		cache = &noopCache{}
	}
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		cache:   cache,
		client:  client,
		baseUrl: base_url,
	}
}

func (c *Client) WithBaseURL(newBaseUrl string) *Client {
	return &Client{
		cache:   c.cache,
		client:  c.client,
		baseUrl: c.baseUrl,
	}
}

func (c *Client) sanitizeBaseURL(endpoint string) string {
	_, after, found := strings.Cut(endpoint, c.baseUrl)
	if found {
		endpoint = after
	}
	return c.baseUrl + endpoint
}

func doUncached[T any](c *Client, endpoint string) (*T, error) {
	endpoint = c.sanitizeBaseURL(endpoint)

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	v := new(T)
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func do[T any](c *Client, endpoint string) (*T, error) {
	_, after, found := strings.Cut(endpoint, c.baseUrl)
	if found {
		endpoint = after
	}

	value := new(T)
	found, err := c.cache.Get(endpoint, value)
	if err != nil {
		return nil, err
	}
	if found {
		return value, nil
	}

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, value)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(endpoint, value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
