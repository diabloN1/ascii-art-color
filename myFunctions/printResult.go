package myFunctions

import "fmt"

func PrintResult(result []string) {
	for i := range result {
		fmt.Println(result[i])
	}
}
