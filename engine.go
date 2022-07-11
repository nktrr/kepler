package main

import (
	"log"
	"net/http"
)

func request(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}
