package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("(C) Bilbo Backends")
		fmt.Println("Usage: go run main.go --tbform=PCode.tbform --twin=PCode.twin")
		fmt.Println("Usage: go run main.go --twin=PCode.twin --tbform=PCode.tbform")
		return
	}
}
