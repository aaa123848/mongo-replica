package main

import (
	"time"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func String(length int) string {
	return StringWithCharset(length, charset)
}
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
  
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}