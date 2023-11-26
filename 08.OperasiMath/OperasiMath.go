/*
!<==== Operasi Matematika ====>
? Aritmatika
* +  ==> penjumlahan
* - ==>	pengurangan
* * ==> perkalian
* / ==> pembagian
* % ==> sisa bagi

? Augmented Assignments
*	aritmatika == Augmented Assignments
* a = a + 10 == a += 10
* a = a - 10 == a -= 10
* a = a * 10 == a *= 10
* a = a / 10 == a /= 10
* a = a % 10 == a %= 10

? Unary Operator
* ++ ==> a = a + 1
* -- ==> a = a - 1
* - ==> Negative
* + ==> Positive
* ! ==> Boolean kebalik


*/

package main

import "fmt"

func main() {
	// ! Aritmatika
	var a,b int32
	a = 10
	b = 5

	//  ! Augmented Assignments
	var c int32 = 22
	c += 10

	// ! Unary Operator
	var d int32 = 1
	d++
	d++
	
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)
	fmt.Println("a + b =",a+b) //? Penjumlahan
	fmt.Println("a - b =",a-b) //? pengurangan
	fmt.Println("a * b =",a*b) //? perkalian
	fmt.Println("a / b =",a/b) //? pembagian
	fmt.Println("a % b =",a%b) //? sisa bagi
}