package pokeapi_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var responses = map[string][]byte{
	"status": []byte(`{"message": "ok"}`),

	"berry?limit=3&offset=0": []byte(`{"count":64,"next":"https://pokeapi.co/api/v2/berry?offset=3&limit=3","previous":null,"results":[{"name":"cheri","url":"https://pokeapi.co/api/v2/berry/1/"},{"name":"chesto","url":"https://pokeapi.co/api/v2/berry/2/"},{"name":"pecha","url":"https://pokeapi.co/api/v2/berry/3/"}]}`),

	"berry?limit=1&offset=1": []byte(`{"count":64,"next":"https://pokeapi.co/api/v2/berry?offset=2&limit=1","previous":"https://pokeapi.co/api/v2/berry?offset=0&limit=1","results":[{"name":"chesto","url":"https://pokeapi.co/api/v2/berry/2/"}]}`),

	"berry?limit=1&offset=63": []byte(`{"count":64,"next":null,"previous":"https://pokeapi.co/api/v2/berry?offset=62&limit=1","results":[{"name":"rowap","url":"https://pokeapi.co/api/v2/berry/64/"}]}`),

	"berry?limit=1&offset=64": []byte(`{"count":64,"next":null,"previous":"https://pokeapi.co/api/v2/berry?offset=63&limit=1","results":[]}`),
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

var serverURL = ""

func TestMain(m *testing.M) {
	server := setUpServer(&serverURL)
	httptest.NewRequest(http.MethodGet, serverURL+"status", nil)
	code := m.Run()
	server.Close()
	os.Exit(code)
}
