# Go Embedded Structs Example

Пример использования **композиции (embedded struct)** в Go как аналога наследования.

## Описание

- `Human` — базовая структура с полями и методами.
- `Action` — включает (`embed`) `Human`, поэтому наследует его поля и методы.
- Методы можно переопределять в `Action` (они затеняют унаследованные).
- Доступ к оригинальному методу родителя возможен через `a.Human.Method()`.

## Код

```go
package main

import "fmt"

type Human struct {
    Name string
    Age  int
}

func (h Human) Talk() {
    fmt.Printf("Меня зовут %s, мне %d.\n", h.Name, h.Age)
}

func (h *Human) Birthday() {
    h.Age++
}

type Action struct {
    Human
    CurrentDoing string
}

func (a Action) Do() {
    fmt.Printf("%s делает: %s\n", a.Name, a.CurrentDoing)
}

func (a Action) Talk() {
    fmt.Printf("Action[%s]: я занят — %s!\n", a.Name, a.CurrentDoing)
}

func main() {
    a := Action{
        Human:        Human{Name: "Аружан", Age: 20},
        CurrentDoing: "пишу код",
    }

    a.Talk()
    a.Human.Talk()
    a.Birthday()
    a.Do()

    fmt.Println("Возраст теперь:", a.Age)
}
