package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TaylorCoons/custodire/api/src/requestcontext"
	"github.com/TaylorCoons/custodire/api/src/sdk/user"
	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "POST", Path: "/user", Handler: CreateUser},
	{Method: "DELETE", Path: "/user/:username", Handler: DeleteUser},
	{Method: "POST", Path: "/user/:username/device", Handler: RegisterDevice},
	{Method: "DELETE", Path: "/user/:username/device/:deviceId", Handler: DeleteDevice},
	{Method: "POST", Path: "/user/:username/authenticate", Handler: Authenticate},
	{Method: "GET", Path: "/user/:username/device/:deviceId/account", Handler: ListAccounts},
	{Method: "GET", Path: "/user/:username/device/:deviceId/account/:id", Handler: GetAccount},
	{Method: "POST", Path: "/user/:username/device/:deviceId/account/account", Handler: CreateAccount},
	{Method: "PUT", Path: "/user/:username/device/:deviceId/account/account/:id", Handler: UpdateAccount},
	{Method: "DELETE", Path: "/user/:username/device/:deviceId/account/account/:id", Handler: DeleteAccount},
}

func ResolveDI(ctx context.Context, w http.ResponseWriter) requestcontext.Data {
	val := ctx.Value(requestcontext.Key)
	if val == nil {
		http.Error(w, "Internal Error. Please contact server admin.", http.StatusInternalServerError)
		panic("couldn't find context")
	}
	di, ok := val.(requestcontext.Data)
	if !ok {
		http.Error(w, "Internal Error. Please contact server admin.", http.StatusInternalServerError)
		panic("couldn't cast context")
	}
	return di
}

func CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	di := ResolveDI(ctx, w)

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
			di.Logger.Info(fmt.Sprintf("username: %s already exists", userInput.Username))
			http.Error(w, "Username already exists.", http.StatusConflict)
			return
		}
		di.Logger.Error("failed to create user")
		http.Error(w, "Internal Error. Please contact server admin.", http.StatusInternalServerError)
		return
	}
}

func DeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	di := ResolveDI(ctx, w)

	username := p["username"]
	err := user.DeleteUser(di.Db, username)
	if err != nil {
		if _, ok := err.(*user.UserDoesNotExistError); ok {
			di.Logger.Info(fmt.Sprintf("username: %s does not exist", username))
			http.Error(w, "Username does not exist.", http.StatusNotFound)
			return
		}
		di.Logger.Error("failed to delete user")
		http.Error(w, "Internal Error. Please contact server admin.", http.StatusInternalServerError)
		return
	}
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
