//go:build test

package main

import "fmt"

func main() {
	k := 0

	for k < 100 {
		k = k + 1
	}

	for i := 0; i < 5; i++ {
		k = k + 2
	}

	fmt.Println(k)
}
