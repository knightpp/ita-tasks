package main

import (
	"flag"
	"fmt"

	tasks "github.com/knightpp"
	"github.com/knightpp/utils"
)

var nFlag = flag.Uint("n", 0, "any natural number")

func main() {
	flag.Parse()
	var n uint = *nFlag
	if !utils.IsFlagPassed("n") {
		n = utils.PromptReadUint("Enter N (uint): ")
	}
	merPrimes := tasks.Task559(n)
	fmt.Printf("Mersenne primes that less than %v\n%v\n",
		n, merPrimes)
}
