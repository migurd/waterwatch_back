package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
