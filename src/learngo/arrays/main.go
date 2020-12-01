// Arrays and Slices
package main

import "fmt"

func main() {
	names := [5]string{"nico", "lynn", "dal"}
	fmt.Println(names)
	names[3] = "alala"
	names[4] = "adfa"
	fmt.Println(names)

	names2 := []string{"nico", "lynn", "dal"}
	names2 = append(names2, "flynn")
	fmt.Println(names)

}
