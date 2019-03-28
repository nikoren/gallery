package views

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"path/filepath"
)

// NewView - creates new view by searching for all layouts in layoutsDir , combining them
// with the provided  contentTemplates , parses the combined list and setting them as
// ContentTemplate, also you should pick the name of the template that will be wrapping the
// ContentTemplate by providing layoutName
func NewView(layoutName string, layoutsDir string,  contentTemplates ...string) *View {

	// find all layouts
		foundLayouts, err := filepath.Glob(
			fmt.Sprintf("%s/*.gohtml",layoutsDir)) // (matches []string, err error)
		if err != nil{
			log.Fatalf("Couldn't glob %s, ERROR: %v", layoutsDir, err.Error())
		}

	// aggregate all templates
	contentTemplates = append(contentTemplates, foundLayouts...)

	// parse all templates
	parsedContentTemplates, err := template.ParseFiles(contentTemplates...)
	if err != nil{
		log.Panicf("Couldn't parse templates, ERROR: %s", err )
	}

	// create view and set the main template to render(LayoutName)
	v := View{
		LayoutName:      layoutName,
		ContentTemplate: parsedContentTemplates,
	}
	return &v
}

type View struct {
	LayoutName      string
	ContentTemplate *template.Template
}

func(v *View)Render(w io.Writer, data interface{}){

	err := v.ContentTemplate.ExecuteTemplate(w, v.LayoutName, data)
	if err != nil{
		log.Panicf("Couldn't execute template correctly, ERROR: %s", err.Error())
	}
}

