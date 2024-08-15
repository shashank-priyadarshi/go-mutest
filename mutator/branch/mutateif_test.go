package branch

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorIf(t *testing.T) {
	test.Mutator(
		t,
		MutatorIf,
		"../../testdata/branch/mutateif.go",
		2,
	)
}
