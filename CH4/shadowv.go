package main

import "fmt"

func main() {

	// In Go it is pretty similiar how blocks work, here is an example of a shadow varaible.const
	x := 10
	if x > 5 {
		fmt.Println(x) // Inside block can access outside variable
		x := 5         // Inside variable is different, this is a shadow variable, same name diff variable
		fmt.Println(x)
	}
	fmt.Println(x) // 10 5 10

	// Shadow linter can detect shadow variables
	// go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

	// Universe block
	// Go is a small languge, 25 keywords. Words like int, string, or constants like true or false.These are predeclared identifiers and are declared in the universe block.
	// They are not included in the keyword list so they can be shadowed in other scopes. NEVER redifine any of the identifiers in the universe block.
}
