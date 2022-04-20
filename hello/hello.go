package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"golang.org/x/example/stringutil"
	"rsc.io/quote"
)

func main() {
	// fmt.Println("hello")

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	message, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})

	fmt.Println(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stringutil.Reverse("Hello"))

	fmt.Println("Hello End.")

}
