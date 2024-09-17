package myFunctions

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ColorHandling(color string) string {
	if color == "" {
		color = "\033[0m"
	} else if color == "red" {
		color = "\033[31m"
	} else if color == "green" {
		color = "\033[32m"
	} else if color == "yellow" {
		color = "\033[33m"
	} else if color == "blue" {
		color = "\033[34m"
	} else if color == "magenta" {
		color = "\033[35m"
	} else if color == "cyan" {
		color = "\033[36m"
	} else if color == "white" {
		color = "\033[37m"
	} else if len(color) > 5 && color [:4] == "rgb(" && color[len(color)-1] == ')' {
		rgb := strings.Split(color[:len(color)-1], ", ")
		r, _ := strconv.Atoi(rgb[0])
		g, _ := strconv.Atoi(rgb[1])
		b, _ := strconv.Atoi(rgb[2])
		color = RGBColor(r, g, b)
	} else {
		log.Println("This color does not exist :", color)
		os.Exit(1)
	}
	return color
}
