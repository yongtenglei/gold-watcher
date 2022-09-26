package main

import (
	"net/http"
	"os"
	"testing"
)

func TestConfig_downloadFile(t *testing.T) {
	client := testApp.Client

	defer func() {
		testApp.Client = client
	}()

	testApp.Client = &http.Client{}

	err := testApp.downloadFile("https://goldprice.org/charts/gold_3d_b_g_cny_x.png", "./test.png")
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat("./test.png")
	if err != nil {
		t.Error(err)
	}

	err = os.Remove("./test.png")
	if err != nil {
		t.Error(err)
	}
}
