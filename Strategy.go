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

func (d *duck) SetFlyBehaviour(flyable Flyable) {
	d.flyable = flyable
}

func (d *duck) SetQuackBehaviour(quackable Quackable) {
	d.quackable = quackable
}

func (d duck) PerformQuack() {
	d.quackable.Quack()
}

func (d duck) PerformFly() {
	d.flyable.Fly()
}

func (d duck) Display() {}

//Ducks:

type JetpackedDuck struct {
	duck duck
}

func NewJetpackedDuck() *JetpackedDuck {
	return &JetpackedDuck{duck: duck{quackable: NewNormalQuack(), flyable: NewFlyRocketPowered()}}
}

func (jd JetpackedDuck) Display() {
	fmt.Println("It's a jetpacked duck!")
}

type RubberDuck struct {
	duck duck
}

func NewRubberDuck() *RubberDuck {
	return &RubberDuck{duck: duck{quackable: NewSqueakSound(), flyable: NewFlyNoWay()}}
}

func (rd RubberDuck) Display() {
	fmt.Println("It's a rubber duck!")
}

type WoodenDuck struct {
	duck duck
}

func NewWoodenDuck() *WoodenDuck {
	return &WoodenDuck{duck: duck{quackable: NewMutedQuack(), flyable: NewFlyNoWay()}}
}

func (wd WoodenDuck) Display() {
	fmt.Println("It's a wooden duck!")
}

//Quack Behaviour Implementation:

type NormalQuack struct{}

func NewNormalQuack() Quackable {
	return NormalQuack{}
}

func (nq NormalQuack) Quack() {
	fmt.Println("Quack!")
}

type SqueakSound struct{}

func NewSqueakSound() Quackable {
	return SqueakSound{}
}

func (ss SqueakSound) Quack() {
	fmt.Println("Squeak!")
}

type MutedQuack struct{}

func NewMutedQuack() Quackable {
	return MutedQuack{}
}

func (mq MutedQuack) Quack() {
	fmt.Println("...")
}

//Fly Behaviour Implementation:

type FlyWithWings struct{}

func NewFlyWithWings() Flyable {
	return FlyWithWings{}
}

func (fww FlyWithWings) Fly() {
	fmt.Println("Flying with wings!")
}

type FlyRocketPowered struct{}

func NewFlyRocketPowered() Flyable {
	return FlyRocketPowered{}
}

func (frp FlyRocketPowered) Fly() {
	fmt.Println("Flying with rocket power!")
}

type FlyNoWay struct{}

func NewFlyNoWay() Flyable {
	return FlyNoWay{}
}

func (fnw FlyNoWay) Fly() {
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
