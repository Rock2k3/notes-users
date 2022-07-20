package adapters

import (
	"database/sql"
	"github.com/Rock2k3/notes-users/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

type postgresDb struct {
	config *config.AppConfig
}

func NewPostgresDb(c *config.AppConfig) *postgresDb {
	return &postgresDb{config: c}
}

func (d postgresDb) GetDb() *sql.DB {
	db, err := sql.Open("postgres", d.config.DatasourceUrl())

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// return the connection
	return db
}

func (d postgresDb) Init() error {
	db := d.GetDb()
	defer db.Close()

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver)
	err := m.Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Println(err)
			return nil
		}
		return err
	}

	return nil
}
