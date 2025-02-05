package pokeapi_test

import (
	"fmt"
	"testing"

	"github.com/nerdwave-nick/pokeapi-go"
)

func TestBerry(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Limit  int
		Offset int
		Reply  pokeapi.NamedAPIResourceList
	}{
		{
			Limit:  3,
			Offset: 0,
			Reply: pokeapi.NamedAPIResourceList{
				Count:    64,
				Next:     "https://pokeapi.co/api/v2/berry?offset=3&limit=3",
				Previous: "",
				Results: []pokeapi.NamedAPIResource{
					{
						Name: "cheri",
						URL:  "https://pokeapi.co/api/v2/berry/1/",
					},
					{
						Name: "chesto",
						URL:  "https://pokeapi.co/api/v2/berry/2/",
					},
					{
						Name: "pecha",
						URL:  "https://pokeapi.co/api/v2/berry/3/",
					},
				},
			},
		},
		{
			Limit:  1,
			Offset: 1,
			Reply: pokeapi.NamedAPIResourceList{
				Count:    64,
				Next:     "https://pokeapi.co/api/v2/berry?offset=2&limit=1",
				Previous: "https://pokeapi.co/api/v2/berry?offset=0&limit=1",
				Results: []pokeapi.NamedAPIResource{
					{
						Name: "chesto",
						URL:  "https://pokeapi.co/api/v2/berry/2/",
					},
				},
			},
		},
		{
			Limit:  1,
			Offset: 63,
			Reply: pokeapi.NamedAPIResourceList{
				Count:    64,
				Next:     "",
				Previous: "https://pokeapi.co/api/v2/berry?offset=62&limit=1",
				Results: []pokeapi.NamedAPIResource{
					{
						Name: "rowap",
						URL:  "https://pokeapi.co/api/v2/berry/64/",
					},
				},
			},
		},
		{
			Limit:  1,
			Offset: 64,
			Reply: pokeapi.NamedAPIResourceList{
				Count:    64,
				Next:     "",
				Previous: "https://pokeapi.co/api/v2/berry?offset=63&limit=1",
				Results:  []pokeapi.NamedAPIResource{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry?offset=%d&limit=%d", tt.Offset, tt.Limit), func(t *testing.T) {
			client := pokeapi.NewClient(nil, nil).WithBaseURL(serverURL)
			list, err := client.Berries(tt.Limit, tt.Offset)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}

func TestBerries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		ID    string
		Reply pokeapi.Berry
	}{
		{
			ID: "1",
			Reply: pokeapi.Berry{
				Firmness: pokeapi.NamedAPIResource{
					Name: "soft",
					URL:  "https://pokeapi.co/api/v2/berry-firmness/2/",
				},
				Flavors: []pokeapi.BerryFlavorMap{
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "spicy",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/1/",
						},
						Potency: 10,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "dry",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/2/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sweet",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/3/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "bitter",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/4/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sour",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/5/",
						},
						Potency: 0,
					},
				},
				GrowthTime: 3,
				ID:         1,
				Item: pokeapi.NamedAPIResource{
					Name: "cheri-berry",
					URL:  "https://pokeapi.co/api/v2/item/126/",
				},
				MaxHarvest:       5,
				Name:             "cheri",
				NaturalGiftPower: 60,
				NaturalGiftType: pokeapi.NamedAPIResource{
					Name: "fire",
					URL:  "https://pokeapi.co/api/v2/type/10/",
				},
				Size:        20,
				Smoothness:  25,
				SoilDryness: 15,
			},
		},
		{
			ID: "32",
			Reply: pokeapi.Berry{
				Firmness: pokeapi.NamedAPIResource{
					Name: "very-soft",
					URL:  "https://pokeapi.co/api/v2/berry-firmness/1/",
				},
				Flavors: []pokeapi.BerryFlavorMap{
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "spicy",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/1/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "dry",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/2/",
						},
						Potency: 30,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sweet",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/3/",
						},
						Potency: 10,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "bitter",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/4/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sour",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/5/",
						},
						Potency: 0,
					},
				},
				GrowthTime: 15,
				ID:         32,
				Item: pokeapi.NamedAPIResource{
					Name: "pamtre-berry",
					URL:  "https://pokeapi.co/api/v2/item/157/",
				},
				MaxHarvest:       15,
				Name:             "pamtre",
				NaturalGiftPower: 70,
				NaturalGiftType: pokeapi.NamedAPIResource{
					Name: "steel",
					URL:  "https://pokeapi.co/api/v2/type/9/",
				},
				Size:        244,
				Smoothness:  35,
				SoilDryness: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry/%s", tt.ID), func(t *testing.T) {
			client := pokeapi.NewClient(nil, nil).WithBaseURL(serverURL)
			list, err := client.Berry(tt.ID)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}

func TestBerryFirmnesses(t *testing.T) {
	t.Parallel()
	tests := []struct {
		ID    string
		Reply pokeapi.Berry
	}{
		{
			ID: "1",
			Reply: pokeapi.Berry{
				Firmness: pokeapi.NamedAPIResource{
					Name: "soft",
					URL:  "https://pokeapi.co/api/v2/berry-firmness/2/",
				},
				Flavors: []pokeapi.BerryFlavorMap{
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "spicy",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/1/",
						},
						Potency: 10,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "dry",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/2/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sweet",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/3/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "bitter",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/4/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sour",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/5/",
						},
						Potency: 0,
					},
				},
				GrowthTime: 3,
				ID:         1,
				Item: pokeapi.NamedAPIResource{
					Name: "cheri-berry",
					URL:  "https://pokeapi.co/api/v2/item/126/",
				},
				MaxHarvest:       5,
				Name:             "cheri",
				NaturalGiftPower: 60,
				NaturalGiftType: pokeapi.NamedAPIResource{
					Name: "fire",
					URL:  "https://pokeapi.co/api/v2/type/10/",
				},
				Size:        20,
				Smoothness:  25,
				SoilDryness: 15,
			},
		},
		{
			ID: "32",
			Reply: pokeapi.Berry{
				Firmness: pokeapi.NamedAPIResource{
					Name: "very-soft",
					URL:  "https://pokeapi.co/api/v2/berry-firmness/1/",
				},
				Flavors: []pokeapi.BerryFlavorMap{
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "spicy",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/1/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "dry",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/2/",
						},
						Potency: 30,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sweet",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/3/",
						},
						Potency: 10,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "bitter",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/4/",
						},
						Potency: 0,
					},
					{
						Flavor: pokeapi.NamedAPIResource{
							Name: "sour",
							URL:  "https://pokeapi.co/api/v2/berry-flavor/5/",
						},
						Potency: 0,
					},
				},
				GrowthTime: 15,
				ID:         32,
				Item: pokeapi.NamedAPIResource{
					Name: "pamtre-berry",
					URL:  "https://pokeapi.co/api/v2/item/157/",
				},
				MaxHarvest:       15,
				Name:             "pamtre",
				NaturalGiftPower: 70,
				NaturalGiftType: pokeapi.NamedAPIResource{
					Name: "steel",
					URL:  "https://pokeapi.co/api/v2/type/9/",
				},
				Size:        244,
				Smoothness:  35,
				SoilDryness: 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry/%s", tt.ID), func(t *testing.T) {
			client := pokeapi.NewClient(nil, nil).WithBaseURL(serverURL)
			list, err := client.Berry(tt.ID)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}
