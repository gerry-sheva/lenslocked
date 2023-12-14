package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
		tpl.Execute(w, r, questions)
	}
}
