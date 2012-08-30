package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"fmt"
)

var (
	dir = flag.String("path", "./", "the path of the project")
)

type Trans struct {
}

type Package ast.Package
func (self Package) Show() {
	f := "// package %s\n"
	fmt.Printf(f, self.Name)
}

type ReturnStmt ast.ReturnStmt
func (self ReturnStmt) Show() string {
	s := "return "
	exps := []string{}
		
	for _, xp := range self.Results {
		exps = append(exps, ShowExpr(xp))
	}
	if len(exps) > 1 {
		return s + "[" + strings.Join(exps, ",") + "]"
	}
	return s + exps[0]
}

type BasicLit ast.BasicLit
func (self BasicLit) Show() string {
	return self.Value
}

type BinaryExpr ast.BinaryExpr
func (self BinaryExpr) Show() string {
	return ShowExpr(self.X) + self.Op.String() + ShowExpr(self.Y) 
}

type SelectorExpr ast.SelectorExpr
func (self SelectorExpr) Show() string {
	return ShowExpr(self.X) + "." + Ident(*self.Sel).Show() 
}

type CompositeLit ast.CompositeLit
func (self CompositeLit) Show() string {
	exps := []string{}
	for _, el := range self.Elts {
		exps = append(exps, ShowExpr(el))
	}
	args := strings.Join(exps, ", ")

	result := ""
	if self.Type != nil {
		switch self.Type.(type) {
		case *ast.ArrayType:
			return result + "["+args+"]"
		}
		
	}	
	return "new " + ShowExpr(self.Type) + "(" + args + ")"
	
}

type ArrayType ast.ArrayType
func (self ArrayType) Show() string {
	log.Println(ShowExpr(self.Elt))
	log.Println(ShowExpr(self.Len))
	return "asdf"
}

// -----------------------------------------------------------------------------
func ShowExpr(e ast.Expr) string {
	switch (e).(type) {
	case *ast.CallExpr:
		return CallExpr(*e.(*ast.CallExpr)).Show()		
	case *ast.Ident:
		return Ident(*e.(*ast.Ident)).Show()
	case *ast.BasicLit:
		return BasicLit(*e.(*ast.BasicLit)).Show()
	case *ast.BinaryExpr:
		return BinaryExpr(*e.(*ast.BinaryExpr)).Show()
	case *ast.SelectorExpr:
		return SelectorExpr(*e.(*ast.SelectorExpr)).Show()
	case *ast.CompositeLit:
		return CompositeLit(*e.(*ast.CompositeLit)).Show()	
	case *ast.ArrayType:
		return ArrayType(*e.(*ast.ArrayType)).Show()
	case *ast.StructType:
		return StructType(*e.(*ast.StructType)).Show()

	}	
	return "unhandled Expr in func ShowExpr: " + fmt.Sprintf("%T", e)
}

type ExprList []ast.Expr
func (self ExprList) Show() string {
	exps := []string{}
	for _, e := range self {
		exps = append(exps, ShowExpr(e))
	}
	return strings.Join(exps, ", ")
}

type CallExpr ast.CallExpr
func (self CallExpr) Show() string {
	const ce = "%s(%s)"
	return fmt.Sprintf(ce, 
		ShowExpr(self.Fun.(ast.Expr)),
		ExprList(self.Args).Show(),
	)
}

type ExprStmt ast.ExprStmt
func (self ExprStmt) Show() string {	
	switch (self.X).(type) {
	case *ast.CallExpr:
		return ShowExpr(self.X)
	}
	return "ExprStmt.Show fails to handle: " + fmt.Sprintf("%T", self.X)
}

type IfStmt ast.IfStmt
func (self IfStmt) Show() string {
	// The first %s is the initialization statement: http://goo.gl/ae5MV
	// it is optional.
	var result string

    // initialization statement; or nil	
	// TODO rethink this, possibly needs create a to go into closure
	// downside, maybe that would occlude scope locals.
	// downside, slow - maybe v8 would optimize it away.
	// might need to have have a cur_scope var up top to provide
	// lookup ability.  The Trans visitor would be resposible for 
	// updating that. could get ugly fast.
	if self.Init != nil {
		result += ShowStmt(self.Init) 
	}

	// condition
	result += "if (" + ShowExpr(self.Cond) + ") "
	result += BlockStmt(*self.Body).Show()

	// else branch; or nil		
	if self.Else != nil {
		result += "else" +  ShowStmt(self.Else) 
	}
	return result
}

type ForStmt ast.ForStmt
func (self ForStmt) Show() string {	
    // initialization statement or nil
	// Init ignore this beast until IfStmt knows how to do it
	temp := `for (%s %s %s)`
	init := ";"
	if self.Init != nil {
		init = ShowStmt(self.Init) 
	}
		
	// condition; or nil
	cond := ";"
	if self.Cond != nil {
		cond = ShowExpr(self.Cond) + ";"
	}

    // post iteration statement; or nil
	post := ""
	if self.Post != nil {
		post = ShowStmt(self.Post)
		post = post[:len(post)-1] // TODO this is bad.
		// What I'm doing is trimming the ; off the end for the forloop is
		// correct.  How to do this better?
	}

	body := ShowStmt(self.Body)
	
	return fmt.Sprintf(temp, init, cond, post) + body
}

var __cur_symb = 0
func gensym() string {
	__cur_symb += 1
	return fmt.Sprintf("__symb__%d", __cur_symb)
}

func SingleAssign(stmt AssignStmt) string {
	return ShowExpr(stmt.Lhs[0]) + " = " + ShowExpr(stmt.Rhs[0]) 
}

type AssignStmt ast.AssignStmt
func (self AssignStmt) Show() string {
	// js doesn't have multiple assign. So
	// val, err = expr()
	// goes to 
	// gensym = expr()
	// val = gensym[0]
	// err = gensym[1]
	
	switch {
	case len(self.Lhs) == 1 && len(self.Rhs) == 1:		
		return "var " + SingleAssign(self)		
	case len(self.Rhs) != 1:
		// return MultiAssign(self)
		return "MultiAssign not implemented yet"
	}
	return "AssignStmt missing case: " + fmt.Sprint(self)
}

type IncDecStmt ast.IncDecStmt
func (self IncDecStmt) Show() string {
	return ShowExpr(self.X) + self.Tok.String()
}

type RangeStmt ast.RangeStmt
func (self RangeStmt) Show() string {
	const temp = "for (var %s in %s) { \n var %s = %s[%s];"
	
	body := ShowStmt(self.Body)
	key := self.Key
	val := self.Value
	exp := ShowExpr(self.X)
	return fmt.Sprintf(temp, key, exp, val, exp, key) + body[1:]
}

type BranchStmt ast.BranchStmt
func (self BranchStmt) Show() string {
	return self.Tok.String() +";"
}

func ShowStmt(s ast.Stmt) string {
	switch (s).(type) {
	case *ast.ReturnStmt: 			
		return ReturnStmt(*(s.(*ast.ReturnStmt))).Show() + ";"
	case *ast.ExprStmt:     
		return ExprStmt(*(s.(*ast.ExprStmt))).Show() + ";"
	case *ast.IfStmt:
		return IfStmt(*(s.(*ast.IfStmt))).Show()
	case *ast.BlockStmt:
		return BlockStmt(*s.(*ast.BlockStmt)).Show() 
	case *ast.ForStmt:
		return ForStmt(*s.(*ast.ForStmt)).Show()
	case *ast.AssignStmt:
		return AssignStmt(*s.(*ast.AssignStmt)).Show() + ";"
	case *ast.IncDecStmt:
		return IncDecStmt(*s.(*ast.IncDecStmt)).Show() + ";"	
	case *ast.RangeStmt:
		return RangeStmt(*s.(*ast.RangeStmt)).Show()
	case *ast.BranchStmt:
		return BranchStmt(*s.(*ast.BranchStmt)).Show()
	}
	return "unhandled Stmt in func ShowStmt: " + fmt.Sprintf("%T", s)
}

type StmtList []ast.Stmt
func (self StmtList) Show() string {
	stmts := []string{}
	for _, s := range self {
		stmts = append(stmts, ShowStmt(s))
	}
	return strings.Join(stmts, "\n")
}

type BlockStmt ast.BlockStmt
func (self BlockStmt) Show() string {
	return fmt.Sprintf("{\n%s\n}", StmtList(self.List).Show())
}

type FuncType ast.FuncType
func (self FuncType) Show() string {
	return fmt.Sprintf("%s", FieldList(*self.Params).Show())
}

type FuncDecl ast.FuncDecl
func (self FuncDecl) Show() string {
	f := "function %s %s %s"		
	return fmt.Sprintf(f, 
		Ident(*self.Name).Show(), 
		FuncType(*self.Type).Show(),
		BlockStmt(*self.Body).Show(),
		//Recv(self.Recv).Show(),
		) 
} 

type Field ast.Field
func (self Field) Show() string {
	xs := []string{}
	for _, f := range self.Names {
		xs = append(xs, Ident(*f).Show())
	}
	return strings.Join(xs, ", ")
}

type Fields []*ast.Field
func (self Fields) Show() string {
	xs := []string{}
	for _, f := range self {
		xs = append(xs, Field(*f).Show())
	}
	return strings.Join(xs, ", ")
}

type FieldList ast.FieldList
func (self FieldList) Show() string {
	return fmt.Sprintf("(%s)", Fields(self.List).Show())
}

type Ident ast.Ident
func (self Ident) Show() string {
	return fmt.Sprint(self.Name)
}

type StructList ast.FieldList
func (self StructList) Show() string {
	flds := []string{}	
	for _, fld := range self.List {
		for _, n := range fld.Names {
			flds = append(flds, "this." + n.String() + "=" + n.String() + ";")
		}
	}
	return strings.Join(flds, "\n")
}

type StructType ast.StructType
func (self StructType) Show() string {
	const class = "function %s %s{%s}"
	args := FieldList(*self.Fields).Show()
	flds := StructList(*self.Fields).Show()
	return fmt.Sprintf(class, "asdf", args, flds)
}

type TypeSpec ast.TypeSpec
func (self TypeSpec) Show() string {
	return self.Name.String() + ShowExpr(self.Type)
}

func (self Trans) Visit(node ast.Node) (w ast.Visitor) {
	if node != nil {
		//log.Printf("%T", node)
		switch node.(type) {
		case *ast.FuncDecl:		
			f := FuncDecl(*node.(*ast.FuncDecl))
			fmt.Println(f.Show())
		case *ast.Package:		
			pkg := Package(*node.(*ast.Package))
			pkg.Show()
		case *ast.StructType:
			str := StructType(*node.(*ast.StructType))
			fmt.Println(str.Show())
		case *ast.TypeSpec:
			ts := TypeSpec(*node.(*ast.TypeSpec))
			fmt.Println(ts.Show())
		default: 
			//log.Println("Not handled")
		}
	}
	return self
}

func filter(info os.FileInfo) bool {
	return strings.HasSuffix(info.Name(), "go") && (!strings.HasPrefix(info.Name(), "."))
}

func parse(dir string) map[string]*ast.Package {
	fset := token.NewFileSet()
	// parse each file in directory
	pks, err := parser.ParseDir(fset, dir, filter, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return pks
}

func trans(pks map[string]*ast.Package) {
	ts := Trans{}
	for _, pk := range pks {
		ast.Walk(ts, pk)
	}	
}

func main() {
	flag.Parse()
	// build the files in the directory.
	log.Println(*dir)
	pks := parse(*dir)
	trans(pks)
}