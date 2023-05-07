package postgres_catalog

import "os"

type CatalogSource string

const (
	directory          string        = "src/commons/log/logger/postgres/catalog/"
	PG_INSERT_SESSION  CatalogSource = "PG_INSERT_SESSION.sql"
	PG_INSERT_REGISTER CatalogSource = "PG_INSERT_REGISTER.sql"
)

func GetSource(code CatalogSource) string {
	path := string(directory) + string(code)
	scriptByte, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(scriptByte)
}