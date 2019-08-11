package common

import (
	"math/rand"
	"time"
)

var (
	generator *rand.Rand
)

func init() {
	seed := time.Now().UnixNano()
	generator = rand.New(rand.NewSource(seed))
}

// NewIntUID crate a int ID which will be created once by this method
func NewIntUID() int64 {
	return generator.Int63()
}
