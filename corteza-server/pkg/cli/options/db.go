package options

import (
	"time"
	sqldsn "github.com/platformsh/config-reader-go/v2/sqldsn"
	psh "github.com/platformsh/config-reader-go/v2"
)

type (
	DBOpt struct {
		DSN      string        `env:"DB_DSN"`
		Logger   bool          `env:"DB_LOGGER"`
		MaxTries int           `env:"DB_MAX_TRIES"`
		Delay    time.Duration `env:"DB_CONN_ERR_DELAY"`
		Timeout  time.Duration `env:"DB_CONN_TIMEOUT"`
	}
)

func DB(pfix string) (o *DBOpt) {
	// Creating a psh.RuntimeConfig struct
	config, err := psh.NewRuntimeConfig()
	if err != nil {
		panic("Not in a Platform.sh Environment.")
	}

	// Accessing the database relationship Credentials struct
	credentials, err := config.Credentials("db")
	if err != nil {
		panic(err)
	}

	// Using the sqldsn formatted credentials package
	formatted, err := sqldsn.FormattedCredentials(credentials)
	if err != nil {
		panic(err)
	}

	o = &DBOpt{
		DSN:      formatted,
		Logger:   false,
		MaxTries: 100,
		Delay:    5 * time.Second,
		Timeout:  1 * time.Minute,
	}

	fill(o, pfix)

	return
}
