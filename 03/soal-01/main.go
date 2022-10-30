package main

import "fmt"

func main() {

	var a, b, c int
	var medianValue int

	fmt.Println("Masukkan nilai A : ")
	fmt.Scanln(&a)
	fmt.Println("Masukkan nilai B : ")
	fmt.Scanln(&b)
	fmt.Println("Masukkan nilai C : ")
	fmt.Scanln(&c)

	var median = 'A'
	medianValue = a

	if a > b {
		if b > c { // a > b > c
			median = 'B'
			medianValue = b
		} else {
			if a > c { // a > c > b
				median = 'C'
				medianValue = c
			}
		}
	} else { // b > a
		if b > c {
			if c > a { // b > c > a
				median = 'C'
				medianValue = c
			}
		} else { // c > b > a
			median = 'B'
			medianValue = b
		}
	}

	fmt.Println("Nilai median adalah", string(median), "dengan nilai = ", medianValue)
}
