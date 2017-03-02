package mastodon

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// PG contains the postgres connection pool.
var PG *Postgres

// Postgres wraps the Postgres client so that functionality may be added if needed.
type Postgres struct {
	*sqlx.DB
}

// InitializePostgres initializes the connections to the Postgres host.
func InitializePostgres(config *Config) (*Postgres, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
	)

	// Make the connection.
	connection, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Ping the server to make sure we're connected.
	if err = connection.Ping(); err != nil {
		return nil, err
	}

	// Set the connection limits.
	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxConnections)

	return &Postgres{connection}, nil
}
