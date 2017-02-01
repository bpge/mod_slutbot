package slutbot

import (
	"fmt"
)

func GeneratePat(target string) string {
	adjectives := []WeightedString{
		{"lovingly", 5},
		{"caringly", 10},
		{"indifferently", 1},
		{"with great care", 1},
		{"furiously", 5},
	}
	return fmt.Sprintf("*pets %s %s*", target, GetRandomWeightedString(adjectives))
}