package tools

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const randomIDLength = 16

// GenerateRandomID generate a random identifier made of letters only
var GenerateRandomID = func() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, randomIDLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
