package dashboard

import (
	"github.com/go-martini/martini"
	"github.com/zanefinner/web-scraper/accounts"
)

//UseHandlers makes all routes work
func UseHandlers() {
	server := martini.Classic()
	server.Get("/dashboard", func() string {
		//present dashboard
		return "dashboard"
	}, accounts.AuthMiddleware)
}
