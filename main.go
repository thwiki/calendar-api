package main

import (
	"github.com/aerogo/aero"
)

var TTL int32 = 86400
var Api = "https://thwiki.cc"
var Memcached = "127.0.0.1:11211"

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	app.Get("/", func(ctx aero.Context) error {
		return ctx.String("webcome to calendar api")
	})

	app.Get("/events/:start/:end", func(ctx aero.Context) error {
		response := ctx.Response()
		response.SetHeader("Content-Security-Policy", "default-src 'none'")
		response.SetHeader("X-Content-Type-Options", "nosniff")
		response.SetHeader("X-Frame-Options", "SAMEORIGIN")
		response.SetHeader("Referrer-Policy", "same-origin")

		start, err := SanitizeDate(ctx.Get("start"))
		if err != nil {
			return ctx.Error(400, err)
		}
		end, err := SanitizeDate(ctx.Get("end"))
		if err != nil {
			return ctx.Error(400, err)
		}

		events, err := GetEvents(start, end)

		if err != nil {
			return ctx.Error(503, err)
		}
		response.SetHeader("Content-Type", "application/json; charset=utf-8")

		return ctx.Bytes(events)
	})

	return app
}
