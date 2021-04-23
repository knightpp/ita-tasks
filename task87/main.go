package main

import (
	"flag"
	"fmt"

	tasks "github.com/knightpp"
	"github.com/knightpp/utils"
)

var nFlag = flag.Uint("n", 0, "any natural number")
var mFlag = flag.Uint("m", 0, "number of digits to sum from the end")

func main() {
	flag.Parse()
	var n, m uint = *nFlag, *mFlag
	if !utils.IsFlagPassed("n") && !utils.IsFlagPassed("m") {
		n = utils.PromptReadUint("Enter N (natural number):   ")
		m = utils.PromptReadUint("Enter M (number of digits): ")
	}
	sum := tasks.Task87(n, m)
	fmt.Printf("Sum of the last %v digits of %v is %v\n",
		m, n, sum)
}
