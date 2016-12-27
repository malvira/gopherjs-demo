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

func NewContext() Context {
	var ctx Context
	ctx.doc = dom.GetWindow().Document()
	ctx.root = ctx.doc.DocumentElement()
	return ctx
}

func (ctx Context) Append(e dom.Element) {
	ctx.root.AppendChild(e)
}

func (ctx Context) NewElement (tag string) dom.Element {
	e := ctx.doc.CreateElement(tag)
	return e
}

func (ctx Context) NewButton() Button {
	o := make(chan bool)
	el := ctx.NewElement("button")
	button := Button{el.(*dom.HTMLButtonElement), o}
	button.AddEventListener("click", false,  func (e dom.Event) {
		go func() { button.out <- true }()
	})
	return button
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
	fmt.Println("Golang frontend")

	// should make this an idiomatic initialization.
	ctx := NewContext()

	clickme := ctx.NewButton()
	ctx.Append(clickme)
	
	clickme.SetTextContent("Click Me!")

	// buttons are automatically wired up to output channels
	go func() {
		for {
			b := <- clickme.out
			print(b)
			print("button clicked")
		}
	} ()
	
}
