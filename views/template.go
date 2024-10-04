package views

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path"

	"github.com/LENSLOCKED/context"
	"github.com/LENSLOCKED/models"
	"github.com/gorilla/csrf"
)

type public interface {
	Public()
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	// filepath.Base() нужен для того, чтобы корректно считывал pattern[0],
	// а именно без этого метода путь galleries/new.gohtml читался как new.gohtml,
	// что приводило к ошибке, но filepath.Base() берет на себя все заморочки с
	// доп директориями и также нужно дополнить файл fs.go
	tpl := template.New(path.Base(pattern[0]))

	tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implement")
			},
			"currentUser": func() (template.HTML, error) {
				return "", fmt.Errorf("currentUser not implement")
			},
			"errors": func() []string {
				return nil
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template %v", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
	}
	errMsgs := errMessage(errs...)
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"errors": func() []string {
				return errMsgs
			},
		},
	)
	w.Header().Set("Content-Type", "text/html")
	var buf bytes.Buffer
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("error executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func errMessage(errs ...error) []string {
	var errMessage []string
	for _, err := range errs {
		var puberr public
		if errors.As(err, &puberr) {
			errMessage = append(errMessage, err.Error())
		} else {
			fmt.Println(err)
			errMessage = append(errMessage, "Something went wrong")
		}
	}
	return errMessage
}
