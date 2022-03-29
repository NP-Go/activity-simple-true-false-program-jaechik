package main

import "fmt"


var input = 10

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessge := ""
	secretValue := 88

	//Insert your code from here
for input := ""; input < secretValue; input ++ {
	fmt.Println("Too Low, try again next time!");
} else if input := ""; input > secretValue; input ++ {
	fmt.Println("Too High, try again next time!");
} else if input == secretValue {
	fmt.Println("Well Done! Your guess is correct");
}

	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	compare(guess)
}
