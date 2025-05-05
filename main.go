package main

import (
	"fmt"
	"htestp/dsl/parser"
	"htestp/dsl/scanner"
)

func main() {

	source :=
		`
StaticNode("ID:123432", GET, "https://github.com", ExistConstraint("ID.users.name", STRING), ExistConstraint("Users", ARRAY),
	StaticNode("ID:123432", GET, "https://github.com", ExistConstraint("ID.users.name", STRING)), StaticNode("ID:123432", GET, "https://github.com", ExistConstraint("ID.users.name", STRING))
)
		`

	s := scanner.CreateScanner(source)

	p := parser.CreateParser(s.Scan())
	p.Parse()

	fmt.Printf("END!")

	//program end
}
