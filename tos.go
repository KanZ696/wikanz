package main

import "fmt"

func main() {
    var pilihan int
    var jumlah int
    var isMember string

    fmt.Println("======= KAFE KOPI KITA =======")
    fmt.Println("1. Kopi Espresso   (Rp 18.000)")
    fmt.Println("2. Cafe Latte      (Rp 24.000)")
    fmt.Println("3. Matcha Latte    (Rp 26.000)")
    fmt.Print("Pilih menu (1-3): ")
    fmt.Scanln(&pilihan)

    fmt.Print("Jumlah cup: ")
    fmt.Scanln(&jumlah)

    harga := 0
    if pilihan == 1 {
        harga = 18000
    } else if pilihan == 2 {
        harga = 24000
    } else if pilihan == 3 {
        harga = 26000
    }

    subtotal := harga * jumlah

    fmt.Print("Apakah punya kartu member? (y/n): ")
    fmt.Scanln(&isMember)

    diskon := 0
    if isMember == "y" || isMember == "Y" {
        diskon = subtotal * 10 / 100 
    }

    totalBayar := subtotal - diskon

    fmt.Println("\n--- Rincian Tagihan ---")
    fmt.Printf("Subtotal    : Rp %d\n", subtotal)
    fmt.Printf("Diskon (10%): Rp %d\n", diskon)
    fmt.Printf("Total Bayar : Rp %d\n", totalBayar)
}