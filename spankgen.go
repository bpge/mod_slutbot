package slutbot

import (
	"fmt"
	"math"
	"math/rand"
)

func GenerateSpank(target string) string {
	descriptor := []WeightedString{
		{"very bad", 5},
		{"bad", 10},
		{"good", 10},
		{"very good", 5},
	}
	tools := []WeightedString{
		{"her bare hand", 100},
		{"a rolled up newspaper", 100},
		{"a bamboo stick", 100},
		{"a birch branch", 100},
		{"a rolled up newspaper", 100},
		{"a slipper", 100},
		{"a spatula", 100},
		{"a hair brush", 100},
		{"a baseball bat :fire:", 50},
		{"a cat of 9 tails", 50},
		{"a Razer™ Ornata membrane mechanical keyboard", 50},
		{"a carpet beater :fire:", 50},
		{"a :diamond_shape_with_a_dot_inside: ", 5},
		{"the ***master sword***", 1},
	}
	numSpanks := NumSpanks()
	return fmt.Sprintf("*You've been %s, %s - spanks you %d times with %s*", GetRandomWeightedString(descriptor), target, numSpanks, GetRandomWeightedString(tools))

}

func NumSpanks() int {
	return int(math.Ceil(10 / rand.Float64() / 3))
}
