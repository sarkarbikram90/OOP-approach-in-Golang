# OOP-Approach-in-Go(lang)
Object-oriented programming is a programming paradigm which uses the idea of “objects” to represent data and methods. 
![Go Image](https://user-images.githubusercontent.com/54825429/170810126-89484a80-e117-4842-819c-4a1f08af43ab.JPG)
Go is a post-OOP programming language that borrows its structure (packages, types, functions) from the Algol/Pascal/Modula language family.
In Go, object-oriented patterns are useful for structuring a program in a clear and understandable way.

## Go does not strictly support object orientation but is a lightweight object Oriented language. 
Object Oriented Programming in Golang is different from other languages like C++ or Java due to factors mentioned below:

# Go does not support custom types through classes but *structs*. *Structs* in Golang are user-defined types that hold just the state and not the behavior. 
Structs can be used to represent a complex object comprising more than one key-value pairs. We can add functions to the struct that can add behavior to it.
# *Encapsulation* is hiding sensitive data from users. In Go, encapsulation is implemented by capitalizing fields, methods, and functions which makes them public. 
When the structs, fields, or functions are made public, they are exported on a package level.
# When a class acquires the properties of its superclass then we can say it is *Inheritance*.
Here, subclass/child class are the terms used for the class which acquire properties. For this, one must use a struct to achieve inheritance in Golang. 
Users have to compose using structs to form other objects.
# Interfaces are types that have multiple methods. Objects that implement all the methods of the interface automatically implement the interface, i.e., interfaces are satisfied implicitly. 
By treating objects of different types in a consistent way, as long as they stick to one interface.
