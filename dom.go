package main

import "golang.org/x/net/html"

type CNode struct {
	Node  *html.Node
	Class string
}

func parseNodesToCNode(node *html.Node) CNode {
	cnode := CNode{}

	for c := cnode.Node.FirstChild; c != nil; c = c.NextSibling {

	}
	return cnode
}
