package types

// The `NodeSpec` type is a structure that defines a monitored node in the
// cluster.
type NodeSpec struct {
	Name       string     `yaml:"name"`                 // Hostname of the node.
	Categories []Category `yaml:"categories,omitempty"` // Categories the node is a part of.
}

// Function to check if a node is in a set of categories.
func (node NodeSpec) InCatrgories(categories []Category) bool {
	// If list of categories is empty, assume ALL categories.
	if len(categories) == 0 {
		return true
	}

	// Check if the node is in any of the given categories.
	for _, category := range categories {
		for _, nodeCategory := range node.Categories {
			if category == nodeCategory || category == CategoryALL {
				return true
			}
		}
	}

	// If execution reached this point, no matching category found.
	return false
}
