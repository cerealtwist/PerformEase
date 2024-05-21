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
	KaryawanID, Tipe, Durasi, Bulan, Tahun int
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
		fmt.Println("1. Add Log Pekerjaan Karyawan")
		fmt.Println("2. Update Log Pekerjaan Karyawan")
		fmt.Println("3. Delete Log Pekerjaan Karyawan")
		fmt.Println("4. Show All Log Pekerjaan")
		fmt.Println("5. Show Log Pekerjaan Karyawan (Status)")
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)

		if choice == 1 {
			addPekerjaan(&T, &nlog)
		} else if choice == 2 {
			updatePekerjaan(&T, &nlog)
		} else if choice == 3 {
			deletePekerjaan(&T, &nlog)
		} else if choice == 4 {
			showAllPekerjaan(T, nlog)
		} else if choice == 5 {
			showLogKaryawan(T, nlog)
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
			// geser array ke kiri utk overwrite array terhapus.
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

// Helper function untuk cek apakah terdapat karyawan dengan id tertentu
func findKaryawan(A arrKaryawan, n int, KaryawanID int) bool {
	for i := 0; i < n; i++ {
		if A[i].ID == KaryawanID {
			return true
		}
	}
	return false
}

// BAGIAN CRUD KARYAWAN

func addPekerjaan(T *arrPekerjaan, nlog *int) {
	var KaryawanID, Tipe, Durasi, Bulan, Tahun int

	if *nlog >= NMAX {
		fmt.Println("Jumlah log pekerjaan mencapai batas.")
		return
	}

	fmt.Print("Masukkan ID karyawan: ")
	fmt.Scan(&KaryawanID)

	// Panggil helper function untuk validasi Id karyawan
	if !findKaryawan(A, n, KaryawanID) {
		fmt.Println("ID Karyawan tidak ditemukan.")
		return
	}

	fmt.Print("Masukkan tipe pekerjaan: ")
	fmt.Scan(&Tipe)
	fmt.Print("Masukkan durasi pekerjaan (dalam jam): ")
	fmt.Scan(&Durasi)
	fmt.Print("Masukkan bulan pekerjaan (1-12): ")
	fmt.Scan(&Bulan)
	fmt.Print("Masukkan tahun pekerjaan: ")
	fmt.Scan(&Tahun)

	for i := 0; i < *nlog; i++ {
		if T[i].KaryawanID == KaryawanID && T[i].Tipe == Tipe && T[i].Bulan == Bulan && T[i].Tahun == Tahun {
			fmt.Println("Log pekerjaan sudah ada.")
			return
		}
	}

	T[*nlog] = Pekerjaan{KaryawanID: KaryawanID, Tipe: Tipe, Durasi: Durasi, Bulan: Bulan, Tahun: Tahun}
	*nlog++
	fmt.Println("Log pekerjaan berhasil ditambah.")

}

func updatePekerjaan(T *arrPekerjaan, nlog *int) {
	var KaryawanID, Tipe, newDurasi, Bulan, Tahun int

	fmt.Print("Masukkan ID Karyawan: ")
	fmt.Scan(&KaryawanID)
	fmt.Print("Masukkan tipe pekerjaan yang ingin diupdate: ")
	fmt.Scan(&Tipe)
	fmt.Print("Masukkan bulan pekerjaan (1-12): ")
	fmt.Scan(&Bulan)
	fmt.Print("Masukkan tahun pekerjaan: ")
	fmt.Scan(&Tahun)

	for i := 0; i < *nlog; i++ {
		if T[i].KaryawanID == KaryawanID && T[i].Tipe == Tipe && T[i].Bulan == Bulan && T[i].Tahun == Tahun {
			fmt.Print("Masukkan durasi baru (dalam jam): ")
			fmt.Scan(&newDurasi)
			T[i].Durasi = newDurasi
			fmt.Println("Log berhasil terupdate.")
			return
		}
	}
	fmt.Println("Log pekerjaan tidak ditemukan.")
}

func deletePekerjaan(T *arrPekerjaan, nlog *int) {
	var KaryawanID, Tipe, Bulan, Tahun int
	var idx int = -1 // variabel utk simpan indeks log entry yang akan dihapus jika ditemukan

	// ARRAY -1 BERARTI INVALID INDEX (sifat array non-negative), BUAT MENGINDIKASIKAN NOT FOUND

	fmt.Print("Enter employee ID: ")
	fmt.Scan(&KaryawanID)
	fmt.Print("Enter work type to delete: ")
	fmt.Scan(&Tipe)
	fmt.Print("Masukkan bulan pekerjaan (1-12): ")
	fmt.Scan(&Bulan)
	fmt.Print("Masukkan tahun pekerjaan: ")
	fmt.Scan(&Tahun)

	for i := 0; i < *nlog && idx == -1; i++ {
		if T[i].KaryawanID == KaryawanID && T[i].Tipe == Tipe && T[i].Bulan == Bulan && T[i].Tahun == Tahun {
			idx = i // jika ditemukan, ubah index sesuai i array
		}
	}

	if idx != -1 { // jika idx bukan -1 (berarti ditemukan)
		// geser array ke kiri utk overwrite array terhapus.
		for j := idx; j < *nlog-1; j++ {
			T[j] = T[j+1]
		}
		*nlog--
		fmt.Println("Log pekerjaan berhasil dihapus..")
	} else {
		fmt.Println("Log pekerjaan tidak ditemukan.")
	}
}

// UNTUK MELIHAT SEMUA LOG PEKERJAAN (FOR TESTING)
func showAllPekerjaan(T arrPekerjaan, nlog int) {
	fmt.Println("\n======================================")
	fmt.Println("List Log Pekerjaan:")
	fmt.Println("======================================")
	if nlog == 0 {
		fmt.Println("Tidak ada log pekerjaan.")
	} else {
		for i := 0; i < nlog; i++ {
			fmt.Printf("Karyawan ID: %d, Tipe: %d, Durasi: %d jam, Bulan: %d, Tahun: %d\n", T[i].KaryawanID, T[i].Tipe, T[i].Durasi, T[i].Bulan, T[i].Tahun)
		}
	}
	fmt.Println("======================================")
}

// MELIHAT KETUNTASAN LOG PEKERJAAN KARYAWAN BERDASARKAN TANGGAL (BULAN & TAHUN)
func showLogKaryawan(T arrPekerjaan, nlog int) {
	var totalJam int = 160
	var KaryawanID, Bulan, Tahun int
	var t1, t2, t3 int

	fmt.Print("Masukkan ID Karyawan: ")
	fmt.Scan(&KaryawanID)
	fmt.Print("Masukkan Bulan (1-12): ")
	fmt.Scan(&Bulan)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&Tahun)

	for i := 0; i < nlog; i++ {
		if T[i].KaryawanID == KaryawanID && T[i].Bulan == Bulan && T[i].Tahun == Tahun {
			if T[i].Tipe == 1 {
				t1 += T[i].Durasi
			} else if T[i].Tipe == 2 {
				t2 += T[i].Durasi
			} else if T[i].Tipe == 3 {
				t3 += T[i].Durasi
			}
		}
	}

	fmt.Printf("Rekap Log Pekerjaan untuk Karyawan ID: %d pada Bulan: %d Tahun: %d\n", KaryawanID, Bulan, Tahun)
	fmt.Printf("Tipe 1: %d jam\n", t1)
	fmt.Printf("Tipe 2: %d jam\n", t2)
	fmt.Printf("Tipe 3: %d jam\n", t3)

	// Cek apakah ketentuan proporsi pekerjaan terpenuhi
	if t1 >= totalJam*25/100 && t2 <= totalJam*50/100 &&
		t2 >= totalJam*10/100 && t3 >= totalJam*10/100 {
		fmt.Println("Karyawan memenuhi ketentuan proporsi pekerjaan.")
	} else {
		fmt.Println("Karyawan tidak memenuhi ketentuan proporsi pekerjaan.")
	}
}
