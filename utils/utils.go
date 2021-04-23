package utils

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
func PromptReadUint(msg string) uint {
START:
	var str string
	fmt.Print(msg)
	_, err := fmt.Scanln(&str)
	if err != nil {
		if err == io.EOF {
			os.Exit(0)
		}
		fmt.Printf("IO Error: %v\n", err)
		// io.ReadAll(os.Stdin)
		goto START
	}
	num, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		goto START
	}
	return uint(num)
}
