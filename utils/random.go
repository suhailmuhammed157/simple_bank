package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	alphabets = "abcdefghijklmnopqrs"
)

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)
	for i := 0; i < n; i++ {
		char := alphabets[rand.Intn(k-1)]
		sb.WriteByte(char)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(1, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "AED"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	emailUser := RandomString(6)
	domain := "@email.com"
	return fmt.Sprint(emailUser + domain)
}
