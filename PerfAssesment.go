package main

import (
	"fmt"
)

// Main Menu
func mainMenu() {
	var choice int

	for {
		fmt.Println("======================================")
		fmt.Println("Employee Performance Assesment")
		fmt.Println("======================================")
		fmt.Println("1. Kelola Data Karyawan")
		fmt.Println("2. Kelola Log Pekerjaan Karyawan")
		fmt.Println("3. Lihat Rekap Data")
		fmt.Println("0. Exit")
		fmt.Println("======================================")
		fmt.Print("Pilih menu: ")

		fmt.Scan(&choice)

		if choice == 1 {
			menuKaryawan()
		} else if choice == 2 {
			// to Log
		} else if choice == 3 {
			// to Rekap
		} else if choice == 0 {
			break // break (exit program)
		} else {
			fmt.Println("Pilihan Invalid.")
		}

	}
}

// Submenu Karyawan
func menuKaryawan() {
	var choice int

	for {
		fmt.Println("======================================")
		fmt.Println("Kelola Data Karyawan")
		fmt.Println("======================================")
		fmt.Println("1. Add Data Karyawan")
		fmt.Println("2. Update Data Karyawan")
		fmt.Println("3. Delete Data Karyawan")
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)

	}
}

func menuLogPekerjaan() {
	var choice int

	for {
		fmt.Println("======================================")
		fmt.Println("Kelola Log Pekerjaan Karyawan")
		fmt.Println("======================================")
		fmt.Println("1. Tambah Log Pekerjaan Karyawan")
		fmt.Println("2. Ubah Log Pekerjaan Karyawan")
		fmt.Println("3. Hapus Log Pekerjaan Karyawan")
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)
	}
}

func main() {
	mainMenu()
}
