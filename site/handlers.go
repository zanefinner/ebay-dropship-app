package site

import "github.com/go-martini/martini"

//UseHandlers makes all routes work
func UseHandlers() {
	server := martini.Classic()
	server.Get("/", func() string {
		return "home page"
	})
	server.Run()
}
