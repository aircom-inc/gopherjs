package util

import (
	"go/ast"

	"golang.org/x/tools/go/types"
)

func RemoveParens(e ast.Expr) ast.Expr {
	for {
		p, isParen := e.(*ast.ParenExpr)
		if !isParen {
			return e
		}
		e = p.X
	}
}

func IsJsPackage(pkg *types.Package) bool {
	return pkg != nil && pkg.Path() == "github.com/gopherjs/gopherjs/js"
}

func IsJsObject(t types.Type) bool {
	named, isNamed := t.(*types.Named)
	return isNamed && IsJsPackage(named.Obj().Pkg()) && named.Obj().Name() == "Object"
}

func SetType(info *types.Info, t types.Type, e ast.Expr) ast.Expr {
	info.Types[e] = types.TypeAndValue{Type: t}
	return e
}

func NewIdent(name string, t types.Type, info *types.Info, pkg *types.Package) *ast.Ident {
	ident := ast.NewIdent(name)
	info.Types[ident] = types.TypeAndValue{Type: t}
	obj := types.NewVar(0, pkg, name, t)
	info.Uses[ident] = obj
	return ident
}
