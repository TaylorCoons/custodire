package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/TaylorCoons/custodire/api/src/requestcontext"
	"github.com/TaylorCoons/custodire/api/src/sdk/user"
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
	val := ctx.Value(requestcontext.Key)
	if val == nil {
		fmt.Fprintln(os.Stderr, "coulnd't find context.")
		http.Error(w, "Internal Error. Please check server logs.", http.StatusInternalServerError)
		return
	}
	di, ok := val.(requestcontext.Data)
	if !ok {
		fmt.Fprintln(os.Stderr, "couldn't cast context.")
		http.Error(w, "Internal Error. Please check server logs.", http.StatusInternalServerError)
		return
	}
	userInput := user.UserModelInput{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&userInput)
	if err != nil {
		di.Logger.Info(fmt.Sprintf("unable to demarshal user data: %s", err.Error()))
		http.Error(w, "Unable to demarshal user data. Check the request JSON.", http.StatusBadRequest)
		return
	}
	err = user.CreateUser(di.Db, userInput)
	if err != nil {
		if _, ok := err.(*user.UserExistsError); ok {
			di.Logger.Info(fmt.Sprintf("username: %s already exists.", userInput.Username))
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
		di.Logger.Error("Failed to create user")
		http.Error(w, "Internal Error. Please check server logs.", http.StatusInternalServerError)
		return
	}
}

func DeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func RegisterDevice(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func DeleteDevice(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func Authenticate(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func ListAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func GetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func CreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func UpdateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}

func DeleteAccount(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Write([]byte("not implimented!"))
}
