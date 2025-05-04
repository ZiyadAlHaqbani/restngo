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

	s := scanner.CreateScanner(source)
	s.Scan()

	fmt.Printf("%s", s.ToString())
	testScanner := scanner.CreateScanner()
	fmt.Printf("%+v", testScanner.ToString() == s.ToString())
	//program end
}
