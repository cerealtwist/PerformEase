package main

import (
	"fmt"
)

const NMAX int = 3       // blm pasti pake atau enggak, just keep it this way dulu dah
const TotalJAM int = 160 // total jam kerja 160

type Karyawan struct {
	ID   int
	Nama string
}

type Pekerjaan struct {
	KaryawanID, Tipe, Durasi int
}

type arrKaryawan [NMAX]Karyawan
type arrPekerjaan [NMAX]Pekerjaan

// buat global scope agar saat keluar dari submenu tidak terinisialisasi ulang
var A arrKaryawan
var n int

var T arrPekerjaan
var nlog int

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
	var choice int

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
		fmt.Println("4. Tampilkan Log Pekerjaan Karyawan")
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)

		if choice == 1 {
			addPekerjaan(&T, &nlog)
		} else if choice == 2 {
			// to edit
		} else if choice == 3 {
			// to delete
		} else if choice == 4 {
			showLogPekerjaan(T, nlog)
		} else if choice == 0 {
			break
		} else {
			fmt.Println("Pilihan Invalid.")
		}
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
	fmt.Println("Karyawan berhasil ditambah.")
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
	fmt.Println("======================================")1
	if n == 0 {
		fmt.Println("Tidak ada karyawan.")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("ID: %d, Nama: %s\n", A[i].ID, A[i].Nama)
		}
	}
	fmt.Println("======================================")
}

// BAGIAN CRUD KARYAWAN

func addPekerjaan(T *arrPekerjaan, nlog *int) {
	var KaryawanID, Tipe, Durasi int

	if *nlog >= NMAX {
		fmt.Println("Jumlah log pekerjaan mencapai batas.")
		return
	}

	fmt.Print("Masukkan ID karyawan: ")
	fmt.Scan(&KaryawanID)
	fmt.Print("Masukkan tipe pekerjaan: ")
	fmt.Scan(&Tipe)
	fmt.Print("Masukkan durasi pekerjaan (dalam jam): ")
	fmt.Scan(&Durasi)

	for i := 0; i < *nlog; i++ {
		if T[i].KaryawanID == KaryawanID && T[i].Tipe == Tipe {
			fmt.Println("Log pekerjaan untuk karyawan ini dan tipe ini sudah ada.")
			return
		}
	}

	T[*nlog] = Pekerjaan{KaryawanID: KaryawanID, Tipe: Tipe, Durasi: Durasi}
	*nlog++
	fmt.Println("Log pekerjaan berhasil ditambah.")

}

func showLogPekerjaan(T arrPekerjaan, nlog int) {
	fmt.Println("======================================")
	fmt.Println("List Log Pekerjaan:")
	fmt.Println("======================================")
	if nlog == 0 {
		fmt.Println("Tidak ada log pekerjaan.")
	} else {
		for i := 0; i < nlog; i++ {
			fmt.Printf("Karyawan ID: %d, Tipe: %d, Durasi: %d jam\n", T[i].KaryawanID, T[i].Tipe, T[i].Durasi)
		}
	}
	fmt.Println("======================================")
}
