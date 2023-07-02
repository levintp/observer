package types

// The `AgentSpec` type is a structure that defines the configuration
// specification of the agent component.
type AgentSpec struct {
	Port       int    `yaml:"port" env:"AGENT_PORT" flag:"agent-port" default:"1016"`                                  // Connection port to the agent.
	LogFile    string `yaml:"log_file" env:"AGENT_LOG_FILE" flag:"log-file" default:"/var/log/observer/observerd.log"` // Output file of the agent log.
	LogLevel   string `yaml:"log_level" env:"AGENT_LOG_LEVEL" flag:"log-level" default:"info"`                         // Level of the agent log.
	BufferSize int    `yaml:"buffer_size" env:"AGENT_BUFFER_SIZE" flag:"buffer-size" default:"4096"`                   // Samples buffer size.
}
