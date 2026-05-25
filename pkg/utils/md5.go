package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func MD5WithSalt(password, salt string) string {
	return MD5(password + salt)
}

func GenerateSalt() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
