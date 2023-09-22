package main

import (
	"fmt"
	"os"
)

type peserta struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var dataPeserta = []peserta{
	{
		nama:      "Budi",
		alamat:    "Jakarta",
		pekerjaan: "software developer",
		alasan:    "Golang memiliki performa yang sangat baik",
	},
	{
		nama:      "Andry",
		alamat:    "Depok",
		pekerjaan: "backend developer",
		alasan:    "Golang tergolong static typed",
	},
	{
		nama:      "Indra",
		alamat:    "Jakarta",
		pekerjaan: "software developer",
		alasan:    "Ingin mempelajari stack teknologi terbaru",
	},
	{
		nama:      "Anto",
		alamat:    "Bekasi",
		pekerjaan: "Akuntan",
		alasan:    "Ingin terjun ke dunia pemrogragman",
	},
	{
		nama:      "Heri",
		alamat:    "Tangerang",
		pekerjaan: "Wiraswasta",
		alasan:    "Memperdalam ilmu di bidang pemrograman",
	},
}

func main() {

	argsRaw := os.Args
	if len(argsRaw) < 2 {
		fmt.Println("Tambahkan Argumen Pada Terminal, Contoh : go run main.go Budi")
		return
	} else if len(argsRaw) > 2 {
		fmt.Println("Maksimal Argumen Adalah Satu")
		return
	}

	inputSearch := argsRaw[1]

	for _, v := range dataPeserta {
		if v.nama == inputSearch {
			fmt.Println("Nama : ", v.nama)
			fmt.Println("Alamat : ", v.alamat)
			fmt.Println("Pekerjaan : ", v.pekerjaan)
			fmt.Println("alasan : ", v.alasan)
			return
		}
	}

	fmt.Println("Data Tidak Ditemukan")

}
