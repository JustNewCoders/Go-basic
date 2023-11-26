/*
todo: start
!<==== Operasi Perbandingan ====>
* Operasi perbandingan adalah operasi untuk membandingkan dua buah data
* Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
* Jika hasil operasinya adalah benar, maka nilainya adalah true
* Jika hasil operasinya adalah salah, maka nilainya adalah false

? Operator Perbandingan
* > ==>Lebih Dari
* < ==>Kurang Dari
* >= ==>Lebih Dari Sama Dengan
* <= ==>Kurang Dari Sama Dengan
* == ==>Sama Dengan
* != ==> Tidak Sama Dengan


todo: end
*/

package main

import "fmt"

func main() {
	var name1,name2 string
	name1 = "wiliam"
	name2 = "wiliam"

	var result bool = name1 == name2

	fmt.Println(result)
}