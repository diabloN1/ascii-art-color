package myFunctions

func ColorResult(str string, flags, params []string, result [][]string) [][]string {
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