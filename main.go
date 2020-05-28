package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/zanefinner/ebay-dropship-app/accounts"
	"github.com/zanefinner/ebay-dropship-app/api"
)

func main() {
	server := martini.Classic()
	server.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true}))
	//Landing page
	server.Get("/", func(params martini.Params, renderer render.Render) {
		renderer.Redirect("http://localhost:3000/hello.html")
	})
	//Account
	server.Group("/accounts", func(r martini.Router) {
		r.Get("/login", accounts.PresentLoginForm)
		r.Post("/login", binding.Bind(accounts.LoginForm{}), accounts.Match)
		r.Post("/new", accounts.New)
		r.Put("/update/:id", accounts.Update)
		r.Delete("/delete/:id", accounts.Del)
	}, accounts.AuthMiddleware)

	//Dashboard
	server.Get("/dashboard", func(r render.Render) {
		r.HTML(http.StatusOK, "dashboard", nil)
	})

	//Api for AJAX
	server.Group("/api", func(link martini.Router) {
		link.Get("/ebay", api.Use)
	})

	server.Run()
}
