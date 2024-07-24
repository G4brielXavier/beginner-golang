package main

import "fmt"

// Variaveis de escopo Global
var vr_int = 10          // Type: int
var vr_float = 2.2       // Type: float
var vr_string = "Hi, Go" // Type: str
var vr_bool = false      // Type: bool

func main() {

	// Operador curto de declaração
	msg := "Hi, Go"

	fmt.Println(msg)

	fmt.Println("variable integer: ", vr_int)
	fmt.Println("variable floating: ", vr_float)
	fmt.Println("variable string: ", vr_string)
	fmt.Println("variable boolean: ", vr_bool)

}
