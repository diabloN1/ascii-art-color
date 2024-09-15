package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func PrintErr() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
}

func main() {
	if len(os.Args) < 3 {
		PrintErr()
		return
	}
	str := os.Args[len(os.Args)-1]
	standard, err := Read("standard.txt")
	if err != nil {
		return
	}
	asciiChars := BytesToAsciiMap([]byte(standard))
	flags, params, err := HandleFlags(os.Args[1:len(os.Args)-1])
	if err != nil {
		PrintErr()
		return
	}
	result, err := makeResultChars(str, asciiChars)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = ColorResult(str, flags, params, result)
	res := MakeResult(result)
	PrintResult(res)
}

func RGBColor(r, g, b int) string {
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColorHandling(color string) string {
	if color == "" {
		color = "\033[0m"
	} else if color == "red" {
		color = "\033[31m"
	} else if color == "green" {
		color = "\033[32m"
	} else if color == "yellow" {
		color = "\033[33m"
	} else if color == "blue" {
		color = "\033[34m"
	} else if color == "magenta" {
		color = "\033[35m"
	} else if color == "cyan" {
		color = "\033[36m"
	} else if color == "white" {
		color = "\033[37m"
	} else if color [:4] == "rgb(" && color[len(color)-1] == ')' {
		rgb := strings.Split(color[:len(color)-1], ", ")
		r, _ := strconv.Atoi(rgb[0])
		g, _ := strconv.Atoi(rgb[1])
		b, _ := strconv.Atoi(rgb[2])
		color = RGBColor(r, g, b)
	}
	return color
}

func ColorResult(str string, flags, params []string, result [][]string) [][]string {
	for i := range flags {
		color := ColorHandling(flags[i])
		if params[i] == "" {
			for charSlice := 0; charSlice < len(result); charSlice++ {
				if result[charSlice][0] == "\n" {
					continue
				}
				for line := 0; line < len(result[charSlice]); line++ {
					// }
					fmt.Print(result[charSlice][line])
					if len(result[charSlice][line]) > 8 && result[charSlice][line][1:6] == color {
						continue
					} else if len(result[charSlice][line]) > 7 && result[charSlice][line][1] == '\033' {
						result[charSlice][line] = RemoveInIndex(result[charSlice][line], 0, 5) 
						result[charSlice][line] = RemoveInIndex(result[charSlice][line], len(result[charSlice][line])-6, 5)
					}
					result[charSlice][line] = AddInIndex(result[charSlice][line], color, 1)
					result[charSlice][line] = AddInIndex(result[charSlice][line], "\033[37m", len(result[charSlice][line])-1)
				}
			}
		} else {
			countReturn := 0
			for j := 0; j <= len(str) - len(params[i]); j++ {
				if str[j] == '\\' && str[j+1] == 'n' {
					j++
					countReturn++
					continue
				}
				if str[j:j+len(params[i])] == params[i] {
					for charSlice := j-countReturn; charSlice < j+len(params[i])-countReturn; charSlice++ {
						for line := 0; line < len(result[charSlice]); line++ {
							if len(result[charSlice][line]) > 8 && result[charSlice][line][1:6] == color {
								continue
							} else if len(result[charSlice][line]) > 7 && result[charSlice][line][1] == '\033' {
								result[charSlice][line] = RemoveInIndex(result[charSlice][line], 0, 5) 
								result[charSlice][line] = RemoveInIndex(result[charSlice][line], len(result[charSlice][line])-6, 5)
							}
							result[charSlice][line] = AddInIndex(result[charSlice][line], color, 1)
							result[charSlice][line] = AddInIndex(result[charSlice][line], "\033[37m", len(result[charSlice][line])-1)
						}
					}
				}
			}
		}
	}
	return result
}

func RemoveInIndex(s string, index, count int) string {
	sRunes := []rune(s)
	validRunes := []rune{}
	validRunes = append(validRunes, sRunes[:index]...)
	validRunes = append(validRunes, sRunes[index+count:]...)
	return string(validRunes)
}

func AddInIndex(s string, char string, index int) string {
		sRunes := []rune(s)
		charRunes := []rune(char)
		validRunes := []rune{}
		validRunes = append(validRunes, sRunes[:index]...)
		validRunes = append(validRunes, charRunes...)
		validRunes = append(validRunes, sRunes[index:]...)
		return string(validRunes)
}

func HandleFlags(args []string) ([]string, []string, error) {
	flags := []string{}
	params := []string{}
	for i := 0; i < len(args); i++ {
		if i < len(args) - 1 && IsFlag(args[i]) && !IsFlag(args[i+1]) {
			flags = append(flags, strings.Split(args[i], "--color=")[1])
			params = append(params, args[i+1])
			i++
		} else if IsFlag(args[i]) {
			flags = append(flags, strings.Split(args[i], "--color=")[1])
			params = append(params, "")
		} else {
			return flags, params, fmt.Errorf("Found not a valid flag")
		}
	}
	return flags, params, nil
}

func IsFlag(s string) bool {
	is := false
	if len(s) > 8 && s[:8] == "--color=" {
		is = true
	}
	return is
}

func PrintResult(result []string) {
	for i := range result {
		fmt.Println(result[i])
	}
}

func makeResultChars(str string, asciiChars map[int][]string) ([][]string, error) {
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
			chars = append(chars, asciiChars[int(str[i])][1:])
		}
	}
	return chars, nil
}

func MakeResult(chars [][]string) []string {
	result := []string{"","","","","","","",""}
	lineToWrite := 0
	for i := 0; i < len(chars); i++ {
		if chars[i][0] == "\n" {
			if len(chars) - 1 == i || chars[i+1][0] == "\n" { //
				result = append(result,	"")
				lineToWrite++
			} else {
				newLineAscii := []string{"","","","","","","",""}
				result = append(result,	newLineAscii...)
				lineToWrite += 8
			}
			continue
		}
		asciiChar := chars[i]
		for j, line := range asciiChar {
			result[lineToWrite+j] += line[1:len(line)-1]
		}
	}
	return result
}

func BytesToAsciiMap(style []byte) map[int][]string {
	chars := make(map[int][]string)
	line := 1
	next := 9
	lineS := ""
	char := []string{}
	nbrChar := 32
	for i := 1; i < len(style); i++ {
		if i < len(style) - 1 {
		 	if style[i] == '\n' {
				char = append(char, lineS)
				lineS = ""
				line++

			} else if line == next+1 {
				next += 9
				chars[nbrChar] = char //[1:len(char)-2]
				nbrChar++
				char = []string{}
				continue
			}
			lineS += string(style[i])
		} else {
			lineS += string(style[i])
			chars[nbrChar] = char
		}
	}
	return chars
}

func Read(fileName string) (string, error) {
	
	//Open File.
    file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening file :", fileName)
		return "", err
	}

	defer file.Close()
	
	//Get file info.
    fileInfo, err := file.Stat()
    if err != nil {
        log.Println("Error getting file stats:", err)
		return "", err
    }

    //Get file size.
    fileSize := fileInfo.Size()
    data := make([]byte, fileSize)

	//Reading data.
    _, err = file.Read(data)
    if err != nil {
		log.Println("Error reading the file:", err)
		return "", err
    }
	return string(data), nil
}
