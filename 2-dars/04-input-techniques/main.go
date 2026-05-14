// ==================================================
// MAVZU: Input olishning 4 xil usuli
// ==================================================
// Nazariya:
//   Foydalanuvchidan ma'lumot olishning bir necha usuli bor.
//   Har birining O'Z foydasi va kamchiligi:
//
//   1) fmt.Scan(&x)     — bo'sh joy yoki yangi qatorgacha o'qiydi.
//                          Bir nechta o'zgaruvchini bir qatorda olishga qulay.
//
//   2) fmt.Scanln(&x)   — Scan'ga o'xshash, lekin YANGI QATORDA TO'XTAYDI.
//                          Faqat bitta qatordan o'qish kerak bo'lsa.
//
//   3) fmt.Scanf("%d %s", &x, &y) — format bilan.
//                          C dasturlash kabi, lekin Go'da kam ishlatiladi.
//
//   4) bufio.Scanner    — eng ISHONCHLI usul. Butun qatorni
//                          (bo'sh joylari bilan birga) o'qiydi.
//                          Production'da AYNAN shu ishlatiladi.
//
// Muammo: fmt.Scan "Ali Ahmad" deb yozilsa, "Ali" oladi va "Ahmad" qoladi.
//         To'liq ismni olish uchun bufio.Scanner kerak.
//
// Real misol:
//   Bank ATM ilovasi: foydalanuvchi to'la ismini, PIN kodini va
//   pul miqdorini kiritadi. Hammasi alohida qatorda.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// --- USUL 1: bufio.Scanner (TAVSIYA QILINADI) ---
	// Avval ushbu usulni o'rganamiz, chunki u eng kuchli va xavfsiz.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("To'liq ismingizni kiriting: ")
	scanner.Scan() // Yangi qatorni bossanizgacha kutadi
	toliqIsm := scanner.Text() // Bo'sh joylari bilan butun qatorni oladi
	// strings.TrimSpace — boshi/oxiridagi probelni olib tashlash
	toliqIsm = strings.TrimSpace(toliqIsm)
	fmt.Println("Salom,", toliqIsm)

	// Yoshni olish va songa o'tkazish
	fmt.Print("Yoshingizni kiriting: ")
	scanner.Scan()
	yoshStr := strings.TrimSpace(scanner.Text())
	yosh, xato := strconv.Atoi(yoshStr)
	if xato != nil {
		fmt.Println("Yosh — son bo'lishi kerak!")
		return // Dasturdan chiqib ketamiz
	}

	// Pul miqdori (float64)
	fmt.Print("Pul miqdorini kiriting (so'm): ")
	scanner.Scan()
	pulStr := strings.TrimSpace(scanner.Text())
	pul, xato := strconv.ParseFloat(pulStr, 64)
	if xato != nil {
		fmt.Println("Pul — son bo'lishi kerak!")
		return
	}

	fmt.Println("\n=== Sizning ma'lumotlaringiz ===")
	fmt.Printf("Ism: %s\nYosh: %d\nBalans: %.2f so'm\n", toliqIsm, yosh, pul)

	// 12% soliq
	soliq := pul * 0.12
	sof := pul - soliq
	fmt.Printf("12%% soliq: %.2f so'm\n", soliq)
	fmt.Printf("Sof: %.2f so'm\n", sof)

	// --- USUL 2: fmt.Scan (BIR NECHTA QIYMAT BIR QATORDA) ---
	// Pastdagi kod izoh ostida — uni ko'rib chiqing.
	// Agar bir necha sonni bir qatorda olishni xohlasangiz, qulay:
	/*
		var a, b int
		fmt.Print("Ikkita son kiriting (bo'sh joy bilan): ")
		fmt.Scan(&a, &b)
		fmt.Println("Yig'indi:", a+b)
	*/

	// --- USUL 3: fmt.Scanln ---
	/*
		var x int
		fmt.Print("Bitta son: ")
		fmt.Scanln(&x)
		fmt.Println(x)
	*/

	// --- USUL 4: fmt.Scanf — kam ishlatiladi ---
	/*
		var sana string
		var temp float64
		fmt.Print("Sana va harorat (2024-01-15 -5.5): ")
		fmt.Scanf("%s %f", &sana, &temp)
		fmt.Println(sana, temp)
	*/

	// --- BUFIO BILAN VALIDATSIYA (CYCLE — keyingi darsda) ---
	// Hozircha — agar foydalanuvchi xato kiritsa, dastur to'xtaydi.
	// 3-darsda if/else bilan retry qilamiz.
	// 4-darsda esa for sikli bilan to'g'ri kiritmaguncha qayta so'raymiz.
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + bufio.Scanner — interaktiv CLI dasturlar, foydalanuvchi inputi.
// + fmt.Scan — testlar, oddiy "olympiad" tipidagi masalalar.
// + fmt.Scanf — kamdan-kam, faqat aniq formatdagi ma'lumot uchun.
//
// QACHON QO'LLANMAYDI:
// - Web/HTTP'da fmt.Scan ishlatilmaydi (request body bo'ladi).
// - JSON/CSV'da Scan'lar emas, maxsus parser'lar.
//
// XAVFSIZLIK:
// - Foydalanuvchi inputiga DOIM ishonchsizlik bilan qarang.
// - Doim validatsiya qiling (strconv xatosini ushlash, uzunlik tekshirish).
// ==================================================
