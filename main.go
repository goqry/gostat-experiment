package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var _ = reflect.Array

func test() {
	ast.Bad.String()
}

type Counter struct {
	primitive, selected int64
}

func Walk(b interface{})          {}
func walkIdentList(b interface{}) {}
func walkExprList(b interface{})  {}
func walkDeclList(b interface{})  {}
func walkStmtList(b interface{})  {}

func (c *Counter) Visit(node ast.Node) ast.Visitor {
	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		// nothing to do

	case *ast.CommentGroup:
		for _, c := range n.List {
			Walk(c)
		}

	case *ast.Field:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		walkIdentList(n.Names)
		Walk(n.Type)
		if n.Tag != nil {
			Walk(n.Tag)
		}
		if n.Comment != nil {
			Walk(n.Comment)
		}

	case *ast.FieldList:
		for _, f := range n.List {
			Walk(f)
		}

	// Expressions
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
		// nothing to do

	case *ast.Ellipsis:
		if n.Elt != nil {
			Walk(n.Elt)
		}

	case *ast.FuncLit:
		Walk(n.Type)
		Walk(n.Body)

	case *ast.CompositeLit:
		if n.Type != nil {
			Walk(n.Type)
		}
		walkExprList(n.Elts)

	case *ast.ParenExpr:
		Walk(n.X)

	case *ast.SelectorExpr:
		Walk(n.X)
		Walk(n.Sel)

	case *ast.IndexExpr:
		Walk(n.X)
		Walk(n.Index)

	case *ast.SliceExpr:
		Walk(n.X)
		if n.Low != nil {
			Walk(n.Low)
		}
		if n.High != nil {
			Walk(n.High)
		}
		if n.Max != nil {
			Walk(n.Max)
		}

	case *ast.TypeAssertExpr:
		Walk(n.X)
		if n.Type != nil {
			Walk(n.Type)
		}

	case *ast.CallExpr:
		Walk(n.Fun)
		walkExprList(n.Args)

	case *ast.StarExpr:
		Walk(n.X)

	case *ast.UnaryExpr:
		Walk(n.X)

	case *ast.BinaryExpr:
		Walk(n.X)
		Walk(n.Y)

	case *ast.KeyValueExpr:
		Walk(n.Key)
		Walk(n.Value)

	// Types
	case *ast.ArrayType:
		if n.Len != nil {
			Walk(n.Len)
		}
		Walk(n.Elt)

	case *ast.StructType:
		Walk(n.Fields)

	case *ast.FuncType:
		if n.Params != nil {
			Walk(n.Params)
		}
		if n.Results != nil {
			Walk(n.Results)
		}

	case *ast.InterfaceType:
		Walk(n.Methods)

	case *ast.MapType:
		Walk(n.Key)
		Walk(n.Value)

	case *ast.ChanType:
		Walk(n.Value)

	// Statements
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		Walk(n.Decl)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		Walk(n.Label)
		Walk(n.Stmt)

	case *ast.ExprStmt:
		Walk(n.X)

	case *ast.SendStmt:
		Walk(n.Chan)
		Walk(n.Value)

	case *ast.IncDecStmt:
		Walk(n.X)

	case *ast.AssignStmt:
		walkExprList(n.Lhs)
		walkExprList(n.Rhs)

	case *ast.GoStmt:
		Walk(n.Call)

	case *ast.DeferStmt:
		Walk(n.Call)

	case *ast.ReturnStmt:
		walkExprList(n.Results)

	case *ast.BranchStmt:
		if n.Label != nil {
			Walk(n.Label)
		}

	case *ast.BlockStmt:
		walkStmtList(n.List)

	case *ast.IfStmt:
		if n.Init != nil {
			Walk(n.Init)
		}
		Walk(n.Cond)
		Walk(n.Body)
		if n.Else != nil {
			Walk(n.Else)
		}

	case *ast.CaseClause:
		walkExprList(n.List)
		walkStmtList(n.Body)

	case *ast.SwitchStmt:
		if n.Init != nil {
			Walk(n.Init)
		}
		if n.Tag != nil {
			Walk(n.Tag)
		}
		Walk(n.Body)

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			Walk(n.Init)
		}
		Walk(n.Assign)
		Walk(n.Body)

	case *ast.CommClause:
		if n.Comm != nil {
			Walk(n.Comm)
		}
		walkStmtList(n.Body)

	case *ast.SelectStmt:
		Walk(n.Body)

	case *ast.ForStmt:
		if n.Init != nil {
			Walk(n.Init)
		}
		if n.Cond != nil {
			Walk(n.Cond)
		}
		if n.Post != nil {
			Walk(n.Post)
		}
		Walk(n.Body)

	case *ast.RangeStmt:
		if n.Key != nil {
			Walk(n.Key)
		}
		if n.Value != nil {
			Walk(n.Value)
		}
		Walk(n.X)
		Walk(n.Body)

	// Declarations
	case *ast.ImportSpec:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		if n.Name != nil {
			Walk(n.Name)
		}
		Walk(n.Path)
		if n.Comment != nil {
			Walk(n.Comment)
		}

	case *ast.ValueSpec:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		walkIdentList(n.Names)
		if n.Type != nil {
			Walk(n.Type)
		}
		walkExprList(n.Values)
		if n.Comment != nil {
			Walk(n.Comment)
		}

	case *ast.TypeSpec:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		Walk(n.Name)
		Walk(n.Type)
		if n.Comment != nil {
			Walk(n.Comment)
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		for _, s := range n.Specs {
			Walk(s)
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		if n.Recv != nil {
			Walk(n.Recv)
		}
		Walk(n.Name)
		Walk(n.Type)
		if n.Body != nil {
			Walk(n.Body)
		}

	// Files and packages
	case *ast.File:
		if n.Doc != nil {
			Walk(n.Doc)
		}
		Walk(n.Name)
		walkDeclList(n.Decls)
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		for _, f := range n.Files {
			Walk(f)
		}
	}

	if node != nil {
		fmt.Printf("%+v\n", node)
	}

	return c
}

func main() {
	basePath := "."
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}
	fileCount := int64(0)
	c := Counter{}
	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return nil
		}

		fileData, err := ioutil.ReadFile(path)
		if err != nil {
			return nil
		}
		_ = fileData
		fileCount++

		ast.Walk(&c, file)
		return nil
	})

	fmt.Printf("processed %d files, found %d func tokens out of %d tokens in %q\n", fileCount, c.selected, c.primitive, basePath)
}
