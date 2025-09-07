package main

import "fmt"

// Human Родитель
type Human struct {
	Name string
	Age  int
}

func (h Human) Talk() {
	fmt.Printf("Меня зовут %s, мне %d.\n", h.Name, h.Age)
}

func (h *Human) Birthday() { // меняет состояние — поэтому указатель
	h.Age++
}

// Action «Потомок»: встраиваем Human
type Action struct {
	Human        // <- embedded
	CurrentDoing string
}

func (a Action) Do() {
	fmt.Printf("%s делает: %s\n", a.Name, a.CurrentDoing)
}

// Talk Переопределение: свой метод с тем же именем «затеняет» поднятый
func (a Action) Talk() {
	fmt.Printf("Action[%s]: я занят — %s!\n", a.Name, a.CurrentDoing)
}

func main() {
	a := Action{
		Human:        Human{Name: "Аян", Age: 20},
		CurrentDoing: "Сплю",
	}

	// Методы Human доступны прямо на Action
	a.Talk()       // вызов переопределённого метода Action.Talk
	a.Human.Talk() // явный вызов «родительского» Human.Talk

	a.Birthday() // поднятый метод (*Human).Birthday
	a.Do()

	fmt.Println("Возраст теперь:", a.Age) // поле тоже поднято
}
