// оголошення, ініціалізація та привласнення
package main

import "fmt"

type (
	myInt int64 // –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807, 8 байт (64 бита)
	color string
)

var (
	globalX = `x`
	globaly = `y`
	globalz = `z`
)

const (
	ColorBlack  color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

var myIntVar myInt = 100

var someInterface1 any = 150.2
var someInterface2 interface{} = `Abadalabada`

func main() {
	var someFlag bool
	someFlag = true
	fmt.Println(someFlag)

	// several
	// var width, height int = 100, 50
	// var (
	// 	x    = 1
	// 	y    = 2
	// 	z    = 3
	// 	form = `square`
	// )
	// short
	// name, age := "Maks", 33
	// namePointer := &name
	// fmt.Println(namePointer)

	// one
	// var size = `big`
	// isCool := true
	// volume := 1.1

	// pointer
	// hidden := new(rune)
	// *hidden = '😊'
	// fmt.Printf("Emoji %c, i.e %s, lives in %p\n", *hidden, string(*hidden), hidden)

	// hidden = nil
	// fmt.Printf("Smile is gone: %p\n", hidden)

	// nope
	// volume = false
	//newDefinition = 0.1

	// using
	// fmt.Printf(`%d %d\n`, width, height)
	// fmt.Printf(`%d %d %d %s\n`, x, y, z, form)

	// fmt.Printf(`My name is %s and i'm %d\n`, name, age)
	// fmt.Printf(`%s %t %f\n`, size, isCool, volume)
}
