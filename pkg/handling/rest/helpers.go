package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// decodeBody takes a Request and a intended decode target, hydrating <t> from
// the Request body if possible. Returns an error otherwise.
func decodeBody(r *http.Request, t interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(t)
}

func encodeBody(w http.ResponseWriter, t interface{}) error {
	return json.NewEncoder(w).Encode(t)
}

func respond(w http.ResponseWriter, status int, data interface{}) {
	hasData := data != nil
	SetCommonHeaders(w, hasData)

	w.WriteHeader(status)

	if hasData {
		encodeBody(w, data)
	}
}

func respondErr(w http.ResponseWriter, status int, args ...interface{}) {
	respond(w, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func respondHTTPErr(w http.ResponseWriter, status int) {
	respondErr(w, status, http.StatusText(status))
}

// NOTE exported so that main can use it on OPTIONS req handling...
func SetCommonHeaders(w http.ResponseWriter, d bool) {
	// TODO strict CORS handling?
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// NOTE in an Event-Sourced application there is only ever GET and POST
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	if d {
		w.Header().Set("Content-Type", "application/json")
	}
}

// getPostBody requires a mocked object acceptable for http.NewRequest
type mockBody struct{ Reader strings.Reader }

func (b *mockBody) Read(a []byte) (int, error) { return b.Reader.Read(a) }
func (b *mockBody) Close() error               { return nil }

// getPostBody abstracts the boilerplate logic of generating a useable "body" for an `http.NewRequest("POST", ...)`
func getPostBody(d string) *mockBody {
	r := strings.NewReader(d)
	return &mockBody{Reader: *r}
}
