// Object oriented approach is based on classes and objects
// Golang does not support classes
// instead of classes, Golang has structs
// structs are user defined types
// unlike classes structs hold only the state and not the behaviour
// In Golang it is dealt through methods

package main

import "fmt"

type number struct {
	value int32
}

func (numberOne number) isEqual(numberTwo number) bool {
	return numberOne.value == numberTwo.value
}

func main() {
	num1 := number{value: 2}
	num2 := number{value: 3}
	fmt.Println("Is 2 equal 3?: ", num1.isEqual(num2))
}


