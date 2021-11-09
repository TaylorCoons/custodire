package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TaylorCoons/custodire/api/src/dbstartup"
	"github.com/TaylorCoons/custodire/api/src/logger"
	"github.com/TaylorCoons/custodire/api/src/requestcontext"
	"github.com/TaylorCoons/custodire/api/src/routes"
	"github.com/TaylorCoons/custodire/api/src/settings"

	server "github.com/TaylorCoons/gorouter"
)

func main() {
	logger, err := logger.New("custodire.log")
	if err != nil {
		panic(err.Error())
	}
	var appSettings settings.AppSettings = settings.GetSettings(logger)
	db := OpenConnectionPool(logger, appSettings)
	defer db.Close()
	InitializeDatabase(logger, db)
	startServer(logger, appSettings, db)
}

func OpenConnectionPool(logger *logger.Logger, appSettings settings.AppSettings) *sql.DB {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		appSettings.DbConnection.Username,
		appSettings.DbConnection.Password,
		appSettings.DbConnection.Host,
		appSettings.DbConnection.Database,
	)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err.Error())
	}
	return db
}

func InitializeDatabase(logger *logger.Logger, db *sql.DB) {
	err := dbstartup.InitializeTables(db)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err.Error())
	}
}

func startServer(logger *logger.Logger, appSettings settings.AppSettings, db *sql.DB) {
	compiledRoutes := server.CompileRoutes(routes.Routes)
	var diInjector server.MiddlewareFunc = func(w http.ResponseWriter, r *http.Request, p server.PathParams, h server.HandlerFunc) {
		var rCtx requestcontext.Data = requestcontext.Data{
			Logger: logger,
			Db:     db,
		}
		ctx := context.WithValue(context.Background(), requestcontext.Ctx(requestcontext.Key), rCtx)
		h(ctx, w, r, p)
	}
	server := server.Server{
		CompiledRoutes: compiledRoutes,
		Middleware:     diInjector,
	}
	bind := fmt.Sprintf(":%s", appSettings.Port)
	logger.Info(fmt.Sprintf("Listening on port: %s", appSettings.Port))
	err := http.ListenAndServe(bind, server)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}
}
