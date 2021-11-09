package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TaylorCoons/custodire/api/src/dbstartup"
	"github.com/TaylorCoons/custodire/api/src/requestcontext"
	"github.com/TaylorCoons/custodire/api/src/routes"
	"github.com/TaylorCoons/custodire/api/src/settings"

	server "github.com/TaylorCoons/gorouter"
)

func main() {
	var appSettings settings.AppSettings = settings.GetSettings()
	db := OpenConnectionPool(appSettings)
	defer db.Close()
	InitializeDatabase(db)
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

func InitializeDatabase(db *sql.DB) {
	err := dbstartup.InitializeTables(db)
	if err != nil {
		panic(err.Error())
	}
}

func startServer(appSettings settings.AppSettings, db *sql.DB) {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	var dbInjector server.MiddlewareFunc = func(w http.ResponseWriter, r *http.Request, p server.PathParams, h server.HandlerFunc) {
		ctx := context.WithValue(context.Background(), requestcontext.Ctx(requestcontext.Key), db)
		h(ctx, w, r, p)
	}
	server := server.Server{
		CompiledRoutes: compiledRoutes,
		Middleware:     dbInjector,
	}
	bind := fmt.Sprintf(":%s", appSettings.Port)
	fmt.Printf("[custodire] Listening on port: %s", appSettings.Port)
	err := http.ListenAndServe(bind, server)
	if err != nil {
		panic(err)
	}
}
