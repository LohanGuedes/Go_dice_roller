package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ""
	again := "Y"
	greet_user()
	for again != "n" {
		fmt.Println("Please enter the dices you want to roll:")
		input = get_line()
		fmt.Println("Input:", input)
		// Roll each dice
		fmt.Println("Do you want to roll again? [Y/n]")
		fmt.Scanf("%s", &again)
	}
}

func greet_user() {
	fmt.Print(
		"===================================================================\n",
		"|              Wellcome to the Go dice roller!                    |\n",
		"===================================================================\n",
		"For help run either 'goroller -h' or simply enter h.\n")
}

func get_line() string {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error ")
	}
	return input
}
