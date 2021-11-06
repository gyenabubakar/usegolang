package views

import (
	"html/template"
	"net/http"
)

func NewView(files ...string) *View {
	files = append(files, "views/components/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{Template: t}
}

func (v *View) RenderWith(w http.ResponseWriter, d interface{}, fn func(error)) {
	err := v.Template.Execute(w, d)

	if err != nil {
		if fn != nil {
			fn(err)
			return
		}

		panic(err)
	}
}

type View struct {
	Template *template.Template
}
