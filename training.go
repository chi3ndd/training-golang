package main

import "fmt"

func main3() {
	// Khai bao biet
	// Cach 1: var <ten bien> kieu du lieu
	var a string
	a = "Ahihi"
	fmt.Println(a)
	var b int
	fmt.Println(b)
	var c bool
	fmt.Println(c)
	// Cach 2: vua khai bao vua gan luon     <tenbien> := giatri
	d := true
	fmt.Println(d)
	e := 1.2
	fmt.Println(e)
	f := 1
	fmt.Println(f)
	// Mo trong: khai bao nhieu bien cung 1 dong: <tenbien1>, <tenbien2> := giatri1, giatri2
	name, age := "Long", 27
	fmt.Println(name)
	fmt.Println(age)
	// Bien da khai bao bang var hoac := thi muon thay doi gia tri, chi can dung =, ko duoc dung :=
	d = false
	b = 10
	// Bo qua bien ko can thiet, dung dau _
	name, _ = A()

	chienIu := "A"
	fmt.Println(chienIu)

	const (
		ChienIu = "B"
		LongIu  = "C"
	)
}

func A() (string, int64) {
	return "long", 27
}
