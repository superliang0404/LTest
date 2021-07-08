package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args[0]

	fmt.Printf("应用名: %s\n", cmd)
}
