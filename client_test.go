package pokeapi_test

import (
	"encoding/json"
	"testing"

	"github.com/nerdwave-nick/pokeapi-go"
)

type mockCache struct {
	LastStoredValue []byte
}

func (c *mockCache) Set(endpoint string, value any) error {
	newValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	c.LastStoredValue = newValue
	return nil
}

func (c *mockCache) Get(endpoint string, value any) (bool, error) {
	if c.LastStoredValue == nil {
		return false, nil
	}

	err := json.Unmarshal(c.LastStoredValue, value)
	if err != nil {
		return true, err
	}

	return true, nil
}

func TestClientDo(t *testing.T) {
	t.Run("set to cache", func(t *testing.T) {
		cache := mockCache{}
		client := pokeapi.NewWithBaseURL(&cache, nil, serverURL)
		list, err := client.Berry("1")
		if err != nil {
			t.Fatal(err)
		}
		b := pokeapi.Berry{}
		got, err := cache.Get("", &b)
		if err != nil {
			t.Fatal(err)
		}
		if !got {
			t.Fatal("couldn't get from cache")
		}
		diffCompare(t, list, b)
	})
	t.Run("get from cache", func(t *testing.T) {
		madeUpCachedBerry := pokeapi.Berry{
			Firmness: pokeapi.NamedAPIResource{
				Name: "wrong",
				URL:  "made up",
			},
			Flavors: []pokeapi.BerryFlavorMap{
				{
					Flavor: pokeapi.NamedAPIResource{
						Name: "spicy",
						URL:  "https://pokeapi.co/api/v2/berry-flavor/1/",
					},
					Potency: 10,
				},
			},
			GrowthTime: 9,
			ID:         3,
			Item: pokeapi.NamedAPIResource{
				Name: "wrong item",
				URL:  "some made up item url",
			},
			MaxHarvest:       2,
			Name:             "aaa",
			NaturalGiftPower: 3,
			NaturalGiftType: pokeapi.NamedAPIResource{
				Name: "none",
				URL:  "weak natural gift link",
			},
			Size:        0,
			Smoothness:  5,
			SoilDryness: 5,
		}
		bytes, err := json.Marshal(madeUpCachedBerry)
		if err != nil {
			t.Fatal(err)
		}

		cache := mockCache{LastStoredValue: bytes}
		client := pokeapi.NewWithBaseURL(&cache, nil, serverURL)
		res, err := client.Berry("1")
		if err != nil {
			t.Fatal(err)
		}
		diffCompare(t, res, madeUpCachedBerry)
	})
}

func TestClientDoUncached(t *testing.T) {
	t.Run("don't set to cache", func(t *testing.T) {
		cache := mockCache{LastStoredValue: nil}
		client := pokeapi.NewWithBaseURL(&cache, nil, serverURL)
		_, err := client.Berries(1, 1)
		if err != nil {
			t.Fatal(err)
		}
		if cache.LastStoredValue != nil {
			t.Fatal("couldn't get from cache")
		}
	})
	t.Run("don't get from cache", func(t *testing.T) {
		cache := mockCache{LastStoredValue: []byte(`{"count":23,"next":"","previous":"","results":[]}`)}
		client := pokeapi.NewWithBaseURL(&cache, nil, serverURL)
		res, err := client.Berries(1, 1)
		if err != nil {
			t.Fatal(err)
		}
		if res.Count == 23 {
			t.Fatal("got value from cache instead of server")
		}
	})
}
