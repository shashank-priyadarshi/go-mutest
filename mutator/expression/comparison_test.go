package expression

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorComparison(t *testing.T) {
	test.Mutator(
		t,
		MutatorComparison,
		"../../testdata/expression/comparison.go",
		4,
	)
}
