package main

import (
	"github.com/dave/jennifer/jen"
)

func main() {
	f := jen.NewFile("main")

	f.Func().Id("TestLoadCSV").Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.List(jen.Id("datasets"), jen.Id("err")).Op(":=").Qual("example.com/assignment3-main", "LoadCSV").Call(jen.Lit("C:/Assignment2/dataset/anscombes.csv")),
		jen.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Id("t").Dot("Errorf").Call(jen.Lit("loadCSV() returned an error: %v"), jen.Id("err")),
			jen.Qual("fmt", "Println").Call(jen.Lit("datasets"))),
	)

	// Save the generated code
	err := f.Save("loadcsv_test.go") // The file name should end with _test.go
	if err != nil {
		panic(err)
	}
}
