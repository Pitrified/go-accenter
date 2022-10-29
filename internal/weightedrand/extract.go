package accenter

import (
	"math/rand"

	wiki "example.com/accenter/pkg/wiki"
)

// given a map of infoword
// pick one according to some logic
func ExtractWord(m map[wiki.Word]InfoWord) wiki.Word {
	return pick(m)
}

// Pick a random key in a map.
//
// https://www.reddit.com/r/golang/comments/kiees6/comment/ggs5z6l/?utm_source=share&utm_medium=web2x&context=3
func pick[K comparable, V any](m map[K]V) K {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k
		}
		i--
	}
	panic("unreachable")
}
