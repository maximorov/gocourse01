package main

import (
	"flag"
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
	reg := flag.String("region", ``, "Region")
	meth := flag.String("method", ``, "Method")
	help := flag.Bool("help", false, "Help")
	flag.Parse()

	if *help {
		fmt.Print("Run program with region flag:\n" +
			"main.go -region [e|s|a] \n" +
			"and optional -method [p|t|b]\n" +
			"For example: go run branchflags/main.go --region a --method t")
	} else {
		switch *reg {
		case string(Europe):
			fmt.Printf("Let's go to Europe!\n")
		case string(Asia):
			fmt.Printf("Let's go to Asia!\n")
		case string(America):
			fmt.Printf("Let's go to America!\n")
		default:
			fmt.Printf("No no no. Which region?\n")
		}
		switch *meth {
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

	// exercise
}
