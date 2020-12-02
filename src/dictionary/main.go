package main

import (
	"dictionary/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	value, err := dictionary.Search(word)
	fmt.Println("found", word, "definition:", value)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	err = dictionary.Update("hihi", "Bowing")
	if err != nil {
		fmt.Println(err)
	}
	wordValue, _ := dictionary.Search(word)
	fmt.Println(wordValue)

	err = dictionary.Update(word, "Bowing")
	if err != nil {
		fmt.Println(err)
	}
	wordValue, _ = dictionary.Search(word)
	fmt.Println(wordValue)

	// Delete
	dictionary.Add("hihi", "hihihi")
	fmt.Println(dictionary)
	err = dictionary.Delete("hih")
	fmt.Println(err)
	err = dictionary.Delete("hihi")
	fmt.Println(err)
	fmt.Println(dictionary)
}
