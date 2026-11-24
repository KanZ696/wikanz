package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Struct untuk menyimpan rincian transaksi
type Transaksi struct {
    NamaMenu   string
    HargaSatuan int
    Jumlah     int
    Subtotal   int
    Diskon     int
    TotalAkhir int
    UangBayar  int
    Kembalian  int
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("==============================================")
    fmt.Println("        SISTEM KASIR KAFE KOPI SENJA          ")
    fmt.Println("==============================================")
    fmt.Println("DAFTAR MENU:")
    fmt.Println("  1. Americano Coffee     : Rp 18.000")
    fmt.Println("  2. Caramel Macchiato    : Rp 25.000")
    fmt.Println("  3. Matcha Green Tea     : Rp 23.000")
    fmt.Println("  4. Butter Croissant     : Rp 15.000")
    fmt.Println("----------------------------------------------")

    var menuPilihan int
    var namaMenu string
    var hargaSatuan int

    // 1. Validasi input menu
    for {
        fmt.Print("Pilih nomor menu (1-4): ")
        input, _ := reader.ReadString('\n')
        val, err := strconv.Atoi(strings.TrimSpace(input))

        if err != nil || val < 1 || val > 4 {
            fmt.Println("[Error] Pilihan tidak valid! Masukkan angka antara 1 sampai 4.")
            continue
        }
        menuPilihan = val
        break
    }

    switch menuPilihan {
    case 1:
        namaMenu = "Americano Coffee"
        hargaSatuan = 18000
    case 2:
        namaMenu = "Caramel Macchiato"
        hargaSatuan = 25000
    case 3:
        namaMenu = "Matcha Green Tea"
        hargaSatuan = 23000
    case 4:
        namaMenu = "Butter Croissant"
        hargaSatuan = 15000
    }

    // 2. Validasi input jumlah pesanan
    var jumlahPesanan int
    for {
        fmt.Print("Masukkan jumlah pesanan (1-50): ")
        input, _ := reader.ReadString('\n')
        val, err := strconv.Atoi(strings.TrimSpace(input))

        if err != nil || val < 1 || val > 50 {
            fmt.Println("[Error] Jumlah harus berupa angka 1 - 50!")
            continue
        }
        jumlahPesanan = val
        break
    }

    subtotal := hargaSatuan * jumlahPesanan

    // 3. Pengecekan Diskon Member
    fmt.Print("Apakah pelanggan memiliki kartu Member? (y/t): ")
    inputMember, _ := reader.ReadString('\n')
    inputMember = strings.ToLower(strings.TrimSpace(inputMember))

    diskon := 0
    if inputMember == "y" {
        diskon = subtotal * 10 / 100 // Diskon 10%
    }
    totalAkhir := subtotal - diskon

    // 4. Validasi pembayaran uang tunai
    var uangBayar int
    for {
        fmt.Printf("\nTotal yang harus dibayar: Rp %s\n", formatRibuan(totalAkhir))
        fmt.Print("Masukkan nominal uang pembayaran: Rp ")
        input, _ := reader.ReadString('\n')
        val, err := strconv.Atoi(strings.TrimSpace(input))

        if err != nil || val < totalAkhir {
            fmt.Println("[Error] Uang pembayaran kurang atau input tidak valid! Silakan bayar pas atau lebih.")
            continue
        }
        uangBayar = val
        break
    }

    kembalian := uangBayar - totalAkhir

    transaksi := Transaksi{
        NamaMenu:   namaMenu,
        HargaSatuan: hargaSatuan,
        Jumlah:     jumlahPesanan,
        Subtotal:   subtotal,
        Diskon:     diskon,
        TotalAkhir: totalAkhir,
        UangBayar:  uangBayar,
        Kembalian:  kembalian,
    }

    // 5. Cetak Struk Pembayaran
    fmt.Println("\n==============================================")
    fmt.Println("             STRUK TRANSAKSI KAFE             ")
    fmt.Println("==============================================")
    fmt.Printf(" Menu Pesanan : %s\n", transaksi.NamaMenu)
    fmt.Printf(" Harga Satuan : Rp %s\n", formatRibuan(transaksi.HargaSatuan))
    fmt.Printf(" Jumlah Item  : %d porsi\n", transaksi.Jumlah)
    fmt.Printf(" Subtotal     : Rp %s\n", formatRibuan(transaksi.Subtotal))
    fmt.Printf(" Diskon       : Rp %s\n", formatRibuan(transaksi.Diskon))
    fmt.Println("----------------------------------------------")
    fmt.Printf(" TOTAL BAYAR  : Rp %s\n", formatRibuan(transaksi.TotalAkhir))
    fmt.Printf(" TUNAI        : Rp %s\n", formatRibuan(transaksi.UangBayar))
    fmt.Printf(" KEMBALIAN    : Rp %s\n", formatRibuan(transaksi.Kembalian))
    fmt.Println("==============================================")
    fmt.Println("      Terima Kasih & Selamat Menikmati!       ")
}

// Fungsi format angka ribuan
func formatRibuan(n int) string {
    s := strconv.Itoa(n)
    if len(s) <= 3 {
        return s
    }
    var parts []string
    for len(s) > 3 {
        parts = append([]string{s[len(s)-3:]}, parts...)
        s = s[:len(s)-3]
    }
    if len(s) > 0 {
        parts = append([]string{s}, parts...)
    }
    return strings.Join(parts, ".")
}