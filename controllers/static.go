package controllers

import (
	"html/template"
	"net/http"

	"github.com/gerry-sheva/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version",
			Answer:   "Yes",
		},
		{
			Question: "Is there a paid version",
			Answer:   "Yes",
		},
		{
			Question: "Is there a free version",
			Answer:   "Yes",
		},
		{
			Question: "Is there a paid version",
			Answer:   "Yes",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
