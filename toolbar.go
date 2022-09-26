package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"goldwatcher/repository"
	"strconv"
	"time"
)

func (app *Config) getToolBar() *widget.Toolbar {
	tb := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()

		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {

		}),
	)
	return tb
}

func (app *Config) addHoldingsDialog() dialog.Dialog {
	addAmountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

	app.AddHoldingsPurchaseAmountEntry = addAmountEntry
	app.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	dateValidator := func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}
		return nil
	}

	amountValidator := func(s string) error {
		if _, err := strconv.Atoi(s); err != nil {
			return err
		}
		return nil
	}

	priceValidator := func(s string) error {
		if _, err := strconv.ParseFloat(s, 32); err != nil {
			return err
		}
		return nil
	}

	purchaseDateEntry.Validator = dateValidator
	purchasePriceEntry.Validator = priceValidator
	addAmountEntry.Validator = amountValidator

	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	// create a dialog
	addForm := dialog.NewForm(
		"Add gold",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Amount in g", Widget: addAmountEntry},
			{Text: "Purchase Price", Widget: purchasePriceEntry},
			{Text: "Purchase Date", Widget: purchaseDateEntry},
		},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(addAmountEntry.Text)
				purchaseDate, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)
				purchasePrice, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)

				purchasePrice = purchasePrice * 100

				_, err := app.DB.InsertHolding(repository.Holdings{
					Amount:        amount,
					PurchaseDate:  purchaseDate,
					PurchasePrice: int(purchasePrice),
				})
				if err != nil {
					app.ErrorLog.Println(err)
				}

				app.refreshHoldingsTable()
			}
		},
		app.MainWindow,
	)

	// size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
