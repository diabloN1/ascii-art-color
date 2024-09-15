package myFunctions

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
