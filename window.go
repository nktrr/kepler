package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"sync"
)

type AppWindow struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow
	SearchBar         *gtk.Entry
	Box               *gtk.Box
}

func createApplication(engine *Engine, wg *sync.WaitGroup) {
	const appID = "org.gtk.example"
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application.", err)
	}

	application.Connect("activate", func() {
		appWindow, err := gtk.ApplicationWindowNew(application)
		if err != nil {
			log.Fatal("Could not create application window.", err)
		}
		appWindow.SetTitle("Basic Application.")
		appWindow.SetDefaultSize(1000, 800)
		box := setupBox(gtk.ORIENTATION_VERTICAL)
		appWindow.Add(box)
		engine.Window.Box = box
		appWindow.Show()
	})
	engine.Window.Application = application
	wg.Done()
}

func startApplication(application *gtk.Application) {
	application.Run(os.Args)
}

func setupBox(orient gtk.Orientation) *gtk.Box {
	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	return box
}

func setupSearch(engine Engine) {
	println("setup search")
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create search", err)
	}
	engine.Window.Box.Add(entry)
	engine.Window.SearchBar = entry
	engine.Window.SearchBar.Connect("activate", func() {
		println("searchbarclicked")
	})
	engine.Window.ApplicationWindow.ShowAll()
}
