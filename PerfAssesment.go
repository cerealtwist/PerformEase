package main

import (
	"fmt"
)

const NMAX int = 3 // blm pasti pake atau enggak, just keep it this way dulu dah

type Karyawan struct {
	ID   int
	Nama string
}

type Pekerjaan struct {
	KaryawanID, Tipe, Durasi int
}

type arrKaryawan [NMAX]Karyawan
type arrPekerjaan []Pekerjaan

// PROGRAM MAIN
func main() {
	mainMenu()
}

// Main Menu
func mainMenu() {
	var choice int

	for {
		fmt.Println("\n======================================")
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
			menuLogPekerjaan()
		} else if choice == 3 {
			menuRekap()
		} else if choice == 0 {
			break // break (exit program)
		} else {
			fmt.Println("Pilihan Invalid.")
		}

	}
}

// Submenu Karyawan
func menuKaryawan() {
	var choice, n int
	var A arrKaryawan

	for {
		fmt.Println("\n======================================")
		fmt.Println("Kelola Data Karyawan")
		fmt.Println("======================================")
		fmt.Println("1. Add Data Karyawan")
		fmt.Println("2. Update Data Karyawan")
		fmt.Println("3. Delete Data Karyawan")
		fmt.Println("4. Show List Karyawan") // Buat testing aja ngecek apakah data karyawan masuk
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)

		if choice == 1 {
			addKaryawan(&A, &n)
		} else if choice == 2 {
			updateKaryawan(&A, &n)
		} else if choice == 3 {
			deleteKaryawan(&A, &n)
		} else if choice == 4 {
			showListKaryawan(A, n)
		} else if choice == 0 {
			break
		} else {
			fmt.Println("Pilihan Invalid.")
		}

	}
}

func menuLogPekerjaan() {
	var choice int

	for {
		fmt.Println("\n======================================")
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

func menuRekap() {
	var choice int

	for {
		fmt.Println("\n======================================")
		fmt.Println("Lihat Rekap Data")
		fmt.Println("======================================")
		fmt.Println("1. Lihat Rekap Log Pekerjaan Karyawan")
		fmt.Println("2. Lihat Rekap Pekerjaan per Tipe")
		fmt.Println("3. Lihat Rekap Aktivitas Berdasarkan Durasi")
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)
	}
}

// BAGIAN CRUD KARYAWAN

func addKaryawan(A *arrKaryawan, n *int) {
	var Nama string

	if *n >= NMAX {
		fmt.Println("Jumlah karyawan mencapai batas.")
		return
	}

	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scan(&Nama)

	A[*n] = Karyawan{ID: *n + 1, Nama: Nama}
	*n++
}

func updateKaryawan(A *arrKaryawan, n *int) {
	var ID int
	var newNama string

	fmt.Print("Masukkan ID karyawan:")
	fmt.Scan(&ID)

	for i := 0; i < *n; i++ {
		if A[i].ID == ID {
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&newNama)
			A[i].Nama = newNama
			fmt.Println("Nama terupdate.")
			return
		}
	}
	fmt.Println("Karyawan tidak ditemukan.")
}

func deleteKaryawan(A *arrKaryawan, n *int) {
	var ID int

	fmt.Print("Masukkan ID karyawan:")
	fmt.Scan(&ID)

	for i := 0; i < *n; i++ {
		if A[i].ID == ID {
			// geser elemen ke kiri utk overwrite elemen terhapus.
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
				A[j].ID = j + 1
			}
			*n--
			fmt.Println("Karyawan dihapus.")
			return
		}
	}
	fmt.Println("Karyawan tidak ditemukan.")
}

func showListKaryawan(A arrKaryawan, n int) {
	fmt.Println("\n======================================")
	fmt.Println("List Karyawan:")
	fmt.Println("======================================")
	if n == 0 {
		fmt.Println("Tidak ada karyawan.")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("ID: %d, Nama: %s\n", A[i].ID, A[i].Nama)
		}
	}
	fmt.Println("======================================")
}
