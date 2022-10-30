package main

import "fmt"

func main() {
	var totalRobot = 50
	var numA, numB, numC int

	var totalHead = 74
	var totalLeg = 122

	for a := 0; a <= totalRobot; a++ {
		for b := 0; b+a <= totalRobot; b++ {
			for c := 0; c+b+a <= totalRobot; c++ {

				var numLeg = (2 * a) + (3 * b) + (4 * c)
				var numHead = a + (2 * b) + (3 * c)
				if numHead == totalHead && numLeg == totalLeg {
					numA = a
					numB = b
					numC = c
					fmt.Println("Jumlah robot A: ", numA, " Total Kepala A: ", numA, " Total Kaki A: ", numA*2)
					fmt.Println("Jumlah robot B: ", numB, " Total Kepala B: ", numB*2, " Total Kaki B: ", numB*3)
					fmt.Println("Jumlah robot C: ", numC, " Total Kepala C: ", numC*3, " Total Kaki C: ", numC*4)
					fmt.Println("Total robot: ", numA+numB+numC)
					fmt.Println()
				}
			}
		}
	}
}
