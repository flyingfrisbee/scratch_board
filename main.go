package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// pen bisa tambah krg dsb + - & / * yes
	// fmt.Scanln() buat nerima input string & operator yes
	// cari tau operator yang dipanggil apa (if)
	// string di convert to int and float64
	// return result fmt.println
	// semua harus di loop

	for {
		var angka1 string
		var angka2 string
		var operator string

		fmt.Println("masukkan angka pertama: ")
		fmt.Scanln(&angka1)

		fmt.Println("masukkan angka kedua: ")
		fmt.Scanln(&angka2)

		fmt.Println("masukkan operator: ")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			angkaSatu, err := strconv.Atoi(angka1)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}

			angkaDua, err := strconv.Atoi(angka2)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}
			fmt.Println(add(angkaSatu, angkaDua))

		case "-":
			angkaSatu, err := strconv.Atoi(angka1)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}

			angkaDua, err := strconv.Atoi(angka2)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}
			fmt.Println(substract(angkaSatu, angkaDua))
		case "%":
			angkaSatu, err := strconv.Atoi(angka1)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}

			angkaDua, err := strconv.Atoi(angka2)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}
			fmt.Println(modulus(angkaSatu, angkaDua))

		case "/":
			angkaSatu, err := strconv.ParseFloat(angka1, 64)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}

			angkaDua, err := strconv.ParseFloat(angka2, 64)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}
			fmt.Println(divide(angkaSatu, angkaDua))

		case "*":
			angkaSatu, err := strconv.ParseFloat(angka1, 64)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}

			angkaDua, err := strconv.ParseFloat(angka2, 64)
			if err != nil { //if error is not non exist
				log.Println(err)
				return
			}
			fmt.Println(multiply(angkaSatu, angkaDua))
		}
	}
}

func add(value1, value2 int) int {
	return value1 + value2
}

func substract(value1, value2 int) int {
	return value1 - value2
}

func modulus(value1, value2 int) int {
	return value1 % value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}
