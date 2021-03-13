package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	i, err := openAndCount("example.txt", "bla123")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println("Calls count:", i)
}

func openAndCount(filename string, funcName string) (int, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		panic(err)
	}

	count := 0
	startCount := false
	positionStart := 0
	positionEnd := 0

	cb := func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GoStmt:
			if startCount && positionStart < int(n.Pos()) && positionEnd > int(n.End()) {
				count++
			}
			return false
		case *ast.Ident:
			if x.Obj != nil && x.Obj.Kind == ast.Fun && x.Name == funcName {
				startCount = true
				return true
			}
			if positionEnd < int(n.Pos()) {
				startCount = false
			}
		case *ast.BlockStmt:
			if startCount {
				positionStart = int(x.Lbrace)
				positionEnd = int(x.Rbrace)
			}

		}

		return true

	}

	ast.Inspect(f, cb)
	return count, nil
}
