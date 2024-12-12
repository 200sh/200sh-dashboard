package layout

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

type LandingPageProps struct {
	Title            string
	Description      string
	IsLoggedIn       bool
	ShowActionButton bool
	OptionalScripts  []Node
}

func LandingPage(props LandingPageProps, children ...Node) Node {
	headNodes := []Node{
		Meta(Charset("UTF-8")),
		Link(Rel("icon"), Href("/static/favicon.ico")),
		Link(Rel("stylesheet"), Href("/static/css/output.css")),
	}
	headNodes = append(headNodes, props.OptionalScripts...)

	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,
		Language:    "en",
		Head:        headNodes,
		Body: []Node{Class("min-h-screen h-screen flex flex-col"),
			header(props.IsLoggedIn, props.ShowActionButton),
			Main(Class("flex flex-col items-center justify-center max-w-full font-roboto"),
				Group(children),
			),
		},
	})
}

func header(isLoggedIn bool, showActionButton bool) Node {
	return Header(Class("bg-white"),
		Nav(Class("mx-auto flex max-w-7xl items-center justify-center gap-x-6 p-6 lg:px-16 pt-6 lg:pt-16"),
			Div(Class("flex lg:flex-1"),
				A(Href("/"),
					Img(Class("block w-14 h-14"), Src("/200sh-logo.svg")),
				),
			),
			Div(Class("flex flex-1 items-center justify-end gap-x-6 text-lg"),
				If(isLoggedIn && showActionButton,
					A(Href("/dashboard"),
						Class("rounded-md px-3.5 py-2.5 text-black transition hover:text-black/70"),
						Text("Dashboard "), Span(Class("text-primary"), Aria("hidden", "true"), Raw(`&rarr;`)),
					),
				),
				If(!isLoggedIn && showActionButton,
					A(Href("/login"), Class("rounded-md bg-primary text-sm px-3.5 py-2.5 font-semibold"),
						Text("Get Started")),
				),
			),
		),
	)
}
