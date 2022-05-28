// Polymorphism is the ability to take on several different forms
// Objects of different classes a common point
// in many languages polymorphism is applied through interfaces
// Golang also has interfaces

package main

import "fmt"

type Animal interface {
	Speak() // common point of all animals

}

type Cat struct{}
type Dog struct{}
type Duck struct{}

// all animals can speak but different sounds
// cat, dog, duck implement Speak() method of the animal interface

func (cat Cat) Speak() {
	fmt.Println("Meow")
}

func (dog Dog) Speak() {
	fmt.Println("Woof")
}

func (duck Duck) Speak() {
	fmt.Println("Quack")
}

func main() {
	cat := Cat{}
	dog := Dog{}
	duck := Duck{}

	cat.Speak()
	dog.Speak()
	duck.Speak()
}
