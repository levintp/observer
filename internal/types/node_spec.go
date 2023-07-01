package types

// The `NodeSpec` type is a structure that defines a monitored node in the
// cluster.
type NodeSpec struct {
	Name       string   `yaml:"name"`                 // Hostname of the node.
	Categories []string `yaml:"categories,omitempty"` // Categories the node is a part of.
}
