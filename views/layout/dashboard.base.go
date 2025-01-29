package layout

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/components"
	lucide "github.com/eduardolat/gomponents-lucide"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
	"strings"
)

type DashboardBaseProps struct {
	Title           string
	Description     string
	CurrentPath     string
	HankoApiUrl     string
	User            *models.User
	OptionalScripts []Node
}

func (p DashboardBaseProps) isCurrentPage(page string) bool {
	return p.CurrentPath == page
}

func (p DashboardBaseProps) isCurrentPageOrChild(page string) bool {
	return strings.HasPrefix(p.CurrentPath, page)
}

func DashboardBase(props DashboardBaseProps, children ...Node) Node {
	headNodes := []Node{
		Meta(Charset("UTF-8")),
		Link(Rel("icon"), Href("/static/favicon.ico")),
		Link(Rel("stylesheet"), Href("/static/css/output.css")),
		Script(Defer(), Src("https://cdn.jsdelivr.net/npm/alpinejs@3.14.7/dist/cdn.min.js")),
		Script(Type("module"),
			Rawf(`
    import { register } from 'https://esm.run/@teamhanko/hanko-elements';

    const { hanko } = await register("%s");

    document.getElementById("logout-link")
      .addEventListener("click", (event) => {
          event.preventDefault();
          hanko.user.logout();
      });

    hanko.onUserLoggedOut(() => {
        // successfully logged out, redirect to a page in your application
        document.location.href = "/"
    })`, props.HankoApiUrl)),
	}
	headNodes = append(headNodes, props.OptionalScripts...)

	initials := strings.Fields(props.User.Name)

	return HTML5(HTML5Props{
		Title:       "200.sh - " + props.Title,
		Description: props.Description,
		Language:    "en",
		Head:        headNodes,
		Body: []Node{Class("h-screen bg-cover bg-gradient-to-br from-slate-50 to-emerald-100"), Div(
			// TODO; Off canvas menu for mobile, show/hide base on off-canvas menu state.
			Div(),

			// Static Sidebar for desktop
			Div(Class("hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col"),
				// Sidebar component swap this element with another sidebar if needed
				Div(Class("flex grow flex-col gap-y-5 overflow-y-auto px-6 pb-4"),
					// Header logo top left
					Div(Class("flex h-16 shrink-0 items-center justify-center lg:mt-8 lg:mb-3"),
						A(Href("/dashboard"),
							Img(Class("block h-16"), Src("/static/200sh-logo.svg"), Alt("200.sh logo")),
						),
					),

					// Sidebar menu items
					Nav(Class("flex flex-1 flex-col"),
						Ul(Class("flex flex-1 flex-col gap-y-7"), Role("list"),
							Li(
								Ul(Class("-mx-2 space-y-1"), Role("list"),
									// Menu List Items
									// 1. Dashboard
									sidebarItem(
										"Dashboard",
										"/dashboard",
										lucide.House,
										props.isCurrentPage("/dashboard"),
									),

									// 2. Monitors
									sidebarItem(
										"Monitors",
										"/dashboard/monitors",
										lucide.MonitorUp,
										props.isCurrentPageOrChild("/dashboard/monitors"),
									),
								),

								// Optional Settings button at bottom if needed later
								//Li(Class("mx-auto"),
								//lucide.Cog(),
								//	A(Href("/dashboard/settings"),
								//		Class("group -mx-2 flex gap-x-3 rounded-md text-sm/6 font-semibold text-gray-700 hover:border-gray-50"),
								//		Text("Settings"),
								//		),
								//	),
							),
						),
					),
				),
			),

			// Content with header
			Div(Class("lg:pl-72"),
				// Navbar
				Nav(Class("hidden lg:flex flex-row items-center justify-between p-5 lg:px-10 lg:py-5"),
					H1(Class("text-2xl font-semibold text-gray-700 leading-tight"),
						Text(props.Title),
					),
					Div(Class("flex items-center gap-x-4"),
						// TODO: Add Contact Us button here
						components.DropDown(
							Div(Class("flex items-center gap-x-2 cursor-pointer rounded-full"),
								Div(
									Img(Class("h-10 w-10 rounded-full object-cover"),
										Alt(props.User.Name),

										Src(fmt.Sprintf("https://ui-avatars.com/api/?name=%s+%s&color=223D30&background=9ACD32", initials[0], initials[1])),
									),
								),
								Div(Class("flex-col hidden lg:flex"),
									Div(Class("text-sm font-semibold text-gray-700"),
										Text(props.User.Name),
									),
									Div(Class("text-xs text-gray-700"),
										Text(props.User.Email),
									),
								),
							),

							// Dropdown items
							A(Class("w-full flex items-center justify-center gap-x-2 hover:bg-primary/60 text-sm rounded-xl p-2"),
								Href("/dashboard/profile"),
								lucide.User(),
								Text("Profile"),
							),

							A(Class("w-full flex items-center justify-center gap-x-2 hover:bg-primary/60 text-sm rounded-xl p-2"),
								Href("/dashboard/billing"),
								lucide.CreditCard(),
								Text("Billing"),
							),

							// Separator
							Br(),
							// Logout
							Button(Class("w-full flex items-center justify-center gap-x-2 hover:bg-primary/60 text-sm rounded-xl p-2"), ID("logout-link"),
								lucide.LogOut(),
								Text("Log Out"),
							),
						),
					),
				),

				// Main content
				Main(Class("py-10"),
					Div(Class("px-4 sm:px-6 lg:px-8"),
						Group(children),
					),
				),
			),
		)},
	})
}

func sidebarItem(name string, href string, icon func(children ...Node) Node, active bool) Node {
	return Li(
		A(Href(href),
			Classes{
				"group flex justify-start items-center gap-x-4 rounded-xl p-2 text-sm/6 font-semibold": true,
				"text-gray-700 bg-primary hover:bg-primary/60":                                         active,
				"text-gray-700 hover:bg-primary/60":                                                    !active,
			},
			icon(Classes{
				"size-8 rounded-xl p-1.5 bg-slate-50": true,
			}),
			Text(name),
		),
	)
}
