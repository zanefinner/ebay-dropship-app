package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/zanefinner/web-scraper/accounts"
	"github.com/zanefinner/web-scraper/api"
)

func main() {
	server := martini.Classic()
	server.Use(render.Renderer())
	//Landing page
	server.Get("/", func() string {
		return "home page"
	})

	//Account
	server.Group("/accounts", func(r martini.Router) {
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
		link.Get("/products", api.ListProducts)
	})
	server.Run()
}
