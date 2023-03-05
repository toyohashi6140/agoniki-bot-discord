package utils

import (
	"math/rand"
	"time"
)

func RandNumber(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}
