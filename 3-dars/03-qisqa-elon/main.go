// ==================================================
// MAVZU: if ichidagi qisqa e'lon (short statement)
// ==================================================
// Nazariya:
//   Go'da `if` operatori ICHIDA o'zgaruvchi e'lon qilib, KEYIN o'sha
//   o'zgaruvchini tekshirish mumkin. Bu — Go'ning eng go'zal idiomalaridan biri.
//
//   Sintaksis:
//     if x := funksiya(); shart {
//         // x faqat shu blok ichida mavjud
//     } else {
//         // x bu yerda HAM mavjud
//     }
//     // x bu yerda MAVJUD EMAS
//
//   Foydasi:
//     1. O'zgaruvchining QAMROVI (scope) tor — keraksiz joyda mavjud bo'lmaydi.
//     2. Kod ixcham va o'qilishi oson.
//     3. Error handling Go'da SHU PATTERN orqali yozilad:
//        if err := f(); err != nil { ... }
//
// Real misol:
//   strconv.Atoi natijasini darhol tekshirish — eng ko'p uchraydigan misol.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// --- 1. ODDIY MISOL ---
	if yosh := 25; yosh >= 18 {
		fmt.Println("Voyaga yetgan, yoshi:", yosh)
	}
	// fmt.Println(yosh) — XATO! `yosh` bu yerda mavjud emas.

	// --- 2. ERROR HANDLING ASOSIY PATTERN ---
	// Bu — Go'ning eng ko'p uchraydigan shakli.
	javob := "42"
	if son, xato := strconv.Atoi(javob); xato != nil {
		fmt.Println("Parse xato:", xato)
	} else {
		fmt.Println("Son:", son)
	}
	// Diqqat — `son` va `xato` HAR IKKI tarmoqda ham mavjud (if va else).

	// --- 3. FUNKSIYA NATIJASINI TEKSHIRISH ---
	if natija := 100 * 50 / 4; natija > 1000 {
		fmt.Println("Katta natija:", natija)
	}

	// --- 4. AMALIY MISOL: foydalanuvchi inputini validate qilish ---
	input := "  ALI  "

	// strings.TrimSpace yordamida tozalash va length tekshirish
	if tozalangan := strings.TrimSpace(input); len(tozalangan) == 0 {
		fmt.Println("Bo'sh kiritma!")
	} else if len(tozalangan) < 3 {
		fmt.Println("Juda qisqa:", tozalangan)
	} else {
		fmt.Println("OK:", tozalangan)
	}

	// --- 5. ENG MUHIM PATTERN: nested error handling ---
	pul := "12500.75"

	if narx, xato := strconv.ParseFloat(pul, 64); xato != nil {
		fmt.Println("Pul noto'g'ri:", xato)
	} else if narx <= 0 {
		fmt.Println("Pul musbat bo'lishi kerak.")
	} else if narx > 1_000_000_000 {
		fmt.Println("Pul juda katta.")
	} else {
		// barcha tekshiruvlardan o'tdi
		fmt.Printf("To'lov qabul qilindi: %.2f\n", narx)
	}

	// --- 6. switch ICHIDA HAM SHU PATTERN ISHLAYDI ---
	// (avvalgi 02-switch faylida ham ko'rganmiz)
	switch hozir := 14; { // hozir = soat 14
	case hozir < 6:
		fmt.Println("Tun")
	case hozir < 12:
		fmt.Println("Ertalab")
	case hozir < 18:
		fmt.Println("Kunduz")
	default:
		fmt.Println("Kech")
	}

	// --- 7. NIMA UCHUN KERAK? Scope'ning ahamiyati ---
	// YOMON kod (o'zgaruvchi tashqarida — keraksiz scope'da):
	yosh2, xato1 := strconv.Atoi("abc")
	if xato1 != nil {
		fmt.Println("Parse muvaffaqiyatsiz:", xato1)
	} else {
		fmt.Println("Yosh:", yosh2)
	}
	// Hatto bu yerdan keyin ham `yosh2` va `xato1` MAVJUD — keraksiz scope.

	// YAXSHI kod (o'zgaruvchi faqat kerak joyda):
	if y, x := strconv.Atoi("123"); x == nil {
		fmt.Println("Yosh:", y)
	}
	// `y` va `x` bu yerda mavjud emas — toza kod, kichikroq scope.

	// --- 8. CHUQURROQ: "comma-ok" idiom bilan ---
	// Bu pattern map va channel'da ham uchraydi (8 va 27 darslarda).
	xaritaMisol := map[string]int{"olma": 100, "anor": 200}
	if narx, mavjud := xaritaMisol["olma"]; mavjud {
		fmt.Println("Olma narxi:", narx)
	}
	if _, mavjud := xaritaMisol["banan"]; !mavjud {
		fmt.Println("Banan yo'q.")
	}
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + DOIM error handling'da (idiomatik Go).
// + Funksiya natijasini darhol tekshirib, bir blokda ishlatish.
// + Map'da kalitning mavjudligini tekshirish.
// + Type assertion (12-darsda).
//
// QOIDA:
//   Agar o'zgaruvchi if'dan TASHQARIDA kerak bo'lmasa, har doim shu shaklni ishlating.
//   Bu Go ruhi: scope toraytirish — bug'larni kamaytiradi.
// ==================================================
