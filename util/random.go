package util

import "math/rand"

func RandUint32(min, max int) uint32 {
	n := min + rand.Intn(max-min)
	return uint32(n)
}
