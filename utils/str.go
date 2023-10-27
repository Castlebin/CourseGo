package utils

import (
	"math/rand"
	"time"
)

var Aalphabes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Numbers = "0123456789"

var RandChars = Numbers + Aalphabes
var RandCharsLength = len(RandChars)

func RandString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = RandChars[r.Intn(RandCharsLength)]
	}
	return string(bytes)
}
