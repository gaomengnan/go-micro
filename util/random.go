package util

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generate integer

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate string

func RandomString(n int) string {
	var sb = bytes.Buffer{}
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generates a random owner name

func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney generate a random amount of money

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency code

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
