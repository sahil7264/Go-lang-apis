package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	a := 1
	b := 2.5

	fmt.Print(a+int(b), " ")

	//random number generator
	// exit program -> panic() log.fatal()

	//Math/rand
	// rand.Seed(time.Now().UnixNano()) // entire algo can be driven by seeder
	// fmt.Println(rand.Intn(5) + 1)

	// Crypto/rand
	// my, _ := rand.Int(rand.Reader, big.NewInt(5))

}

