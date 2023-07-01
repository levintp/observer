package types

// The `ModuleSpec` type is a structure that defines the specification of a
// module.
type ModuleSpec struct {
	Plugin    string                 `yaml:"plugin"`                  // Name of the module plugin.
	Timeout   int                    `yaml:"timeout" default:"1"`     // Maximum execution time.
	Function  string                 `yaml:"function" default:"exec"` // Symbol name of the function to run within the module.
	Arguments map[string]interface{} `yaml:"arguments,omitempty"`     // Arguments to pass to the module.
}
