package main

import (
	"fmt"
	"htestp/DSL/scanner"
)

func main() {

	source :=
		`
StaticNode("https://github.com", GET, ExistConstraint(),
	StaticNode()
)
		`

	scanner := scanner.CreateScanner(source)
	scanner.Scan()

	fmt.Printf("%s", scanner.ToString())
	testScanner := scanner.CreateScanner(scanner.ToString())
	testScanner.Scan()
	fmt.Printf("%+v", testScanner.ToString() == scanner.ToString())
	//program end
}
