package main

import (
	"fmt"
	"os"
	"ascii-art-color/myFunctions"
)

func main() {
	if len(os.Args) < 3 {
		myFunctions.PrintErr()
		return
	}
	str := os.Args[len(os.Args)-1]
	standard, err := myFunctions.Read("standard.txt")
	if err != nil {
		return
	}
	asciiChars := myFunctions.BytesToAsciiMap([]byte(standard))
	flags, params, err := myFunctions.HandleFlags(os.Args[1:len(os.Args)-1])
	if err != nil {
		myFunctions.PrintErr()
		return
	}
	result, err := myFunctions.MakeResultChars(str, asciiChars)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = myFunctions.ColorResult(str, flags, params, result)
	res := myFunctions.MakeResult(result)
	myFunctions.PrintResult(res)
}
