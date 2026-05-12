// =============================================================================
// MAVZU 3: MA'LUMOT TURLARI (DATA TYPES)
// =============================================================================
//
// REAL HAYOTDAN MISOL:
//   Tasavvur qiling, siz bank ilovasini yozyapsiz:
//     - Hisob raqami         -> int64 (juda katta son, 19 ta raqam bo'lishi mumkin)
//     - Foydalanuvchi yoshi  -> int (kichikroq son, 32-bit yetadi)
//     - Hisobdagi pul        -> float64 (puxta hisob uchun aslida `decimal` paket
//                                ishlatamiz, lekin hozircha float64 misol uchun)
//     - To'liq ismi          -> string
//     - Hisob aktivmi        -> bool
//     - Karta raqami         -> []byte yoki string (xavfsizlik uchun byte)
//
// NAZARIYA — GO'NING ASOSIY TURLARI:
//
//   1. SONLI (Numeric) turlar:
//      - int, int8, int16, int32, int64
//        (int — platformaga bog'liq: 32-bit yoki 64-bit OS'da mos keladi)
//      - uint, uint8, uint16, uint32, uint64  (manfiy bo'lmagan)
//      - float32, float64  (kasr sonlar)
//      - byte  -> uint8 ning aliasi  (1 bayt)
//      - rune  -> int32 ning aliasi (Unicode kod nuqtasi)
//
//   2. STRING (string) — UTF-8 da kodlangan baytlar ketma-ketligi
//
//   3. BOOL (bool) — faqat `true` yoki `false`
//
//   4. COMPLEX turlar (complex64, complex128) — kamdan-kam ishlatiladi
//
// MUHIM QOIDA: AVTOMATIK TUR KONVERTATSIYASI YO'Q!
//   Boshqa tillarda int va float qo'shsa ishlaydi, Go'da esa AYNAN
//   kerakli tur ko'rsatilishi shart. Bu xato qilishni kamaytiradi.
// =============================================================================

package main

import (
	"fmt"
	"math"
)

func main() {
	// -----------------------------------------------------------------
	// 1) BUTUN SONLAR (Integers)
	// -----------------------------------------------------------------
	var yosh int = 25                       // platformaga bog'liq
	var aholiSoni int32 = 37_500_000        // 32-bit yetarli
	var hisobRaqami int64 = 9_860_001_010   // katta son
	var yoshlar uint8 = 200                 // 0..255 oraliqda
	fmt.Println("Yosh:", yosh, "| Aholi:", aholiSoni, "| Hisob:", hisobRaqami, "| uint8:", yoshlar)

	// Maksimal qiymatlar:
	fmt.Println("int8 max:", math.MaxInt8, "| int64 max:", math.MaxInt64)

	// -----------------------------------------------------------------
	// 2) KASR SONLAR (float)
	// -----------------------------------------------------------------
	var pi float64 = 3.14159265358979
	var temperatura float32 = 36.6
	balans := 1_500_000.75 // : float64 (default)
	fmt.Println("Pi:", pi, "| Harorat:", temperatura, "| Balans:", balans)

	// EHTIYOT BO'LING: float aniq emas
	// 0.1 + 0.2 == 0.3  NATIJASI false (chunki binary representation)
	fmt.Println("0.1 + 0.2 =", 0.1+0.2, "| 0.3 ga teng mi?", 0.1+0.2 == 0.3)

	// -----------------------------------------------------------------
	// 3) MANTIQIY (bool)
	// -----------------------------------------------------------------
	var aktivMi bool = true
	emailTasdiqlanganMi := false
	fmt.Println("Aktiv:", aktivMi, "| Email tasdiqlanganmi:", emailTasdiqlanganMi)

	// -----------------------------------------------------------------
	// 4) STRING (matnlar)
	// -----------------------------------------------------------------
	ism := "Ali Valiev"
	salom := "Salom, dunyo!"
	fmt.Println(salom, "| Ism:", ism)
	fmt.Println("Ism uzunligi (BAYT):", len(ism))

	// String konkatenatsiya
	familiya := "Karimov"
	tolaIsm := "Aziz " + familiya
	fmt.Println("To'liq ism:", tolaIsm)

	// Multi-line string — backtick `` ishlatiladi
	xat := `Salom Ali,
Bizning saytimizga xush kelibsiz!
Hurmat bilan, MaxMarket.`
	fmt.Println(xat)

	// -----------------------------------------------------------------
	// 5) RUNE va BYTE
	//    - byte = uint8 (1 ta bayt, 0..255)
	//    - rune = int32 (Unicode kod nuqtasi)
	//    Bu farq O'ZBEK harflari uchun MUHIM!
	// -----------------------------------------------------------------
	matn := "Salom"
	fmt.Println("len(\"Salom\"):", len(matn)) // 5

	ozbekcha := "O'zbek"
	fmt.Println("len(\"O'zbek\"):", len(ozbekcha))
	// O' va bek -> kutilgan: 6, lekin O' bir nechta bayt bo'lishi mumkin

	// Rune slice'iga aylantirib aniq belgilar sonini olamiz
	runelar := []rune(ozbekcha)
	fmt.Println("Belgilar (rune) soni:", len(runelar))

	// Birinchi belgini chiqarish
	birinchiBelgi := runelar[0]
	fmt.Printf("Birinchi belgi: %c (kod: %d)\n", birinchiBelgi, birinchiBelgi)

	// -----------------------------------------------------------------
	// 6) TUR KONVERTATSIYASI (Type Conversion)
	//    Go AVTOMATIK qilmaydi — qo'lda yozish kerak: T(qiymat)
	// -----------------------------------------------------------------
	var butun int = 100
	var kasr float64 = float64(butun) // int -> float64
	var orqaga int = int(kasr)        // float64 -> int (kasr qismi tushiriladi)
	fmt.Println("int:", butun, "| float:", kasr, "| orqaga int:", orqaga)

	// Quyidagi qator XATO beradi:
	// summa := butun + kasr  // ❌ mismatched types int and float64

	// To'g'risi:
	summa := float64(butun) + kasr
	fmt.Println("Yig'indi:", summa)

	// -----------------------------------------------------------------
	// 7) TURNI ANIQLASH — %T verb'i orqali
	// -----------------------------------------------------------------
	fmt.Printf("yosh turi: %T\n", yosh)
	fmt.Printf("balans turi: %T\n", balans)
	fmt.Printf("ism turi: %T\n", ism)
	fmt.Printf("aktivMi turi: %T\n", aktivMi)
}

// =============================================================================
// AMALIY VAZIFA:
//   1. 3 ta o'zgaruvchi yarating: a, b, c — turli turlardan.
//   2. Ularning turini Printf("%T") orqali chiqaring.
//   3. int va float64'ni qo'shing — xato olasiz. Keyin to'g'rilang.
//   4. O'zingizning ismingizni `[]rune` ga aylantiring va belgilar sonini chiqaring.
//   5. Bir string'ni ikkinchisiga `+` orqali ulang va natijani chiqaring.
// =============================================================================
