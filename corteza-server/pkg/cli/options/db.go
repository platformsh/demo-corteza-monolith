package options

import (
	"time"

	sqldsn "github.com/platformsh/config-reader-go/sqldsn"
	psh "github.com/platformsh/config-reader-go"
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
	var dsn string
	config, err := psh.NewRuntimeConfig()
	if err != nil {
		dsn = "corteza:corteza@tcp(db:3306)/corteza?collation=utf8mb4_general_ci"
	} else {
		credentials, err := config.Credentials("db")
		if err != nil {
			panic(err)
		}
		
		formatted, err := sqldsn.FormattedCredentials(credentials)
		if err != nil {
			panic(err)
		}
		
		dsn = formatted
	}
	
	o = &DBOpt{
		DSN:      dsn,
		Logger:   false,
		MaxTries: 100,
		Delay:    5 * time.Second,
		Timeout:  1 * time.Minute,
	}

	fill(o, pfix)

	return
}
