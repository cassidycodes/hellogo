// Any file with `package main` is consideredthe global scope.
// Two files with this declaration will tell the compler that they're in the same scope
package main

// It is possible to import from a remote repo by providing a link to it in the import statement
import (
	"fmt"
	"stringutils"
)

// Global scope for this package.
const state = "Ontario"

// Global scope for other pacakges.
var Name string

// package name and func name define the entry point. If they match, this package is executible. If they don't, the package can only be imported into other packages.
func main() {
	Name = "Cassidy"
	// String literal. Cannot include interpolation.
	from := `Toronto`
	n := 10
	// Complier sees this `state` as different from the global `state`.
	state := "Not Ontario"

	fmt.Printf("Hello fellow %s gophers!\n", stringutils.Upper(state))
	fmt.Printf("My name is %s and I'm from %s.\n", Name, from)
	fmt.Printf("By the time %d o'clock comes around, we'll know how to code in Go!\n", n)
	fmt.Println("Let's get started!")
}
