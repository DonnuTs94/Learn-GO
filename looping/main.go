package main

import "fmt"

func main() {
	// xs := "123"
	// for i, x := range xs {
	// 	fmt.Println("Index=", i, "Value=", x)
	// }

	// arr := [5]int{10, 20, 30, 40, 50}
	// for _, v := range arr {
	// 	fmt.Println("Value", v)
	// }

	// arr1 := map[byte]int{'a': 0, 'b': 1, 'c': 2}
	// for i, v := range arr1 {
	// 	fmt.Println("Key:", i, "Value:", v)
	// }

	// for range arr1 {
	// 	fmt.Println("Done")
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Number", i)
	// }

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

}
