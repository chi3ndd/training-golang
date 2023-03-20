package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	a := []string{"a", "b", "c", "d"}
	for _, value := range a {
		fmt.Println(value)
	}

	b := map[string]string{"a": "b", "c": "d"}
	for key, value := range b {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
