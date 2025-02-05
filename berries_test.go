package pokeapi_test

import (
	"fmt"
	"testing"

	"github.com/nerdwave-nick/pokeapi-go"
)

func TestBerries(t *testing.T) {
	t.Parallel()
	tests := []TestNamedAPIResourcesListValue{
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
			client := pokeapi.NewWithBaseURL(nil, nil, serverURL)
			list, err := client.Berries(tt.Limit, tt.Offset)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}

func TestBerry(t *testing.T) {
	t.Parallel()
	tests := []TestResource[pokeapi.Berry]{
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
			client := pokeapi.NewWithBaseURL(nil, nil, serverURL)
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
	tests := []TestNamedAPIResourcesListValue{
		{
			Limit:  1,
			Offset: 4,
			Reply: pokeapi.NamedAPIResourceList{
				Count:    5,
				Next:     "",
				Previous: "https://pokeapi.co/api/v2/berry-firmness?offset=3&limit=1",
				Results: []pokeapi.NamedAPIResource{
					{
						Name: "super-hard",
						URL:  "https://pokeapi.co/api/v2/berry-firmness/5/",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry-firmness?limit=%d&offset=%d", tt.Limit, tt.Offset), func(t *testing.T) {
			client := pokeapi.NewWithBaseURL(nil, nil, serverURL)
			list, err := client.BerryFirmnesses(tt.Limit, tt.Offset)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}

func TestBerryFirmness(t *testing.T) {
	t.Parallel()
	tests := []TestResource[pokeapi.BerryFirmness]{
		{
			ID: "2",
			Reply: pokeapi.BerryFirmness{
				Berries: []pokeapi.NamedAPIResource{
					{
						Name: "cheri",
						URL:  "https://pokeapi.co/api/v2/berry/1/",
					},
					{
						Name: "figy",
						URL:  "https://pokeapi.co/api/v2/berry/11/",
					},
					{
						Name: "iapapa",
						URL:  "https://pokeapi.co/api/v2/berry/15/",
					},
					{
						Name: "bluk",
						URL:  "https://pokeapi.co/api/v2/berry/17/",
					},
					{
						Name: "grepa",
						URL:  "https://pokeapi.co/api/v2/berry/25/",
					},
					{
						Name: "tamato",
						URL:  "https://pokeapi.co/api/v2/berry/26/",
					},
					{
						Name: "rabuta",
						URL:  "https://pokeapi.co/api/v2/berry/29/",
					},
					{
						Name: "spelon",
						URL:  "https://pokeapi.co/api/v2/berry/31/",
					},
					{
						Name: "watmel",
						URL:  "https://pokeapi.co/api/v2/berry/33/",
					},
					{
						Name: "passho",
						URL:  "https://pokeapi.co/api/v2/berry/37/",
					},
					{
						Name: "rindo",
						URL:  "https://pokeapi.co/api/v2/berry/39/",
					},
					{
						Name: "chople",
						URL:  "https://pokeapi.co/api/v2/berry/41/",
					},
					{
						Name: "shuca",
						URL:  "https://pokeapi.co/api/v2/berry/43/",
					},
					{
						Name: "payapa",
						URL:  "https://pokeapi.co/api/v2/berry/45/",
					},
					{
						Name: "haban",
						URL:  "https://pokeapi.co/api/v2/berry/49/",
					},
					{
						Name: "lansat",
						URL:  "https://pokeapi.co/api/v2/berry/58/",
					},
					{
						Name: "micle",
						URL:  "https://pokeapi.co/api/v2/berry/61/",
					},
					{
						Name: "jaboca",
						URL:  "https://pokeapi.co/api/v2/berry/63/",
					},
				},
				ID:   2,
				Name: "soft",
				Names: []pokeapi.Name{
					{
						Language: pokeapi.NamedAPIResource{
							Name: "ja-Hrkt",
							URL:  "https://pokeapi.co/api/v2/language/1/",
						},
						Name: "やわらかい",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "zh-Hant",
							URL:  "https://pokeapi.co/api/v2/language/4/",
						},
						Name: "柔軟",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "fr",
							URL:  "https://pokeapi.co/api/v2/language/5/",
						},
						Name: "Tendre",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "es",
							URL:  "https://pokeapi.co/api/v2/language/7/",
						},
						Name: "Blanda",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "en",
							URL:  "https://pokeapi.co/api/v2/language/9/",
						},
						Name: "Soft",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "zh-Hans",
							URL:  "https://pokeapi.co/api/v2/language/12/",
						},
						Name: "柔软",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry-firmness/%s", tt.ID), func(t *testing.T) {
			client := pokeapi.NewWithBaseURL(nil, nil, serverURL)
			list, err := client.BerryFirmness(tt.ID)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}

func TestBerryFlavors(t *testing.T) {
	t.Parallel()
	tests := []TestNamedAPIResourcesListValue{
		{
			Limit:  3,
			Offset: 2,
			Reply: pokeapi.NamedAPIResourceList{
				Count:    5,
				Next:     "",
				Previous: "https://pokeapi.co/api/v2/berry-flavor?offset=0&limit=2",
				Results: []pokeapi.NamedAPIResource{
					{
						Name: "sweet",
						URL:  "https://pokeapi.co/api/v2/berry-flavor/3/",
					},
					{
						Name: "bitter",
						URL:  "https://pokeapi.co/api/v2/berry-flavor/4/",
					},
					{
						Name: "sour",
						URL:  "https://pokeapi.co/api/v2/berry-flavor/5/",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry-firmness?limit=%d&offset=%d", tt.Limit, tt.Offset), func(t *testing.T) {
			client := pokeapi.NewWithBaseURL(nil, nil, serverURL)
			list, err := client.BerryFlavors(tt.Limit, tt.Offset)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}

func TestBerryFlavor(t *testing.T) {
	t.Parallel()
	tests := []TestResource[pokeapi.BerryFlavor]{
		{
			ID: "1",
			Reply: pokeapi.BerryFlavor{
				Berries: []pokeapi.FlavorBerryMap{
					{
						Berry: pokeapi.NamedAPIResource{
							Name: "rowap",
							URL:  "https://pokeapi.co/api/v2/berry/64/",
						},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "leppa", URL: "https://pokeapi.co/api/v2/berry/6/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "oran", URL: "https://pokeapi.co/api/v2/berry/7/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "persim", URL: "https://pokeapi.co/api/v2/berry/8/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "lum", URL: "https://pokeapi.co/api/v2/berry/9/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "razz", URL: "https://pokeapi.co/api/v2/berry/16/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "pinap", URL: "https://pokeapi.co/api/v2/berry/20/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "pomeg", URL: "https://pokeapi.co/api/v2/berry/21/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "qualot", URL: "https://pokeapi.co/api/v2/berry/23/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "hondew", URL: "https://pokeapi.co/api/v2/berry/24/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "nomel", URL: "https://pokeapi.co/api/v2/berry/30/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "belue", URL: "https://pokeapi.co/api/v2/berry/35/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "rindo", URL: "https://pokeapi.co/api/v2/berry/39/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "shuca", URL: "https://pokeapi.co/api/v2/berry/43/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "charti", URL: "https://pokeapi.co/api/v2/berry/47/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "apicot", URL: "https://pokeapi.co/api/v2/berry/57/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "cheri", URL: "https://pokeapi.co/api/v2/berry/1/"},
						Potency: 10,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "chople", URL: "https://pokeapi.co/api/v2/berry/41/"},
						Potency: 15,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "figy", URL: "https://pokeapi.co/api/v2/berry/11/"},
						Potency: 15,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "occa", URL: "https://pokeapi.co/api/v2/berry/36/"},
						Potency: 15,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "tamato", URL: "https://pokeapi.co/api/v2/berry/26/"},
						Potency: 20,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "tanga", URL: "https://pokeapi.co/api/v2/berry/46/"},
						Potency: 20,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "babiri", URL: "https://pokeapi.co/api/v2/berry/51/"},
						Potency: 25,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "starf", URL: "https://pokeapi.co/api/v2/berry/59/"},
						Potency: 30,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "liechi", URL: "https://pokeapi.co/api/v2/berry/53/"},
						Potency: 30,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "spelon", URL: "https://pokeapi.co/api/v2/berry/31/"},
						Potency: 30,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "petaya", URL: "https://pokeapi.co/api/v2/berry/56/"},
						Potency: 30,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "lansat", URL: "https://pokeapi.co/api/v2/berry/58/"},
						Potency: 30,
					},
					{
						Berry:   pokeapi.NamedAPIResource{Name: "enigma", URL: "https://pokeapi.co/api/v2/berry/60/"},
						Potency: 40,
					},
				},
				ContestType: pokeapi.NamedAPIResource{
					Name: "cool",
					URL:  "https://pokeapi.co/api/v2/contest-type/1/",
				},
				ID:   1,
				Name: "spicy",
				Names: []pokeapi.Name{
					{
						Language: pokeapi.NamedAPIResource{
							Name: "ja-Hrkt",
							URL:  "https://pokeapi.co/api/v2/language/1/",
						},
						Name: "からい",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "zh-Hant",
							URL:  "https://pokeapi.co/api/v2/language/4/",
						},
						Name: "辣",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "fr",
							URL:  "https://pokeapi.co/api/v2/language/5/",
						},
						Name: "Épicé",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "en",
							URL:  "https://pokeapi.co/api/v2/language/9/",
						},
						Name: "Spicy",
					},
					{
						Language: pokeapi.NamedAPIResource{
							Name: "zh-Hans",
							URL:  "https://pokeapi.co/api/v2/language/12/",
						},
						Name: "辣",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("berry-flavor/%s", tt.ID), func(t *testing.T) {
			client := pokeapi.NewWithBaseURL(nil, nil, serverURL)
			list, err := client.BerryFlavor(tt.ID)
			if err != nil {
				t.Fatal(err)
			}
			diffCompare(t, list, tt.Reply)
		})
	}
}
