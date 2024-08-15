package arithmetic

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorArithmeticAssignment(t *testing.T) {
	test.Mutator(
		t,
		MutatorArithmeticAssignment,
		"../../testdata/arithmetic/assignment.go",
		11,
	)
}
