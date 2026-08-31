package internal

import "math/rand"

func Random_int(number int) int {
	return rand.Intn(number)
}
