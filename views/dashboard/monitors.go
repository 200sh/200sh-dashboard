package dashboard

import (
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Monitor(currentPath string, hankoApiUrl string, user *models.User) Node {
	props := layout.DashboardBaseProps{
		Title:           "Monitors",
		Description:     "All uptime monitors",
		CurrentPath:     currentPath,
		HankoApiUrl:     hankoApiUrl,
		User:            user,
		OptionalScripts: nil,
	}

	return layout.DashboardBase(props,
		H1(Text("Hello from Monitors")),
	)
}
