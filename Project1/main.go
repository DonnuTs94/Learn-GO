package main

import "fmt"

func main() {
	// fmt.Println("Hi there!")

	// var firstName string = "John"

	// var lastName string = "wick"

	// fmt.Printf("halo %s %s! \n", firstName, lastName)

	// fmt.Print("Hello", firstName, lastName)

	// fmt.Println("Hello", firstName, lastName+"!")

	// fmt.Println("Nilai variable name:", firstName, lastName)
	// fmt.Println("Alamat memori dari:", ptr)
	// fmt.Println("Nilai yang ditunjuk oleh pointer:", *ptr)

	// exist := true
	// fmt.Printf("exist? %t \n", exist)
	// fmt.Println("exist?", exist)

	// message := `Nama saya "John Wick"
	// Salam kenal.
	// Mari belajar "Golang".`
	// fmt.Println(message)

	// const name = "John"
	// _ := "John"

	// const (
	// 	str     = "name"
	// 	int     = 123
	// 	current = true
	// )

	// // fmt.Println(name, nama)

	// fmt.Println(str, int, current)

	// point := 8840.0

	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect! \n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }

	// point := 6

	// switch {
	// case point == 8:
	// 	fmt.Println("Perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("Awesome")
	// 	fallthrough
	// default:
	// 	{
	// 		fmt.Println("not Bad")
	// 		fmt.Println("You need to learn more")
	// 	}
	// }

	point := 3

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("Not bad")
		} else if point == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("You can do it")
			if point == 0 {
				fmt.Println("Try harder!")
			}
		}
	}
}
