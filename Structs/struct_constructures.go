package main

import "fmt"

// Aircraft abstract
type Aircraft struct {
	name  string
	fuel  int
	speed int
}

// FighterJet defined struct
type FighterJet struct {
	weapons []string
	Aircraft
}

// NewFighterJet return a new FighterJet from name, fuel, speed, weapons
func NewFighterJet(name string, fuel int, speed int, weapons []string) *FighterJet {
	return &FighterJet{weapons: weapons,
		Aircraft: Aircraft{
			name:  name,
			fuel:  fuel,
			speed: speed,
		},
	}
}

func main() {

	mig29 := NewFighterJet("MiG-29", 600, 900, []string{"missiles", "minigun"})

	fmt.Println(*mig29)

}
