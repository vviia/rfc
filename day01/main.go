package main

import "fmt"

func main() {

	//hitung maju
	fmt.Println("HITUNG MAJU")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}

	// hitungan mundur
	fmt.Println("HITUNG MUNDUR")
	for i := 20; i > 0; i-- {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
