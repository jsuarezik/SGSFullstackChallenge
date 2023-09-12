package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GetFakeCategory() string {

	defaultCategories := []string{"Electronics", "Home", "Fashion", "Beauty", "Sports", "Kids", "Automotive", "Books"}

	src := rand.NewSource(time.Now().UnixNano())

	r := rand.New(src)

	index := r.Intn(len(defaultCategories))

	return defaultCategories[index]
}

func GetRandomBool() bool {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(2) == 0
}

func GetRandomFloat(min float64, max float64) float64 {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randomValue := min + r.Float64()*(max-min)
	float, err := strconv.ParseFloat(fmt.Sprintf("%.2f", randomValue), 64)
	if err != nil {
		log.Fatalf("Error converting string to float: %v", err)
	}

	return float
}

func GetRandomInt(max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max)
}
