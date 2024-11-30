package html

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Home() g.Node {
	return html.Div(
		g.Text("Hello World!!!"),
	)
}
