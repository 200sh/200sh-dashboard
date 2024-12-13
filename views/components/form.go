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
}

func StyledInput(p StyledInputProps) Node {
	return Input(
		Type(p.Type),
		Name(p.Name),
		ID(p.ID),
		AutoComplete(p.AutoComplete),
		Placeholder(p.Placeholder),
		Class("block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-primary sm:text-sm/6"),
	)
}
