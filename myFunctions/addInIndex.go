package myFunctions

func AddInIndex(s string, char string, index int) string {
	sRunes := []rune(s)
	charRunes := []rune(char)
	validRunes := []rune{}
	validRunes = append(validRunes, sRunes[:index]...)
	validRunes = append(validRunes, charRunes...)
	validRunes = append(validRunes, sRunes[index:]...)
	return string(validRunes)
}
