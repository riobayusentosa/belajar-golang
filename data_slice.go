package main

import "fmt"

func main() {
	names := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	names[5] = "diubah"
	fmt.Println(slice)

	slice[0] = "Mei lagi"
	fmt.Println(names)

	///////////////////////////////////
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysslice := days[5:]
	daysslice[0] = "Sabtu Baru"
	daysslice[1] = "Minggu Baru"

	fmt.Println(days)

	daysslice2 := append(daysslice, "Libur Baru")
	daysslice2[0] = "Ups"
	fmt.Println(daysslice2)
	fmt.Println(days)

	///make slice
	newslice := make([]string, 2, 5)
	newslice[0] = "rio"
	newslice[1] = "bayu"
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	//copy slice
	fromslice := days[:]
	toslice := make([]string, len(fromslice), cap(fromslice))

	copy(toslice, fromslice)
	fmt.Println(toslice)
}
