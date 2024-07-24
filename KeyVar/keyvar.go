package main

import "fmt"

var value_two = 100

func main() {

	value := 20

	print(value)
}

func print(x int) {

	calc := x + value_two
	fmt.Println(calc)

}
