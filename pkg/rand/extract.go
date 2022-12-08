package accenter

import (
	"math/rand"
)

// Pick a random key in a map.
//
// https://www.reddit.com/r/golang/comments/kiees6/comment/ggs5z6l
func Pick[K comparable, V any](m map[K]V) K {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k
		}
		i--
	}
	panic("unreachable")
}

// Pick a random key in a map, according to the weights returned by getWeight.
//
// The weights must sum to totalWeight, set it to 0 to compute it.
func PickMap[K comparable, V any](
	m map[K]V, getWeight func(V) int, totalWeight int,
) K {
	// compute the total if we don't have it
	if totalWeight == 0 {
		for _, v := range m {
			totalWeight += getWeight(v)
		}
	}
	// pick a random number and find in which interval it falls
	i := rand.Intn(totalWeight)
	for k, v := range m {
		i -= getWeight(v)
		if i < 0 {
			return k
		}
	}
	panic("unreachable")
}
