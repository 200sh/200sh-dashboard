package components

import (
	x "github.com/glsubri/gomponents-alpine"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func DeleteConfirmationModal(extraClasses string, content Node) Node {
	return Div(
		Class("relative z-10"),
		x.Show("showDeleteModal"), // This controls visibility
		Div(
			// Backdrop
			Class("fixed inset-0 bg-black/25"),
			x.On("click", "showDeleteModal = false"),
		),
		Div(
			// Modal panel
			Class("fixed inset-0 z-10 w-screen overflow-y-auto"),
			Div(
				Class("flex min-h-full items-center justify-center p-4 text-center"),
				Div(
					Class("relative transform overflow-hidden rounded-lg text-left sm:w-full sm:max-w-lg "+FrostedBg),
					Div(
						Class("px-4 pb-4 pt-5 sm:p-6 sm:pb-4"),
						Div(
							Class("flex items-center justify-center mx-auto h-12 w-12 bg-red-100 rounded-full"),
							content,
						),
						Div(
							Class("mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left"),
							H3(
								Class("text-base/7 font-semibold text-gray-900"),
								Text("Delete Monitor"),
							),
							Div(
								Class("mt-2"),
								Text("Are you sure you want to delete this monitor? This action cannot be undone."),
							),
						),
					),
					Div(
						Class("px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6"),
						Button(
							Type("button"),
							Class("inline-flex w-full justify-center rounded-md bg-red-600/90 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500/90 sm:ml-3 sm:w-auto"),
							ID("confirm-delete-button"),
							Text("Delete"),
						),
						Button(
							Type("button"),
							Class("mt-3 inline-flex w-full justify-center rounded-md bg-white/90 px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50/90 sm:mt-0 sm:w-auto"),
							ID("cancel-delete-button"),
							x.On("click", "showDeleteModal = false"),
							Text("Cancel"),
						),
					),
				),
			),
		),
	)
}
