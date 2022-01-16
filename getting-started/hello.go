package main

import (
	"fmt"

	"github.com/adelgado0723/PracticeGo/greetings"
)

func main() {

	name := "LeeAnne"

	if name == "LeeAnne" {

		for i := 0; i < 2; i++ {

			// fmt.Println(quote.Go())
			// fmt.Println(quote.Hello())
			// fmt.Println(quote.Opt())
			// fmt.Println(quote.Glass())
			fmt.Println(greetings.Hello(name))
		}
	} else {

		fmt.Println("Hi Stranger")
	}
}
