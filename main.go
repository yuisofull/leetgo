package main

import (
	"github.com/yuisofull/leetgo/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
