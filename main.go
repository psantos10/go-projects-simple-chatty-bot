package main

import "fmt"

func main() {
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")
	fmt.Println("Please, remind me your name.")

	var name string
	fmt.Scanln(&name)

	fmt.Printf("What a great name you have, %s!\n", name)

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	var rem3, rem5, rem7 int
	fmt.Scan(&rem3)
	fmt.Scan(&rem5)
	fmt.Scan(&rem7)

	age := (rem3*70 + rem5*21 + rem7*15) % 105

	fmt.Printf("Your age is %d; that's a good time to start programming!\n", age)
}
