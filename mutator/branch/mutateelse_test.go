package branch

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorElse(t *testing.T) {
	test.Mutator(
		t,
		MutatorElse,
		"../../testdata/branch/mutateelse.go",
		1,
	)
}
