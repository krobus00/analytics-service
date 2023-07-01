package bootstrap

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // postgres driver

	"github.com/krobus00/analytics-service/internal/config"
	"github.com/krobus00/analytics-service/internal/utils"
	"github.com/pressly/goose/v3"
)

func StartMigration(actionType string, name string, step *int64) {
	migrationDir := "db/migrations"

	db, err := sql.Open("postgres", config.DatabaseDSN())
	utils.ContinueOrFatal(err)
	err = goose.SetDialect("postgres")
	utils.ContinueOrFatal(err)

	switch actionType {
	case "create":
		err = goose.Create(db, migrationDir, name, "sql")
	case "up":
		err = goose.Up(db, migrationDir)
	case "up-by-one":
		err = goose.UpByOne(db, migrationDir)
	case "up-to":
		err = goose.UpTo(db, migrationDir, *step)
	case "down":
		err = goose.Down(db, migrationDir)
	case "down-to":
		err = goose.DownTo(db, migrationDir, *step)
	case "status":
		err = goose.Status(db, migrationDir)
	case "reset":
		err = goose.Reset(db, migrationDir)
		if err != nil {
			break
		}
		err = goose.Up(db, migrationDir)
	default:
		err = errors.New("invalid command")
	}

	utils.ContinueOrFatal(err)
}
