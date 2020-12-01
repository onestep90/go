package main

import "fmt"

func canIDrink(age int) bool {
	// Variable expression => Reader can know that koreanAge is variable for if condition check.
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true

}

// Use switch.
func canIDrink2(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 50:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(51))
}
