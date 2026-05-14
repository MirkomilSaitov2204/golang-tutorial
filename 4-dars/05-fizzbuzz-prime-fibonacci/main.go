// ==================================================
// MAVZU: Real algoritmik misollar
// ==================================================
// Bu faylda biz for sikllarini ishlatib, dasturlash interviyularida
// va kundalik vazifalarda eng ko'p uchraydigan 5 ta algoritmni yozamiz:
//
//   1. FizzBuzz — interview classic
//   2. Tub sonlar (prime numbers) — 1-100 oralig'ida
//   3. Fibonacci ketma-ketligi — birinchi 15 ta
//   4. Sonni teskari aylantirish (123 → 321)
//   5. Sonlar yig'indisi raqamlari bo'yicha (digital sum)
//
// Bularning hammasi REAL holatlarda kerak bo'ladi:
//   - Hash funksiyalar, checksum (digital sum)
//   - Cryptografiya (prime numbers)
//   - Dinamik dasturlash (Fibonacci)

package main

import "fmt"

func main() {
	// =========================================
	// 1. FIZZBUZZ
	// =========================================
	// Qoidalar:
	//   - 3 ga bo'linsa → "Fizz"
	//   - 5 ga bo'linsa → "Buzz"
	//   - HAR IKKALASIGA bo'linsa → "FizzBuzz" (ya'ni 15 ga)
	//   - Aks holda — son
	//
	// MUHIM: 15 ni AVVAL tekshiring, aks holda Fizz chiqib qoladi.

	fmt.Println("=== 1) FizzBuzz (1-30) ===")
	for i := 1; i <= 30; i++ {
		switch {
		case i%15 == 0:
			fmt.Println(i, "→ FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "→ Fizz")
		case i%5 == 0:
			fmt.Println(i, "→ Buzz")
		default:
			fmt.Println(i)
		}
	}

	// =========================================
	// 2. TUB SONLAR (PRIME NUMBERS) 1-50
	// =========================================
	// Tub son — faqat 1 va o'ziga bo'linadigan son (1 dan kattaroq).
	// Optimizatsiya: i ga sqrt(n) gacha bo'lib ko'rish kifoya.
	// Chunki agar n = a*b va a > sqrt(n), demak b < sqrt(n).

	fmt.Println("\n=== 2) Tub sonlar 2 dan 50 gacha ===")
	for n := 2; n <= 50; n++ {
		tubMi := true
		for i := 2; i*i <= n; i++ { // i*i <= n  — sqrt o'rniga
			if n%i == 0 {
				tubMi = false
				break
			}
		}
		if tubMi {
			fmt.Print(n, " ")
		}
	}
	fmt.Println()
	// 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47

	// =========================================
	// 3. FIBONACCI KETMA-KETLIGI
	// =========================================
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
	// Har bir element — oldingi ikkitasining yig'indisi.

	fmt.Println("\n=== 3) Fibonacci — birinchi 15 ta ===")
	a, b := 0, 1
	fmt.Print(a, " ", b, " ")
	for i := 0; i < 13; i++ {
		c := a + b
		fmt.Print(c, " ")
		a = b
		b = c
	}
	fmt.Println()
	// 0 1 1 2 3 5 8 13 21 34 55 89 144 233 377

	// Eslatma: REKURSIYA bilan ham yozish mumkin (6-darsda),
	// lekin iteratsiya O(n) — ancha tezroq.

	// =========================================
	// 4. SONNI TESKARI AYLANTIRISH
	// =========================================
	// 12345 → 54321
	// Algoritm:
	//   - har safar 10 ga modul olamiz → oxirgi raqam
	//   - natijaga qo'shamiz va 10 ga ko'paytirib siljitamiz
	//   - sonni 10 ga bo'lamiz

	fmt.Println("\n=== 4) Sonni teskari aylantirish ===")
	asl := 1234567
	son := asl
	teskari := 0
	for son > 0 {
		raqam := son % 10
		teskari = teskari*10 + raqam
		son /= 10
	}
	fmt.Printf("%d → %d\n", asl, teskari) // 1234567 → 7654321

	// =========================================
	// 5. RAQAMLAR YIG'INDISI (DIGITAL SUM)
	// =========================================
	// 1234 → 1+2+3+4 = 10
	// Bu Luhn algoritmida (kredit karta validatsiya) va checksum'larda ishlatiladi.

	fmt.Println("\n=== 5) Raqamlar yig'indisi ===")
	karta := 49271234567
	yigindi := 0
	tmp := karta
	for tmp > 0 {
		yigindi += tmp % 10
		tmp /= 10
	}
	fmt.Printf("%d → yig'indi = %d\n", karta, yigindi)

	// =========================================
	// 6. BONUS: PALINDROM TEKSHIRUVCHI
	// =========================================
	// 12321 → palindrom (chap-o'ng bir xil)
	// 12345 → emas

	fmt.Println("\n=== 6) Palindrom tekshiruv ===")
	tekshirish := []int{121, 12321, 12345, 9, 1001, 10}
	for _, x := range tekshirish {
		asl := x
		teskari := 0
		for x > 0 {
			teskari = teskari*10 + x%10
			x /= 10
		}
		if asl == teskari {
			fmt.Printf("%d — palindrom\n", asl)
		} else {
			fmt.Printf("%d — palindrom emas\n", asl)
		}
	}

	// =========================================
	// 7. BONUS: ENG KATTA UMUMIY BO'LUVCHI (GCD)
	// =========================================
	// Yevklid algoritmi — eng tezligi bo'yicha mashhur.
	// Misol: gcd(48, 18) = 6

	fmt.Println("\n=== 7) GCD (Yevklid algoritmi) ===")
	x, y := 48, 18
	a, b = x, y
	for b != 0 {
		a, b = b, a%b
	}
	fmt.Printf("gcd(%d, %d) = %d\n", x, y, a)
}

// ==================================================
// AMALIY VAZIFA:
// ==================================================
// 1. FizzBuzz'ni KENGAYTIRING: agar son 7 ga bo'linsa "Boom" chiqsin.
//    Variantlar: Fizz, Buzz, Boom, FizzBuzz, FizzBoom, BuzzBoom, FizzBuzzBoom.
//
// 2. Birinchi 50 ta tub sonni topib, ulardan ikkita qo'shni tub son
//    o'rtasidagi maksimum farqni hisoblang.
//
// 3. Fibonacci'da: 1000000 (bir million) dan kichik bo'lgan ENG KATTA
//    Fibonacci sonini toping.
//
// 4. Luhn algoritmi: kredit karta raqamining oxirgi raqami checksum.
//    O'qing va o'z funksiyangizni yozing (Google'da Luhn algorithm).
//
// 5. Palindrom: STRING palindrom tekshiruvchini yozing (son emas).
//    Masalan "ana", "kabak" — palindrom; "olma" — emas.
// ==================================================
