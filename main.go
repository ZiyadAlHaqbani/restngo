package main

import (
	"fmt"
	"htestp/dsl/scanner"
)

func main() {

	source :=
		`
StaticNode("https://github.com", GET, ExistConstraint(),
	StaticNode(), 1000
)
		`

	s := scanner.CreateScanner(source)
	s.Scan()

	testScanner := scanner.CreateScanner(s.ToString())
	testScanner.Scan()
	fmt.Printf("%s\n", s.ToString())
	fmt.Printf("%s\n", testScanner.ToString())
	fmt.Printf("%+v", testScanner.ToString() == s.ToString())
	//program end
}
