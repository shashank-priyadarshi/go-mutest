package branch

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorCase(t *testing.T) {
	test.Mutator(
		t,
		MutatorCase,
		"../../testdata/branch/mutatecase.go",
		3,
	)
}
