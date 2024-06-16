package main

import "fmt"

type Driver struct {
	name string
}

type Pilot struct {
	name string
}

type Adapter struct {
	Driver
}

type fly interface {
	Fly()
}

func (a *Adapter) Fly() {
	fmt.Println("Flying on the plain")
}

func (d *Driver) drive() {
	fmt.Println("driving a car")
}

func main() {
	ob := Driver{"Oleg"}
	ob.drive()

	ad := Adapter{ob}
	ad.Fly()

}
