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

// Package Declare the structure of a package
type Package struct {
	Name      string
	Functions []string
}

// Packages Export all the available packages & their functions
var Packages = getAutomatePackages()

// getAutomatePackages Get & return all the available packages & their functions
func getAutomatePackages() []*Package {
	const subPackage = "fiber/packages"
	var packs []*Package

	filepath.Walk(subPackage, func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) == "functions.go" {
			pathSlice := strings.Split(filepath.Dir(path), "\\")
			directory := pathSlice[len(pathSlice)-1]
			newPackage := &Package{
				Name: directory,
			}
			set := token.NewFileSet()
			pack, err := parser.ParseFile(set, path, nil, 0)
			if err != nil {
				fmt.Println("GOTOMATE FATAL ERROR: Unable to parse the file", err)
				os.Exit(1)
			}

			for _, d := range pack.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					newPackage.Functions = append(newPackage.Functions, fn.Name.String())
				}
			}
			packs = append(packs, newPackage)
		}
		return nil
	})
	return packs
}
