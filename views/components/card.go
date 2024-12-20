package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

const FrostedBg = "backdrop-blur-lg bg-white/10 shadow-lg ring-1 ring-black/5"

func Card(children ...Node) Node {
	// Find the class in Children
	return Div(Class("block p-6 rounded-xl backdrop-blur-lg bg-white/10 shadow-lg ring-1 ring-black/5"),
		Group(children),
	)
}
