package main

import "fmt"

/*
	BEFORE
*/

// type kendaraan struct {
// 	totalroda       int
// 	kecepatanperjam int
// }

// type mobil struct {
// 	kendaraan
// }

// func (m *mobil) berjalan() {
// 	m.tambahkecepatan(10)
// }

// func (m *mobil) tambahkecepatan(kecepatanbaru int) {
// 	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
// }

// func main() {
// 	mobilcepat := mobil{}
// 	mobilcepat.berjalan()
// 	mobilcepat.berjalan()
// 	mobilcepat.berjalan()
// 	fmt.Println(mobilcepat)

// 	mobillamban := mobil{}
// 	mobillamban.berjalan()
// 	fmt.Println(mobillamban)
// }

/*
	AFTER
*/

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatan int) {
	m.KecepatanPerJam = m.KecepatanPerJam + kecepatan
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	fmt.Println(mobilCepat)

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
	fmt.Println(mobilLamban)
}
