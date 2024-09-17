package myFunctions

import "fmt"

func MakeResultChars(str string, asciiChars map[int][]string) ([][]string, error) {
	chars := [][]string{}
	for i := 0; i < len(str); i++ {
		//Handle non ascii char
		if str[i] < 32 || str[i] > 126 {
			return chars, fmt.Errorf("A none ascii char has been found !!")
		}
		if i < len(str) - 1 && str[i] == '\\' && str[i+1] == 'n' {
			chars = append(chars, []string{"\n"})
			i++
		} else {
			chars = append(chars, asciiChars[int(str[i])])
		}
	}
	return chars, nil
}
