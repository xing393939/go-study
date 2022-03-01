package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// read file
	// here you can filepath.Walk() for your go files
	gopath := os.ExpandEnv("$GOPATH")
	fname := gopath + "/src/github.com/codegangsta/martini-contrib/web/web.go"

	// read file
	file, err := os.Open(fname)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// read the whole file in
	srcbuf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	src := string(srcbuf)

	// file set
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "lib.go", src, 0)
	if err != nil {
		log.Println(err)
		return
	}

	// main inspection
	ast.Inspect(f, func(n ast.Node) bool {

		switch fn := n.(type) {

		// catching all function declarations
		// other intersting things to catch FuncLit and FuncType
		case *ast.FuncDecl:
			fmt.Print("func ")

			// if a method, explore and print receiver
			if fn.Recv != nil {
				fmt.Printf("(%s)", fields(*fn.Recv))
			}

			// print actual function name
			fmt.Printf("%v", fn.Name)

			// print function parameters
			if fn.Type.Params != nil {
				fmt.Printf("(%s)", fields(*fn.Type.Params))
			}

			// print return params
			if fn.Type.Results != nil {
				fmt.Printf("(%s)", fields(*fn.Type.Results))
			}

			fmt.Println()

		}
		return true
	})
}

func expr(e ast.Expr) (ret string) {
	switch x := e.(type) {
	case *ast.StarExpr:
		return fmt.Sprintf("%s*%v", ret, x.X)
	case *ast.Ident:
		return fmt.Sprintf("%s%v", ret, x.Name)
	case *ast.ArrayType:
		if x.Len != nil {
			log.Println("OH OH looks like homework")
			return "TODO: HOMEWORK"
		}
		res := expr(x.Elt)
		return fmt.Sprintf("%s[]%v", ret, res)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", expr(x.Key), expr(x.Value))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", expr(x.X), expr(x.Sel))
	default:
		fmt.Printf("\nTODO HOMEWORK: %#v\n", x)
	}
	return
}

func fields(fl ast.FieldList) (ret string) {
	pcomma := ""
	for i, f := range fl.List {
		// get all the names if present
		var names string
		ncomma := ""
		for j, n := range f.Names {
			if j > 0 {
				ncomma = ", "
			}
			names = fmt.Sprintf("%s%s%s ", names, ncomma, n)
		}
		if i > 0 {
			pcomma = ", "
		}
		ret = fmt.Sprintf("%s%s%s%s", ret, pcomma, names, expr(f.Type))
	}
	return ret
}
