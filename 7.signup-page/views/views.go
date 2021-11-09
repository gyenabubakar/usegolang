package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	Dir           = "views/"
	ComponentsDir = Dir + "components/"
	ComponentsExt = ".gohtml"
)

func NewView(layout string, files ...string) *View {
	for i, f := range files {
		files[i] = attachFullPathParts(f)
	}
	files = append(files, getComponents()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, d interface{}, fn func(error)) {
	w.Header().Set("Content-Type", "text/html")
	if err := v.Template.ExecuteTemplate(w, v.Layout, d); err != nil {
		if fn != nil {
			fn(err)
			return
		}
		panic(err)
	}
}

func (v *View) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	v.Render(w, nil, func(err error) {
		panic(err)
	})
}

type View struct {
	Template *template.Template
	Layout   string
}

func getComponents() []string {
	files, err := filepath.Glob(ComponentsDir + "*" + ComponentsExt)
	if err != nil {
		panic(err)
	}
	return files
}

func attachFullPathParts(sp string) string {
	return Dir + sp + ComponentsExt
}