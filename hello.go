package main

import "fmt"
import "honnef.co/go/js/dom"

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


// a button sends true when clicked on its output channel
type Button struct {
	*dom.HTMLButtonElement
	out chan bool
}

func main() {
	fmt.Println("Golang frontend")

	ctx := NewContext()

	clickme := ctx.NewButton()
	ctx.Append(clickme)

	span := ctx.NewElement("span")
	ctx.Append(span)
	span.SetTextContent("0")
	
	clickme.SetTextContent("Click Me!")

	// buttons are automatically wired up to output channels
	go func() {
		var i int = 0
		for {
			b := <- clickme.out
			i += 1
			print(b)
			print("button clicked", i)
			span.SetTextContent(fmt.Sprintf("%d", i))
		}
	} ()
	
}
