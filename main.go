// Package main contains the main function.
package main

// This is a comment.

/*
This is a multiline comment.
*/


//Imports loads the packages for use.
import (
	"fmt"

	"github.com/Turtle_Example_Code/turtle"
)

//Global constants/vars are declared first.
const (
	frankName = "Frank"
	millieName = "Millie"
)

// Every Golang program starts with a main function.
func main(){
	frankTurtle := turtle.NewTurtle(frankName, "green")
	millieTurtle := turtle.NewTurtle(millieName, "brown")

	frankTurtle.Direction = "Right"
	millieTurtle.Direction = "Left"

	frankTurtle.Move(20)
	millieTurtle.Move(30)

	fmt.Printf("%s Cordinations X:%d Y:%d Z:%d\n", frankTurtle.Name, frankTurtle.Position.X, frankTurtle.Position.Y, frankTurtle.Position.Z)
	fmt.Printf("%s Cordinations X:%d Y:%d Z:%d\n", millieTurtle.Name, millieTurtle.Position.X, millieTurtle.Position.Y, millieTurtle.Position.Z)

	fmt.Println("")
	// Frank is speaking to Millie
	for i := 0; i < 5; i++ {
		fmt.Printf("%s says: %s\n", frankTurtle.Name, frankTurtle.Phrase)
		switch frankTurtle.Phrase {
		case "Hello World!":
			fmt.Printf("%s says: %s\n", millieTurtle.Name, millieTurtle.Phrase)
			frankTurtle.SetRandomPhrase()
		case "How r u today?":
			fmt.Printf("%s says: %s\n", millieTurtle.Name, millieTurtle.Feeling)
			frankTurtle.SetRandomPhrase()
		case "What's ur favorite color?":
			fmt.Printf("%s says: %s\n", millieTurtle.Name, millieTurtle.Color)
			frankTurtle.SetRandomPhrase()
		case "How big r u?":
			fmt.Printf("%s says: %d\n", millieTurtle.Name, millieTurtle.Size)
			frankTurtle.SetRandomPhrase()
		case "Where r u?":
			fmt.Printf("%s Cordinations X:%d Y:%d Z:%d\n", millieTurtle.Name, millieTurtle.Position.X, millieTurtle.Position.Y, millieTurtle.Position.Z)
			frankTurtle.SetRandomPhrase()
		}
	}

	fmt.Println("")

	// Millie is speaking to Frank.
	for i := 0; i < 5; i++ {
		fmt.Printf("%s says: %s\n", millieTurtle.Name, millieTurtle.Phrase)
		switch millieTurtle.Phrase {
		case "Hello World!":
			fmt.Printf("%s says: %s\n", frankTurtle.Name, frankTurtle.Phrase)
			millieTurtle.SetRandomPhrase()
		case "How r u today?":
			fmt.Printf("%s says: %s\n", frankTurtle.Name, frankTurtle.Feeling)
			millieTurtle.SetRandomPhrase()
		case "What's ur favorite color?":
			fmt.Printf("%s says: %s\n", frankTurtle.Name, frankTurtle.Color)
			millieTurtle.SetRandomPhrase()
		case "How big r u?":
			fmt.Printf("%s says: %d\n", frankTurtle.Name, frankTurtle.Size)
			millieTurtle.SetRandomPhrase()
		case "Where r u?":
			fmt.Printf("%s Cordinations X:%d Y:%d Z:%d\n", frankTurtle.Name, frankTurtle.Position.X, frankTurtle.Position.Y, frankTurtle.Position.Z)
			millieTurtle.SetRandomPhrase()
		}
	}
}