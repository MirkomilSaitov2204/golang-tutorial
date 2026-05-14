// ==================================================
// MAVZU: strconv paketi (String ↔ Son konvertatsiyasi)
// ==================================================
// Nazariya:
//   Foydalanuvchi terminaldan "25" deb yozsa, bu STRING ("25"), son emas!
//   Uni `int`ga o'tkazish uchun `strconv` paketi kerak.
//
//   Asosiy funksiyalar:
//     strconv.Atoi(s)              → int, error    (string → int)
//     strconv.Itoa(i)              → string         (int → string)
//     strconv.ParseFloat(s, 64)    → float64, error
//     strconv.FormatFloat(f,'f',2,64) → string ("123.45")
//     strconv.ParseBool("true")    → bool, error    ("true"/"false"/"1"/"0")
//     strconv.FormatBool(b)        → string
//
//   "i" — int, "f" — float, "b" — bool. Esda saqlab oling:
//     Atoi = ASCII to Integer
//     Itoa = Integer to ASCII
//
// Birinchi tanishuv: ERROR
//   Atoi ikkita qiymat qaytaradi: natija VA xato (error).
//   Agar string noto'g'ri bo'lsa (masalan "abc"), xato qaytadi.
//   Bu Go'ning idiomatik usuli — xatolar interfeysini 14-darsda chuqurroq ko'ramiz.
//
// Real misol:
//   Web forma: foydalanuvchi yoshini kiritadi.
//   HTTP request'dan "25" string kelyapti → uni songa aylantirib,
//   yoshini tekshirish (>= 18).

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// --- 1. STRING → INT (Atoi) ---
	fmt.Println("=== Atoi (string → int) ===")

	yoshStr := "25"
	yosh, xato := strconv.Atoi(yoshStr)
	if xato != nil {
		// Bu blokga tushmaymiz — "25" to'g'ri.
		fmt.Println("Xato:", xato)
	} else {
		fmt.Println("Yosh:", yosh, "tipi:", fmt.Sprintf("%T", yosh))
	}

	// Yomon ma'lumot — xato qaytadi
	yomon := "abc"
	natija, xato2 := strconv.Atoi(yomon)
	fmt.Println("natija:", natija, "xato:", xato2)
	// natija: 0 (zero value), xato: strconv.Atoi: parsing "abc": invalid syntax

	// --- 2. INT → STRING (Itoa) ---
	fmt.Println("\n=== Itoa (int → string) ===")

	tugilganYil := 2000
	yil := strconv.Itoa(tugilganYil)
	xabar := "Siz " + yil + "-yilda tug'ilgansiz."
	fmt.Println(xabar)
	// Eslatma: int va string'ni TO'G'RIDAN-TO'G'RI + qilib bo'lmaydi!

	// --- 3. STRING → FLOAT ---
	fmt.Println("\n=== ParseFloat (string → float64) ===")

	narxStr := "199.99"
	narx, xato3 := strconv.ParseFloat(narxStr, 64) // 64 = float64
	if xato3 == nil {
		fmt.Println("Narx (float):", narx)
		fmt.Printf("Solid qo'shilgan narx: %.2f\n", narx*1.12)
	}

	// --- 4. FLOAT → STRING ---
	fmt.Println("\n=== FormatFloat (float → string) ===")

	pi := 3.14159265
	// 'f' formati: oddiy, 'e' formati: 3.14e+00
	// precision: nechta kasr (2 = ikki kasr)
	piStr := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Println("Pi (string, 2 kasr):", piStr) // "3.14"

	piTolaStr := strconv.FormatFloat(pi, 'f', -1, 64) // -1 = barchasi
	fmt.Println("Pi to'la:", piTolaStr)

	// --- 5. STRING → BOOL ---
	fmt.Println("\n=== ParseBool ===")

	roziStr := "true"
	rozi, _ := strconv.ParseBool(roziStr) // _ = xatoni e'tiborsiz qoldirish
	fmt.Println("Rozi:", rozi)

	// "1", "t", "T", "TRUE", "true", "True" → true
	// "0", "f", "F", "FALSE", "false", "False" → false
	// Boshqa narsalar → xato

	// --- 6. BOOL → STRING ---
	fmt.Println("\n=== FormatBool ===")
	javob := false
	fmt.Println("Javob string:", strconv.FormatBool(javob))

	// --- 7. AMALIY MISOL ---
	fmt.Println("\n=== AMALIY MISOL: yosh tekshirish ===")
	// Tasavvur qiling — HTTP form'dan kelgan ma'lumot
	formYosh := "17"
	parsedYosh, xato := strconv.Atoi(formYosh)
	if xato != nil {
		fmt.Println("Yoshni o'qib bo'lmadi:", xato)
	} else if parsedYosh < 18 {
		fmt.Println("Afsus, faqat 18+ uchun!")
	} else {
		fmt.Println("Marhamat, kirishingiz mumkin.")
	}

	// --- 8. ParseInt (kengaytirilgan Atoi) ---
	fmt.Println("\n=== ParseInt ===")
	// Atoi = ParseInt(s, 10, 0) ning qisqartmasi.
	// ParseInt sizga BASE (10, 2, 16) va BITSIZE (0, 8, 16, 32, 64) berishga ruxsat beradi.

	binar, _ := strconv.ParseInt("1010", 2, 64) // 2-lik sanoq → 10
	fmt.Println("Binar '1010' =", binar)

	hex, _ := strconv.ParseInt("FF", 16, 64) // 16-lik → 255
	fmt.Println("Hex 'FF' =", hex)
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + Foydalanuvchi inputini parse qilish (CLI, web form).
// + HTTP request parametrlari ("?yosh=25").
// + Konfiguratsiya fayli (env var: PORT="8080").
// + CSV/JSON parsing (qisman).
//
// QOIDA — ERROR'NI E'TIBORSIZ QOLDIRMANG:
//   Hech qachon: y, _ := strconv.Atoi(s)  — bu prod kodda yomon!
//   Doim: y, err := strconv.Atoi(s); if err != nil { ... }
// ==================================================
