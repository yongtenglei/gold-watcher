package main

import (
	"bytes"
	"fyne.io/fyne/v2/app"
	"goldwatcher/repository"
	"io"
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {

	a := app.New()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.Client = client
	testApp.DB = repository.NewTestRepository()

	os.Exit(m.Run())
}

var jsonToReturn = `
{
    "ts": 1663001482040,
    "tsj": 1663001475929,
    "date": "Sep 12th 2022, 12:51:15 pm NY",
    "items": [
        {
            "curr": "CNY",
            "xauPrice": 11988.3862,
            "xagPrice": 138.0053,
            "chgXau": 91.0502,
            "chgXag": 7.4458,
            "pcXau": 0.7653,
            "pcXag": 5.703,
            "xauClose": 11897.33596,
            "xagClose": 130.55948
        }
    ]
}`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
