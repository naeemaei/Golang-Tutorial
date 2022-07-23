package main

import "fmt"

type Runner interface {
	Run()
}

type Walker interface {
	Walk()
}

type Shooter interface {
	Shoot()
}

type ShooterRunner interface {
	Runner
	Shooter
}

type Player struct {
	Name     string
	Age      int
	Height   int
	Weight   int
	Position string
}

func main() {

	player1 := &Player{
		Name:     "Alireza",
		Age:      30,
		Height:   180,
		Weight:   70,
		Position: "Forward",
	}

	var runner ShooterRunner = player1

	var shooter ShooterRunner = player1

	runner.Run()

	shooter.Shoot()
}

func (player *Player) Run() {
	fmt.Printf("name: %s, position: %s , Player is running\n",player.Name,player.Position)
}
func (player *Player) Shoot() {
	fmt.Printf("name: %s, position: %s , Player is shooting",player.Name,player.Position)
}