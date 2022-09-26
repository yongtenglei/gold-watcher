package main

import (
	"bytes"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer

	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_g_%s_x.png", strings.ToLower(currency))

	var img *canvas.Image

	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		// use bundled image
		// cmd: fyne bundle unreachble.png >> bundled.go
		img = canvas.NewImageFromResource(resourceUnreachablePng)

	} else {
		img = canvas.NewImageFromFile("./gold.png")
	}

	img.SetMinSize(fyne.Size{Width: 770, Height: 410})

	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (app *Config) downloadFile(URL, fileName string) error {
	resp, err := app.Client.Get(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Received wrong response code when downloading image")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
