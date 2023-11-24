package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	alphabet        = "abcdefghijklmnopqrstuvwxyz"
	persianAlphabet = "الفبپتثجچحخدذرزژسشصضطظعغفقکگلمنوهی"
)

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return int(int64(min) + rand.Int63n(int64(max-min+1)))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomPersianString generates a random string of length n
func RandomPersianString(n int) string {
	var sb strings.Builder
	k := len(persianAlphabet)

	for i := 0; i < n; i++ {
		c := persianAlphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
