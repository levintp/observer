package types

// The `AgentSpec` type is a structure that defines the configuration
// specification of the agent component.
type AgentSpec struct {
	Port     int    `yaml:"port" env:"AGENT_PORT" flag:"agent-port" default:"1016"`                                       // Connection port to the agent.
	LogFile  string `yaml:"log_file" env:"DATABASE_AGENT_FILE" flag:"log-file" default:"/var/log/observer/observerd.log"` // Output file of the agent log.
	LogLevel string `yaml:"log_level" env:"DATABASE_AGENT_LEVEL" flag:"log-level" default:"info"`                         // Level of the agent log.
}
