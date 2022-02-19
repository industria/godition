package splitter // github.com/industria/godition/splitter"

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/industria/godition/dredition"
	"golang.org/x/net/html"
)

type deckSlices struct {
	start int
	end   int
}

type Deck struct {
	Number int
	HTML   string
}

func Split(r io.Reader, notification dredition.Notification) (*[]Deck, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	body, err := body(doc)
	if err != nil {
		return nil, err
	}

	var decks []deckSlices

	rootElement := body.FirstChild
	rootElementClass := classValue(rootElement)
	groups := children(rootElement)
	start := 0
	for i, g := range groups {
		if contains(g, splitter) {
			di := deckSlices{start: start, end: i - 1}
			decks = append(decks, di)
			start = i + 1
		}
	}
	di := deckSlices{start: start, end: len(groups) - 1}
	if di.start < len(groups) && di.end < len(groups) {
		decks = append(decks, di)
	}

	var ds []Deck
	for i, d := range decks {
		deckGroups := groups[d.start : d.end+1]
		html := create(1, deckGroups, rootElementClass, notification)
		deck := Deck{Number: i + 1, HTML: render(html)}
		ds = append(ds, deck)
	}
	return &ds, nil

	// Groups
	// 00 - 03 : Deck 1
	// 04 - 04 : Split
	// 05 - 07 : Deck 2
	// 08 - 08 : Split
	// 09 - 09 : Leder (incl deck 2)
	// 10 - 15 : Deck 3
	// 16 - 16 : Split
	// 17 - 22 : Deck 4
	// 23 - 23 : Split
	// 24 - 29 : Deck 5
	// 30 - 30 : Split
	// 31 - 38 : Deck 6
	// 39 - 39 : Split
	// 40 - 44 : Deck 7
	// 45 - 45 : Split
	// 46 - 53 : Deck 8
	// 54 - 54 : Split
	// 55 - 59 : deck 9
	// 60 - 60 : Split
	// 61 - 74 : Deck 10

}

func create(deckNumber int, groups []*html.Node, class string, n dredition.Notification) *html.Node {
	deck := &html.Node{
		Data: "div",
		Type: html.ElementNode,
	}
	id := fmt.Sprintf("front-%s-%s-%d", n.Data.Product.Name, n.Data.Edition.Name, deckNumber)
	deck.Attr = append(deck.Attr, html.Attribute{Key: "id", Val: id})
	deck.Attr = append(deck.Attr, html.Attribute{Key: "data-product", Val: n.Data.Product.Id})
	deck.Attr = append(deck.Attr, html.Attribute{Key: "data-productname", Val: n.Data.Product.Name})
	deck.Attr = append(deck.Attr, html.Attribute{Key: "data-edition", Val: n.Data.Edition.Id})
	deck.Attr = append(deck.Attr, html.Attribute{Key: "data-editionname", Val: n.Data.Edition.Name})
	deck.Attr = append(deck.Attr, html.Attribute{Key: "data-decknumber", Val: strconv.FormatInt(int64(deckNumber), 10)})
	deck.Attr = append(deck.Attr, html.Attribute{Key: "class", Val: class})

	// Clone with out parent and siblings which can not be appended
	for _, g := range groups {
		clone := &html.Node{
			Data:       g.Data,
			Attr:       g.Attr,
			Type:       g.Type,
			DataAtom:   g.DataAtom,
			Namespace:  g.Namespace,
			FirstChild: g.FirstChild,
			LastChild:  g.LastChild,
		}

		deck.AppendChild(clone)
	}

	return deck
}

func children(node *html.Node) []*html.Node {
	var children []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, child)
	}
	return children
}

func classValue(node *html.Node) string {
	for _, a := range node.Attr {
		if a.Key == "class" {
			return a.Val
		}
	}
	return ""
}

func splitter(node *html.Node) bool {
	if node.Type == html.ElementNode {
		for _, a := range node.Attr {
			if a.Key == "data-placeholder" && a.Val == "split" {
				return true
			}
		}
	}
	return false
}

func contains(node *html.Node, predicate func(*html.Node) bool) bool {
	match := false
	var f func(*html.Node)
	f = func(node *html.Node) {
		if predicate(node) {
			match = true
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(node)
	return match
}

func body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			f(child)
		}
	}
	f(doc)
	if body == nil {
		return nil, errors.New("unable to finde <body> in node tree")
	}
	return body, nil
}

func render(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
