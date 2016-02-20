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

// NextFloat64 returns the next float64.
func NextFloat64() float64 {
	return rand.Float64()
}

// NextFloat64Bounded returns the next float64 bounded by start and end.
func NextFloat64Bounded(start float64, end float64) float64 {
	return rand.Float64()*(end-start) + start
}

// NextFloat32 returns the next float32.
func NextFloat32() float32 {
	return rand.Float32()
}

// NextFloat32Bounded returns the next float32 bounded by start and end.
func NextFloat32Bounded(start float32, end float32) float32 {
	return rand.Float32()*(end-start) + start
}

// NextInt returns the next int.
func NextInt() int {
	return rand.Int()
}

// NextIntBounded returns the next int bounded by start and end.
func NextIntBounded(start int, end int) int {
	return start + rand.Intn(end)
}

// NextIntBounded returns the next int bounded by a maximum.
func NextIntUpperBounded(end int) int {
	return rand.Intn(end)
}

// NextBytes creates an array of random bytes.
func NextBytes(count int) []byte {
	a := make([]byte, count)
	for i := range a {
		a[i] = (byte)(NextInt())
	}
	return a
}
