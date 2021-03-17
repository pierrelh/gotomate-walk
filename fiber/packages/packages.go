package packages

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Packages Export all the available packages & their functions
var Packages = getAutomatePackages()

// getAutomatePackages Get & return all the available packages & their functions
func getAutomatePackages() map[string][]string {
	const subPackage = "fiber/packages"
	packs := make(map[string][]string)

	filepath.Walk(subPackage, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "" && filepath.Base(path) != "packages" {
			packs[filepath.Base(path)] = []string{}
		} else {
			if filepath.Base(path) == "functions.go" {
				set := token.NewFileSet()
				pack, err := parser.ParseFile(set, path, nil, 0)
				if err != nil {
					fmt.Println("GOTOMATE FATAL ERROR: Unable to parse the file", err)
					os.Exit(1)
				}

				for _, d := range pack.Decls {
					if fn, isFn := d.(*ast.FuncDecl); isFn {
						pathSlice := strings.Split(filepath.Dir(path), "\\")
						directory := pathSlice[len(pathSlice)-1]
						packs[directory] = append(packs[directory], fn.Name.String())
					}
				}
			}
		}
		return nil
	})
	return packs
}
