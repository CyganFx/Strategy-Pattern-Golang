//Strategy Pattern with Ducks
package main

import (
	"fmt"
)

type Quackable interface {
	Quack()
}

type Flyable interface {
	Fly()
}

type IDuck interface {
	PerformQuack()
	PerformFly()
	Display()
}

//Make it private, as abstract imitation
type duck struct {
	IDuck
	quackable Quackable
	flyable   Flyable
}

func (duck *duck) SetFlyBehaviour(flyable Flyable) {
	duck.flyable = flyable
}

func (duck *duck) SetQuackBehaviour(quackable Quackable) {
	duck.quackable = quackable
}

func (duck duck) PerformQuack() {
	duck.quackable.Quack()
}

func (duck duck) PerformFly() {
	duck.flyable.Fly()
}

func (duck duck) Display() {}

//Ducks:

type JetpackedDuck struct {
	duck duck
}

func NewJetpackedDuck() *JetpackedDuck {
	return &JetpackedDuck{duck: duck{quackable: NewNormalQuack(), flyable: NewFlyRocketPowered()}}
}

func (jetpackedDuck JetpackedDuck) Display() {
	fmt.Println("It's a jetpacked duck!")
}

type RubberDuck struct {
	duck duck
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{duck: duck{quackable: NewSqueakSound(), flyable: NewFlyNoWay()}}
}

func (rubberDuck RubberDuck) Display() {
	fmt.Println("It's a rubber duck!")
}

type WoodenDuck struct {
	duck duck
}

func NewWoodenDuck() *WoodenDuck {
	return &WoodenDuck{duck: duck{quackable: NewMutedQuack(), flyable: NewFlyNoWay()}}
}

func (woodenDuck WoodenDuck) Display() {
	fmt.Println("It's a wooden duck!")
}

//Quack Behaviour Implementation:

type NormalQuack struct{}

func NewNormalQuack() Quackable {
	return NormalQuack{}
}

func (normalQuack NormalQuack) Quack() {
	fmt.Println("Quack!")
}

type SqueakSound struct{}

func NewSqueakSound() Quackable {
	return SqueakSound{}
}

func (squeakSound SqueakSound) Quack() {
	fmt.Println("Squeak!")
}

type MutedQuack struct{}

func NewMutedQuack() Quackable {
	return MutedQuack{}
}

func (mutedQuack MutedQuack) Quack() {
	fmt.Println("...")
}

//Fly Behaviour Implementation:

type FlyWithWings struct{}

func NewFlyWithWings() Flyable {
	return FlyWithWings{}
}

func (flyWithWings FlyWithWings) Fly() {
	fmt.Println("Flying with wings!")
}

type FlyRocketPowered struct{}

func NewFlyRocketPowered() Flyable {
	return FlyRocketPowered{}
}

func (flyRocketPowered FlyRocketPowered) Fly() {
	fmt.Println("Flying with rocket power!")
}

type FlyNoWay struct{}

func NewFlyNoWay() Flyable {
	return FlyNoWay{}
}

func (flyNoWay FlyNoWay) Fly() {
	fmt.Println("Not flying, sorry!")
}

func main() {
	jetpackDuck := NewJetpackedDuck()
	jetpackDuck.Display()
	jetpackDuck.duck.PerformQuack()
	jetpackDuck.duck.PerformFly()

	jetpackDuck.duck.SetFlyBehaviour(NewFlyWithWings())
	jetpackDuck.duck.SetQuackBehaviour(NewMutedQuack())

	jetpackDuck.duck.PerformQuack()
	jetpackDuck.duck.PerformFly()

	rubberDuck := NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.duck.PerformQuack()
	rubberDuck.duck.PerformFly()

	woodenDuck := NewWoodenDuck()
	woodenDuck.Display()
	woodenDuck.duck.PerformQuack()
	woodenDuck.duck.PerformFly()
}
