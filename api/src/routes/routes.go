package routes

import (
	"context"
	"net/http"

	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/user", Handler: CreateUser},
	{Method: "DELETE", Path: "/user:userId", Handler: DeleteUser},
	{Method: "POST", Path: "/user/:userId/device", Handler: RegisterDevice},
	{Method: "DELETE", Path: "/user/:userId/device/:deviceId", Handler: DeleteDevice},
	{Method: "POST", Path: "/user/:userId/authenticate", Handler: Authenticate},
	{Method: "GET", Path: "/account", Handler: ListAccounts},
	{Method: "GET", Path: "/account:id", Handler: GetAccount},
	{Method: "POST", Path: "/account", Handler: CreateAccount},
	{Method: "PUT", Path: "/account/:id", Handler: UpdateAccount},
	{Method: "DELETE", Path: "/account/:id", Handler: DeleteAccount},
}

func CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func DeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func RegisterDevice(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func DeleteDevice(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func Authenticate(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func ListAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func GetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func CreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func UpdateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func DeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}
