package myFunctions

func RemoveInIndex(s string, index, count int) string {
	sRunes := []rune(s)
	validRunes := []rune{}
	validRunes = append(validRunes, sRunes[:index]...)
	validRunes = append(validRunes, sRunes[index+count:]...)
	return string(validRunes)
}
