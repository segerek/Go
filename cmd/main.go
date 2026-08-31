package main

import (
	"Project1/internal"
	"fmt"
)

func main() {
	var number, random_number int
	fmt.Println("Привет, введи число!")
	number = internal.Read_int()
	random_number = internal.Random_int(number)
	fmt.Printf("Случайное число от 0 до %d = %d\n", number-1, random_number)
}
