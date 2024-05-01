package main

// START OMIT
type Animal interface { // HL
	Speak()
}

type Dog struct{}

func (d Dog) Speak() { println("Woof!") }

type Cat struct{}

func (c Cat) Speak() { println("Meow!") }

func AnimalGreeting(a Animal) { // HL
	a.Speak()
}

func main() {
	var d Dog // Could also say d := Dog{}
	var c Cat

	AnimalGreeting(d)
	AnimalGreeting(c)
}

// END OMIT
