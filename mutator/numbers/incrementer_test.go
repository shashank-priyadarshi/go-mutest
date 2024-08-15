package numbers

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorNumbersIncrementer(t *testing.T) {
	test.Mutator(
		t,
		MutatorNumbersIncrementer,
		"../../testdata/numbers/incrementer.go",
		2,
	)
}
