package main

import (
	"fmt"
	"htestp/dsl/parser"
	"htestp/dsl/scanner"
)

func main() {

	source :=
		`
StaticNode("ID:123432", GET, "https://github.com", ExistConstraint("ID", STRING), ExistConstraint("Users", ARRAY),
	StaticNode()
)
		`

	s := scanner.CreateScanner(source)

	// testScanner := scanner.CreateScanner(s.ToString())
	// testScanner.Scan()
	// fmt.Printf("%s\n", s.ToString())
	// fmt.Printf("%s\n", testScanner.ToString())
	// fmt.Printf("%+v", testScanner.ToString() == s.ToString())

	p := parser.CreateParser(s.Scan())
	p.Parse()
	print("end")
	fmt.Printf("END!")
	//program end
}
