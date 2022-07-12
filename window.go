package main

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"golang.org/x/net/html"
	"log"
)

type AppWindow struct {
	Application       *gtk.Application
	ApplicationWindow *gtk.ApplicationWindow
	SearchBar         *gtk.Entry
	Box               *gtk.Box
	CurrentPage       *gtk.Box
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
	setupPage(&window)
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
		render(window, parse(request(url)))
	})
	window.ApplicationWindow.ShowAll()
}

func setupPage(window *AppWindow) {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create page box:", err)
	}
	window.Box.Add(box)
	window.CurrentPage = box
	window.ApplicationWindow.ShowAll()
}

func render(window *AppWindow, rootNode *html.Node) {
	var currentBox *gtk.Box
	currentBox = window.CurrentPage
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if isSupported(node.Data) {
			switch getType(node.Data) {
			case p:
				addText(currentBox, node)
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(rootNode)
	window.ApplicationWindow.ShowAll()
}

func addText(box *gtk.Box, node *html.Node) {
	if node.FirstChild != nil {
		println("-------------------")
		println(node.FirstChild.Data)
		fmt.Printf("%v", node.Attr)
		println()
		println("-------------------")
		label, _ := gtk.LabelNew(node.FirstChild.Data)
		box.Add(label)
	}

}
