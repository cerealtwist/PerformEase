package main

import (
	"fmt"
)

func mainMenu() {
	var choice int

	for {
		fmt.Println("======================================")
		fmt.Println("Aplikasi Pengolahan Data Kinerja Karyawan")
		fmt.Println("======================================")
		fmt.Println("1. Kelola Data Karyawan")
		fmt.Println("2. Kelola Log Pekerjaan Karyawan")
		fmt.Println("3. Lihat Rekap Data")
		fmt.Println("0. Exit")
		fmt.Println("======================================")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&choice)

	}
}

func main() {
	mainMenu()
}
