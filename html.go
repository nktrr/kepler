package main

const (
	HTML  = "html"
	head  = "head"
	title = "title"
	body  = "body"
	div   = "div"
	li    = "li"
	img   = "img"
	a     = "a"
	span  = "span"
	h1    = "h1"
	p     = "p"
	h2    = "h2"
)

const (
	name = "bebra"
)

func isSupported(tag string) bool {
	if getType(tag) != "" {
		return true
	}
	return false
}

func getType(tag string) string {
	switch tag {
	case "html":
		return HTML
	case "head":
		return head
	case "title:":
		return title
	case "body":
		return body
	case "div":
		return div
	case "li":
		return li
	case "img":
		return img
	case a:
		return a
	case "span":
		return span
	case "h1":
		return h1
	case "p":
		return p
	case "h2":
		return h2
	}
	return ""
}
