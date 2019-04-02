package views

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
	"shortcuts"
)


var(
	LayoutsDir string = "views/layouts"
	TemplateExt string  = "gohtml"
)


// NewView - creates new view by searching for all layouts in LayoutsDir , combining them
// with the provided  contentTemplates , parses the combined list and setting them as
// ContentTemplate, layoutName will be the template that will be wrapping the
// ContentTemplate
func NewView(layoutName string, contentTemplates ...string) *View {

	// find all layouts
	globPattern := fmt.Sprintf("%s/*.%s", LayoutsDir, TemplateExt)
	foundLayouts, err := filepath.Glob(globPattern) // (matches []string, err error)
		log.Debugf("Found %v matching %s", foundLayouts, globPattern)
	if err != nil {
		log.Fatalf("Couldn't glob %s, ERROR: %v", LayoutsDir, err.Error())
	}

	// aggregate all templates
	contentTemplates = append(contentTemplates, foundLayouts...)

	// parse all templates
	parsedContentTemplates, err := template.ParseFiles(contentTemplates...)
	if err != nil {
		log.Panicf("Couldn't parse templates, ERROR: %s", err)
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

func (v *View) Render(w io.Writer, data interface{}) error {
	err := v.ContentTemplate.ExecuteTemplate(w, v.LayoutName, data)
	if err != nil {
		log.Panicf("Couldn't execute template correctly, ERROR: %s", err.Error())
		return err
	}
	return nil
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set(shortcuts.ContentTypeTextHtml())
	if err := v.Render(w, nil); err != nil{
		log.Fatalf("Couldn't render view: %s", err.Error())
	}
}
