package types

// The `DatabaseSpec` type is a structure that defines the configuration
// specification of the database component.
type DatabaseSpec struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" flag:"database-host"`                                                // Hostname of the database.
	Port     int    `yaml:"port" env:"DATABASE_PORT" flag:"database-port" default:"9200"`                                 // Connection port of the database.
	User     string `yaml:"user" env:"DATABASE_USER" flag:"database-user"`                                                // Username used to authenticate with the database.
	Pass     string `yaml:"pass" env:"DATABASE_PASS" flag:"database-pass"`                                                // Password used to authenticate with the database.
	Index    string `yaml:"index" env:"DATABASE_INEDX" flag:"database-index" default:"observer-streams"`                  // The elasticsearch index to write metrics to.
	LogFile  string `yaml:"log_file" env:"DATABASE_LOG_FILE" flag:"log-file" default:"/var/log/observer/observerdbd.log"` // Output file of the database log.
	LogLevel string `yaml:"log_level" env:"DATABASE_LOG_LEVEL" flag:"log-level" default:"info"`                           // Level of the database log.
}
