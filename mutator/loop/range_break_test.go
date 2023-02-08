package loop

import (
	"testing"

	"github.com/visu-suganya/go-mutesting/test"
)

func TestMutatorLoopRangeBreak(t *testing.T) {
	test.Mutator(
		t,
		MutatorLoopRangeBreak,
		"../../testdata/loop/range_break.go",
		2,
	)
}
