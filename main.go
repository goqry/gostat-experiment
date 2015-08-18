package main

import (
	"fmt"
	//"go/ast"
	"go/parser"
	"go/token"
		"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	basePath := "."
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}
	fmt.Println(basePath)
	var lines int64 = 0
	var fileCount int64 = 0
	filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		fset := token.NewFileSet()
		_, err = parser.ParseFile(fset, path, nil, 0)
		//		fmt.Println(path)
		if err != nil {
			return nil
		}

				fileData, err := ioutil.ReadFile(path)
				if err != nil {
					return nil
				}
		fileCount++
		lines += int64(len(strings.Split(string(fileData), "\n")))


		/*ast.Inspect(file, func(node ast.Node) bool {
			if _, ok := node.(*ast.CallExpr); ok {
				count++
				//				startPos, endPos := fset.Position(f.Pos()), fset.Position(f.End())
				//				fmt.Println(string(fileData[startPos.Offset:endPos.Offset]))
			}
			return true
		})*/
		return nil
	})

	fmt.Printf("processed %d files, found %d lines\n", fileCount, lines)
}
