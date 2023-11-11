// database/database.go

package database

import (
	"fmt"
	"sync"
	"github.com/jmoiron/sqlx"
	// Use _ to save mysql driver for sqlx to use it correctly
	_ "github.com/go-sql-driver/mysql"

	"github.com/stephane-nguyen/FamilyTree/server/config"
)

var (
	db         *sqlx.DB
	dbInitOnce sync.Once
)

func CreateMySqlConnection() *sqlx.DB {
	dbInitOnce.Do(func() {
		db = createDatabaseConnection()
	})
	return db
}

func createDatabaseConnection() *sqlx.DB {
	env := config.LoadEnv()
	conn, err := sqlx.Connect("mysql", env.DSN)
	if err != nil {
		panic(fmt.Sprintf("[MySQL] Unable to connect to the database: %v\n", err))
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("[MySQL] Successfully connected to the database")

	return conn
}
