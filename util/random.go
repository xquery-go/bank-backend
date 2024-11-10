package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// set random seed
func init() {
	rand.NewSource(time.Now().UnixNano())
}

// generate random int64 from min -> max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}

	return sb.String()
}

// generate random account owner name
func RandomOwner() string {
	return RandomString(6)
}

// generate random bank balance
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"AUD", "EUR", "USD", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
