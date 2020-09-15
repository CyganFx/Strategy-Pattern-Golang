package main

import "fmt"

type Quackable interface {
	Quack()
}

type Flyable interface {
	Fly()
}

type CustomDuck struct {
	Quackable
	Flyable
}

func NewCustomDuck(quackBehaviour Quackable, flyBehaviour Flyable) CustomDuck {
	return CustomDuck{quackBehaviour, flyBehaviour}
}

func (customDuck CustomDuck) Display() {
	fmt.Println("It's a custom duck!")
}

type JetpackedDuck struct {
	Quackable
	Flyable
}

func NewJetpackedDuck() JetpackedDuck {
	return JetpackedDuck{NewNormalQuack(), NewFlyRocketPowered()}
}

func (jetpackedDuck JetpackedDuck) Display() {
	fmt.Println("It's a jetpacked duck duck!")
}

type RubberDuck struct {
	Quackable
	Flyable
}

func NewRubberDuck() RubberDuck {
	return RubberDuck{NewSqueakSound(), NewFlyNoWay()}
}

func (rubberDuck RubberDuck) Display() {
	fmt.Println("It's a rubber duck!")
}

type WoodenDuck struct {
	Quackable
	Flyable
}

func NewWoodenDuck() WoodenDuck {
	return WoodenDuck{NewMutedQuack(), NewFlyNoWay()}
}

func (woodenDuck WoodenDuck) Display() {
	fmt.Println("It's a wooden duck!")
}

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
	customDuck := NewCustomDuck(NewNormalQuack(), NewFlyWithWings())
	customDuck.Display()
	customDuck.Fly()
	customDuck.Quack()

	jetpackDuck := NewJetpackedDuck()
	jetpackDuck.Display()
	jetpackDuck.Fly()
	jetpackDuck.Quack()

	rubberDuck := NewRubberDuck()
	rubberDuck.Display()
	rubberDuck.Fly()
	rubberDuck.Quack()

	woodenDuck := NewWoodenDuck()
	woodenDuck.Display()
	woodenDuck.Fly()
	woodenDuck.Quack()
}
