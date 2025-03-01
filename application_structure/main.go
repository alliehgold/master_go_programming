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
//unlike "go build" it does not produce an executable
//you may have to run "go mod init" first if you get an error
//"go build" compiles the program but does not execute it
//"go build -o appname" creates an executable with the specific name
//Use "GOOS=<operating_system_name> go build" to build for a specific operating system
//choices are Windows, linux, or darwin (for mac)
//use "GOARCH=<architecturename>" to specify the CPU architecture such as AMD64 or ARM64
//when specifying the name of the application, you do not need a file extension for linux or Mac

//go strongly suggests certain styles
//run "gofmt -w main.go" to have go format your application
//use "gofmt -w -l <directory name>" to format every go file in a directory
//just run "go fmt" to format all files in the current directory
//vscode by default runs gofmt when you save a go file
