package main

import (
	"github.com/aerogo/aero"
)

var (
	Version string = ""

	EventsTTL int32 = 86400
	ParseTTL  int32 = 86400 * 7
	Api             = "https://thwiki.cc"
	Memcached       = "127.0.0.1:11211"
)

func main() {
	app := aero.New()

	if Version == "" {
		app.Config.Ports.HTTP += 100
	}
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	app.Get("/", func(ctx aero.Context) error {
		return ctx.String("welcome to calendar api version " + Version)
	})

	app.Get("/events", func(ctx aero.Context) error {
		return handleEvents(ctx, ctx.Query("start"), ctx.Query("end"))
	})

	app.Get("/events/", func(ctx aero.Context) error {
		return handleEvents(ctx, ctx.Query("start"), ctx.Query("end"))
	})

	app.Get("/events/:start/:end", func(ctx aero.Context) error {
		return handleEvents(ctx, ctx.Get("start"), ctx.Get("end"))
	})

	return app
}

func handleEvents(ctx aero.Context, startStr string, endStr string) error {
	response := ctx.Response()
	response.SetHeader("Content-Security-Policy", "default-src 'none'")
	response.SetHeader("X-Content-Type-Options", "nosniff")
	response.SetHeader("X-Frame-Options", "SAMEORIGIN")
	response.SetHeader("Referrer-Policy", "same-origin")

	start, err := SanitizeDate(startStr)
	if err != nil {
		return ctx.Error(400, err)
	}
	end, err := SanitizeDate(endStr)
	if err != nil {
		return ctx.Error(400, err)
	}

	events, err := GetEvents(start, end)

	if err != nil {
		return ctx.Error(503, err)
	}

	response.SetHeader("Content-Type", "application/json; charset=utf-8")
	return ctx.Bytes(events)
}
