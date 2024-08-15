package arithmetic

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorArithmeticBase(t *testing.T) {
	test.Mutator(
		t,
		MutatorArithmeticBase,
		"../../testdata/arithmetic/base.go",
		5,
	)
}
