package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// array with 10 elements
	// var arr = [10]int{15, 47, 12, 4, 31, 18, 7, 8, 23, 10}
	var arr = [10]int{40, 40, 40, 40, 40, 49, 49, 49, 49, 49}

	var subsets = [][]int{}

	var target = 50
	var result = [][]int{}

	// print the arr
	fmt.Println("Array :", arr)

	length := len(arr)

	// getting all combinations of the array
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int
		for object := 0; object < length; object++ {
			// ignore subset with one element
			if bits.OnesCount(uint(subsetBits)) == 1 {
				continue
			}

			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, arr[object])
			}
		}

		if len(subset) > 0 {
			subsets = append(subsets, subset)
		}
	}

	// getting the sum of each combination
	for _, subset := range subsets {
		sum := 0
		for _, element := range subset {
			sum += element
		}

		if sum == target {
			result = append(result, subset)
		}
	}

	fmt.Printf("Kombinasi semua element yang dapat menghasilkan target %d adalah:", target)
	fmt.Println()
	if len(result) > 0 {

		// print each element in result
		for _, v := range result {
			fmt.Println(v)
		}
	} else {
		fmt.Println("TIDAK ADA")
	}

}
