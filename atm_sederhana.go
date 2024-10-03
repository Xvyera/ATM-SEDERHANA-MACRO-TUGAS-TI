package main

import (
	"fmt"
	"os"
	"time"
)

type Account struct {
	Username string
	Owner    string
	Password string
	Saldo    []float64
}

var account = Account{
	Username: "Qheyla",
	Owner:    "Qheyla Zavyera Valendro",
	Password: "2406356725",
	Saldo:    []float64{3500000},
}

var HistoryTransaksi []string

func getTimestamp() string {
	return time.Now().Format("09-Nov-2006 07.11.09")
}

func login() {
	var username, password string
	fmt.Print("Masukkan Username Anda: ")
	fmt.Scanln(&username)

	fmt.Print("Masukkan Password Anda: ")
	fmt.Scanln(&password)

	if username == account.Username && password == account.Password {
		fmt.Println("<================================================>")
		fmt.Println("|     <3 SELAMAT DATANG DI BANK SPESIAL <3       |")
		fmt.Println("<================================================>")
		return
	} else {
		fmt.Println("Password atau Username Anda Salah, Mohon untuk Coba Lagi!")
		os.Exit(1)
	}
}

func menu() {
	for {
		fmt.Println("<================================================>")
		fmt.Println("|       <3 MENU TRANSAKSI BANK SPESIAL <3        |")
		fmt.Println("<================================================>")
		fmt.Println("|                1. Lihat Saldo                  |")
		fmt.Println("|                2. Tambah Saldo                 |")
		fmt.Println("|                3. Tarik Saldo                  |")
		fmt.Println("|                4. Lihat Akun                   |")
		fmt.Println("|                5. History Transaksi            |")
		fmt.Println("|                6. Keluar                       |")
		fmt.Println("<================================================>")
		fmt.Print("Ayooo mari pilih satu opsi: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Lihatsaldo()
		case 2:
			Tambahsaldo()
		case 3:
			Tariksaldo()
		case 4:
			Lihatakun()
		case 5:
			Historytransaksi()
		case 6:
			fmt.Println("Yahh Anda Keluar dari Program!")
			os.Exit(0)
		default:
			fmt.Println("Yahh Sangat Disayangkan, opsi anda tidak valid! AYOO COBA LAGIII ^3^")
		}
	}
}

func Lihatsaldo() {
	fmt.Println("<================================================>")
	fmt.Println("|                 Jumlah Saldo Anda              |")
	fmt.Println("<================================================>")
	fmt.Printf("\nSaldo Anda Saat ini Berjumlah Rp.%.2f\n\n\n", account.Saldo)

	lanjut()
}

func Tambahsaldo() {
	var tambah float64
	fmt.Println("<===================================================>")
	fmt.Println("|                 Tambah Saldo Anda                 |")
	fmt.Println("<===================================================>")
	fmt.Print("Saldo yang Ingin di Tambahkan Sebesar: Rp. ")
	fmt.Scanln(&tambah)

	if tambah <= 0 {
		fmt.Println("Yahh Jumlahnya Tidak Valid! Tolong Masukkan Jumlah yang Valid Yaaaa!!!")
	} else {
		account.Saldo = append(account.Saldo, account.Saldo[len(account.Saldo)-1]+tambah)
		HistoryTransaksi = append(HistoryTransaksi, fmt.Sprintf("Penambahan saldo : Rp. %.2f", tambah))
		fmt.Printf("Saldo berhasil ditambahkan, Saldo anda sekarang adalah Rp. %.2f.\n", account.Saldo)
	}

	lanjut()
}

func Tariksaldo() {
	var tarik float64
	fmt.Println("<===================================================>")
	fmt.Println("|                 Tarik Saldo Anda                  |")
	fmt.Println("<===================================================>")
	fmt.Print("Saldo yang Ingin di Tarik Sebesar: Rp. ")
	fmt.Scanln(&tarik)

	if tarik <= 0 {
		fmt.Println("Yahh Jumlahnya Tidak Valid! Tolong Masukkan Jumlah Angka yang Positif Yaaaa!!!")
	} else if int(tarik)%50000 != 0 {
		fmt.Println("Wahh Maaf!!! Jumlah yang ingin ditarik harus berkelipatan Rp. 50.000 atau kelipatan Rp. 100.000 nihh kawannn!")
	} else if tarik > account.Saldo[len(account.Saldo)-1] {
		fmt.Println("WADUH GAWAT! Saldonya Tidak Cukup untuk Menarik Uang Sebanyak Ituuu!")
	} else {
		account.Saldo = append(account.Saldo, account.Saldo[len(account.Saldo)-1]-tarik)
		HistoryTransaksi = append(HistoryTransaksi, fmt.Sprintf("Penarikan uang: Rp. %.2f", tarik))
		fmt.Printf("Rp. %.2f Berhasil Ditarik YIPPIEEEE!!! Saldo Baru Anda Adalah Rp. %.2f.\n", tarik, account.Saldo)
	}

	lanjut()
}

func lanjut() {
	fmt.Println("\n KAWAN?!?! Apakah Anda Mau Melanjutkan Transaksi???")
	fmt.Println("\nAYOOO Tekan Apapun untuk Kembali ke Menuuu!\nTekan angka '1' untuk Keluar YAAAA ^^")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println("Yahh Anda Keluar dari program! TRIMAKASI BANYAK TELAH BERTRANSAKSI DI BANK SPESIAL <3333")
		os.Exit(0)
	default:
		menu()
	}
}

func Lihatakun() {
	fmt.Println("<===================================================>")
	fmt.Println("|                 Informasi Akun Anda               |")
	fmt.Println("<===================================================>")
	fmt.Printf("Username: %s\n", account.Username)
	fmt.Printf("Owner: %s\n", account.Owner)
	fmt.Printf("Saldo: Rp. %.2f\n", account.Saldo)

	lanjut()
}

func Historytransaksi() {
	fmt.Println("<===================================================>")
	fmt.Println("|               Informasi Akun Anda                 |")
	fmt.Println("<===================================================>")
	if len(HistoryTransaksi) == 0 {
		fmt.Println("Yahh Tidak Ada Riwayat Transaksi Nihhh!")
	} else {
		for i, transaksi := range HistoryTransaksi {
			fmt.Printf("%d. %s\n", i+1, transaksi)
		}
	}
	lanjut()
}

func main() {

	login()

	menu()

}
