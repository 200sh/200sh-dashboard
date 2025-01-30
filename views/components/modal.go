package components

import (
	. "maragu.dev/gomponents/html"
)

func DeleteConfirmationModal(extraClasses Class, content Node, dataAttributes ...Attribute) Node {
	return Div(
		Class("relative z-10 hidden"),
		DataAttr("dialog", ""),
		Div(
			Class("fixed inset-0 bg-black/25 transition-opacity"),
		),
		Div(
			Class("fixed inset-0 z-10 w-screen overflow-y-auto"),
			Div(
				Class("flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"),
				Div(
					Class("relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg"),
					Div(
						Class("bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4"),
						Div(
							Class("flex items-center justify-center mx-auto h-12 w-12 bg-red-100 rounded-full"),
							Children{
								Text("!"),
							},
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
						Class("bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6"),
						Button(
							Type("button"),
							Class("inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto"),
							ID("confirm-delete-button"),
							Text("Delete"),
						),
						Button(
							Type("button"),
							Class("mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"),
							ID("cancel-delete-button"),
							DataAttr("dialog-close", ""),
							Text("Cancel"),
						),
					),
				),
			),
		),
	)
}
