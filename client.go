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

func (c *Client) doUncached(output any, endpoint string) error {
	endpoint = c.sanitizeBaseURL(endpoint)

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, output)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) do(output any, endpoint string) error {
	endpoint = c.sanitizeBaseURL(endpoint)

	found, err := c.cache.Get(endpoint, output)
	if err != nil {
		return err
	}
	if found {
		return nil
	}

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, output)
	if err != nil {
		return err
	}

	err = c.cache.Set(endpoint, output)
	if err != nil {
		return err
	}
	return nil
}
