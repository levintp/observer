package types

// The `ControllerSpec` type is a structure that defines the configuration
// specification of the controller component.
type ControllerSpec struct {
	Host     string `yaml:"host" env:"CONTROLLER_HOST" flag:"controller-host"`                                               // Hostname of the controller.
	Port     int    `yaml:"port" env:"CONTROLLER_PORT" flag:"controller-port" default:"1139"`                                // Connection port of the controller.
	LogFile  string `yaml:"log_file" env:"CONTROLLER_LOG_FILE" flag:"log-file" default:"/var/log/observer/observerctld.log"` // Output file of the controller log.
	LogLevel string `yaml:"log_level" env:"CONTROLLER_LOG_LEVEL" flag:"log-level" default:"info"`                            // Level of the controller log.
}
