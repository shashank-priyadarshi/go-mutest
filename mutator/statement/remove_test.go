package statement

import (
	"testing"

	"github.com/shashank-priyadarshi/go-mutest/test"
)

func TestMutatorRemoveStatement(t *testing.T) {
	test.Mutator(
		t,
		MutatorRemoveStatement,
		"../../testdata/statement/remove.go",
		17,
	)
}
