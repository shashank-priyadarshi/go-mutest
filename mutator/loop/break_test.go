package loop

import (
	"testing"

	"github.com/visu-suganya/go-mutesting/test"
)

func TestMutatorLoopBreak(t *testing.T) {
	test.Mutator(
		t,
		MutatorLoopBreak,
		"../../testdata/loop/break.go",
		2,
	)
}
