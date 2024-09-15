package myFunctions

import "fmt"

func RGBColor(r, g, b int) string {
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
