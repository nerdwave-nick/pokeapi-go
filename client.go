package pokeapi

import (
	"encoding/json"
	"fmt"
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
	BaseURL string
}

func New(cache Cache, client *http.Client) *Client {
	return NewWithBaseURL(cache, client, base_url)
}

func NewWithBaseURL(cache Cache, client *http.Client, baseUrl string) *Client {
	if cache == nil {
		cache = &noopCache{}
	}
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		cache:   cache,
		client:  client,
		BaseURL: baseUrl,
	}
}

func (c *Client) sanitizeBaseURL(endpoint string) string {
	_, after, found := strings.Cut(endpoint, c.BaseURL)
	if found {
		endpoint = after
	}
	return c.BaseURL + endpoint
}

func (c *Client) doUncached(output any, endpoint string) error {
	endpoint = c.sanitizeBaseURL(endpoint)

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return fmt.Errorf("requesting body from %q: %w", endpoint, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading body from %q: %w", endpoint, err)
	}
	err = json.Unmarshal(body, output)
	if err != nil {
		fmt.Printf("error unmarshalling body:\n%s\n", body)
		return fmt.Errorf("unmarshalling json from %q: %w", endpoint, err)
	}
	return nil
}

func (c *Client) do(output any, endpoint string) error {
	endpoint = c.sanitizeBaseURL(endpoint)

	found, err := c.cache.Get(endpoint, output)
	if err != nil {
		return fmt.Errorf("checking cache for %q: %w", endpoint, err)
	}
	if found {
		return nil
	}

	resp, err := c.client.Get(endpoint)
	if err != nil {
		return fmt.Errorf("requesting body from %q: %w", endpoint, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading body from %q: %w", endpoint, err)
	}
	err = json.Unmarshal(body, output)
	if err != nil {
		fmt.Printf("error unmarshalling body:\n\n%s", body)
		return fmt.Errorf("unmarshalling json from %q: %w", endpoint, err)
	}

	err = c.cache.Set(endpoint, output)
	if err != nil {
		return fmt.Errorf("writing cache for %q: %w", endpoint, err)
	}
	return nil
}
