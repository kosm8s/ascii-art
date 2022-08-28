package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		fmt.Println("You should pass 1 argument. ")
		return
	}
	template, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	if input == "\n" {
		fmt.Println()
		return
	}
	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 127) && input[i] != 10 {
			fmt.Println("Incorrect input.")
			return
		}
	}
	splitted := strings.Split(string(template)[1:], "\n\n")
	if len(splitted) != 95 {
		fmt.Println("Incorrect template.")
		return
	}
	lines := strings.Split(input, "\n")
	res := ""
	for _, line := range lines {
		if line == "" && res != "" {
			res += string('\n')
			continue
		}
		for row := 0; row < 8; row++ {
			for i := 0; i < len(line); i++ {
				temp := strings.Split(splitted[line[i]-32], "\n")[row]
				for j := 0; j < len(temp); j++ {
					res += string(temp[j])
				}
			}
			res += string('\n')
		}
	}
	fmt.Print(res)
}
