package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomPosition() string {
	positions := []string{"Design", "Develop"}
	n := len(positions)

	return positions[rand.Intn(n)]
}

func RandomGender() string {
	genders := []string{"Man", "Women"}
	n := len(genders)

	return genders[rand.Intn(n)]
}

func RandomDob() time.Time {
	return time.Now().UTC()
}

func RandomStatus() string {
	status := []string{"Active", "Inactive"}
	n := len(status)

	return status[rand.Intn(n)]
}
