package main

import "fmt"

func main() {
	var score int
	_, _ = fmt.Scan(&score)

	grade := 'F'
	if score >= 90 {
		grade = 'A'
	} else if score >= 80 {
		grade = 'B'
	} else if score >= 70 {
		grade = 'C'
	} else if score >= 60 {
		grade = 'D'
	}
	fmt.Printf("%c", grade)
}
