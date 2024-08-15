package loop

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorLoopCondition(t *testing.T) {
	test.Mutator(
		t,
		MutatorLoopCondition,
		"../../testdata/loop/condition.go",
		2,
	)
}
