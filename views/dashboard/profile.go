package dashboard

import (
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/components"
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Profile(currentPath string, hankoApiUrl string, user *models.User, errorMsg string) Node {
	props := layout.DashboardBaseProps{
		Title:       "Profile Settings",
		Description: "Manage your account settings",
		CurrentPath: currentPath,
		HankoApiUrl: hankoApiUrl,
		User:        user,
		OptionalScripts: []Node{
			Script(Type("module"),
				Rawf(`
import { register } from "https://esm.run/@teamhanko/hanko-elements";

await register(%s);
`, hankoApiUrl)),
		},
	}

	return layout.DashboardBase(props,
		Div(Class("max-w-2xl mx-auto space-y-8 divide-y divide-gray-200"),
			If(errorMsg != "",
				Div(Class("p-4 mb-4 text-red-800 rounded-lg bg-red-50"),
					Text(errorMsg),
				),
			),
			components.StyledForm(components.StyledFormProps{
				Action:       "/dashboard/profile",
				CancelButton: []Node{Text("")},
				SubmitButton: []Node{Text("Save")},
			},
				"post",

				components.StyledInput(components.StyledInputProps{
					Type:     "text",
					Name:     "name",
					ID:       "name",
					Required: true,
					Label:    "Full Name",
					Value:    user.Name,
					Class:    "w-full",
				}),
			),

			Raw("<hanko-profile />"),
		),
	)
}
