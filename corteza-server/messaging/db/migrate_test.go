// +build migrations

package db

import (
	"os"
	"testing"

	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/titpetric/factory"
	dbLogger "github.com/titpetric/factory/logger"
)

func TestMigrations(t *testing.T) {
	factory.Database.Add("messaging", os.Getenv("MESSAGING_DB_DSN"))
	db := factory.Database.MustGet("messaging")
	db.SetLogger(dbLogger.Default{})

	if err := Migrate(db, logger.Default()); err != nil {
		t.Fatalf("Unexpected error: %#v", err)
	}
}