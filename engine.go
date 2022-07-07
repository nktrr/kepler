package main

import (
	"log"
	"net/http"
	"sync"
)

type Engine struct {
	Window AppWindow
}

func createEngine() Engine {
	var wg sync.WaitGroup
	var engine = Engine{}
	wg.Add(1)
	go createApplication(&engine, &wg)
	wg.Wait()
	println("start app")
	startApplication(engine.Window.Application)
	println("app started")
	setupSearchCallback(engine)
	return engine
}

func startEngine(engine Engine) {

}

func setupSearchCallback(engine Engine) {
	println("setup callback")

	setupSearch(engine)
}

func request(url string) {
	println("request")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		println(resp.Header)
	}
}
