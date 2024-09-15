package myFunctions

func IsFlag(s string) bool {
	is := false
	if len(s) > 8 && s[:8] == "--color=" {
		is = true
	}
	return is
}
