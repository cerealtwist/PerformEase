package main

import (
	"fmt"
)

const NMAX int = 3
const NMAXJob int = 1000 // blm pasti pake atau enggak, just keep it this way dulu dah
const TotalJAM int = 160 // total jam kerja 160

type Karyawan struct {
	ID   int
	Nama string
}

type Pekerjaan struct {
	KaryawanID, Tipe, Durasi, Bulan, Tahun int
}

type arrKaryawan [NMAX]Karyawan
type arrPekerjaan [NMAXJob]Pekerjaan

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
			fmt.Println("Exiting...")
			return
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
			mainMenu()
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
		} else if choice == 0 {
			mainMenu()
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
		fmt.Println("1. Show Log Pekerjaan Karyawan (Status)")
		fmt.Println("2. Show Rekap Pekerjaan per Tipe")
		fmt.Println("3. Show Rekap Aktivitas by Durasi (Sorted)")
		fmt.Println("0. Main Menu")
		fmt.Println("======================================")

		fmt.Scan(&choice)
		if choice == 1 {
			showLogKaryawan(T, nlog)
		} else if choice == 2 {
			showRekapTipe(T, nlog)
		} else if choice == 3 {
			showRekapDurasi(T, nlog)
		} else if choice == 0 {
			mainMenu()
		} else {
			fmt.Println("Pilihan Invalid.")
		}
	}
}

// Helper function untuk mencari indeks karyawan by ID (updated to Binary Search!)
func findKaryawan(A arrKaryawan, n int, KaryawanID int) bool {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if A[mid].ID == KaryawanID {
			return true
		} else if A[mid].ID < KaryawanID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// Helper function untuk sort array karyawan by ID
func sortKaryawan(A *arrKaryawan, n int) {
	for pass := 0; pass < n-1; pass++ {
		idx := pass
		for i := pass + 1; i < n; i++ {
			if A[i].ID < A[idx].ID {
				idx = i
			}
		}
		if idx != pass {
			temp := A[pass]
			A[pass] = A[idx]
			A[idx] = temp
		}
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
	sortKaryawan(A, *n) // Memastikan array terurut setelah penambahan
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
			sortKaryawan(A, *n)
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

// BAGIAN CRUD KARYAWAN

func addPekerjaan(T *arrPekerjaan, nlog *int) {
	var KaryawanID, Tipe, Durasi, Bulan, Tahun int

	if *nlog >= NMAXJob {
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

// FOR REKAP DATA
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
	if t1 >= totalJam*25/100 && t1 <= totalJam*50/100 && t2 <= totalJam*50/100 &&
		t2 >= totalJam*10/100 && t3 >= totalJam*10/100 {
		fmt.Println("Karyawan memenuhi ketentuan proporsi pekerjaan.")
	} else {
		fmt.Println("Karyawan tidak memenuhi ketentuan proporsi pekerjaan.")
	}
}

func showRekapTipe(T arrPekerjaan, nlog int) {
	var Tipe, Bulan, Tahun, totalJam int
	var found bool

	fmt.Print("Masukkan Tipe Pekerjaan: ")
	fmt.Scan(&Tipe)
	fmt.Print("Masukkan Bulan (1-12): ")
	fmt.Scan(&Bulan)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&Tahun)

	fmt.Printf("Rekap Pekerjaan Tipe %d pada Bulan: %d Tahun: %d\n", Tipe, Bulan, Tahun)
	fmt.Println("======================================")

	for i := 0; i < nlog; i++ {
		if T[i].Tipe == Tipe && T[i].Bulan == Bulan && T[i].Tahun == Tahun {
			found = true
			totalJam += T[i].Durasi
			fmt.Printf("Karyawan ID: %d, Durasi: %d jam\n", T[i].KaryawanID, T[i].Durasi)

		}
	}

	if !found {
		fmt.Println("Tidak ada pekerjaan untuk tipe ini pada bulan dan tahun tersebut.")
	} else {
		fmt.Printf("Total Durasi: %d jam\n", totalJam)
	}
	fmt.Println("======================================")

}

func showRekapDurasi(T arrPekerjaan, nlog int) {
	var Bulan, Tahun, sortOrder int

	fmt.Print("Masukkan Bulan (1-12): ")
	fmt.Scan(&Bulan)
	fmt.Print("Masukkan Tahun: ")
	fmt.Scan(&Tahun)
	fmt.Print("Pilih urutan pengurutan: 1 untuk Ascending, 2 untuk Descending: ")
	fmt.Scan(&sortOrder)

	var pekerjaanCount int = 0

	// Mengumpulkan data yang sesuai dengan bulan dan tahun yang diminta
	for i := 0; i < nlog; i++ {
		if T[i].Bulan == Bulan && T[i].Tahun == Tahun {
			// Salin elemen yang sesuai ke bagian awal array
			T[pekerjaanCount] = T[i]
			pekerjaanCount++
		}
	}

	// sort pekerjaan by durasi menggunakan selection sort
	for pass := 0; pass < pekerjaanCount-1; pass++ {
		idx := pass
		for i := pass + 1; i < pekerjaanCount; i++ {
			if (sortOrder == 1 && T[i].Durasi < T[idx].Durasi) || (sortOrder == 2 && T[i].Durasi > T[idx].Durasi) {
				idx = i
			}
		}
		// Pertukaran elemen
		if idx != pass {
			temp := T[pass]
			T[pass] = T[idx]
			T[idx] = temp
		}
	}

	order := "Ascending"
	if sortOrder == 2 {
		order = "Descending"
	}

	fmt.Printf("Rekap Aktivitas Pekerjaan pada Bulan: %d Tahun: %d (Sorted by Durasi - %s)\n", Bulan, Tahun, order)
	fmt.Println("======================================")

	if pekerjaanCount == 0 {
		fmt.Println("Tidak ada aktivitas pekerjaan pada bulan dan tahun tersebut.")
	} else {
		for i := 0; i < pekerjaanCount; i++ {
			fmt.Printf("Karyawan ID: %d, Tipe: %d, Durasi: %d jam\n", T[i].KaryawanID, T[i].Tipe, T[i].Durasi)
		}
	}
	fmt.Println("======================================")
}
