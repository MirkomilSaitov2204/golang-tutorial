// =============================================================================
// MAVZU 2: O'ZGARUVCHILAR (VARIABLES) VA O'ZGARMASLAR (CONSTANTS)
// =============================================================================
//
// REAL HAYOTDAN MISOL:
//   Tasavvur qiling, siz e-commerce sayti yozyapsiz. Sizga kerak:
//     - Foydalanuvchi yoshi (har bir foydalanuvchida boshqacha) -> O'ZGARUVCHI
//     - Foydalanuvchi ismi (har birida boshqacha) -> O'ZGARUVCHI
//     - VIP narxlari uchun chegirma foizi (har doim 15%) -> O'ZGARMAS (const)
//     - Saytning URL'i (hech qachon o'zgarmaydi) -> O'ZGARMAS (const)
//
// NAZARIYA:
//
//   Go — STATIK TIPLI til. Ya'ni o'zgaruvchining turi (int, string va h.k.)
//   kompilyatsiya vaqtida aniqlanadi va keyin O'ZGARMAYDI.
//   Bu xato qilish ehtimolini kamaytiradi.
//
//   O'zgaruvchini e'lon qilishning 4 ta usuli:
//     1. var ism string = "Ali"   -> to'liq forma
//     2. var ism = "Ali"          -> tur avtomatik aniqlanadi (string)
//     3. var ism string           -> faqat e'lon (qiymat "" — bo'sh string)
//     4. ism := "Ali"             -> qisqa shakl, FAQAT funksiya ICHIDA
//
//   ZERO VALUES (default qiymatlar):
//     - int     -> 0
//     - float64 -> 0.0
//     - string  -> "" (bo'sh string)
//     - bool    -> false
//     - pointer/slice/map -> nil
//
//   O'ZGARMAS (const):
//     - Kompilyatsiya vaqtida hisoblanadi
//     - Faqat oddiy turlar (raqam, satr, bool) bo'lishi mumkin
//     - Hech qachon o'zgartirib bo'lmaydi
// =============================================================================

package main

import "fmt"

// Paket darajasidagi o'zgarmaslar — odatda UPPER_CASE yoki PascalCase
const (
	SaytNomi      = "MaxMarket"
	VIPChegirmasi = 15.0 // foizda
	MaxYosh       = 120
)

// Paket darajasidagi o'zgaruvchilar (funksiya tashqarisida)
var dasturVersiyasi = "1.0.0"

func main() {
	// -----------------------------------------------------------------
	// 1-USUL: To'liq forma — var NOMI TURI = QIYMAT
	// -----------------------------------------------------------------
	var foydalanuvchiYoshi int = 25
	var foydalanuvchiIsmi string = "Aziz"
	fmt.Println("1-usul:", foydalanuvchiIsmi, foydalanuvchiYoshi)

	// -----------------------------------------------------------------
	// 2-USUL: Tur ko'rsatilmaydi — Go o'zi aniqlaydi (type inference)
	// -----------------------------------------------------------------
	var balansi = 1500.50 // Go: bu float64
	var aktivMi = true    // Go: bu bool
	fmt.Println("2-usul:", balansi, aktivMi)

	// -----------------------------------------------------------------
	// 3-USUL: Faqat e'lon (qiymat keyinroq beriladi)
	// O'zgaruvchi ZERO VALUE bilan boshlanadi.
	// -----------------------------------------------------------------
	var savatdaMahsulot int
	var promokod string
	fmt.Println("3-usul (zero values):", savatdaMahsulot, "->", promokod, "<-")
	// 0 va bo'sh string (-> <-) — chunki ular zero values

	savatdaMahsulot = 3       // keyin qiymat beramiz
	promokod = "WELCOME2026"  // keyin qiymat beramiz
	fmt.Println("3-usul (keyin):", savatdaMahsulot, promokod)

	// -----------------------------------------------------------------
	// 4-USUL: Qisqa shakl :=  (eng ko'p ishlatiladi)
	// !!! Faqat funksiya ICHIDA ishlatiladi
	// -----------------------------------------------------------------
	mahsulotNomi := "Sichqoncha"
	narxi := 150000
	fmt.Println("4-usul:", mahsulotNomi, narxi, "so'm")

	// -----------------------------------------------------------------
	// MULTIPLE ASSIGNMENT — bir vaqtning o'zida bir nechta o'zgaruvchi
	// -----------------------------------------------------------------
	x, y, z := 10, 20, 30
	fmt.Println("Ko'p qiymat:", x, y, z)

	// Qiymatlarni almashtirish (swap) — bunday tilning chiroyli xususiyati!
	x, y = y, x
	fmt.Println("Swap'dan keyin: x =", x, ", y =", y)

	// -----------------------------------------------------------------
	// O'ZGARMASLAR — const
	// -----------------------------------------------------------------
	const PI = 3.14159
	const Salom = "Salom!"
	fmt.Println("Sayt:", SaytNomi, "| VIP chegirma:", VIPChegirmasi, "%")
	fmt.Println("PI =", PI, "|", Salom)

	// Quyidagi qator XATO beradi (kompilyatsiya bosqichida):
	// PI = 3.14  // ❌  cannot assign to PI

	// -----------------------------------------------------------------
	// REAL MISOL: VIP foydalanuvchi uchun chegirmali narx
	// -----------------------------------------------------------------
	asilNarx := 1_000_000.0 // _ ni o'qish uchun ishlatish mumkin (1 mln)
	chegirma := asilNarx * VIPChegirmasi / 100
	yakuniyNarx := asilNarx - chegirma

	fmt.Printf("\n--- VIP Hisob-kitob ---\n")
	fmt.Println("Asl narx:   ", asilNarx)
	fmt.Println("Chegirma:   ", chegirma)
	fmt.Println("Siz to'laysiz:", yakuniyNarx, "so'm")
	fmt.Println("Versiya:", dasturVersiyasi)
}

// =============================================================================
// AMALIY VAZIFA:
//   1. O'zingiz uchun o'zgaruvchilar yarating:
//        - ismingiz (string), yoshingiz (int), bo'yingiz (float64), talabaMisiz (bool)
//   2. Bularni `:=` shakli bilan yarating.
//   3. Printf yoki Println orqali chiqaring.
//   4. Bir const yarating — `OliyTalabaSinfi` deb va qiymat bering.
//   5. Const'ni o'zgartirib ko'ring — qanday xato ko'rinadi?
// =============================================================================
