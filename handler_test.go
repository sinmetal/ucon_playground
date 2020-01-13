package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/favclip/ucon/v3"
)

func TestImageHandler(t *testing.T) {
	ucon.Orthodox()

	mux := ucon.DefaultMux
	mux.HandleFunc("GET", "/api/image/{id}/s={size}", ImageHandler)

	cases := []struct {
		name string
		url  string
		id   string
		size int
	}{
		{"default", "/api/image/hogeID", "hogeID", 0},
		{"size", "/api/image/hogeID/s=100", "hogeID", 100},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			req, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			mux.ServeHTTP(resp, req)

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if e, g := http.StatusOK, resp.Code; e != g {
				t.Errorf("status want %v got %v. body=%+v", e, g, string(body))
				return
			}
			var respBody ImageHandlerResp
			if err := json.Unmarshal(body, &respBody); err != nil {
				t.Fatal(err)
			}
			if e, g := tt.id, respBody.ID; e != g {
				t.Errorf("id want %v got %v", e, g)
			}
			if e, g := tt.size, respBody.Size; e != g {
				t.Errorf("size want %v got %v", e, g)
			}
		})
	}
}
