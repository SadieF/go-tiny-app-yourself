package utils

import (
	"fmt"
	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func MakeRandomString() string {
	return MakeRandomStringOfLength(6)
}

func MakeRandomStringOfLength(length int) string {
	if length == 0 {
		return ""
	}
	randomIndex := rand.Intn(len(alphabet))
	return fmt.Sprintf("%c%s", alphabet[randomIndex], MakeRandomStringOfLength(length-1))
}
