package main

import "fmt"

func main() {

	var a int = 0
	var x int = 0
	var nilai int = 0
	a = 10
	if a <= 1 {
		fmt.Println("Masukkan angka mulai dari 2.")
	} else {
		if a == 2 {
			fmt.Println(" merupakan bilangan prima. %d", a)
		} else {
			for x = 2; x < a; x + 1 {
				nilai = x
				if nilai%2 == 1 {
					fmt.Println("%d adalah bilangan ganjil\n", nilai)
					if nilai <= 10 {
						if nilai == 3 || nilai%nilai == 0 && nilai%3 != 0 { //formula bilangan prima
							fmt.Println("%d adalah bilangan prima\n", nilai)
						}
					} else if nilai > 10 {
						if nilai == 3 || nilai%nilai == 0 && nilai%3 != 0 && nilai%5 != 0 && nilai%7 != 0 {
							fmt.Println("%d adalah bilangan prima \n", nilai)
						}
					}

				} else {
					fmt.Println("%d adalah bilangan genap\n", nilai)
					if nilai == 2 {
						fmt.Println("%d adalah bilangan prima \n", nilai)
					}
				}

			}
		}
	}
}
