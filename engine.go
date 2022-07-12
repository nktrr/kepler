package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func request(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	println(resp.Status)
	return resp
}

//url for test:  http://cdda.chezzo.com/

func parse(response *http.Response) *html.Node {
	doc, err := html.Parse(response.Body)
	if err != nil {
		log.Fatal("Cannot parse html.", err)
		return nil
	}
	//d := 0
	//var f func(*html.Node)
	//f = func(n *html.Node) {
	//	d++
	//	if n.Type == html.TextNode {
	//		//s := strings.Repeat(" ", d)
	//		//print(s, n.Data)
	//		println(n.Data)
	//		//print("   ")
	//		//fmt.Printf("%v", n.Attr)
	//		//println()
	//	}
	//	for c := n.FirstChild; c != nil; c = c.NextSibling {
	//		f(c)
	//	}
	//	d--
	//}
	//f(doc)
	return doc
}
