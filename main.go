package main

import (
	"fmt"
)

type Ban string

// class mobil
type Mobil struct {
	roda [4]Ban
}

func NewMobil() *Mobil {
	return &Mobil{
		roda: [4]Ban{"karet", "karet", "karet", "karet"},
	}
}

// ganti roda
func (m *Mobil) GantiRoda(index int, jenis Ban) {
	m.roda[index] = jenis
}

// ketuk dan bika pintu kanan
func (m *Mobil) KetukPintuKanan() {
	fmt.Println("Knock")
}
func (m *Mobil) BukaPintuKanan() {
	fmt.Println("Beep")
}

// ketu dan buka pintu kiri
func (m *Mobil) KetukPintuKiri() {
	fmt.Println("Beep")
}
func (m *Mobil) BukaPintuKiri() {
	fmt.Println("Knock")
}

func main() {
	mobil := NewMobil()

	fmt.Println("===================================================")
	fmt.Println("Mobil saat ini memiliki jenis roda sebagai berikut:")
	for i, roda := range mobil.roda {
		fmt.Printf("Roda %d: %s\n", i+1, roda)
	}

	menuExercise := []string{"Ganti Roda 1", "Ganti Roda 2", "Ganti Roda 3", "Ganti Roda 4", "Ketuk Pintu Kanan", "Buka Pintu Kanan", "Ketuk Pintu Kiri", "Buka Pintu Kiri"}

	for {
		fmt.Println("\nMenu:")
		for i, opsi := range menuExercise {
			fmt.Printf("%d. %s\n", i+1, opsi)
		}

		var pilihan int
		fmt.Print("Pilih menu antara 1 sampai 8, atau 0 untuk keluar): ")
		fmt.Scanln(&pilihan)

		if pilihan == 0 {
			break
		}

		switch pilihan {
		case 1, 2, 3, 4:
			var roda int = pilihan - 1
			fmt.Println("------------------------------")
			fmt.Println("Pilihan jenis ban untuk roda:")
			fmt.Println("1. Karet")
			fmt.Println("2. Kayu")
			fmt.Println("3. Besi")
			var jenisBan int
			fmt.Print("Masukkan pilihan jenis ban (1-3): ")
			fmt.Scanln(&jenisBan)
			var ban Ban
			switch jenisBan {
			case 1:
				ban = "karet"
			case 2:
				ban = "kayu"
			case 3:
				ban = "besi"
			default:
				fmt.Println("Pilihan tidak valid.")
				continue
			}
			mobil.GantiRoda(roda, ban)
			fmt.Printf("Roda ke-%d telah diganti dengan ban jenis %s.\n", roda+1, ban)
			fmt.Println("===================================================")
			fmt.Println("Mobil saat ini memiliki jenis roda sebagai berikut:")
			for i, roda := range mobil.roda {
				fmt.Printf("Roda %d: %s\n", i+1, roda)
			}
		case 5:
			mobil.KetukPintuKanan()
		case 6:
			mobil.BukaPintuKanan()
		case 7:
			mobil.KetukPintuKiri()
		case 8:
			mobil.BukaPintuKiri()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
	fmt.Println("Terima kasih!")
}
