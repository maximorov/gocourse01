package main

import (
	"fmt"
)

type region string
type method string

const (
	Europe  region = "e"
	Asia    region = "s"
	America region = "a"

	Plane method = `p`
	Train method = `t`
	Bus   method = `b`
)

func main() {
	fmt.Printf("Choose region")
	var reg string
	_, _ = fmt.Scan(&reg)
	fmt.Printf("Choose method")
	var meth string
	_, _ = fmt.Scan(&meth)

	switch reg {
	case string(Europe):
		fmt.Printf("Let's go to Europe!\n")
	case string(Asia):
		fmt.Printf("Let's go to Asia!\n")
	case string(America):
		fmt.Printf("Let's go to America!\n")
	default:
		fmt.Printf("No no no. Which region?\n")
	}
	switch meth {
	case string(Bus):
		fmt.Printf("Bus is default method!\n")
	case string(Train):
		fmt.Printf("Train? Ok!\n")
	case string(Plane):
		fmt.Printf("Plane? Ok!\n")
	default:
		fmt.Printf("Bus is default method!\n")
	}
}
