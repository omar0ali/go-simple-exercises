package main

import (
	"fmt"
	"os"

	"github.com/omar0ali/go-simple-exercises/math"
)

func main() {
	var total int = 0
	if len(os.Args) == 1 {
		fmt.Println("No args provided.\nUsage: app.go operation <numbers>...\napp.go add 1 2 3")
	} else {
		switch os.Args[1] {
		case "add", "+":
			total = math.Add(os.Args[2:])
		case "sub", "-":
			total = math.Sub(os.Args[2:])
		case "multi", "x":
			total = math.Multiply(os.Args[2:])
		case "division", "/":
			total = math.Division(os.Args[2:])
		case "bSearch":
			ls := os.Args[2:]
			target := math.ConvertToInt(ls[0])
			values := ls[2:]
			fmt.Printf("%d %v\n", target, values)
			index := 0
			total, index = math.BinarySearch(target, values)
			fmt.Printf("Found at Index: %d\n", index)
		default:
			fmt.Println("Usage: app.go operation <numbers> ...\napp.go add 1 2 3")
		}
	}
	fmt.Printf("%v\n", total)
}
