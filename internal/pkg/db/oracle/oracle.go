package oracle

import (
	"b2b-api/internal/config"
	"b2b-api/internal/pkg/logger"
	"database/sql"
	_ "github.com/godror/godror"
	"go.uber.org/fx"
)

var NewOracle = fx.Provide(newOracle)

type dependencies struct {
	fx.In
	Config config.IConfig
	Logger logger.ILogger
}

type oracle struct {
	Oracle *sql.DB
}

type Idb interface {
	GetConnection() *sql.DB
}

func newOracle(dp dependencies) Idb {
	dbConn, err := sql.Open("godror", dp.Config.GetString("api.oracle.tns"))
	if err != nil {
		dp.Logger.Error("%s", "Connection error: "+dp.Config.GetString("api.oracle.tns ::: ")+err.Error())
		return nil
	}
	dp.Logger.Info("%s", "Connection success: "+dp.Config.GetString("api.oracle.tns"))
	if err := dbConn.Ping(); err != nil {
		dp.Logger.Error("%s", "Ping to DB has got an error ::: "+err.Error())
		return nil
	}
	return &oracle{Oracle: dbConn}
}
func (o *oracle) GetConnection() *sql.DB {
	return o.Oracle
}
