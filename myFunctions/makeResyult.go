package myFunctions

func MakeResult(chars [][]string) []string {
	result := []string{}
	lineToWrite := 0
	inWord := false
	letterCount := 0
	for i := 0; i < len(chars); i++ {

		// Handle \n.
		if chars[i][0] == "\n" {
			if !inWord {
				result = append(result, "")
				lineToWrite++
			} else if OnlyNewLinesRemaining(chars[i:]) {
				result = append(result, "")
				lineToWrite++
			}
			inWord = false
			continue
		}
		
		// Prepare Slice to write character.
		if letterCount == 0 {
			newLineAscii := []string{"", "", "", "", "", "", "", ""}
			result = append(result, newLineAscii...)
		} else if !inWord {
			newLineAscii := []string{"", "", "", "", "", "", "", ""}
			result = append(result, newLineAscii...)
			lineToWrite += 8
		}

		//Filling the letter in the result slice.
		letterCount++
		inWord = true
		asciiChar := chars[i]
		for j, line := range asciiChar {
			result[lineToWrite+j] += line
		}
	}
	return result
}

func OnlyNewLinesRemaining(slice [][]string) bool {
	for i := 0; i < len(slice); i++ {
		if len(slice[i]) > 0 && slice[i][0] == "\n" {
			i++
			continue
		} else {
			return false
		}
	}
	return true
}
