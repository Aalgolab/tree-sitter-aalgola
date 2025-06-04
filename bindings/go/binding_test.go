package tree_sitter_aalgola_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_aalgola "github.com/aalgolab/tree-sitter-aalgola/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_aalgola.Language())
	if language == nil {
		t.Errorf("Error loading Aalgola grammar")
	}
}
