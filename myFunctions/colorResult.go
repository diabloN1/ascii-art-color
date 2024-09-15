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