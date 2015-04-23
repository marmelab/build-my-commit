package main

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const randomIDLength = 16

var _generateRandomID = func() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, randomIDLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var generateRandomID = _generateRandomID
