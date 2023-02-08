package arithmetic

import (
	"testing"

	"github.com/visu-suganya/go-mutesting/test"
)

func TestMutatorArithmeticBase(t *testing.T) {
	test.Mutator(
		t,
		MutatorArithmeticBase,
		"../../testdata/arithmetic/base.go",
		5,
	)
}
