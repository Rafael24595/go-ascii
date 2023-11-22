package logger_repository

import (
	"database/sql"
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/log/event"
	"go-ascii/src/commons/log/logger"
	"go-ascii/src/commons/log/logger/postgres/catalog"
	"strings"
	_ "github.com/lib/pq"
)

const LoggerPostgresKey = "LoggerPostgres"

type LoggerPostgres struct {
	dataBase *sql.DB
}

func NewLoggerPostgres(args map[string]string) logger.Logger {
	connStr := getConnectionUri(args)
	dataBase, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = dataBase.Ping(); err != nil {
		panic(err)
	}
	return LoggerPostgres{dataBase: dataBase}
}

func getConnectionUri(args map[string]string) string {
	user := args["ASCII_POSTGRES_USERNAME"]
	password := args["ASCII_POSTGRES_PASSWORD"]
	server := args["ASCII_POSTGRES_SERVER"]
	port := args["ASCII_POSTGRES_PORT"]
	dataBase := args["ASCII_POSTGRES_DB"]

	var connection strings.Builder
	connection.WriteString("postgres://")
	connection.WriteString(user)
	connection.WriteString(":")
	connection.WriteString(password)
	connection.WriteString("@")
	connection.WriteString(server)
	connection.WriteString(":")
	connection.WriteString(port)
	connection.WriteString("/")
	connection.WriteString(dataBase)
	connection.WriteString("?sslmode=disable")
	return connection.String()

}

func (this LoggerPostgres) DependencyName() string {
	return LoggerPostgresKey
}

func (this LoggerPostgres) OnLoad() bool {
	configuration := configuration.GetInstance()
	insertDynStmt := postgres_catalog.GetSource(postgres_catalog.PG_INSERT_SESSION)
	_, err := this.dataBase.Exec(insertDynStmt, configuration.GetSessionId(), configuration.GetTimestamp())
    if err != nil {
		panic(err)
	}
	return true
}

func (this LoggerPostgres) OnExit() bool {
	return true
}

func (this LoggerPostgres) Log(event log_event.LogEvent) string {
	insertDynStmt := postgres_catalog.GetSource(postgres_catalog.PG_INSERT_REGISTER)
    _, err := this.dataBase.Exec(insertDynStmt, event.GetSessionId(), event.GetCategory(), event.GetFamily(), event.GetMessage(), event.GetTimestamp().UnixMilli())
    if err != nil {
		panic(err)
	}
	return ""
}
