package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	greeting := flag.String("greeting", "Hello", "a greeting message")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go [options] <name>")
		os.Exit(1)
	}

	name := args[0]
	fmt.Println(*greeting, name,"Welcome to Devops!")
	user := os.Getenv("USER")
	fmt.Println("Your user is:", user)
}