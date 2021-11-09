package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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
	fmt.Println("This called 1")
	val := ctx.Value(requestcontext.Key)
	if val == nil {
		// TODO: Log & Return 500
		panic(errors.New("couldn't find context. "))
	}
	db, ok := val.(*sql.DB)
	if !ok {
		// TODO: Log & Return 500
		panic(errors.New("couldn't cast context. "))
	}
	userInput := user.UserModelInput{}
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		// TODO: Log & Return 500
		panic(err.Error())
	}
	err = user.CreateUser(db, userInput)
	if err != nil {
		// TODO: Log & Return appropriate error
		panic(err.Error())
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
