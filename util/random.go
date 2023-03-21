package util

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// fmt.Println(letters)
	fmt.Println(string(letters))

	result := make([]byte, n)
	// fmt.Println(result)
	// fmt.Println(string(result))
	rand.Seed(time.Now().Unix())
	for k := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	// fmt.Println(result)
	fmt.Println(string(result))
	return string(result)
}

func main() {
	RandomString(10)
}
