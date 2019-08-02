package main

import (
	"math/rand"
	"time"
)

const (
	consonants = "bcdfghjklmnpqrstvwxyz"
	vowels     = "aeiou"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomFromCharset(charset string) byte {
	return charset[seededRand.Intn(len(charset))]
}

// GenerateIdentifier constructs a random human readable indentifier of a specific length
func GenerateIdentifier(length int) string {
	b := make([]byte, length)
	for i := range b {
		if i%2 == 0 {
			b[i] = randomFromCharset(vowels)
		} else {
			b[i] = randomFromCharset(consonants)
		}
	}
	return string(b)
}
