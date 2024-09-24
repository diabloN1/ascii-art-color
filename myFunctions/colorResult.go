package myFunctions

import "fmt"

// go run . --color=red "lo\n" "lo\nllllo\n".
func ColorResult(str string, flags, params []string, result [][]string) [][]string {
	str = remakeMewLines(str)
	for i := range params {
		params[i] = remakeMewLines(params[i])
	}
	for i := range flags {
		color := ColorHandling(flags[i])
		if params[i] == "" {
			for charSlice := 0; charSlice < len(result); charSlice++ {
				if result[charSlice][0] == "\n" {
					continue
				}
				for line := 0; line < len(result[charSlice]); line++ {
					if len(result[charSlice][line]) > 8 && result[charSlice][line][0:5] == color {
						continue
					} else if len(result[charSlice][line]) > 7 && result[charSlice][line][0] == '\033' {
						result[charSlice][line] = result[charSlice][line][5:]
						result[charSlice][line] = result[charSlice][line][:len(result[charSlice][line])-5]
					}
					result[charSlice][line] = color + result[charSlice][line] + "\033[37m"
				}
			}
		} else {
			for j := 0; j <= len(str) - len(params[i]); j++ {
				fmt.Println(params)
				fmt.Println(str[j:j+len(params[i])])
				if str[j:j+len(params[i])] == params[i] {
					fmt.Println("ok")
					for charSlice := j; charSlice < j+len(params[i]); charSlice++ {
						for line := 0; line < len(result[charSlice]); line++ {
							if result[charSlice][0] == "\n" {
								
								continue
							}
							if len(result[charSlice][line]) > 8 && result[charSlice][line][0:5] == color {
								continue
							} else if len(result[charSlice][line]) > 7 && result[charSlice][line][0] == '\033' {
								result[charSlice][line] = result[charSlice][line][5:]
								result[charSlice][line] = result[charSlice][line][:len(result[charSlice][line])-5]
							}
							result[charSlice][line] = color + result[charSlice][line] + "\033[37m"
						}
					}
				}
			}
		}
	}
	return result
}

func remakeMewLines(str string) string {
	newStr := ""
	for i:=0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == '\\' && str[i+1] == 'n' {
			newStr+="\n"
			i++
		} else {
			newStr+=string(str[i])
		}
	}
	return newStr
}