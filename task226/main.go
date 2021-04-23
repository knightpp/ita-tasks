package main

import (
	"flag"
	"fmt"

	tasks "github.com/knightpp"
	"github.com/knightpp/utils"
)

var nFlag = flag.Uint("n", 0, "any natural number")
var mFlag = flag.Uint("m", 0, "any natural number")

func main() {
	flag.Parse()
	var n, m uint = *nFlag, *mFlag
	if !utils.IsFlagPassed("n") && !utils.IsFlagPassed("m") {
		m = utils.PromptReadUint("Enter M (uint): ")
		n = utils.PromptReadUint("Enter N (uint): ")
	}
	nums := tasks.Task226(m, n)
	fmt.Printf("Common denominators less than %v\n%v\n",
		n*m, nums)
}
