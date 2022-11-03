package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/desktop/{version}", func(ctx iris.Context) {
		platform := ctx.URLParam("platform")
		arch := ctx.URLParam("arch")
		version := ctx.Params().Get("version")
		ctx.Redirect(fmt.Sprintf("https://registry.npmmirror.com/-/binary/cypress/%s/%s-%s/cypress.zip", version, platform, arch), iris.StatusPermanentRedirect)
	})

	app.Listen(":8090")
}
