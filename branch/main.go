package main

import (
	"fmt"
)

func main() {
	name, age := "Maks", 33

	secret := new(string)
	*secret = "I'm fat"

	if secret == nil {
		fmt.Printf("My name is %s and my only secret that i'm %d\n", name, age)
	} else if *secret == `Nope` {
		fmt.Print("Nope!")
	} else {
		fmt.Printf("My name is %s and I have a secret\n", *secret)
	}

	// exercise
}
