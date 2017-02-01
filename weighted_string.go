package slutbot

import "math/rand"


type WeightedString struct {
	Value  string
	Weight int
}

func GetRandomWeightedString(ws []WeightedString) string {
	total := 0
	for _, w := range ws {
		total = w.Weight + total
	}
	target := rand.Intn(total-1) + 1
	counter := 0
	for _, w := range ws {
		if counter <= target && counter+w.Weight > target {
			return w.Value
		}
		counter = w.Weight + counter
	}
	return ws[len(ws)-1].Value
}