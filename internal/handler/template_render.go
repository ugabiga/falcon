package handler

import (
	"html/template"
	"net/http"
)

const (
	pageRoot      = "web/src/page"
	componentRoot = "web/src/component"
	layoutPage    = "web/src/layout/index.html"
)

func RenderPage(w http.ResponseWriter, data any, filenames ...string) error {
	for i, file := range filenames {
		filenames[i] = pageRoot + file
	}
	filenames = append(filenames, layoutPage)

	t := template.Must(template.ParseFiles(
		filenames...,
	))

	if err := t.Execute(w, data); err != nil {
		return err
	}

	return nil
}

func RenderComponent(w http.ResponseWriter, data any, file string) error {
	file = componentRoot + file

	t := template.Must(template.ParseFiles(
		file,
	))

	if err := t.Execute(w, data); err != nil {
		return err
	}

	return nil
}
