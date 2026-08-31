package internal

import (
	"fmt"
)

func Read_int() int {
	var number int = 1
	fmt.Scanln(&number)
	return number
}
