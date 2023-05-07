package postgres_repository

import (
	"database/sql"
	"go-ascii/src/commons/configurator/configuration"
	"go-ascii/src/commons/dto"
	"go-ascii/src/commons/log/event"
	"go-ascii/src/infrastructure/repository"
	"go-ascii/src/infrastructure/repository/log/postgres/catalog"
	"strings"
)

const RepositoryLogPostgresKey = "RepositoryLogPostgres"

type RepositoryLogPostgres struct {
	dataBase *sql.DB
}

func NewRepositoryLogPostgres(args map[string]string) repository.RepositoryLog {
	connStr := getConnectionUri(args)
	dataBase, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = dataBase.Ping(); err != nil {
		panic(err)
	}
	return RepositoryLogPostgres{dataBase: dataBase}
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

func (this RepositoryLogPostgres) DependencyName() string {
	return RepositoryLogPostgresKey
}

func (this RepositoryLogPostgres) OnLoad() bool {
	return true
}

func (this RepositoryLogPostgres) OnExit() bool {
	return true
}

func (this RepositoryLogPostgres) FindAll() (logs []log_event.LogEvent) {
	configuration := configuration.GetInstance()
	query := postgres_catalog.GetSource(postgres_catalog.PG_FIND_ALL_REGISTER)
	rows, err := this.dataBase.Query(query, configuration.GetSessionId())
    if err != nil {
		panic(err)
	}
	defer rows.Close()

	return this.rowsToLogEvent(rows)
}

func (this RepositoryLogPostgres) Find(category string) []log_event.LogEvent{
	query := postgres_catalog.GetSource(postgres_catalog.PG_FIND_CATEGORY_REGISTER)
	rows, err := this.dataBase.Query(query, category)
    if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	return this.rowsToLogEvent(rows)
}

func (this RepositoryLogPostgres) rowsToLogEvent(rows *sql.Rows) (logs []log_event.LogEvent) {
	dtos := make([]dto.InfoLogResponse, 0)

	for rows.Next() {
		dto := dto.InfoLogResponse{}
		err := rows.Scan(&dto.Id, &dto.SessionId, &dto.Category, &dto.Message, &dto.Timestamp)
		if err != nil {
			panic(err)
		}
		dtos = append(dtos, dto)
	}

	logs = make([]log_event.LogEvent, 0)

	for _, dto := range dtos {
		log := log_event.NewLogEventDB(dto.Id, dto.SessionId, dto.Category, dto.Message, dto.Timestamp)

		logs = append(logs, log)
	}

	return
}