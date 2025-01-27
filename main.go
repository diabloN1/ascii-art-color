package main

import (
    "ascii-art-color/myFunctions"
    "fmt"
    "os"
    "strings"
)

func main() {
    var InputFlags = make(map[string]bool)
    banner := "standard"
    if len(os.Args) == 1 {
        myFunctions.PrintErr()
        return
    } 
	if HasBanner(os.Args[len(os.Args)-1]) {
        InputFlags["bannerFlag"] = true 
        banner =  os.Args[len(os.Args)-1]
        os.Args = os.Args[:len(os.Args)-1]
    } 
	if strings.HasPrefix(os.Args[1], "--color=") {
        InputFlags["colorFlag"] = true
    } 

    str := os.Args[len(os.Args)-1]


    standard, err := myFunctions.Read(banner + ".txt")
    if err != nil {
        return
    }

    asciiChars := myFunctions.BytesToAsciiMap([]byte(standard))
    var flags, params []string
    if InputFlags["colorFlag"] {
        flags, params, err = myFunctions.HandleFlags(os.Args[1:len(os.Args)-1])
        if err != nil {
            myFunctions.PrintErr()
            return
        }
    }

    result, err := myFunctions.MakeResultChars(str, asciiChars)
    if err != nil {
        fmt.Println(err)
        return
    }
    if InputFlags["colorFlag"]  {
        result = myFunctions.ColorResult(str, flags, params, result)
    }
    res := myFunctions.MakeResult(result)
    myFunctions.PrintResult(res)
}


func HasBanner(str string) bool {
     return str == "standard" || str == "thinkertoy" ||  str == "shadow"
}