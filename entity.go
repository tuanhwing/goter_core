package gotercore

import (
	"os"
)

///GetDBName return database name
func GetDBName() string {
	return os.Getenv("DB_NAME")
}
