package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TaylorCoons/custodire/api/routes"
	"github.com/TaylorCoons/custodire/api/settings"

	server "github.com/TaylorCoons/gorouter"
)

func main() {
	var appSettings settings.AppSettings = settings.GetSettings()
	db := OpenConnectionPool(appSettings)
	defer db.Close()
	startServer(appSettings, db)
}

func OpenConnectionPool(appSettings settings.AppSettings) *sql.DB {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		appSettings.DbConnection.Username,
		appSettings.DbConnection.Password,
		appSettings.DbConnection.Host,
		appSettings.DbConnection.Database,
	)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func startServer(appSettings settings.AppSettings, db *sql.DB) {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	bind := fmt.Sprintf(":%s", appSettings.Port)
	fmt.Printf("[custodire] Listening on port: %s", appSettings.Port)
	err := http.ListenAndServe(bind, server)
	if err != nil {
		panic(err)
	}
}
