// example of a function in Golang

package main

import "fmt"

func isEqual(numberOne int32, numberTwo int32) bool {
	return numberOne == numberTwo
}

func main() {
	fmt.Println("Is 2 equal to 3?: ", isequal(2, 3))
}
