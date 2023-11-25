/*
todo: start
!<==== Konversi Tipe Data ====>
* Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
* Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

todo: end
*/

package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	nilai16 := int16(nilai32)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
}