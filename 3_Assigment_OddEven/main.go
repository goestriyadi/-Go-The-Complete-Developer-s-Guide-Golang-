package main

import "fmt"

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range n {
		if n[i]%2 == 0 {
			fmt.Println(n[i], " is even")
		} else {
			fmt.Println(n[i], " is odd")
		}
	}
}
