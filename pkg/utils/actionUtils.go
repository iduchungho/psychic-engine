package utils

import (
	"math/rand"
	"time"
)

func RandNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100000)
}
