package something

import "fmt"

// Not exported because function name start with lowercase.
func sayBye() {
	fmt.Println("Bye")
}

// Exported because function name starts with uppercase.
func SayHello() {
	fmt.Println("Hello")
}
