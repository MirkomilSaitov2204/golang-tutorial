// ==================================================
// MAVZU: Operatorlar (Operators)
// ==================================================
// Nazariya:
//   Operator — bu bir yoki bir nechta qiymat ustida amal bajaruvchi belgi.
//   Go'da 4 ta asosiy guruh bor:
//
//   1) ARIFMETIK: +, -, *, /, %  (modul — qoldiq)
//      Eslatma: butun son bo'lishi (int/int) qoldiqsiz: 7/2 = 3
//      Float bo'lsa to'liq: 7.0/2.0 = 3.5
//
//   2) TAQQOSLASH: ==, !=, <, >, <=, >=
//      Natija — har doim bool (true/false)
//
//   3) MANTIQIY: && (va), || (yoki), ! (emas)
//      Short-circuit: agar && da chap noto'g'ri bo'lsa, o'ng tekshirilmaydi
//
//   4) BITWISE (bit darajasidagi): &, |, ^, <<, >>
//      Kamdan-kam ishlatiladi (flag'lar, encryption, optimizatsiya).
//
// Real misol:
//   E-commerce: mahsulot narxi 100ming, 20% chegirma.
//   Yangi narx = 100000 - (100000 * 20 / 100) = 80000.
//   VIP mijoz VA balansi > 500ming bo'lsa qo'shimcha chegirma.
//
// Operatorlar prioriteti (yuqoridan pastga):
//   1. *  /  %  <<  >>  &
//   2. +  -  |  ^
//   3. ==  !=  <  <=  >  >=
//   4. &&
//   5. ||
//   Ya'ni: 2 + 3*4 = 14 (3*4 birinchi)

package main

import "fmt"

func main() {
	// --- 1. ARIFMETIK ---
	fmt.Println("=== ARIFMETIK ===")

	// Mahsulot narxi va chegirma hisoblash
	narx := 100000
	chegirmaFoiz := 20
	chegirmaSummasi := narx * chegirmaFoiz / 100 // 20000
	yangiNarx := narx - chegirmaSummasi          // 80000
	fmt.Printf("Eski narx: %d, chegirma: %d, yangi narx: %d\n", narx, chegirmaSummasi, yangiNarx)

	// Modul (qoldiq) — juft/toq tekshirish uchun ham ishlatiladi
	son := 17
	fmt.Printf("17 / 5 = %d (butun bo'lish)\n", son/5)
	fmt.Printf("17 %% 5 = %d (qoldiq)\n", son%5)

	// Butun va float farqi — diqqat!
	fmt.Println("7 / 2 (int) =", 7/2)         // 3
	fmt.Println("7.0 / 2.0 (float) =", 7.0/2) // 3.5

	// Qisqa shakldagi operatorlar
	balans := 1000
	balans += 500 // balans = balans + 500
	balans -= 200 // balans = balans - 200
	balans *= 2   // balans = balans * 2
	fmt.Println("Yakuniy balans:", balans) // (1000+500-200)*2 = 2600

	// --- 2. TAQQOSLASH ---
	fmt.Println("\n=== TAQQOSLASH ===")

	yosh := 19
	fmt.Println("yosh >= 18:", yosh >= 18) // true (voyaga yetgan)
	fmt.Println("yosh == 21:", yosh == 21) // false
	fmt.Println("yosh != 18:", yosh != 18) // true

	// String'lar ham taqqoslanishi mumkin
	ism1 := "Ali"
	ism2 := "Vali"
	fmt.Println("Ali == Vali:", ism1 == ism2) // false

	// --- 3. MANTIQIY ---
	fmt.Println("\n=== MANTIQIY ===")

	vip := true
	hisobBalans := 600000
	// VIP VA balansi 500ming dan ko'p bo'lsa — qo'shimcha chegirma
	qoshimchaChegirma := vip && hisobBalans > 500000
	fmt.Println("Qo'shimcha chegirma huquqi:", qoshimchaChegirma) // true

	// YOKI — yosh < 18 yoki student bo'lsa, chegirma
	yoshMijoz := 16
	student := false
	yoshlarChegirmasi := yoshMijoz < 18 || student
	fmt.Println("Yoshlar chegirmasi:", yoshlarChegirmasi) // true (yosh < 18)

	// EMAS (NOT)
	tizimYopiq := false
	fmt.Println("Tizim ochiqmi?:", !tizimYopiq) // true

	// Short-circuit misoli — agar chap false bo'lsa, o'ng tekshirilmaydi
	// (foydali: nil pointerni tekshirishda crash bo'lmaydi)
	a := 0
	b := 10
	if a != 0 && b/a > 2 { // a==0 sababli b/a hech qachon hisoblanmaydi
		fmt.Println("Bu chiqmaydi")
	} else {
		fmt.Println("Short-circuit ishladi — crash bo'lmadi")
	}

	// --- 4. BITWISE (qisqacha) ---
	fmt.Println("\n=== BITWISE (qisqacha) ===")
	// 5 = 0101
	// 3 = 0011
	fmt.Println("5 & 3 =", 5&3) // 0001 = 1 (AND)
	fmt.Println("5 | 3 =", 5|3) // 0111 = 7 (OR)
	fmt.Println("5 ^ 3 =", 5^3) // 0110 = 6 (XOR)
	fmt.Println("5 << 1 =", 5<<1) // 1010 = 10 (chapga siljish — 2 ga ko'paytirish)
	fmt.Println("5 >> 1 =", 5>>1) // 0010 = 2 (o'ngga siljish — 2 ga bo'lish)

	// --- 5. PRIORITET MISOLI ---
	fmt.Println("\n=== PRIORITET ===")
	natija := 2 + 3*4 // 14, chunki * birinchi
	fmt.Println("2 + 3*4 =", natija)

	natija2 := (2 + 3) * 4 // 20, qavs birinchi
	fmt.Println("(2 + 3) * 4 =", natija2)

	// Mantiqiy operatorlar prioritet:  && | ||
	// Quyidagi: yosh >= 18 BIRINCHI, vip BIRINCHI, keyin &&, keyin ||
	natija3 := yosh >= 18 && vip || student
	fmt.Println("yosh>=18 && vip || student:", natija3) // true
}

// ==================================================
// QACHON QO'LLANADI / QO'LLANMAYDI:
// ==================================================
// + Arifmetik: hisob-kitob — narx, soliq, statistika.
// + Mantiqiy: shartlarni birlashtirib tekshirish (auth, validation).
// + Taqqoslash: filter, sort, validation.
// - Bitwise: kundalik kodda kam — performance-critical yoki low-level joyda.
//
// ESLATMA — float bilan ehtiyot:
//   0.1 + 0.2 == 0.3 → false (floating point xato)
//   Pul bilan ishlash: int (tiyinda saqlash) yoki maxsus decimal kutubxona.
// ==================================================
