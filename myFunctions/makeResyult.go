package myFunctions

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
