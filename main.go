package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	min := 1
	max := 2

	if len(os.Args) >= 3 {
		var err1, err2 error
		min, err1 = strconv.Atoi(os.Args[1])
		max, err2 = strconv.Atoi(os.Args[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Please provide two valid integers.")
			os.Exit(1)
		}

		if min > max {
			fmt.Println("The first number must be less than or equal to the second number.")
			os.Exit(1)
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(max-min+1) + min

	fmt.Printf("%d\n", randomNumber)
}
