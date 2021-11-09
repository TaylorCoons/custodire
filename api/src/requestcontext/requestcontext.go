package requestcontext

import (
	"database/sql"

	"github.com/TaylorCoons/custodire/api/src/logger"
)

type Ctx string

const (
	Key Ctx = "RequestContext"
)

type Data struct {
	Logger *logger.Logger
	Db     *sql.DB
}
