package auth

import (
	"github.com/200sh/200sh-dashboard/views/components"
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func UserSetup() Node {
	props := layout.LandingPageProps{
		Title:            "200.sh - Setup Your Account",
		Description:      "200.sh - The setup form for creating your account",
		IsLoggedIn:       true,
		ShowActionButton: false,
	}

	return layout.LandingPage(props,
		Form(Class("mt-20"),
			Action("/user/setup"), Method("post"),
			Div(Class("space-y-12"),
				// Form Section
				Div(Class("border-b border-gray-900/10 pb-12"),
					// Title for section
					H2(Class("text-base/7 font-semibold text-gray-900"), Text("Personal Information")),

					// Description for the form section
					P(Class("mt-1 text-sm/6 text-gray-600"),
						Text("Personal Information for your personal user account"),
					),

					// Form label and inputs
					Div(Class("mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6"),
						Div(Class("sm:col-span-3"),
							Label(For("first-name"),
								Class("block text-sm/6 font-medium text-gray-900"),
								Text("First Name")),
							Div(Class("mt-2"),
								components.StyledInput(components.StyledInputProps{
									Type:         "text",
									Name:         "first-name",
									ID:           "first-name",
									AutoComplete: "given-name",
								}),
							),
						),
						Div(Class("sm:col-span-3"),
							Label(For("last-name"),
								Class("block text-sm/6 font-medium text-gray-900"),
								Text("Last Name"),
							),

							Div(Class("mt-2"),
								components.StyledInput(components.StyledInputProps{
									Type:         "text",
									Name:         "last-name",
									ID:           "last-name",
									AutoComplete: "family-name",
								}),
							),
						),
					),
				),
			),

			// Action buttons
			Div(Class("mt-6 flex items-center justify-end gap-x-6"),
				Button(Type("button"), Class("text-sm/6 font-semibold text-gray-900"),
					Text("Cancel")),

				Button(Type("submit"), Class("rounded-md bg-primary px-3 py-2 text-sm font-semibold"),
					Text("Save")),
			),
		),
	)
}
