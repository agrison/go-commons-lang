// Package randUtils provides some utilities regarding random numbers.
package randUtils

import (
	"math/rand"
	"time"
)

// Seed inits the random seed using the current UTC Unix Nano time.
func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NextFloat64() float64 {
	return rand.Float64()
}

func NextFloat64Bounded(start float64, end float64) float64 {
	return rand.Float64()*(end-start) + start
}

func NextFloat32() float32 {
	return rand.Float32()
}

func NextFloat32Bounded(start float32, end float32) float32 {
	return rand.Float32()*(end-start) + start
}

func NextInt() int {
	return rand.Int()
}

func NextIntBounded(start int, end int) int {
	return start + rand.Intn(end)
}

func NextIntUpperBounded(end int) int {
	return rand.Intn(end)
}

// NextBytes creates an array of random bytes.
func NextBytes(count int) []byte {
	a := make([]byte, count)
	for i, _ := range a {
		a[i] = (byte)(NextInt())
	}
	return a
}
