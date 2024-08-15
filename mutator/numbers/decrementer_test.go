package numbers

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorNumbersDecrementer(t *testing.T) {
	test.Mutator(
		t,
		MutatorNumbersDecrementer,
		"../../testdata/numbers/decrementer.go",
		2,
	)
}
