package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

func main() {
	src := `package p
const c = 1.0
var X = f(3.14)*2 + c

func add(x,y int)int{
	return x + y
}

type Op struct{
	A  string
	B  string
}

func(o *Op)Add(x,y int)int{
	return x + y
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	// Print the AST.
	//f.Decls = f.Decls[:len(f.Decls)-1]
	ast.Print(fset, f)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	fmt.Printf("src.go : \n%s", buf.Bytes())
}
