package html

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	// . "maragu.dev/gomponents/html"
)

type PageProps struct {
	Title       string
	Description string
}

func page(props PageProps, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "en",
		Head:        []Node{},
		Body:        []Node{Group(children)},
	})
}
