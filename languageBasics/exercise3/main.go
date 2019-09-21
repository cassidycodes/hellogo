package main

import (
	"fmt"
	"strings"
)

const allProverbs = `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`

func main() {
	fmt.Println("-------------------------------------------------------------------------------------------------------")
	fmt.Println("\n\n\n")
	proverbs := strings.Split(allProverbs, "\n")

	for i, proverb := range proverbs {
		var counts string
		fmt.Printf("%d. %s\n", i+1, proverb)
		for _, char := range strings.Split(proverb, "") {
			count := fmt.Sprintf("'%s'%d=,", char, strings.Count(proverb, char))
			counts = fmt.Sprintf("%s %s", counts, count)
		}
		fmt.Println(counts)
	}
}
