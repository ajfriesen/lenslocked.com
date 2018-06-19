package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is the directory where all layout files will be palced
	LayoutDir string = "views/layouts/"

	// TemplateExt is the extension the layout files will have
	TemplateExt string = ".gohtml"
)

// layoutFiles will get all files in LayoutDir with the extension TemplateExt
// Then it is returning all files
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// NewView get's all layout files and a number of files and stores it in "files"
// Then it will everything will be parsed as template and stored
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
