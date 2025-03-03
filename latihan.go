package main

import (
    "fmt"
    "math"
)

func kuadrat(awal int, akhir int) { 
    for i := awal; i <= akhir; i++ {
        akar := math.Sqrt(float64(i)) // Hitung akar kuadrat
        if akar == float64(int(akar)) { // Cek apakah akar kuadratnya bilangan bulat
            fmt.Println(i, "adalah kuadrat sempurna")
        }
    }
}

func main() {
    var d1, d2 int
    fmt.Print("Masukkan angka awal dan akhir: ")
    fmt.Scan(&d1, &d2)
    
    kuadrat(d1, d2)
}

