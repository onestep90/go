package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// Naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer. Executed after function is finished.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("onestep")
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("onestep")
	fmt.Println(totalLength2)

	_, upperName2 := lenAndUpper("onestep")
	fmt.Println(upperName2)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	totalLength3, upperName3 := lenAndUpper2("onestep")
	fmt.Println(totalLength3, upperName3)
}
