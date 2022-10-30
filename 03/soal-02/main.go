package main

import "fmt"

func main() {
	// a, b, c is a side of a triangle
	var a, b, c int
	var triangleType string

	fmt.Println("Masukkan nilai A : ")
	fmt.Scanln(&a)
	fmt.Println("Masukkan nilai B : ")
	fmt.Scanln(&b)
	fmt.Println("Masukkan nilai C : ")
	fmt.Scanln(&c)

	// Check the triangle type
	if a == b && b == c {
		triangleType = "SAMA SISI"
	} else if a == b && b != c || a == c && c != b || b == c && c != a {
		triangleType = "SAMA KAKI"
	} else if a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b {
		triangleType = "SIKU-SIKU"
	} else {
		triangleType = "SEMBARANG"
	}

	fmt.Println("Segitiga dengan sisi-sisi", a, b, c, "bertipe", triangleType)

}
