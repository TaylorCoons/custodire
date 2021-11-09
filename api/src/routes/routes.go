package routes

import (
	"context"
	"net/http"

	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/user", Handler: CreateUser},
	{Method: "DELETE", Path: "/user/:userId", Handler: DeleteUser},
	{Method: "POST", Path: "/user/:userId/device", Handler: RegisterDevice},
	{Method: "DELETE", Path: "/user/:userId/device/:deviceId", Handler: DeleteDevice},
	{Method: "POST", Path: "/user/:userId/authenticate", Handler: Authenticate},
	{Method: "GET", Path: "/user:userId/device/:deviceId/account", Handler: ListAccounts},
	{Method: "GET", Path: "/user:userId/device/:deviceId/account/:id", Handler: GetAccount},
	{Method: "POST", Path: "/user:userId/device/:deviceId/account/account", Handler: CreateAccount},
	{Method: "PUT", Path: "/user:userId/device/:deviceId/account/account/:id", Handler: UpdateAccount},
	{Method: "DELETE", Path: "/user:userId/device/:deviceId/account/account/:id", Handler: DeleteAccount},
}

func CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func DeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func RegisterDevice(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func DeleteDevice(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func Authenticate(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func ListAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func GetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func CreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func UpdateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}

func DeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not Implimented!"))
}
