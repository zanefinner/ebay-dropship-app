package dashboard

import (
	"github.com/go-martini/martini"
	"github.com/zanefinner/web-scraper/accounts"
)

//UseHandlers makes all routes work
func UseHandlers() {
	server := martini.Classic()
	server.Group("/dashboard", func(r martini.Router) {
		r.Post("/home", Home)
	}, accounts.AuthMiddleware)
}
