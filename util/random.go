package util

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func RandomUsername() string {
	var adjectives = []string{"happy", "funny", "silly", "lively", "clever", "brave", "calm", "crazy", "gentle", "graceful", "loving", "proud", "wild", "zealous"}
	var animals = []string{"dog", "cat", "bird", "fish", "lion", "tiger", "bear", "fox", "elephant", "panda", "koala", "monkey", "zebra", "giraffe"}

	rand.Seed(time.Now().UnixNano())

	adjective := adjectives[rand.Intn(len(adjectives))]
	animal := animals[rand.Intn(len(animals))]
	number := rand.Intn(100)

	return adjective + animal + strconv.Itoa(number)
}
