package postgres

import (
	"database/sql"
	"log"

	"anilkhadka.com.np/task-management/conf"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func CreateDBConnection() *sql.DB {
	connectionStr := conf.EnvConfigs.DatabaseSource

	conn, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	}

	runMigrations()

	return conn
}

func runMigrations() {
	migrationURL := "file://database/postgres/migrations"

	migration, err := migrate.New(migrationURL, conf.EnvConfigs.DatabaseSource)

	if err != nil {
		log.Println(err)
		log.Fatal("cannot create new migrate instance")
	}

	defer migration.Close()

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("db migrated successfully")

}
