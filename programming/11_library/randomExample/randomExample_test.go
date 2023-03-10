package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// random or rng or random number generator

/*
1. random actually is determiistic and probablity calculation algorithm. So, if we know the state, we can guessing the random number
2. in random or rng there is seed which used by RNG to generate random number in every sequence
*/

func TestRandomSeed(t *testing.T) {
	rng_generator := rand.New(rand.NewSource(10)) // i set seed to 10
	for i := 0; i < 10; i++ {
		fmt.Println("Random generator", i, rng_generator.Int())
	}
}

// with seed, the rng which created by system or program always same

/*
Random generator 0 5221277731205826435
Random generator 1 3852159813000522384
Random generator 2 8532807521486154107
Random generator 3 3888302351045490779
Random generator 4 4512466281294657143
Random generator 5 8455134228352958140
Random generator 6 3397244140909564941
Random generator 7 6857276301702941374
Random generator 8 7956601014869229133
Random generator 9 250098600798745966

if i run this again, the result is same
*/

// generate with unique seed, so the result always difference each time we run it or in other computer
// to achieve it, we can use unix time (or other alternative)

func TestRandomSeedUnique(t *testing.T) {
	fmt.Println(time.Now().UTC().UnixNano())
	rng_generator := rand.New(rand.NewSource(time.Now().UTC().UnixNano())) // set seed to unique value, for example, unix time
	for i := 0; i < 10; i++ {
		fmt.Println("Random generator", i, rng_generator.Int())
	}
}

// how about using float?

func TestRandomSeedUniqueFloat(t *testing.T) {
	fmt.Println(time.Now().UTC().UnixNano())
	rng_generator := rand.New(rand.NewSource(time.Now().UTC().UnixNano())) // set seed to unique value, for example, unix time
	for i := 0; i < 10; i++ {
		fmt.Println("Random generator", i, rng_generator.Float32()) // or float64
	}
}

// how about using string?

func TestRandomSeedUniqueStr(t *testing.T) {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rngStr := make([]rune, 12) // make string from rune
	for i := range rngStr {
		rngStr[i] = letters[randomizer.Intn(len(letters))]
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Random generator", i, string(rngStr)) // generate string
	}
}

// get single random number

func TestRandomSeedSingleInt(t *testing.T) {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println("Random generator", i, randomizer.Intn(100)) // generate string
	}
}
