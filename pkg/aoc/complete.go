package aoc

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
)

func Complete(year string, day string) {
	// Delete input and sample files
	os.Remove("input")
	os.Remove("sample")

	// Copy solve.go to correct location
	os.MkdirAll(fmt.Sprintf("%s", year), 0644)
	os.Rename("pkg/aoc/solve.go", fmt.Sprintf("%s/day%s.go", year, day))

	// Reset the solve.go file using jen
	generateFile()

	fmt.Printf("Completed puzzle archived to %s/day%s.go 🗄️\n", year, day)
}

func generateFile() error {
	f := NewFile("aoc")
	f.Func().Id("Solve").Params(Id("filename").String()).Block(
		Id("file").Op(",").Id("_").Op(":=").Qual("os", "ReadFile").Call(Id("filename")),
		Id("data").Op(":=").Qual("strings", "Split").Call(Qual("strings", "TrimSpace").Call(String().Call(Id("file"))), Lit("\n")),
		Qual("fmt", "Println").Call(Id("data")),
	)

	f.Save("pkg/aoc/solve.go")

	return nil
}
