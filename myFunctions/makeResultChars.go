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
			// Create a copy of the slice from asciiChars to solve dependecy of 2 chars if they are the same.
			asciiCharCopy := make([]string, len(asciiChars[int(str[i])]))
			copy(asciiCharCopy, asciiChars[int(str[i])])
			chars = append(chars, asciiCharCopy)
		}
	}
	return chars, nil
}
