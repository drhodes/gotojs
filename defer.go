package main

import (
	"go/ast"
	"fmt"
)

type deferLocator struct {
	found *bool
}
func (self deferLocator) Visit(n ast.Node) ast.Visitor {	
	switch n.(type) {
	case *ast.DeferStmt:
		*self.found = true
	}	
	return self
}

func generateDeferClosure(f FuncDecl, missingBody string) string {
	const bodyWrapper = `{
	var __defer_stack = [];
	var __retvals = function() %s();

	while (__defer_stack.length != 0) {
	 	__defer_stack.pop()();
    }
	return __retvals;
    }
    `

	dl := deferLocator{new(bool)}
	ast.Walk(dl, f.Body)
	if *dl.found {
		newbody := fmt.Sprintf(bodyWrapper, ShowStmt(f.Body))
		return fmt.Sprintf(missingBody, newbody)
	}
	return fmt.Sprintf(missingBody, ShowStmt(f.Body))
}

















