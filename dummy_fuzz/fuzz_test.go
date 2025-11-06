package main

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func FuzzDummy(f *testing.F) {
	magicLen := rand.IntN(42)

	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) == magicLen {
			fmt.Println("magic length!")
		} else {
			fmt.Println("not magic length...")
		}
	})
}
