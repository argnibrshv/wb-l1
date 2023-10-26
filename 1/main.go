package main

import "fmt"

// определяем структуру Human  и 4 метода для нее
type Human struct {
	name string
}

func (h Human) eat(food string) {
	fmt.Printf("%s eat %s\n", h.name, food)
}

func (h Human) Sleep() {
	fmt.Printf("%s sleep\n", h.name)
}

func (h Human) SameName() {
	fmt.Println("Called Human SameName func")
}

func (h Human) CallSameName() {
	h.SameName()
}

// определяем струкруту Action со встроенным полем типа Human.
type Action struct {
	Human
}

func (a Action) SameName() {
	fmt.Println("Called Action SameName func")
}

func main() {
	action := Action{
		Human{
			name: "Bob",
		},
	}

	//	Любые поля или методы, объявленные во встроенном поле, повышаются
	//	до содержащей его структуры и могут быть вызваны непосредственно в ней.
	action.eat("someFood") // Bob eat someFood
	action.Sleep()         // Bob sleep

	// 	Если вмещающая структура имеет поля или методы с таким же именем, как
	// 	у встроенного поля, то для обращения к этим затененным полям и методам
	// 	следует использовать тип встроенного поля.
	action.SameName()       // Called Action SameName func
	action.Human.SameName() // Called Human SameName func

	// Методы встроенного поля не знают о том, что они являются встроенными.
	// Если метод встроенного поля будет вызывать другой метод этого встроенного
	// поля и вмещающая структура будет обладать методом с таким же именем, то
	// метод встроенного поля не будет вызывать метод вмещающей структуры,
	// а вызовет метод встроенного поля.
	// Метод CallSameName вызывает метод SameName встроенного поля Human,
	// а не метод SameName вмещающей структуры Action.
	action.CallSameName() // Called Human SameName func
}
