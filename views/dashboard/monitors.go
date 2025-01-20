package dashboard

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/components"
	"github.com/200sh/200sh-dashboard/views/layout"
	lucide "github.com/eduardolat/gomponents-lucide"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Monitor(currentPath string, hankoApiUrl string, user *models.User, monitors []models.Monitor) Node {
	props := layout.DashboardBaseProps{
		Title:           "Monitors",
		Description:     "All uptime monitors",
		CurrentPath:     currentPath,
		HankoApiUrl:     hankoApiUrl,
		User:            user,
		OptionalScripts: nil,
	}

	var cardComp Node
	if len(monitors) < 1 {
		cardComp = NoMonitor()
	} else {
		cardComp = ListMonitors(monitors)
	}

	return layout.DashboardBase(props,
		Div(Class("flex justify-center items-center"),
			Div(Class("w-fit"),
				components.Card(cardComp),
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

func ListMonitors(monitors []models.Monitor) Node {
	return Div(Class("min-w-96"),
		// Heading
		Div(Class("-ml-4 -mt-2 flex flex-wrap items-center justify-between sm:flex-nowrap"),
			Div(Class("ml-4 mt-2"),
				H3(Class("text-base font-semibold text-gray-900"), Text("")),
			),
			Div(Class("ml-4 mt-2 shrink-0"),
				A(Class("flex flex-row items-center rounded-md bg-primary hover:bg-primary/60 px-3 py-2 text-sm font-semibold shadow-sm"),
					Href("/dashboard/monitors/new"),
					lucide.Plus(Class("-ml-0.5 mr-2 size-4")),
					Text("New Monitor"),
				),
			),
		),

		Ul(
			Role("list"),
			Class("mt-8 divide-y divide-gray-100"),
			Map(monitors, func(m models.Monitor) Node {
				return MonitorListItem(m)
			}),
		),
	)
}

func MonitorListItem(monitor models.Monitor) Node {
	return Li(
		Class("relative flex justify-between gap-x-6 rounded-lg px-4 py-5 hover:bg-primary/60 sm:px-6 lg:px-8"),
		Div(
			Class("flex min-w-0 gap-x-4"),
			//Img(
			//	Class("size-12 flex-none rounded-full bg-gray-50"),
			//	Src(fmt.Sprintf("https://ui-avatars.com/api/?name=%s+%s&color=223D30&background=9ACD32")),
			//	Alt(""),
			//),
			Div(
				Class("min-w-0 flex-auto"),
				P(
					Class("text-sm/6 font-semibold text-gray-900"),
					A(
						Href(fmt.Sprintf("/dashboard/monitors/%d", monitor.Id)),
						// The <span> with absolute classes can just be included inline if needed:
						Span(
							Class("absolute inset-x-0 -top-px bottom-0"),
						),
						Text(monitor.Url),
					),
				),
				//P(
				//	Class("mt-1 flex text-xs/5 text-gray-500"),
				//	A(
				//		Href("mailto:"+user.Email),
				//		Class("relative truncate hover:underline"),
				//		Text(user.Email),
				//	),
				//),
			),
		),
		Div(
			Class("flex shrink-0 items-center gap-x-4"),
			// Hide/show certain text if in "sm" screens or bigger
			//Div(
			//	Class("hidden sm:flex sm:flex-col sm:items-end"),
			//	P(
			//		Class("text-sm/6 text-gray-900"),
			//		Text("test"),
			//	),
			//),
			lucide.ArrowRight(),
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

func ViewMonitor(currentPath string, hankoApiUrl string, user *models.User, monitor *models.Monitor) Node {
	props := layout.DashboardBaseProps{
		Title:           fmt.Sprintf("Monitor - %s", monitor.Url),
		Description:     "",
		CurrentPath:     currentPath,
		HankoApiUrl:     hankoApiUrl,
		User:            user,
		OptionalScripts: []Node{
			Script(Defer(), Src("/static/js/view-monitor.js")),
			Script(Src("https://cdn.jsdelivr.net/npm/apexcharts")),
		},
	}

	return layout.DashboardBase(props,
		Div(Class("flex flex-col items-center"),
			// Header with URL, edit and delete buttons
			Div(Class("flex justify-between items-center w-full p-4 bg-white shadow-md rounded-md"),
				Div(Class("text-lg font-semibold"), Text(monitor.Url)),
				Div(Class("flex gap-2"),
					A(Href(fmt.Sprintf("/dashboard/monitors/%d/edit", monitor.Id)),
						Class("px-4 py-2 bg-blue-500 text-white rounded-md"),
						Text("Edit"),
					),
					Button(
						Class("px-4 py-2 bg-red-500 text-white rounded-md"),
						ID("delete-button"),
						Data("monitor-id", fmt.Sprintf("%d", monitor.Id)),
						Text("Delete"),
					),
				),
			),

			// Latency graph container
			Div(Class("w-full mt-4 p-4 bg-white shadow-md rounded-md"),
				Div(ID("latency-graph"), Class("w-full h-64")),
			),
		),
	)
}
