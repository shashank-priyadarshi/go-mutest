package arithmetic

import (
	"testing"

	"github.com/JekaMas/go-mutesting/test"
)

func TestMutatorArithmeticAssignInvert(t *testing.T) {
	test.Mutator(
		t,
		MutatorArithmeticAssignInvert,
		"../../testdata/arithmetic/assign_invert.go",
		5,
	)
}
