package postgres_catalog

import "os"

type CatalogSource string

const (
	directory          string        = "src/infrastructure/repository/log/postgres/catalog/"
	PG_FIND_ALL_REGISTER  CatalogSource = "PG_FIND_ALL_REGISTER.sql"
)

func GetSource(code CatalogSource) string {
	path := string(directory) + string(code)
	scriptByte, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(scriptByte)
}