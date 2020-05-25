package accounts

import "github.com/go-martini/martini"

//UseHandlers makes all routes work
func UseHandlers() {
	server := martini.Classic()
	server.Group("/accounts", func(r martini.Router) {
		r.Post("/new", NewAccount)
		r.Put("/update/:id", UpdateAccount)
		r.Delete("/delete/:id", DelAccount)
	}, AuthMiddleware)
}
