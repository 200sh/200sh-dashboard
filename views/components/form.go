package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type StyledInputProps struct {
	Type         string
	Name         string
	ID           string
	AutoComplete string
	Placeholder  string
	Required     bool
	Label        string
	Value        string
}

func StyledInput(p StyledInputProps) Node {
	return Div(
		Label(Class("block text-sm/6 font-medium text-gray-900"), For(p.ID), Text(p.Label)),

		Div(Class("mt-2"),
			Div(Class("flex items-center rounded-md pl-3 bg-white outline outline-1 -outline-offset-1 outline-gray-300 focus-within:outline focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-primary"),
				Input(
					Type(p.Type),
					Name(p.Name),
					ID(p.ID),
					If(p.AutoComplete != "", AutoComplete(p.AutoComplete)),
					Placeholder(p.Placeholder),
					If(p.Required, Required()),
					Value(p.Value),
					Class("block min-w-0 grow w-full rounded-md bg-white pl-1 pr-3 py-1.5 text-base text-gray-900 placeholder:text-gray-400 focus:outline focus:outline-0 sm:text-sm/6"),
				),
			),
		),
	)
}

type PrefixStyledInputProps struct {
	Type         string
	Name         string
	ID           string
	AutoComplete string
	Placeholder  string
	Required     bool
	Label        string
	Prefix       string
}

func PrefixStyledInput(p PrefixStyledInputProps) Node {
	return Div(
		Label(Class("block text-sm/6 font-medium text-gray-900"), For(p.ID), Text(p.Label)),

		Div(Class("mt-2"),
			Div(Class("flex items-center rounded-md pl-3 bg-white outline outline-1 -outline-offset-1 outline-gray-300 focus-within:outline focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-primary"),
				Div(Class("shrink-0 select-none text-base text-gray-500 sm:text-sm/6"), Text(p.Prefix)),
				Input(
					Type(p.Type),
					Name(p.Name),
					ID(p.ID),
					If(p.AutoComplete != "", AutoComplete(p.AutoComplete)),
					Placeholder(p.Placeholder),
					If(p.Required, Required()),
					Class("block min-w-0 grow w-full rounded-md bg-white pl-1 pr-3 py-1.5 text-base text-gray-900 placeholder:text-gray-400 focus:outline focus:outline-0 sm:text-sm/6"),
				),
			),
		),
	)
}

func StyledFormSectionH2(t string) Node {
	return H2(Class("text-base/7 font-semibold text-gray-900 pb-8"), Text(t))
}

func StyledFormSectionDescription(t string) Node {
	return P(Class("mt-1 text-sm/6 text-gray-600 "),
		Text(t),
	)
}

func StyledFormSection(children ...Node) Node {
	return Div(Class("space-y-12"),
		// Form Section
		Div(Class("pb-8"),
			Group(children),
		),
	)
}

type StyledFormProps struct {
	Action       string
	CancelButton []Node
	SubmitButton []Node
}

func StyledForm(p StyledFormProps, method string, children ...Node) Node {
	// Default action buttons
	if p.CancelButton == nil {
		p.CancelButton = []Node{
			Text("Cancel"),
		}
	}

	if p.SubmitButton == nil {
		p.SubmitButton = []Node{
			Text("Submit"),
		}
	}

	return Form(Class("mt-2"),
		Action(p.Action), Method(method),

		Group(children),

		// Action buttons
		Div(Class("mt-6 flex items-center justify-end gap-x-6"),
			Button(
				AdditionalNodes(
					p.CancelButton,
					Type("button"),
					ID("cancel-button"),
					Class("text-sm/6 font-semibold text-gray-900"),
				)...,
			),

			Button(
				AdditionalNodes(
					p.SubmitButton,
					Type("submit"),
					ID("submit-button"),
					Class("rounded-md bg-primary px-3 py-2 text-sm font-semibold hover:bg-primary/60"),
				)...,
			),
		),
	)
}

func AdditionalNodes(nodes []Node, additional ...Node) []Node {
	return append(
		nodes,
		additional...,
	)
}
