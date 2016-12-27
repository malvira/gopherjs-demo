package main

import "fmt"
//import "github.com/gopherjs/gopherjs/js"
//import "github.com/gopherjs/jquery"
import "honnef.co/go/js/dom"

//convenience:
//var jQuery = jquery.NewJQuery

type Context struct {
	doc dom.Document
	root dom.Element
}

func (ctx Context) NewElement (tag string) dom.Element {
	e := ctx.doc.CreateElement(tag)
	return e
}


// a button
// sends true on down
// false on up
type Button struct {
	*dom.HTMLButtonElement
	out chan bool
}

// a box that changes color

func main() {
	fmt.Println("blahadfadsf")

	var ctx Context
	ctx.doc = dom.GetWindow().Document()
	ctx.root = ctx.doc.DocumentElement()

	butt := make(chan bool)
	el := ctx.NewElement("button")
	clickme := Button{el.(*dom.HTMLButtonElement), butt}

	ctx.root.AppendChild(clickme)
	clickme.SetTextContent("Click Me!")
	clickme.AddEventListener("click", false,  func (e dom.Event) {
		go func() { clickme.out <- true }()
	})

	go func() {
		for {
			b := <- butt
			print(b)
			print("button clicked")
		}
	} ()
	
}
