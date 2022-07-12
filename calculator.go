package main

import (
	"flag"
	"fmt"
)

func tambah(x float64, y float64) {
	var result float64 = x + y
	fmt.Printf("%.0f + %.0f = %.0f \n", x, y, result)
}

func kurang(x float64, y float64) {
	var result float64 = x - y
	fmt.Printf("%.0f - %.0f = %.0f \n", x, y, result)
}

func kali(x float64, y float64) {
	var result float64 = x * y
	fmt.Printf("%.0f x %.0f = %.0f \n", x, y, result)
}

func bagi(x float64, y float64) {
	var result float64 = x / y
	fmt.Printf("%.1f : %.1f = %.1f \n", x, y, result)
}

func main() {
	fmt.Printf("Task 1 - Calculator \n \n")

	var angka1 *float64 = flag.Float64("angka1", 0, "tulis angka 1")
	var opt *string = flag.String("opt", "", "tulis operator")
	var angka2 *float64 = flag.Float64("angka2", 0, "tulis angka 2")

	flag.Parse()

	var a1 float64 = *angka1
	var a2 float64 = *angka2

	switch *opt {
	case "+":
		tambah(a1, a2)
		break
	case "-":
		kurang(a1, a2)
		break
	case "*":
		kali(a1, a2)
		break
	case "/":
		bagi(a1, a2)
		break
	default:
		fmt.Println("Masukkan input dengan benar, contoh \n -angka1=int -opt=\"simbol operator\" -angka2=int ")
	}

	fmt.Println("\n \n Task Done")

}
