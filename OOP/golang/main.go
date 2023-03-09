package main

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name + " โฮ่งโฮ่ง")
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name + " เมี๊ยวเมี๊ยว")
}

func main() {
	dog := Dog{Name: "Rufus"}
	cat := Cat{Name: "Garfield"}

	animals := []Animal{dog, cat}

	for _, animal := range animals {
		animal.Speak()
	}
}
