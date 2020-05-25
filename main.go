package main

import (
	"github.com/go-martini/martini"
	"github.com/zanefinner/web-scraper/accounts"
	"github.com/zanefinner/web-scraper/api"
)

func main() {
	server := martini.Classic()

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
	server.Get("/dashboard", func() string {
		//present dashboard
		return "dashboard"
	})

	//Api for AJAX
	server.Group("/api", func(link martini.Router) {
		link.Get("/products", api.ListProducts)
	})
	server.Run()
}
