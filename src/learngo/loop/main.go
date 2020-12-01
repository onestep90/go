package main

import "fmt"

func superAdd(numbers ...int) int {
	fmt.Println(numbers)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
