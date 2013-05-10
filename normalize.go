package main

import (
	"go/ast"
	"log"
)

// javascript objects uses "this" for all receivers, 
// go methods may name any identifier for recievers.
// the code below renames all idents in go methods that 
// refer to the receiver to "this".

type namechanger string
func (m namechanger) Visit(n ast.Node) ast.Visitor {

	switch n.(type) {
	case *ast.Ident:
		log.Println(n)

		f := n.(*ast.Ident)
		if f.Name == string(m) {
			f.Name = "this"
		}
	}
	return m
}

type normalizer struct{}
func (m normalizer) Visit(n ast.Node) ast.Visitor {
	switch n.(type) {
	case *ast.FuncDecl:
		f := n.(*ast.FuncDecl)
		// if we're in a method
		if f.Recv != nil { 
			// dive into the idents and change all names of the 
			// receiver to "this"
			fld := f.Recv.List[0]
			rname := fld.Names[0].Name			
			ast.Walk(namechanger(rname), n)
		}
	}
	return m
}










