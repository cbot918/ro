package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var lg = fmt.Println

func main() {
	bytes, err := os.ReadFile("/dev/stdin")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.TrimRight(string(bytes), "\n")

	fmt.Println("\t.global main")
	fmt.Println("main:")
	fmt.Printf("\tmovl $%s, %%eax\n", contents)
	fmt.Println("\tret")

}
