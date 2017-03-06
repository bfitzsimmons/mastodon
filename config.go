package mastodon

// Config holds the config. data for the Postgres connection.
type Config struct {
	Host               string
	Port               int
	Database           string
	User               string
	Password           string
	SSLMode            string
	MaxConnections     int
	MaxIdleConnections int
	ConnectionTimeout  int
}
