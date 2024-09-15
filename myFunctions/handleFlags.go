package myFunctions

import (
	"strings"
	"fmt"
)

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
