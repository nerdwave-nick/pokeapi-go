package pokeapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"slices"
	"testing"

	"github.com/andreyvit/diff"
)

var responses = map[string][]byte{
	"status": []byte(`{"message": "ok"}`),

	"berry?limit=3&offset=0": []byte(`{"count":64,"next":"https://pokeapi.co/api/v2/berry?offset=3&limit=3","previous":null,"results":[{"name":"cheri","url":"https://pokeapi.co/api/v2/berry/1/"},{"name":"chesto","url":"https://pokeapi.co/api/v2/berry/2/"},{"name":"pecha","url":"https://pokeapi.co/api/v2/berry/3/"}]}`),

	"berry?limit=1&offset=1": []byte(`{"count":64,"next":"https://pokeapi.co/api/v2/berry?offset=2&limit=1","previous":"https://pokeapi.co/api/v2/berry?offset=0&limit=1","results":[{"name":"chesto","url":"https://pokeapi.co/api/v2/berry/2/"}]}`),

	"berry?limit=1&offset=63": []byte(`{"count":64,"next":null,"previous":"https://pokeapi.co/api/v2/berry?offset=62&limit=1","results":[{"name":"rowap","url":"https://pokeapi.co/api/v2/berry/64/"}]}`),

	"berry?limit=1&offset=64": []byte(`{"count":64,"next":null,"previous":"https://pokeapi.co/api/v2/berry?offset=63&limit=1","results":[]}`),

	"berry/1": []byte(`{"firmness":{"name":"soft","url":"https://pokeapi.co/api/v2/berry-firmness/2/"},"flavors":[{"flavor":{"name":"spicy","url":"https://pokeapi.co/api/v2/berry-flavor/1/"},"potency":10},{"flavor":{"name":"dry","url":"https://pokeapi.co/api/v2/berry-flavor/2/"},"potency":0},{"flavor":{"name":"sweet","url":"https://pokeapi.co/api/v2/berry-flavor/3/"},"potency":0},{"flavor":{"name":"bitter","url":"https://pokeapi.co/api/v2/berry-flavor/4/"},"potency":0},{"flavor":{"name":"sour","url":"https://pokeapi.co/api/v2/berry-flavor/5/"},"potency":0}],"growth_time":3,"id":1,"item":{"name":"cheri-berry","url":"https://pokeapi.co/api/v2/item/126/"},"max_harvest":5,"name":"cheri","natural_gift_power":60,"natural_gift_type":{"name":"fire","url":"https://pokeapi.co/api/v2/type/10/"},"size":20,"smoothness":25,"soil_dryness":15}`),

	"berry/32": []byte(`{"firmness":{"name":"very-soft","url":"https://pokeapi.co/api/v2/berry-firmness/1/"},"flavors":[{"flavor":{"name":"spicy","url":"https://pokeapi.co/api/v2/berry-flavor/1/"},"potency":0},{"flavor":{"name":"dry","url":"https://pokeapi.co/api/v2/berry-flavor/2/"},"potency":30},{"flavor":{"name":"sweet","url":"https://pokeapi.co/api/v2/berry-flavor/3/"},"potency":10},{"flavor":{"name":"bitter","url":"https://pokeapi.co/api/v2/berry-flavor/4/"},"potency":0},{"flavor":{"name":"sour","url":"https://pokeapi.co/api/v2/berry-flavor/5/"},"potency":0}],"growth_time":15,"id":32,"item":{"name":"pamtre-berry","url":"https://pokeapi.co/api/v2/item/157/"},"max_harvest":15,"name":"pamtre","natural_gift_power":70,"natural_gift_type":{"name":"steel","url":"https://pokeapi.co/api/v2/type/9/"},"size":244,"smoothness":35,"soil_dryness":8}`),
}

func setUpServer(urlPointer *string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUrl := r.URL.RequestURI()
		if resp, ok := responses[reqUrl]; ok {
			w.Header().Add("content-type", "application/json")
			w.Header().Add("content-type", "charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		}
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
