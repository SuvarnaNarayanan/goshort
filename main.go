package main

import (
	"github.com/labstack/echo/v4"
	"goshort/m/v2/api"
	"goshort/m/v2/api/impl"
	"net"
)

func main() {
	server := impl.NewGoShortApiServer()
	e := echo.New()

	serverHandler := api.NewStrictHandler(server, nil)
	api.RegisterHandlers(e, serverHandler)

	e.Logger.Fatal(e.Start(net.JoinHostPort("localhost", "8080")))
}
