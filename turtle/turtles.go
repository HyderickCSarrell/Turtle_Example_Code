package turtle

import (
	"fmt"
	"math/rand"
)

var (
	questions = []string{"How r u today?", "What's ur favorite color?", "How big r u?", "Where r u?"}
)

type Turtle struct{
	Name string
	Color string
	Position *pos
	Size int
	Speed int
	Direction string
	Phrase string
	Feeling string
	Miscellaneous map[string]string
}

type pos struct{
	X int
	Y int
	Z int
}

func NewTurtle(name string, color string) *Turtle {
	turtle := &Turtle{
		Name: name,
		Color: color,
		Position: &pos{
			X: 0,
			Y: 0,
			Z:0,
		},
		Size: 10,
		Speed: 5,
		Direction: "Standing",
		Phrase: "Hello World!",
		Feeling: "Good",
		Miscellaneous: map[string]string{"Fav_Color": "blue", "Toes": "Four"},
	}

	return turtle
}

func (t *Turtle) Move(steps int) {

	distance := steps * t.Speed

	switch t.Direction {
	case "Right":
		t.Position.X += distance
	case "Left":
		t.Position.X -= distance
	case "Forward":
		t.Position.Y += distance
	case "Backward":
		t.Position.Y -= distance
	case "Up":
		t.Position.Z += distance
	case "Down":
		t.Position.Z -= distance
	}
}

func (t *Turtle) Speak(phrase string) {
	fmt.Printf("%s says %s", t.Name, phrase)
}

func (t *Turtle) SpeakPhrase() {
	fmt.Printf("%s says %s", t.Name, t.Phrase)
}

func (t *Turtle) SetRandomPhrase() {
	t.Phrase = questions[rand.Intn(len(questions))]

}