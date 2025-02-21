// Declaring what package the file belongs to
package main

//import packages. packages must be in quotes
//all imported packages must be used or the program will not compile
import "fmt"

// camelcase is the common convention for go variables
const secondsInHours = 3600

// the main.go requires a main function to run
func main() {
	fmt.Println("Hello, World!")

	//you can let go infer types using :=
	distance := 60.8 //the distance in kilometers

	//%f is called a verb and tells the Printf function how to format and print the output
	fmt.Printf("The distance in miles is %f \n", distance*0.621)
}

//you can use "go run main.go" to compile and run main.go
