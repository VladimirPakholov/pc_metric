package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"pc_metric/internal/db"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration() error {

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error get current dir: %v", err)
	}

	pathWorkDir := os.Getenv("WORK_DIR_PATH")

	if pathWorkDir == "" {
		pathWorkDir = cwd
	}

	pathToMigration := filepath.Join(pathWorkDir, "migrations/pg")

	absMigrationPath, err := filepath.Abs(pathToMigration)

	if err != nil {
		return fmt.Errorf("Failed to get absolute path for %s: %w", absMigrationPath, err)
	}

	migration, err := migrate.New("file://"+filepath.ToSlash(absMigrationPath), db.BuildPostgreDSN())
	if err != nil {
		return err
	}
	defer migration.Close()
	errM := migration.Up()
	if errM != nil && errM != migrate.ErrNoChange {
		return errM
	}
	return nil
}
