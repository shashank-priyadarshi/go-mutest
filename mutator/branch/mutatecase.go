package branch

import (
	"go/ast"
	"go/types"

	"github.com/visu-suganya/go-mutesting/astutil"
	"github.com/visu-suganya/go-mutesting/mutator"
)

func init() {
	mutator.Register("branch/case", MutatorCase)
}

// MutatorCase implements a mutator for case clauses.
func MutatorCase(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.CaseClause)
	if !ok {
		return nil
	}

	old := n.Body

	if !mutator.CheckForPanic(old) {
		return nil
	}
	return []mutator.Mutation{
		{
			Change: func() {
				n.Body = []ast.Stmt{
					astutil.CreateNoopOfStatements(pkg, info, n.Body),
				}
			},
			Reset: func() {
				n.Body = old
			},
		},
	}
}
