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
		Reply  *pokeapi.NamedAPIResourceList
	}{
		{
			Limit:  3,
			Offset: 0,
			Reply: &pokeapi.NamedAPIResourceList{
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
			Reply: &pokeapi.NamedAPIResourceList{
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
			Reply: &pokeapi.NamedAPIResourceList{
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
			Reply: &pokeapi.NamedAPIResourceList{
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
			if list.Count != tt.Reply.Count {
				t.Fatalf("different count, want: %d, got: %d", tt.Reply.Count, list.Count)
			}
			if list.Next != tt.Reply.Next {
				t.Fatalf("different next, want: %q, got: %q", tt.Reply.Next, list.Next)
			}
			if list.Previous != tt.Reply.Previous {
				t.Fatalf("different previous, want: %q, got: %q", tt.Reply.Previous, list.Previous)
			}
			if len(list.Results) != len(tt.Reply.Results) {
				t.Fatalf("different result count, want: %d, got: %d", len(tt.Reply.Results), len(list.Results))
			}
			for i := range list.Results {
				if list.Results[i].Name != tt.Reply.Results[i].Name {
					t.Fatalf("different result.i.Name, want: %q, got: %q", tt.Reply.Results[i].Name, list.Results[i].Name)
				}
				if list.Results[i].URL != tt.Reply.Results[i].URL {
					t.Fatalf("different result.i.URL, want: %q, got: %q", tt.Reply.Results[i].URL, list.Results[i].URL)
				}
			}
		})
	}
}
