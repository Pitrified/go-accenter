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
func PickMap[K comparable, V any](
	m map[K]V, getWeight func(V) int, totalWeight int,
) K {
	i := rand.Intn(totalWeight)
	for k, v := range m {
		i -= getWeight(v)
		if i < 0 {
			return k
		}
	}
	panic("unreachable")
}
