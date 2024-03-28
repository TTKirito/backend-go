package utils

import (
	"fmt"
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

func RandomTime() time.Time {
	return time.Now().UTC()
}

func RandomStatus() string {
	status := []string{"Active", "Inactive"}
	n := len(status)

	return status[rand.Intn(n)]
}

func RandomEmegency() bool {
	return rand.Intn(2) == 1
}

func RandomEventType() string {
	eventTypes := []string{"Event", "Meeting"}
	n := len(eventTypes)

	return eventTypes[rand.Intn(n)]
}

func RandomVisitType() string {
	visitTypes := []string{"Office", "Online"}
	n := len(visitTypes)
	return visitTypes[rand.Intn(n)]
}

type Coordinates struct {
	Lat    string
	Long   string
	Street string
}

func RandomLatLong() Coordinates {
	coordinates := []Coordinates{{"14.0583", "1082772", "Hanoi"}, {"15.8700", "100.9925", "Thailand"}}
	n := len(coordinates)

	return coordinates[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
