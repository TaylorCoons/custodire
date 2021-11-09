package routes

import (
	"net/http"

	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/user", Handler: CreateUser},
	{Method: "DELETE", Path: "/user:userId", Handler: DeleteUser},
	{Method: "POST", Path: "/user/:userId/device", Handler: RegisterDevice},
	{Method: "DELETE", Path: "/user/:userId/device/:deviceId", Handler: DeleteDevice},
	{Method: "POST", Path: "/user/:userId/authenticate", Handler: Autheticate},
	{Method: "GET", Path: "/account", Handler: ListAccounts},
	{Method: "GET", Path: "/account:id", Handler: GetAccount},
	{Method: "POST", Path: "/account", Handler: CreateAccount},
	{Method: "PUT", Path: "/account/:id", Handler: UpdateAccount},
	{Method: "DELETE", Path: "/account/:id", Handler: DeleteAccount},
}

func CreateUser(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func RegisterDevice(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func DeleteDevice(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func Authenticate(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func ListAccounts(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func GetAccount(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func CreateAccount(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func UpdateAccount(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}

func DeleteAccount(w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("Not Implimented!"))
}
