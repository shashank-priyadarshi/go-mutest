package loop

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorLoopBreak(t *testing.T) {
	test.Mutator(
		t,
		MutatorLoopBreak,
		"../../testdata/loop/break.go",
		2,
	)
}
