package main

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
)

func main() {
	f := jen.NewFile("main")

	f.ImportName("example.com/assignment3/pkg/dataloading", "dataloading")
	f.ImportName("example.com/assignment3/pkg/regression", "regression")

	f.Func().Id("main").Params().Block(
		jen.Id("startTime").Op(":=").Qual("time", "Now").Call(),
		jen.For(jen.Id("i").Op(":=").Lit(0), jen.Id("i").Op("<").Lit(100), jen.Id("i").Op("++")).Block(
			jen.List(jen.Id("datasets"), jen.Id("err")).Op(":=").Qual("dataloading", "LoadCSV").Call(jen.Lit("C:/Assignment2/dataset/anscombes.csv")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Qual("log", "Fatalf").Call(jen.Lit("Error loading CSV file: %v"), jen.Id("err")),
			),
			jen.For(jen.List(jen.Id("name"), jen.Id("data")).Op(":=").Range().Id("datasets")).Block(
				jen.If(jen.Id("err").Op(":=").Qual("regression", "PlotRegression").Call(jen.Id("name"), jen.Id("data")), jen.Id("err").Op("!=").Nil()).Block(
					jen.Qual("log", "Fatalf").Call(jen.Lit("could not plot dataset %s: %v"), jen.Id("name"), jen.Id("err")),
				),
			),
		),
		jen.Id("elapsed").Op(":=").Qual("time", "Now").Call().Dot("Sub").Call(jen.Id("startTime")),
		jen.Qual("fmt", "Printf").Call(jen.Lit("Execution took %v\n"), jen.Id("elapsed")),
	)

	err := f.Save("main_gen.go")
	if err != nil {
		fmt.Printf("Failed to save the generated code: %s\n", err)
		os.Exit(1)
	}
}
