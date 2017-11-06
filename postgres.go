package mastodon

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// InitializePostgres initializes the connections to the Postgres host.
func InitializePostgres(config *Config) (*sqlx.DB, error) {
	dataSourceName := generateDataSourceName(config)

	// Make the connection.
	connection, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Set the connection limits.
	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxConnections)

	return connection, nil
}

// InitializePostgresWithContext initializes the connections to the Postgres host -- with context.
func InitializePostgresWithContext(ctx context.Context, config *Config) (*sqlx.DB, error) {
	dataSourceName := generateDataSourceName(config)

	// Make the connection.
	connection, err := sqlx.ConnectContext(ctx, "postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Set the connection limits.
	connection.SetMaxIdleConns(config.MaxIdleConnections)
	connection.SetMaxOpenConns(config.MaxConnections)

	return connection, nil
}

func generateDataSourceName(config *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
		config.ConnectTimeout,
	)
}
