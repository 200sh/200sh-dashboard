package dashboard

import (
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(currentPath string, hankoApiUrl string) Node {
	props := layout.DashboardBaseProps{
		Title:           "Dashboard",
		Description:     "200.sh - A global uptime dashboard",
		CurrentPath:     currentPath,
		HankoApiUrl:     hankoApiUrl,
		OptionalScripts: nil,
	}
	return layout.DashboardBase(props,
		H1(Text("Welcome to the dashboard")),
	)
}
