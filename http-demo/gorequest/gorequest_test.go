package gorequest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/parnurzeal/gorequest"
)

func TestGet(t *testing.T) {
	const case1Empty = "/"
	const case2SetHeader = "/set_header"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != gorequest.GET {
			t.Errorf("Expected method %q; got %q", gorequest.GET, r.Method)
		}
		if r.Header == nil {
			t.Error("Expected non-nil request Header")
		}
		switch r.URL.Path {
		case case1Empty:
			t.Logf("caee %v ", case1Empty)
		case case2SetHeader:
			t.Logf("case %v ", case2SetHeader)
			if r.Header.Get("API-Key") != "fookey" {
				t.Errorf("Expected 'API-Key' == %q; got %q", "fookey", r.Header.Get("API-Key"))
			}
		default:
			t.Errorf("No testing for this case yet : %q", r.URL.Path)
		}
	}))

	defer ts.Close()

	t.Logf("url: %s\n", ts.URL)

	gorequest.New().Get(ts.URL + case1Empty).
		End()
	gorequest.New().Get(ts.URL+case2SetHeader).
		Set("API-Key", "fookey").
		End()
}

type (
	heyYou struct {
		Hey string `json:"hey"`
	}
)

func TestEndStruct(t *testing.T) {
	var resStruct heyYou
	expStruct := heyYou{Hey: "you"}
	serverOutput, err := json.Marshal(expStruct)
	if err != nil {
		t.Errorf("Unexpected errors: %s", err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(serverOutput)
	}))
	defer ts.Close()

	// Callback
	{
		resp, bodyBytes, errs := gorequest.New().Get(ts.URL).EndStruct(&resStruct, func(resp gorequest.Response, v interface{}, body []byte, errs []error) {
			if len(errs) > 0 {
				t.Fatalf("Unexpected errors: %s", errs)
			}
			if resp.StatusCode != 200 {
				t.Fatalf("Expected StatusCode=200, actual StatusCode=%v", resp.StatusCode)
			}
			if !reflect.DeepEqual(expStruct, resStruct) {
				resBytes, _ := json.Marshal(resStruct)
				t.Errorf("Expected body=%s, actual bodyBytes=%s", serverOutput, string(resBytes))
			}
			if !reflect.DeepEqual(body, serverOutput) {
				t.Errorf("Expected bodyBytes=%s, actual bodyBytes=%s", serverOutput, string(body))
			}
		})

		if len(errs) > 0 {
			t.Fatalf("Unexpected errors: %s", errs)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Expected StatusCode=200, actual StatusCode=%v", resp.StatusCode)
		}
		if !reflect.DeepEqual(bodyBytes, serverOutput) {
			t.Errorf("Expected bodyBytes=%s, actual bodyBytes=%s", serverOutput, string(bodyBytes))
		}
	}

	// No callback.
	{
		resp, bodyBytes, errs := gorequest.New().Get(ts.URL).EndStruct(&resStruct)
		if len(errs) > 0 {
			t.Errorf("Unexpected errors: %s", errs)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected StatusCode=200, actual StatusCode=%v", resp.StatusCode)
		}
		if !reflect.DeepEqual(expStruct, resStruct) {
			resBytes, _ := json.Marshal(resStruct)
			t.Errorf("Expected body=%s, actual bodyBytes=%s", serverOutput, string(resBytes))
		}
		if !reflect.DeepEqual(bodyBytes, serverOutput) {
			t.Errorf("Expected bodyBytes=%s, actual bodyBytes=%s", serverOutput, string(bodyBytes))
		}
	}
}
