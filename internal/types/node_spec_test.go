package types_test

import (
	"testing"

	"github.com/levintp/observer/internal/types"
)

func TestInCategories(t *testing.T) {
	node := types.NodeSpec{"node01", []types.Category{"alpha", "beta", "gamma"}}

	// Test with single matching category.
	if !node.InCatrgories([]types.Category{"delta", "beta"}) {
		t.Error("node falsely report it is not in a category")
	}

	// Test without any matching categories.
	if node.InCatrgories([]types.Category{"delta", "eplison"}) {
		t.Error("node falsely report is it in a category")
	}

	// Test implicit ALL category.
	if !node.InCatrgories([]types.Category{}) {
		t.Error("node falsely report it is not in the ALL category (implicit)")
	}

	// Test explicit ALL category.
	if !node.InCatrgories([]types.Category{types.CategoryALL}) {
		t.Error("node falsely report it is not in the ALL category (explicit)")
	}
}
