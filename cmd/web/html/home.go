package html

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home() Node {
	return page(PageProps{},
		Form(Action("/send_page"), Method("post"),
			P(
				Input(ID("radio_html"), Type("radio"), Name("page-type"), Value("html")), Text("html"),
				Input(ID("radio-link"), Type("radio"), Name("page-type"), Value("link")), Text("link"),
			),
			P(
				Textarea(Name("page"), Rows("20"), Cols("100")),
			),
			P(
				Input(Type("submit"), Value("send")),
			),
		))
}

func HomeWith(index int, message bool) Node {
	return page(PageProps{},
		Form(Action("/send_page"), Method("post"),
			P(
				Input(ID("radio_html"), Type("radio"), Name("page-type"), Value("html")), Text("html"),
				Input(ID("radio-link"), Type("radio"), Name("page-type"), Value("link")), Text("link"),
			),
			P(
				Textarea(Name("page"), Rows("20"), Cols("100")),
			),
			P(
				Input(Type("submit"), Value("send")),
			),
		),
		Form(Action(fmt.Sprintf("/gif/%d", index)), Method("post"),
			P(
				If(message, Text("not ready")),
				Input(Type("submit"), Value("download gif")),
			),
		),
	)
}
