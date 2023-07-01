package types

// The `APISpec` type is a structure that defines the configuration
// specification of the REST API component.
type APISpec struct {
	Host     string `yaml:"host" env:"API_HOST" flag:"api-host"`                                                           // Hostname of the API.
	Port     int    `yaml:"port" env:"API_PORT" flag:"api-port" default:"1086"`                                            // RPC port of the API.
	RestPort int    `yaml:"rest_port" env:"API_REST_PORT" flag:"api-rest-port" default:"1291"`                             // REST port of the API.
	LogFile  string `yaml:"log_file" env:"DATABASE_API_FILE" flag:"log-file" default:"/var/log/observer/observerapid.log"` // Output file of the API log.
	LogLevel string `yaml:"log_level" env:"DATABASE_API_LEVEL" flag:"log-level" default:"info"`                            // Level of the API log.
}
