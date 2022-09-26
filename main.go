package main

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"goldwatcher/repository"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                            fyne.App
	InfoLog                        *log.Logger
	ErrorLog                       *log.Logger
	MainWindow                     fyne.Window
	PriceContainer                 *fyne.Container
	Client                         *http.Client
	ToolBar                        *widget.Toolbar
	PriceChartContainer            *fyne.Container
	DB                             repository.Repository
	Holdings                       [][]any
	HoldingsTable                  *widget.Table
	AddHoldingsPurchaseAmountEntry *widget.Entry
	AddHoldingsPurchaseDateEntry   *widget.Entry
	AddHoldingsPurchasePriceEntry  *widget.Entry
}

var myApp Config

func main() {
	// app
	fyneApp := app.NewWithID("rey.com.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.Client = &http.Client{}

	// logger
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// connection to database
	sqlDB, err := myApp.connectSQL()
	if err != nil {
		log.Panicln(err)
	}

	// create table
	myApp.setupDB(sqlDB)

	// window
	myApp.MainWindow = fyneApp.NewWindow("Gold Watcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	//myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	// UI
	myApp.makeUI()

	// run app
	myApp.MainWindow.ShowAndRun()
}

func (app *Config) connectSQL() (db *sql.DB, err error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
		app.InfoLog.Println("db in: ", path)
	}

	db, err = sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		log.Println(err)
	}
}
