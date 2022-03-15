package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := ""
	again := "Y"
	greet_user()
	// TODO: Ask if wants to log
	// log the program
	for again != "n" {
		fmt.Println("Please enter the dices you want to roll:")
		input = get_line()
		if input == "q" {
			break
		}
		fmt.Println("To roll:", input)
		roll_dices(input)
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
	input = strings.TrimSuffix(input, "\n")
	if err != nil {
		fmt.Println("Error ")
	}
	return input
}

func roll_dices(input string) {
	inp := strings.Split(input, " ")
	for _, dice := range inp {
		fmt.Println(dice)
		// TODO: Separate number of dices from number of faces. Example: XdY
		// TODO: Roll X times a dice with Y faces. Final result can be from 1~Y
		// TODO: Check for modifier
	}
}
