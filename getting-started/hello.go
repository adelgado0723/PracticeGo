package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	name := "LeeAnne"

	if name == "LeeAnne" {

		for i := 0; i < 2; i++ {

			fmt.Println(quote.Go())
			fmt.Println(quote.Hello())
			fmt.Println(quote.Opt())
			fmt.Println(quote.Glass())
		}
	} else {

		fmt.Println("Hi Stranger")
	}
}
