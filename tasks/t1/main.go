package main

import "fmt"

type Human struct {
	Name     string
	Lastname string
	Age      uint16
	ID       string
}

type Action struct {
	Voice string
	Human
}

func NewHuman(name string, lastname string, age uint16, id string) *Human {
	h := &Human{
		Name:     name,
		Lastname: lastname,
		Age:      age,
		ID:       id,
	}
	return h
}

func (h *Human) FullName() string {
	return fmt.Sprintf("%s %s", h.Name, h.Lastname)
}

func NewAction(h *Human, voice string) Action {
	return Action{
		Human: *h,
		Voice: voice,
	}
}

func (a *Action) Do() {
	fmt.Printf("A person named %s (Age: %d) says: %s\n", a.FullName(), a.Age, a.Voice)
}

func main() {
	aaa := NewAction(NewHuman("Vasya", "Pupkin", 25, "aye228test"), "Wake Up Samirai, I pissed my bed")
	aaa.Do()
}
