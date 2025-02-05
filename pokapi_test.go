package pokeapi_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"slices"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/nerdwave-nick/pokeapi-go"
)

var responses = map[string][]byte{
	"/status": []byte(`{"message": "ok"}`),

	"/berry?limit=3&offset=0": []byte(`{"count":64,"next":"https://pokeapi.co/api/v2/berry?offset=3&limit=3","previous":null,"results":[{"name":"cheri","url":"https://pokeapi.co/api/v2/berry/1/"},{"name":"chesto","url":"https://pokeapi.co/api/v2/berry/2/"},{"name":"pecha","url":"https://pokeapi.co/api/v2/berry/3/"}]}`),

	"/berry?limit=1&offset=1": []byte(`{"count":64,"next":"https://pokeapi.co/api/v2/berry?offset=2&limit=1","previous":"https://pokeapi.co/api/v2/berry?offset=0&limit=1","results":[{"name":"chesto","url":"https://pokeapi.co/api/v2/berry/2/"}]}`),

	"/berry?limit=1&offset=63": []byte(`{"count":64,"next":null,"previous":"https://pokeapi.co/api/v2/berry?offset=62&limit=1","results":[{"name":"rowap","url":"https://pokeapi.co/api/v2/berry/64/"}]}`),

	"/berry?limit=1&offset=64": []byte(`{"count":64,"next":null,"previous":"https://pokeapi.co/api/v2/berry?offset=63&limit=1","results":[]}`),

	"/berry/1": []byte(`{"firmness":{"name":"soft","url":"https://pokeapi.co/api/v2/berry-firmness/2/"},"flavors":[{"flavor":{"name":"spicy","url":"https://pokeapi.co/api/v2/berry-flavor/1/"},"potency":10},{"flavor":{"name":"dry","url":"https://pokeapi.co/api/v2/berry-flavor/2/"},"potency":0},{"flavor":{"name":"sweet","url":"https://pokeapi.co/api/v2/berry-flavor/3/"},"potency":0},{"flavor":{"name":"bitter","url":"https://pokeapi.co/api/v2/berry-flavor/4/"},"potency":0},{"flavor":{"name":"sour","url":"https://pokeapi.co/api/v2/berry-flavor/5/"},"potency":0}],"growth_time":3,"id":1,"item":{"name":"cheri-berry","url":"https://pokeapi.co/api/v2/item/126/"},"max_harvest":5,"name":"cheri","natural_gift_power":60,"natural_gift_type":{"name":"fire","url":"https://pokeapi.co/api/v2/type/10/"},"size":20,"smoothness":25,"soil_dryness":15}`),

	"/berry/32": []byte(`{"firmness":{"name":"very-soft","url":"https://pokeapi.co/api/v2/berry-firmness/1/"},"flavors":[{"flavor":{"name":"spicy","url":"https://pokeapi.co/api/v2/berry-flavor/1/"},"potency":0},{"flavor":{"name":"dry","url":"https://pokeapi.co/api/v2/berry-flavor/2/"},"potency":30},{"flavor":{"name":"sweet","url":"https://pokeapi.co/api/v2/berry-flavor/3/"},"potency":10},{"flavor":{"name":"bitter","url":"https://pokeapi.co/api/v2/berry-flavor/4/"},"potency":0},{"flavor":{"name":"sour","url":"https://pokeapi.co/api/v2/berry-flavor/5/"},"potency":0}],"growth_time":15,"id":32,"item":{"name":"pamtre-berry","url":"https://pokeapi.co/api/v2/item/157/"},"max_harvest":15,"name":"pamtre","natural_gift_power":70,"natural_gift_type":{"name":"steel","url":"https://pokeapi.co/api/v2/type/9/"},"size":244,"smoothness":35,"soil_dryness":8}`),

	"/berry-firmness?limit=1&offset=4": []byte(`{"count":5,"next":null,"previous":"https://pokeapi.co/api/v2/berry-firmness?offset=3&limit=1","results":[{"name":"super-hard","url":"https://pokeapi.co/api/v2/berry-firmness/5/"}]}`),

	"/berry-firmness/2": []byte(`{"berries":[{"name":"cheri","url":"https://pokeapi.co/api/v2/berry/1/"},{"name":"figy","url":"https://pokeapi.co/api/v2/berry/11/"},{"name":"iapapa","url":"https://pokeapi.co/api/v2/berry/15/"},{"name":"bluk","url":"https://pokeapi.co/api/v2/berry/17/"},{"name":"grepa","url":"https://pokeapi.co/api/v2/berry/25/"},{"name":"tamato","url":"https://pokeapi.co/api/v2/berry/26/"},{"name":"rabuta","url":"https://pokeapi.co/api/v2/berry/29/"},{"name":"spelon","url":"https://pokeapi.co/api/v2/berry/31/"},{"name":"watmel","url":"https://pokeapi.co/api/v2/berry/33/"},{"name":"passho","url":"https://pokeapi.co/api/v2/berry/37/"},{"name":"rindo","url":"https://pokeapi.co/api/v2/berry/39/"},{"name":"chople","url":"https://pokeapi.co/api/v2/berry/41/"},{"name":"shuca","url":"https://pokeapi.co/api/v2/berry/43/"},{"name":"payapa","url":"https://pokeapi.co/api/v2/berry/45/"},{"name":"haban","url":"https://pokeapi.co/api/v2/berry/49/"},{"name":"lansat","url":"https://pokeapi.co/api/v2/berry/58/"},{"name":"micle","url":"https://pokeapi.co/api/v2/berry/61/"},{"name":"jaboca","url":"https://pokeapi.co/api/v2/berry/63/"}],"id":2,"name":"soft","names":[{"language":{"name":"ja-Hrkt","url":"https://pokeapi.co/api/v2/language/1/"},"name":"やわらかい"},{"language":{"name":"zh-Hant","url":"https://pokeapi.co/api/v2/language/4/"},"name":"柔軟"},{"language":{"name":"fr","url":"https://pokeapi.co/api/v2/language/5/"},"name":"Tendre"},{"language":{"name":"es","url":"https://pokeapi.co/api/v2/language/7/"},"name":"Blanda"},{"language":{"name":"en","url":"https://pokeapi.co/api/v2/language/9/"},"name":"Soft"},{"language":{"name":"zh-Hans","url":"https://pokeapi.co/api/v2/language/12/"},"name":"柔软"}]}`),

	"/berry-flavor?limit=3&offset=2": []byte(`{"count":5,"next":null,"previous":"https://pokeapi.co/api/v2/berry-flavor?offset=0&limit=2","results":[{"name":"sweet","url":"https://pokeapi.co/api/v2/berry-flavor/3/"},{"name":"bitter","url":"https://pokeapi.co/api/v2/berry-flavor/4/"},{"name":"sour","url":"https://pokeapi.co/api/v2/berry-flavor/5/"}]}`),

	"/berry-flavor/1": []byte(`{"berries":[{"berry":{"name":"rowap","url":"https://pokeapi.co/api/v2/berry/64/"},"potency":10},{"berry":{"name":"leppa","url":"https://pokeapi.co/api/v2/berry/6/"},"potency":10},{"berry":{"name":"oran","url":"https://pokeapi.co/api/v2/berry/7/"},"potency":10},{"berry":{"name":"persim","url":"https://pokeapi.co/api/v2/berry/8/"},"potency":10},{"berry":{"name":"lum","url":"https://pokeapi.co/api/v2/berry/9/"},"potency":10},{"berry":{"name":"razz","url":"https://pokeapi.co/api/v2/berry/16/"},"potency":10},{"berry":{"name":"pinap","url":"https://pokeapi.co/api/v2/berry/20/"},"potency":10},{"berry":{"name":"pomeg","url":"https://pokeapi.co/api/v2/berry/21/"},"potency":10},{"berry":{"name":"qualot","url":"https://pokeapi.co/api/v2/berry/23/"},"potency":10},{"berry":{"name":"hondew","url":"https://pokeapi.co/api/v2/berry/24/"},"potency":10},{"berry":{"name":"nomel","url":"https://pokeapi.co/api/v2/berry/30/"},"potency":10},{"berry":{"name":"belue","url":"https://pokeapi.co/api/v2/berry/35/"},"potency":10},{"berry":{"name":"rindo","url":"https://pokeapi.co/api/v2/berry/39/"},"potency":10},{"berry":{"name":"shuca","url":"https://pokeapi.co/api/v2/berry/43/"},"potency":10},{"berry":{"name":"charti","url":"https://pokeapi.co/api/v2/berry/47/"},"potency":10},{"berry":{"name":"apicot","url":"https://pokeapi.co/api/v2/berry/57/"},"potency":10},{"berry":{"name":"cheri","url":"https://pokeapi.co/api/v2/berry/1/"},"potency":10},{"berry":{"name":"chople","url":"https://pokeapi.co/api/v2/berry/41/"},"potency":15},{"berry":{"name":"figy","url":"https://pokeapi.co/api/v2/berry/11/"},"potency":15},{"berry":{"name":"occa","url":"https://pokeapi.co/api/v2/berry/36/"},"potency":15},{"berry":{"name":"tamato","url":"https://pokeapi.co/api/v2/berry/26/"},"potency":20},{"berry":{"name":"tanga","url":"https://pokeapi.co/api/v2/berry/46/"},"potency":20},{"berry":{"name":"babiri","url":"https://pokeapi.co/api/v2/berry/51/"},"potency":25},{"berry":{"name":"starf","url":"https://pokeapi.co/api/v2/berry/59/"},"potency":30},{"berry":{"name":"liechi","url":"https://pokeapi.co/api/v2/berry/53/"},"potency":30},{"berry":{"name":"spelon","url":"https://pokeapi.co/api/v2/berry/31/"},"potency":30},{"berry":{"name":"petaya","url":"https://pokeapi.co/api/v2/berry/56/"},"potency":30},{"berry":{"name":"lansat","url":"https://pokeapi.co/api/v2/berry/58/"},"potency":30},{"berry":{"name":"enigma","url":"https://pokeapi.co/api/v2/berry/60/"},"potency":40}],"contest_type":{"name":"cool","url":"https://pokeapi.co/api/v2/contest-type/1/"},"id":1,"name":"spicy","names":[{"language":{"name":"ja-Hrkt","url":"https://pokeapi.co/api/v2/language/1/"},"name":"からい"},{"language":{"name":"zh-Hant","url":"https://pokeapi.co/api/v2/language/4/"},"name":"辣"},{"language":{"name":"fr","url":"https://pokeapi.co/api/v2/language/5/"},"name":"Épicé"},{"language":{"name":"en","url":"https://pokeapi.co/api/v2/language/9/"},"name":"Spicy"},{"language":{"name":"zh-Hans","url":"https://pokeapi.co/api/v2/language/12/"},"name":"辣"}]}`),
}

type TestNamedAPIResourcesListValue struct {
	Limit  int
	Offset int
	Reply  pokeapi.NamedAPIResourceList
}
type TestResource[T any] struct {
	ID    string
	Reply T
}

func setUpServer(urlPointer *string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUrl := r.URL.RequestURI()
		if resp, ok := responses[reqUrl]; ok {
			w.Header().Add("content-type", "application/json")
			w.Header().Add("content-type", "charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(resp)
			if err != nil {
				fmt.Printf("something went wrong writing body: %v\n", err)
			}

			return
		}
		fmt.Printf("could not find output from %q\n", reqUrl)
		http.NotFound(w, r)
	}))
	*urlPointer = server.URL + "/"
	return server
}

func diffCompare(t *testing.T, got any, wanted any) {
	listBytes, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	wantedBytes, err := json.Marshal(wanted)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(listBytes, wantedBytes) {
		t.Fatal(
			diff.CharacterDiff(
				string(wantedBytes),
				string(listBytes),
			),
		)
	}
}

var serverURL = ""

func TestMain(m *testing.M) {
	server := setUpServer(&serverURL)
	httptest.NewRequest(http.MethodGet, serverURL+"status", nil)
	code := m.Run()
	server.Close()
	os.Exit(code)
}
