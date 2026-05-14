// ==================================================
// MAVZU: Tur konvertatsiyasi (Type Conversion)
// ==================================================
// Nazariya:
//   Go — statik tipli til. Ya'ni `int` ni `float64`ga AVTOMATIK
//   o'zgartira olmaysiz. Buni o'zingiz, ANIQ ko'rsatib qilishingiz kerak.
//
//   Sintaksis: T(value)  →  yangiQiymat := turi(eski_qiymat)
//
//   Misollar:
//     var x int = 10
//     var y float64 = float64(x)  // OK
//     var z int = int(3.99)       // 3 (kasr qismi tashlanadi!)
//
// Nima uchun?
//   - Tip xatolarini kompilyatsiyada tutadi (runtime'da emas).
//   - "Magic conversion" yo'q → kod aniq va ishonchli.
//
// Real misol:
//   Foydalanuvchi yoshi (int) va o'rtacha umr (float64) bilan ishlash:
//     yosh := 25
//     ortachaUmr := 73.5
//     qolganYil := ortachaUmr - float64(yosh)  // explicit conversion KERAK
//
// Aniqlik (precision) yo'qotish:
//   float → int : kasr qismi yo'qoladi (truncation, yaxlitlanmaydi!)
//   int64 → int32 : agar son katta bo'lsa, "kesilib" buziladi (overflow)

package main

import "fmt"

func main() {
	// --- 1. SONLI TURLAR ORASIDA ---
	fmt.Println("=== SONLI TURLAR ===")

	yosh := 25 // int
	ortachaUmr := 73.5 // float64

	// XATO bo'lardi: ortachaUmr - yosh  (turlari farq qiladi)
	// To'g'ri:
	qolganYil := ortachaUmr - float64(yosh)
	fmt.Printf("Taxminiy qolgan yil: %.1f\n", qolganYil)

	// float → int: kasr qismi YO'QOLADI (yaxlitlanmaydi!)
	narx := 199.99
	narxInt := int(narx)
	fmt.Println("199.99 → int =", narxInt) // 199, 200 emas!

	// Manfiy son: -3.7 → -3 (nolga tomon kesiladi)
	// Eslatma: int(-3.7) shaklida YOZILMAYDI — Go untyped konstantani
	// ruxsat etmaydi. Avval o'zgaruvchiga olib, keyin konvertatsiya:
	manfiy := -3.7
	fmt.Println("int(-3.7) =", int(manfiy)) // -3

	// --- 2. ANIQLIK YO'QOTISH (Overflow) ---
	fmt.Println("\n=== OVERFLOW ===")

	var katta int32 = 130
	var kichik int8 = int8(katta) // int8 max = 127
	// 130 int8'ga sig'maydi → "wrap" bo'ladi
	fmt.Println("130 → int8 =", kichik) // -126 (overflow!)

	// Bu kompilyatsiyada xato bermaydi — RUNTIME natija buziladi.
	// Doim diapazonni hisobga oling.

	// --- 3. byte va rune ---
	fmt.Println("\n=== byte va rune ===")
	// byte = uint8 alias (0-255), bitta ASCII belgini saqlaydi
	// rune = int32 alias, bitta Unicode belgini saqlaydi

	var b byte = 'A'
	var r rune = 'Ж' // Kirill harfi — ASCII'ga sig'maydi
	fmt.Printf("byte 'A' = %d, rune 'Ж' = %d\n", b, r)

	// Son → harf
	harf := rune(65) // 65 → 'A'
	fmt.Printf("rune(65) = %c\n", harf)

	// Harf → son (ASCII/Unicode kod)
	kod := int('z')
	fmt.Println("'z' kodi:", kod) // 122

	// --- 4. STRING va SON ---
	fmt.Println("\n=== STRING va SON ===")
	// MUHIM: string(65) — "65" emas, bu rune 65 → "A" beradi!
	xato := string(65)
	fmt.Println("string(65) =", xato) // "A", "65" EMAS

	// "25" stringini sonta o'tkazish — bu yerda T(x) ISHLAMAYDI!
	// strconv paketi kerak. Buni 03-strconv mavzusida ko'ramiz.

	// String — bayt'lar to'plami. Konvertatsiya orqali kirish mumkin:
	salom := "Hello"
	birinchiBayt := salom[0] // byte 'H' = 72
	fmt.Printf("salom[0] = %d (%c)\n", birinchiBayt, birinchiBayt)

	// String → []byte va []rune
	baytlar := []byte(salom)
	runelar := []rune(salom)
	fmt.Println("baytlar:", baytlar)
	fmt.Println("runelar:", runelar)

	// --- 5. AMALIY MISOL: O'rtacha baho ---
	fmt.Println("\n=== AMALIY MISOL ===")
	// 3 ta talaba balli (int), o'rtacha (float64) hisoblash
	bal1, bal2, bal3 := 85, 92, 78
	yigindi := bal1 + bal2 + bal3
	// XATO bo'lardi: ortacha := yigindi / 3 (int chiqadi, kasr yo'qoladi)
	ortacha := float64(yigindi) / 3.0
	fmt.Printf("O'rtacha bal: %.2f\n", ortacha) // 85.00
}

// ==================================================
// QACHON QO'LLANADI / QO'LLANMAYDI:
// ==================================================
// + Hisob-kitobda turlarni birlashtirish (int + float).
// + Eng kichik turdan kattaroq turga — xavfsiz.
// + ASCII/Unicode bilan ishlash (byte/rune).
// - Katta turdan kichikga — overflow xavfli, doim diapazonni tekshiring.
// - string(son) — bu Atoi EMAS! Buni unutmang. strconv.Itoa ishlating.
// ==================================================
