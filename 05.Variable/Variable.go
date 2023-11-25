/*
todo: start
!<==== Variable ====>
* Variable adalah tempat untuk menyimpan data
* Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
* Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
* Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

? Tipe Data Variable
* Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
* Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

? Kata Kunci Var
* Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
* Asalkan saat membuat variable kita langsung menginisialisasi datanya
* Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut

? Deklarasi Multiple Variable
* Di Go-Lang kita bisa membuat variable secara sekaligus banyak
* Code yang dibuat akan lebih bagus dan mudah dibaca

? Constant
* Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
* Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
* Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

todo: end
*/

package main

import "fmt"

func main() {
	var name string
	name = "Wiliam Fery Sagala"
	fmt.Println("Name:", name)
	var age = 21
	fmt.Println("Umur:" , age)
	isFemale := false
	if isFemale {
		fmt.Println("Gender:Perempuan")
	}else {
		fmt.Println("Gender:Laki-Laki")
	}
	var (
		userid = 12324
		birhtDay = "16/08/2002"
	)

	fmt.Println("User Id:", userid)
	fmt.Println("TTL:", birhtDay)

	const PI float32 = 3.14 // Constant tidak dapat di ubah
	fmt.Println("PI =",PI)
}