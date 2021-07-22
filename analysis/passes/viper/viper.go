package viper

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "viper",
	Doc:  "check for viper.Sub and viper.UnmarshalKey usage",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var viperPackage = "github.com/spf13/viper"

func run(pass *analysis.Pass) (interface{}, error) {
	// Fast path: bypass check if file doesn't use viper.
	if !hasImport(pass.Pkg, viperPackage) {
		return nil, nil
	}

	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeTypes := []ast.Node{
		(*ast.Ident)(nil),
	}

	i.Preorder(nodeTypes, func(n ast.Node) {
		if identifier, ok := n.(*ast.Ident); ok {
			if function, ok := pass.TypesInfo.Uses[identifier].(*types.Func); ok {
				switch function.Name() {
				case "Sub", "UnmarshalKey":
					if function.Pkg().Path() == viperPackage {
						pass.ReportRangef(identifier, "avoid usage of viper.%s as it doesn't work well with environment variables", function.Name())
					}
				}
			}
		}
	})

	return nil, nil
}

func hasImport(pkg *types.Package, path string) bool {
	for _, imp := range pkg.Imports() {
		if imp.Path() == path {
			return true
		}
	}
	return false
}
