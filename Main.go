package main

import (
	"./serv"
	"fmt"
	"time"
	"math/rand"
	"os"
)

func main() {
	const MAX_STEP int = 500

	Server := serv.InitServer(35)

	file1, err := os.Create("file1.txt")
	file2, err := os.Create("file2.txt")
    check(err)

	defer file1.Close()
	defer file2.Close()

	for i := 0; i < MAX_STEP; i++ {
		seed := rand.New(rand.NewSource(time.Now().UnixNano()))

		fmt.Printf("%d / %d\n", i, MAX_STEP)
		startTime := makeTimestamp()
		result := serv.Fibon(seed.Intn(40))
		file1.WriteString(fmt.Sprintf("%d %d\n", result, makeTimestamp() - startTime))
	}

	for i := 0; i < MAX_STEP; i++ {
		seed := rand.New(rand.NewSource(time.Now().UnixNano()))

		fmt.Printf("%d / %d\n", i, MAX_STEP)
		startTime := makeTimestamp()
		result := Server.GetFibonacciNumber(seed.Intn(40))
		file2.WriteString(fmt.Sprintf("%d %d\n", result, makeTimestamp() - startTime))
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func makeTimestamp() int64 {
    return time.Now().UnixNano()
}
