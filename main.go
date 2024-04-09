package main

import (
	"evarolang/morestrings"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	name := "Shikinami"
	fmt.Print(name)
	morestrings.Hr(1)

	_age := -18 // mean integer
	fmt.Print(_age)
	morestrings.Hr(1)

	var age uint = 18 // uint , cannot have a negative number value
	fmt.Print(age)
	morestrings.Hr(1)

	var _name string
	_name = "Rei"
	fmt.Print(_name)
	morestrings.Hr(1)

	float := 10.2
	fmt.Print(float)
	morestrings.Hr(1)

	var text [2]string
	text[0] = "Hello"             // First index
	text[1] = name                // Second index
	fmt.Println(text[0], text[1]) // Separate index
	fmt.Println(text)             // Raw array

	primes := [6]int{2, 3, 5, 7}
	fmt.Println(primes)

	/*
		Comment
	*/
	// Also comment

	hello("Rei")                             // normal function
	fmt.Println("Upper: " + uppercase(name)) // function to uppercase a string
	gbr, ex := goodbyeMultipleReturn(name)   // multiple return values
	gbp := goodbyeMultipleParam(name, "💀")   // multiple return values

	fmt.Println(gbr, ex)
	fmt.Println(gbp)

	fmt.Println("Reversed word: " + morestrings.Reverse("ieR")) // Importing packages from my module
	fmt.Println(cmp.Diff("1 deletion", "1 addition"))           // Importing packages from remote modules

	fmt.Println(typeCasting(name, age))

	handleInput()

}

// Funciton description
func hello(name string) {
	fmt.Printf("Hello %v\n\n", name)
}

func uppercase(text string) string {
	return strings.ToUpper(text + "\n")
}

func goodbyeMultipleReturn(name string) (text, text2 string) {
	text = " Goodbye,\n"
	text2 = "See ya " + name + "\n"
	return
}

func goodbyeMultipleParam(name, emoji string) (text string) {
	text = "Goodbye " + name + " " + emoji + "\n"
	return
}

func typeCasting(name interface{}, age uint) (text string) {
	text = "Hi " + name.(string) + " are you " + strconv.Itoa(int(age)) + " y.o?\n"
	// `strconv.Itoa(int(age))` convert the age variable, which is of type uint, to a string.
	return
}

func handleInput() {
	var x, y int

	fmt.Print("[Step 1/2] Type first numbers: ")
	fmt.Scan(&x)

	if x == 0 {
		fmt.Println("The first input is not a valid number")
		return
	}

	fmt.Print("[Step 2/2] Type second numbers: ")
	fmt.Scan(&y)

	if y == 0 {
		fmt.Println("The second input is not a valid number")
		return
	}

	fmt.Println("Your numbers are:", x, "and", y)
}
