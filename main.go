package main

import (
	"fmt"
	"htestp/DSL/parser"
)

func main() {

	source :=
		`
StaticNode("https://github.com", GET, ExistConstraint(),
	StaticNode()
)
		`

	scanner := parser.CreateScanner(source)
	scanner.Scan()

	fmt.Printf("%s", scanner.ToString())
	testScanner := parser.CreateScanner(scanner.ToString())
	testScanner.Scan()
	fmt.Printf("%+v", testScanner.ToString() == scanner.ToString())
	//program end
}
