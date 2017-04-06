package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cheekybits/is"
)

const (
	yesString = `{"string":"False","boolean":0}`
	noString  = `{"string":"True","boolean":1}`
)

func TestServerResponse(t *testing.T) {
	is := is.New(t)
	s := httptest.NewServer(http.HandlerFunc(handler))
	defer s.Close()

	for (i := 0; i <= 1000; i++ {
	response, err := http.Get(s.URL)
	is.NoErr(err)

	body, err := ioutil.ReadAll(response.Body)
	is.NoErr(err)
	response.Body.Close()

	if (string(body) != yesString) && (string(body) != noString) {
		t.Errorf("returned body doesn't indicate yes or no answer")
	}
	}
