package components

import (
	x "github.com/glsubri/gomponents-alpine"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DropDown(button Node, panelEls ...Node) Node {
	return Div(
		x.Data(`{
open: false,
toggle() {
	if (this.open) {
		return this.close()
	}

	this.$refs.button.focus()

	this.open = true
},
close(focusAfter) {
	if (!this.open) return
	
	this.open = false

	focusAfter && focusAfter.focus()
}
}`),
		x.On("keydown.escape.prevent.stop", "close($refs.button)"),
		x.On("focusin.window", "! $refs.panel.contains($event.target) && close()"),
		x.Id("['dropdown-button']"),
		Class("relative"),

		// Button
		Button(
			x.Ref("button"),
			x.On("click", "toggle()"),
			Attr(":aria-expanded", "open"),
			Attr(":aria-controls", "$id('dropdown-button')"),
			Type("button"),
			Class("relative flex items-center whitespace-nowrap justify-center gap-2 px-3 py-2 rounded-full hover:bg-gray-100 text-gray-800"),

			button,
		),

		// Panel
		Div(
			x.Ref("panel"),
			x.Show("open"),
			x.Transition("origin.top.left"),
			x.On("click.outside", "close($refs.button)"),
			Attr(":id", "$id('dropdown-button')"),
			x.Cloak(),
			Class("absolute left-0 gap-y-2 min-w-48 rounded-lg shadow-sm mt-2 z-10 origin-top-left  p-1.5 outline-none bg-white border border-gray-200"),

			Group(panelEls),
		),
	)
}
