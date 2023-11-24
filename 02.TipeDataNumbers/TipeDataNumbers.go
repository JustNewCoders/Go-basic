package main

/*
todo: start
!<==== Tipe data numbers ====>
* ada dua jenis tipe data number :
	* integer
	* floating point

? Integer
__________________________________________________________________
|	Tipe Data		|		Nilai Minimum				|		Nilai Maksimum				|
------------------------------------------------------------------
|	int8				|	-128									|	127											|
|	int16				|	-32768								|	32767										|
|	int32				|	-2147483648						|	2147483647							|
|	int64				|	-9223372036854775808	|	9223372036854775808			|
|	uint8				|	0											|	255											|
|	uint16			|	0											|	65535										|
|	uint32			|	0											|	4294967295							|
|	uint64			|	0											|	18446744073709551615		|
------------------------------------------------------------------

? Floating Point
_______________________________________________________________________
|	Tipe Data		|		Nilai Minimum									|		Nilai Maksimum		|
-----------------------------------------------------------------------
|	float32			|	1.18×10^-38											|	3.4×10^38						|
|	float64			|	2.24×10^-308										|	1.80×10^308					|
|	complex64		|	complex numbers with float32 real and imagenry parts	|
|	complex128	|	complex numbers with float32 real and imagenry parts	|
-----------------------------------------------------------------------

? Alias
_______________________________
|	Alias		|		Alias untuk			|
-------------------------------
| byte		|		uint8						|
| rune		|		int32						|
| int			|		minimal int32		|
| uint		|		minimal uint32	|
-------------------------------

todo: end
*/

import "fmt"

func main() {
	fmt.Println("Satu =", 1)
	fmt.Println("dua =", 2)
	fmt.Println("Satu koma 2 =", 1.2)
}