/*
todo: start
!<==== Type Declaration ====>
* Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
* Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
todo: end
*/

package main

import "fmt"

func main() {

	type Names string
	var firstname Names = "Wiliam"
	var midlename Names = "Fery"
	var lastname Names = "Sagala"

	fmt.Println("Name:", firstname,midlename,lastname)
	
}