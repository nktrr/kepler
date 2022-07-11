package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type AppWindow struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow
	SearchBar         *gtk.Entry
	Box               *gtk.Box
}

func createApplication(window AppWindow) {
	const appID = "org.gtk.example"
	gtk.Init(nil)
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application.", err)
	}
	window.Application = application
	setupWindow(&window)
	setupBox(&window, gtk.ORIENTATION_VERTICAL)
	setupSearch(&window)
	gtk.Main()
}

func setupWindow(window *AppWindow) {
	win, err := gtk.ApplicationWindowNew(window.Application)
	if err != nil {
		println("Could not create application window")
	}
	win.SetTitle("KEPLER")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(1000, 800)
	window.ApplicationWindow = win
}

func setupBox(window *AppWindow, orient gtk.Orientation) {
	box, err := gtk.BoxNew(orient, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	window.ApplicationWindow.Add(box)
	window.Box = box
}

func setupSearch(window *AppWindow) {
	println("setup search")
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create search", err)
	}
	window.Box.Add(entry)
	window.SearchBar = entry
	window.SearchBar.Connect("activate", func() {
		println(window.SearchBar.GetText())
		url, _ := window.SearchBar.GetText()
		request(url)
	})
	window.ApplicationWindow.ShowAll()
}
