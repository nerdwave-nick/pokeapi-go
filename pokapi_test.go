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
	"/status":                          []byte(`{"message": "ok"}`),
	"/berry/1":                         Berry1,
	"/berry/32":                        Berry32,
	"/berry-firmness/2":                BerryFirmness2,
	"/berry-firmness?limit=1&offset=4": BerryFirmnessL1O4,
	"/berry-flavor/1":                  BerryFlavor1,
	"/berry-flavor?limit=3&offset=2":   BerryFlavorL3O2,
	"/berry?limit=1&offset=1":          BerryL1O1,
	"/berry?limit=1&offset=63":         BerryL1O63,
	"/berry?limit=1&offset=64":         BerryL1O64,
	"/berry?limit=3&offset=0":          BerryL3O0,
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
	if got == nil || wanted == nil {
		if got == nil {
			t.Fatal("got is nil")
		}
		t.Fatal("wated is nil")
	}
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
