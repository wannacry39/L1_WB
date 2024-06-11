package main

import "fmt"

type Human struct {
	Name   string
	Age    uint
	Weight float32
	Height uint
}

func (h *Human) Talk() {
	fmt.Println("Hello, friend")
}

func (h *Human) DoWork() {
	fmt.Println("I'm working right now")
}

type Action struct {
	Human //Встраиваем структуру Human
}

func main() {

	act := Action{Human{}}
	act.Talk()
	act.DoWork()

}
