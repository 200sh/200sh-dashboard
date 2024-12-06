package pages

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"os"
)

type PageTemplate struct {
	pages map[string]*template.Template
}

func LoadPageTemplates() *PageTemplate {
	// Find all pages
	p := map[string]*template.Template{}
	dir, err := os.ReadDir("./templates/pages")
	if err != nil {
		log.Fatalln("Could not find pages directory")
	}

	for _, entry := range dir {
		fmt.Println(entry)
		tl := template.Must(template.ParseGlob("./templates/layouts/*.html"))
		tl = template.Must(tl.ParseGlob("./templates/components/*.html"))
		if !entry.IsDir() {
			p[entry.Name()] = template.Must(tl.ParseFiles("./templates/pages/" + entry.Name()))
		}
	}

	return &PageTemplate{pages: p}
}

func (t *PageTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.pages[name].ExecuteTemplate(w, "base", data)
}
