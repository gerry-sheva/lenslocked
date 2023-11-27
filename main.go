package main

import (
	"fmt"
	"net/http"

	"github.com/gerry-sheva/lenslocked/controllers"
	"github.com/gerry-sheva/lenslocked/templates"
	"github.com/gerry-sheva/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)

	fmt.Println("server is ready at port 3000")
	http.ListenAndServe(":3000", r)
}
