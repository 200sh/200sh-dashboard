package auth

import (
	"github.com/200sh/200sh-dashboard/views/layout"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Login(hankoApiUrl string) Node {
	props := layout.LandingPageProps{
		Title:            "200.sh - Login",
		Description:      "200.sh - Login Page to access the uptime monitor",
		IsLoggedIn:       false,
		ShowActionButton: false,
		OptionalScripts: []Node{
			Script(Type("module"),
				Rawf(`
import {register} from "https://esm.run/@teamhanko/hanko-elements";

const {hanko} = await register("%s");

hanko.onSessionCreated(() => {
	// successfully logged in, redirect to a page in your application
	document.location.href = "/dashboard";
});`, hankoApiUrl),
			),
		},
	}

	return layout.LandingPage(props,
		Div(Class("flex items-center justify-center lg-mt-5 lg:mt-20"),
			Raw(`<hanko-auth />`),
		),
	)
}
