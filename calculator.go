package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Printf("Task 1 - Calculator \n \n")

	var angka1 *float64 = flag.Float64("angka1", 0, "tulis angka 1")
	var opt *string = flag.String("opt", "", "tulis operator")
	var angka2 *float64 = flag.Float64("angka2", 0, "tulis angka 2")

	flag.Parse()

	var result float64
	switch *opt {
	case "+":
		result = *angka1 + *angka2
		fmt.Printf("%.0f %s %.0f = %.0f \n", *angka1, *opt, *angka2, result)
		break
	case "-":
		result = *angka1 - *angka2
		fmt.Printf("%.0f %s %.0f = %.0f \n", *angka1, *opt, *angka2, result)
		break
	case "*":
		result = *angka1 * *angka2
		fmt.Printf("%.0f %s %.0f = %.0f \n", *angka1, *opt, *angka2, result)
		break
	case "/":
		result = *angka1 / *angka2
		fmt.Printf("%.1f %s %.1f = %.1f \n", *angka1, *opt, *angka2, result)
		break
	default:
		fmt.Println("Masukkan input dengan benar, contoh \n -angka1=int -opt=\"simbol operator\" -angka2=int ")
	}

	fmt.Println("\n \n Task Done")

}
