// +build migrations

package db

import (
	"os"
	"testing"

	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/platformsh-upstream-forks/factory"
	dbLogger "github.com/platformsh-upstream-forks/factory/logger"
)

func TestMigrations(t *testing.T) {
	factory.Database.Add("compose", os.Getenv("COMPOSE_DB_DSN"))
	db := factory.Database.MustGet("compose")
	db.SetLogger(dbLogger.Default{})

	if err := Migrate(db, logger.Default()); err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}
}
