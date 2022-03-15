package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ""
	greet_user()
	for input != "q\n" {
		input = get_input(input)
	}
}

func greet_user() {
	fmt.Print(
		"===================================================================\n",
		"|              Wellcome to the Go dice roller!                    |\n",
		"===================================================================\n",
		"For help run either 'goroller -h' or simply enter h.\n",
		"For quit enter q.\n")
}

func get_input(input string) string {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error ")
	}
	return input
}
