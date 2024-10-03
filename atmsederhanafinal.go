package main

import (
 "fmt"
 "os"
)

// Data Mahasiswa
type Mahasiswa struct {
 Username string
 NPM      string
 Saldo    int
}

// Global variables
var saldoAwal = 3500000
var transaksiHistory []string

// Database Mahasiswa (username dan npm diperbarui)
var mahasiswaData = []Mahasiswa{
 {Username: "denendro", NPM: "2406357942", Saldo: saldoAwal},
}

// Fungsi login untuk verifikasi username dan password
func login(username, npm string) (*Mahasiswa, bool) {
 for i, m := range mahasiswaData {
  if m.Username == username && m.NPM == npm {
   return &mahasiswaData[i], true
  }
 }
 return nil, false
}

// Function untuk melihat informasi akun
func LihatAkun(mahasiswa *Mahasiswa) {
 fmt.Println("=== Informasi Akun ===")
 fmt.Printf("Username: %s\n", mahasiswa.Username)
 fmt.Printf("NPM: %s\n", mahasiswa.NPM)
 fmt.Println("=====================")
}

// Function untuk melihat saldo
func LihatSaldo(mahasiswa *Mahasiswa) {
 fmt.Println("=== Informasi Saldo ===")
 fmt.Printf("Saldo Anda: Rp %d\n", mahasiswa.Saldo)
 fmt.Println("=====================")
}

// Function untuk menambah saldo
func TambahSaldo(mahasiswa *Mahasiswa, jumlah int) {
 if jumlah > 0 {
  mahasiswa.Saldo += jumlah
  transaksiHistory = append(transaksiHistory, fmt.Sprintf("Tambah Saldo: Rp %d", jumlah))
  fmt.Println("Saldo berhasil ditambahkan!")
 } else {
  fmt.Println("Jumlah saldo yang dimasukkan tidak valid.")
 }
 fmt.Println("=====================")
}

// Function untuk menarik saldo
func TarikSaldo(mahasiswa *Mahasiswa, jumlah int) {
 if jumlah > 0 && jumlah <= mahasiswa.Saldo {
  mahasiswa.Saldo -= jumlah
  transaksiHistory = append(transaksiHistory, fmt.Sprintf("Tarik Saldo: Rp %d", jumlah))
  fmt.Println("Saldo berhasil ditarik!")
 } else if jumlah <= 0 {
  fmt.Println("Jumlah saldo yang dimasukkan tidak valid.")
 } else {
  fmt.Println("Saldo tidak mencukupi.")
 }
 fmt.Println("=====================")
}

// Function untuk melihat histori transaksi
func HistoryTransaksi() {
 fmt.Println("=== Histori Transaksi ===")
 if len(transaksiHistory) == 0 {
  fmt.Println("Belum ada transaksi.")
 } else {
  for _, transaksi := range transaksiHistory {
   fmt.Println(transaksi)
  }
 }
 fmt.Println("=====================")
}

// Main function untuk menjalankan program
func main() {
 var username, npm string

 // Pesan pembuka
 fmt.Println("=== Selamat Datang di ATM ===")
 fmt.Print("Masukkan Username: ")
 fmt.Scanln(&username)
 fmt.Print("Masukkan Password (NPM): ")
 fmt.Scanln(&npm)

 // Verifikasi login
 mahasiswa, valid := login(username, npm)

 if valid {
  fmt.Println("Login berhasil!")
  fmt.Println("=====================")

  for {
   // Menu pilihan
   fmt.Println("=== Menu ATM ===")
   fmt.Println("1. Lihat Informasi Akun")
   fmt.Println("2. Lihat Saldo")
   fmt.Println("3. Tambahkan Saldo")
   fmt.Println("4. Tarik Saldo")
   fmt.Println("5. Histori Transaksi")
   fmt.Println("6. Keluar")

   var pilihan int
   fmt.Print("Pilih menu: ")
   fmt.Scanln(&pilihan)
   fmt.Println("=====================")

   switch pilihan {
   case 1:
    fmt.Println("Anda memilih: Lihat Informasi Akun")
    LihatAkun(mahasiswa)
   case 2:
    fmt.Println("Anda memilih: Lihat Saldo")
    LihatSaldo(mahasiswa)
   case 3:
    fmt.Println("Anda memilih: Tambahkan Saldo")
    var tambah int
    fmt.Print("Masukkan jumlah yang ingin ditambahkan: ")
    fmt.Scanln(&tambah)
    TambahSaldo(mahasiswa, tambah)
   case 4:
    fmt.Println("Anda memilih: Tarik Saldo")
    var tarik int
    fmt.Print("Masukkan jumlah yang ingin ditarik: ")
    fmt.Scanln(&tarik)
    TarikSaldo(mahasiswa, tarik)
   case 5:
    fmt.Println("Anda memilih: Histori Transaksi")
    HistoryTransaksi()
   case 6:
    fmt.Println("Terima kasih telah menggunakan ATM. Sampai jumpa!")
    os.Exit(0)
   default:
    fmt.Println("Pilihan tidak valid, silakan coba lagi.")
    fmt.Println("=====================")
   }
  }
 } else {
  // Pesan jika login gagal
  fmt.Println("Username atau password salah. Program berhenti.")
 }
}