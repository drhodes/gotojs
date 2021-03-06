package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

var (
	dir = flag.String("path", "./", "the path of the project")
)

type Trans struct {
}

type Package ast.Package

func (self Package) Show() {
	//f := "// package %s\n"
	//fmt.Printf(f, self.Name)
}

type ReturnStmt ast.ReturnStmt

func (self ReturnStmt) Show() string {
	s := "return "
	exps := []string{}

	for _, xp := range self.Results {
		exps = append(exps, ShowExpr(xp))
	}
	if len(exps) == 1 {
		return s + exps[0]
	}
	if len(exps) > 1 {
		return s + "[" + strings.Join(exps, ",") + "]"
	}
	return s
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
			return result + "[" + args + "]"
		}
	}
	return "new " + ShowExpr(self.Type) + "(" + args + ")"
}

type ArrayType ast.ArrayType

func (self ArrayType) Show() string {
	log.Println(ShowExpr(self.Elt))
	log.Println(ShowExpr(self.Len))
	return "array-type-show"
}

type StarExpr ast.StarExpr

func (self StarExpr) Show() string {
	// FULLSTOP UNTIL exp/types is done.
	return ShowExpr(self.X)
}

type IndexExpr ast.IndexExpr

func (self IndexExpr) Show() string {
	idx := ShowExpr(self.Index)
	return ShowExpr(self.X) + "[" + idx + "]"
}

type SliceExpr ast.SliceExpr

func (self SliceExpr) Show() string {
	const temp = "%s.slice(%s,%s)"
	x := ShowExpr(self.X)
	low := "0"
	if self.Low != nil {
		low = ShowExpr(self.Low)
	}
	high := fmt.Sprintf("%s.length", x)
	if self.High != nil {
		high = ShowExpr(self.High)
	}
	return fmt.Sprintf(temp, x, low, high)
}

type FuncLit ast.FuncLit

func (self FuncLit) Show() string {
	const temp = "function%s%s"
	t := FuncType(*self.Type).Show()
	b := BlockStmt(*self.Body).Show()
	return fmt.Sprintf(temp, t, b)
}

type UnaryExpr ast.UnaryExpr
func (self UnaryExpr) Show() string {
	return self.Op.String() + ShowExpr(self.X)
}

type ParenExpr ast.ParenExpr
func (self ParenExpr) Show() string {	
	const temp = "(%s)"
	return fmt.Sprintf(temp, ShowExpr(self.X))
}

type TypeAssertExpr ast.TypeAssertExpr
func (self TypeAssertExpr) Show() string {
	const temp = "%s.(%s)"
	return fmt.Sprintf(temp, ShowExpr(self.X), ShowExpr(self.Type))
}

type KeyValueExpr ast.KeyValueExpr
func (self KeyValueExpr) Show() string {
	const temp = "[%s, %s]"
	key := ShowExpr(self.Key)
	val := ShowExpr(self.Value)
	return fmt.Sprintf(temp, key, val)
}

type MapType ast.MapType
func (self MapType) Show() string {
	const temp = "MapObj"
	return fmt.Sprint(temp)
}

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
	case *ast.StarExpr:
		return StarExpr(*e.(*ast.StarExpr)).Show()
	case *ast.IndexExpr:
		return IndexExpr(*e.(*ast.IndexExpr)).Show()
	case *ast.SliceExpr:
		return SliceExpr(*e.(*ast.SliceExpr)).Show()
	case *ast.FuncLit:
		return FuncLit(*e.(*ast.FuncLit)).Show()
	case *ast.UnaryExpr:
		return UnaryExpr(*e.(*ast.UnaryExpr)).Show()
	case *ast.ParenExpr:
		return ParenExpr(*e.(*ast.ParenExpr)).Show()
	case *ast.TypeAssertExpr:
		return TypeAssertExpr(*e.(*ast.TypeAssertExpr)).Show()
	case *ast.KeyValueExpr:
		return KeyValueExpr(*e.(*ast.KeyValueExpr)).Show()
	case *ast.MapType:
		return MapType(*e.(*ast.MapType)).Show()

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
	result := fmt.Sprintf(ce,
		ShowExpr(self.Fun.(ast.Expr)),
		ExprList(self.Args).Show(),
	)
	if strings.HasPrefix(result, "append") {
		return "lib." + result
	}
	return result
}

type ExprStmt ast.ExprStmt

func (self ExprStmt) Show() string {
	return ShowExpr(self.X)
	// switch (self.X).(type) {
	// case *ast.CallExpr:
	// 	return ShowExpr(self.X)
	// }
	// return "ExprStmt.Show fails to handle: " + fmt.Sprintf("%T", self.X)
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
		result += "else" + ShowStmt(self.Else)
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

func SingleAssign(stmt AssignStmt) (bool, string) {
	// safe assumption: len(self.lhs) == len(self.rhs) == 1
	e := stmt.Lhs[0]
	switch e.(type) {
	case *ast.StarExpr:
		log.Panic("Pointers aren't ready yet, can't use em. Sorry.")
	}

	// TODO: if Rhs is callexpr, if so, is append? ok
	// so need to find a way to include js src.
	tok := stmt.Tok.String()
	if tok == ":=" {
		tok = "="
	}

	left := ShowExpr(stmt.Lhs[0])
	hasdot := strings.Contains(left, ".")

	right := ShowExpr(stmt.Rhs[0])
	tok = " " + tok + " "
	return hasdot, fmt.Sprintf("%s %s %s", left, tok, right)
}

type AssignStmt ast.AssignStmt

func (self AssignStmt) Show() string {
	// js doesn't have multiple assign. So
	// val, err = expr()
	// goes to
	// gensym = expr()
	// val = gensym[0]
	// err = gensym[1]
	// if left hand side is . identifier

	switch {
	case len(self.Lhs) == 1 && len(self.Rhs) == 1:
		hasdot, result := SingleAssign(self)
		if hasdot {
			return result
		} else if self.Tok.String() == ":=" {
			return "var " + result
		} else {
			return result
		}
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
	return self.Tok.String() + ";"
}

type SwitchStmt ast.SwitchStmt

func (self SwitchStmt) Show() string {
	cond := ""
	if self.Tag != nil {
		cond = ShowExpr(self.Tag)
	}
	temp := `switch (%s) %s`
	body := ShowStmt(self.Body)
	if self.Tag == nil {
		return fmt.Sprintf(temp, "true", body)
	}
	return fmt.Sprintf(temp, cond, body)
}

type CaseClause ast.CaseClause

func (self CaseClause) Show() string {
	temp := "case %s: %s"
	list := ExprList(self.List).Show()
	body := StmtList(self.Body).Show()
	return fmt.Sprintf(temp, list, body) + "break;"
}

type DeclStmt ast.DeclStmt
func (self DeclStmt) Show() string {
	return "DeclStmt not implemented"
} 

type DeferStmt ast.DeferStmt
func (self DeferStmt) Show() string {
	const temp = "__defer_stack.push(function(){%s})"
	return fmt.Sprintf(temp, ShowExpr(self.Call))
}

type TypeSwitchStmt ast.TypeSwitchStmt
func (self TypeSwitchStmt) Show() string {
	temp := "(typeswitch body:%s init:%s assign:%s)"
	return fmt.Sprintf(temp, ShowStmt(self.Body), 
		ShowStmt(self.Init), ShowStmt(self.Assign))
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
	case *ast.SwitchStmt:
		return SwitchStmt(*s.(*ast.SwitchStmt)).Show()
	case *ast.CaseClause:
		return CaseClause(*s.(*ast.CaseClause)).Show()
	case *ast.DeclStmt:
		return DeclStmt(*s.(*ast.DeclStmt)).Show()
	case *ast.DeferStmt:
		return DeferStmt(*s.(*ast.DeferStmt)).Show()	
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt(*s.(*ast.TypeSwitchStmt)).Show()

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
	ident := Ident(*self.Name).Show()
	ftype := FuncType(*self.Type).Show()
	//bstmt := BlockStmt(*self.Body).Show()
	recv := ""

	if self.Recv != nil {
		if len(self.Recv.List) != 0 {
			// Point.prototype.Add = function
			f := "%s.%s = function %s %s "
			typ := self.Recv.List[0]
			recv = ShowExpr(typ.Type) + ".prototype"
			result := fmt.Sprintf(f, recv, ident, ftype, "%s")//bstmt)
			return generateDeferClosure(self, result)
		}
	}
	f := "var %s = function %s %s %s"
	result := fmt.Sprintf(f, ident, recv, ftype, "%s")//bstmt)
	return generateDeferClosure(self, result)
}

type Field ast.Field

func (self Field) Show() string {
	xs := []string{}
	for _, f := range self.Names {

		id := Ident(*f).Show()
		if id != "" {
			xs = append(xs, Ident(*f).Show())
		}
	}
	return strings.Join(xs, ", ")
}

type Fields []*ast.Field

func (self Fields) Show() string {
	xs := []string{}
	for _, f := range self {
		x := Field(*f).Show()
		if x != "" {
			xs = append(xs, x)
		}
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
			flds = append(flds, "this."+n.String()+"="+n.String()+";")
		}
	}
	return strings.Join(flds, "\n")
}

type StructType ast.StructType

func (self StructType) Show() string {
	const class = "function %s%s{%s}"
	args := FieldList(*self.Fields).Show()
	flds := StructList(*self.Fields).Show()
	return fmt.Sprintf(class, "__STRUCT__NAME__", args, flds)
}

type TypeSpec ast.TypeSpec

func (self TypeSpec) Show() string {
	result := ""
	switch self.Type.(type) {
	case *ast.StructType:
		result = StructType(*self.Type.(*ast.StructType)).Show()
		return strings.Replace(result, "__STRUCT__NAME__", self.Name.String(), -1)
	}
	return result
}

type ValueSpec ast.ValueSpec

func (self ValueSpec) Show(dec string) string {
	const temp = "%s %s %s = %s"
	result := []string{}

	for i := range self.Names {
		t := ""
		if self.Type != nil {
			t = ShowExpr(self.Type)
		}
		name := self.Names[i].String()
		val := ShowExpr(self.Values[i])
		r := fmt.Sprintf(temp, dec, t, name, val) + ";"
		result = append(result, r)
	}
	return strings.Join(result, "\n")
}

type GenDecl ast.GenDecl

func (self GenDecl) Show() string {
	result := []string{}
	for _, spec := range self.Specs {
		switch spec.(type) {
		case *ast.ValueSpec:
			vs := ValueSpec(*spec.(*ast.ValueSpec))
			result = append(result, vs.Show(self.Tok.String()))
		case *ast.ImportSpec:
			continue //result = append(Result, "")
		case *ast.TypeSpec:
			ts := TypeSpec(*spec.(*ast.TypeSpec))
			result = append(result, ts.Show())
		default:
			return "unhandled GenDecl in func GenDecl.Show: " + fmt.Sprintf("%T", spec)
		}
	}
	return strings.Join(result, "\n")
}

func (self Trans) Visit(node ast.Node) (w ast.Visitor) {
	if node != nil {
		switch node.(type) {
		case *ast.FuncDecl:
			f := FuncDecl(*node.(*ast.FuncDecl))
			fmt.Println(f.Show())

		case *ast.Package:
			pkg := Package(*node.(*ast.Package))
			pkg.Show()

		// case *ast.TypeSpec:
		// 	ts := TypeSpec(*node.(*ast.TypeSpec))
		// 	fmt.Println(ts.Show())

		case *ast.GenDecl:
			fmt.Println(GenDecl(*node.(*ast.GenDecl)).Show())
		default:
			//log.Println("Not handled")
		}
	}
	return self
}

func filter(info os.FileInfo) bool {
	if strings.Contains(info.Name(), "console.go") {
		return false
	}
	if strings.HasSuffix(info.Name(), "go") && (!strings.HasPrefix(info.Name(), ".")) {
		return true
	}
	return false
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
		ast.Walk(normalizer{}, pk)		
		ast.Walk(ts, pk)
	}
	fmt.Println("main();")
}

func main() {
	flag.Parse()
	// build the files in the directory.
	log.Println(*dir)
	pks := parse(*dir)
	fmt.Println("var lib = require('../../lib.js');")
	trans(pks)
}
