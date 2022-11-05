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
