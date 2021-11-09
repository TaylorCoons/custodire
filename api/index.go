package main

import (
	"fmt"
	"net/http"

	"github.com/TaylorCoons/custodire/api/routes"
	"github.com/TaylorCoons/custodire/api/util"

	server "github.com/TaylorCoons/gorouter"
)

func main() {
	startServer()
}

func startServer() {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	port := util.GetPort()
	bind := fmt.Sprintf(":%d", port)
	fmt.Printf("[custadire] Listening on port: %d", port)
	err := http.ListenAndServe(bind, server)
	if err != nil {
		panic(err)
	}
}
