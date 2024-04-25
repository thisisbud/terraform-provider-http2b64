package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLDoesnotExist(t *testing.T) {

	// Test for a Dodgy URL
	Nothing, statusCode, err := GetFile("http://a.url.that.doesnt.exist.com")
	if Nothing != "" {
		t.Errorf("Response Should be Empty but got '%v'", Nothing)
	}
	if statusCode != 0 {
		t.Errorf("Response Code Should be 0 but got '%v'", statusCode)
	}
	if err == nil {
		t.Error("Should have got an error, but got nil error ")
	}
}

func TestOKResponse(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, "OK")
	}))
	defer ts.Close()

	Response, statusCode, err := GetFile(ts.URL)
	if Response != "T0sK" {
		t.Errorf("Response should be 'T0sK' but was '%v'", Response)
	}
	if statusCode != 200 {
		t.Errorf("Response Code should be '200' but was '%d'", statusCode)
	}

	if err != nil {
		t.Errorf("Response Error should be nil but was '%v'", err)
	}

}

func Test404Response(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "404 Not found")
	}))
	defer ts.Close()

	_, statusCode, err := GetFile(ts.URL)

	if statusCode != 404 {
		t.Errorf("Response Code should be '404' but was '%d'", statusCode)
	}

	if err != nil {
		t.Errorf("Response Error should be nil but was '%v'", err)
	}

}
