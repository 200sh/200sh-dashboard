package dashboard

import (
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/components"
	"github.com/200sh/200sh-dashboard/views/layout"
	lucide "github.com/eduardolat/gomponents-lucide"
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
		Div(Class("flex justify-center items-center"),
			Div(Class("w-fit"),
				components.Card(NoMonitor()),
			),
		),
	)
}

func NewMonitor(hankoApiUrl string, user *models.User) Node {
	props := layout.DashboardBaseProps{
		Title:       "Monitors",
		Description: "All uptime monitors",
		CurrentPath: "/dashboard/monitors",
		HankoApiUrl: hankoApiUrl,
		User:        user,
		OptionalScripts: []Node{
			Script(Defer(), Src("/static/js/forms.js")),
			Script(Defer(), Src("/static/js/new-monitor.js")),
		},
	}

	return layout.DashboardBase(props,
		components.Card(
			components.StyledForm(components.StyledFormProps{
				Action:       "/dashboard/monitors/new",
				CancelButton: []Node{Text("Cancel"), Data("link", "/dashboard/monitors")},
				SubmitButton: []Node{Text("Save")},
			}, "post",
				components.StyledFormSection(
					components.StyledFormSectionH2("General"),

					components.StyledInput(components.StyledInputProps{
						Type:  "url",
						Name:  "monitor-url",
						ID:    "monitor-url",
						Label: "Url",
						Value: "https://",
					}),
				),
			),
		),
	)
}

func NoMonitor() Node {
	return Div(Class("flex flex-col justify-center items-center text-center"),
		lucide.MonitorUp(Class("mx-auto size-12 text-gray-400")),
		H3(Class("mt-2 text-sm font-semibold text-gray-900"), Text("No Monitors")),
		P(Class("mt-1 text-sm text-gray-500"), Text("Get started by creating a new monitor")),
		Div(Class("mt-6"),
			A(Class("flex flex-row items-center rounded-md bg-primary hover:bg-primary/60 px-3 py-2 text-sm font-semibold shadow-sm"),
				Href("/dashboard/monitors/new"),
				lucide.Plus(Class("-ml-0.5 mr-2 size-4")),
				Text("New Monitor"),
			),
		),
	)
}
