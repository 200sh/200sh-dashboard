package landing

import (
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(isLoggedIn bool) Node {
	props := layout.LandingPageProps{
		Title:            "200.sh - Global view of your applications uptime",
		Description:      "200.sh - Global view of your application's uptime. Be safe knowing your services are accessible globally",
		IsLoggedIn:       isLoggedIn,
		ShowActionButton: true,
	}

	return layout.LandingPage(props,
		H1(Text("Hello world")),
	)
}
