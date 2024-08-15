package expression

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorRemoveTerm(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveTerm,
		"../../testdata/expression/remove.go",
		6,
	)
}
