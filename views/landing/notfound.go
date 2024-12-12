package landing

import (
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func NotFound(isLoggedIn bool) Node {
	props := layout.LandingPageProps{
		Title:            "200.sh - Not found",
		Description:      "200.sh - The page you tried to access could not be found",
		IsLoggedIn:       isLoggedIn,
		ShowActionButton: true,
	}

	return layout.LandingPage(props,
		Div(Class("flex flex-col justify-center items-center lg:mt-20"),
			H1(Class("text-xl font-bold"), Text("404 - Not Found 😢")),
			P(Text("The page you tried to access could not be found")),
		),
	)
}
