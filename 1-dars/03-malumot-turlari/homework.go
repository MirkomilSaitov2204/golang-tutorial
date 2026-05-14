package main

// AMALIY VAZIFA:
//   1. 3 ta o'zgaruvchi yarating: a, b, c — turli turlardan.
//   2. Ularning turini Printf("%T") orqali chiqaring.
//   3. int va float64'ni qo'shing — xato olasiz. Keyin to'g'rilang.
//   4. O'zingizning ismingizni `[]rune` ga aylantiring va belgilar sonini chiqaring.
//   5. Bir string'ni ikkinchisiga `+` orqali ulang va natijani chiqaring.

import (
	"fmt" 
)

func main(){
	var a int = 10
	var b float64 = 3.14
	var c bool = true


	fmt.Printf("%T", a)
	fmt.Printf("%T", b)
	fmt.Printf("%T", c)

	// int va float64'ni qo'shish
	// xato: summa := a + b  // ❌ mismatched types int and float64

	// To'g'rilangan versiya:
	summa := float64(a) + b
	fmt.Println("Yig'indi:", summa)

	// O'zingizning ismingizni []rune ga aylantirish
	ism := "Sardor"
	runelar := []rune(ism)
	fmt.Printf("Ism: %s, Belgilar soni: %d\n", ism, len(runelar))

	// Stringlarni + orqali ulash
	str1 := "Salom"
	str2 := "Dunyo"
	ulangan := str1 + " " + str2
	fmt.Println("Ulangan string:", ulangan)

}