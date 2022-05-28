// Encapsulation is used to hide private information from other classes
// some languages have private keyword 
// but Golang does not 
// In Golang, Private information is kept hidden by naming them with lowercase
// If a struct, field or function's name starts with a lowercase, 
// it means its private and other packages cant access it

package model

type Director struct { // prublic 
	Name string // public
	address string // private
}

type movie struct { // private
	Name string // public
	Details string // public
	director string // private
}
