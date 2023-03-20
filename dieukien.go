package main

import "fmt"

func main4() {
	a := 3
	b := 2
	c := 3
	// Dieu kieu co ban
	if a == 1 {
		fmt.Println("Dung")
	} else {
		fmt.Println("Sai")
	}
	// Dieu kien phuc tap
	// and: && , or: ||
	if a == 1 && (b == 2 && c == 3) {
		fmt.Println("Dung Dung")
	}
	// Dieu kien lien tiep
	if a != 1 {
		fmt.Println("A khac 1")
	} else if b != 2 {
		fmt.Println("B khac 2")
	} else if c == 3 {
		fmt.Println("C bang 3")
	}
	/*
		Switch
	*/
	switch a {
	case 1:
		fmt.Println("A bang 1")
		break
	case 3:
		fmt.Println("A bang 2")
	default:
		fmt.Println("Cha dung gia tri nao")
	}
}
