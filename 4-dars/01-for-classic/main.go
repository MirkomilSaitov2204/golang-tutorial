// ==================================================
// MAVZU: for klassik shakli
// ==================================================
// Nazariya:
//   Go'da FAQAT bitta sikl operatori bor — `for`.
//   `while`, `do-while` YO'Q. Lekin `for` 3 xil shaklga ega:
//
//     1) KLASSIK (C-style):
//        for init; shart; post {
//            // tana
//        }
//        Misol: for i := 0; i < 10; i++ { ... }
//
//     2) FAQAT SHART (while-like):
//        for shart {
//            // tana
//        }
//
//     3) INFINITE (cheksiz, break bilan to'xtatiladi):
//        for {
//            // tana
//        }
//
//   Bu darsda KLASSIK shaklini chuqur o'rganamiz.
//
// Klassik shakli komponentlari:
//   - init: faqat BIR MARTA, siklning boshida bajariladi.
//   - shart: HAR ITERATSIYA boshida tekshiriladi.
//   - post: HAR ITERATSIYA oxirida bajariladi.
//
//   i := 0  →  init
//   i < 10  →  shart (true bo'lsa, davom etadi)
//   i++     →  post (har iteratsiyadan keyin oshiradi)
//
// Real misol:
//   Loyihaning DB jadvalidagi 100 ta foydalanuvchini ko'rib chiqish:
//     for i := 0; i < 100; i++ { processUser(users[i]) }

package main

import "fmt"

func main() {
	// --- 1. ENG ODDIY for ---
	fmt.Println("=== 1 dan 5 gacha ===")
	for i := 1; i <= 5; i++ {
		fmt.Println("Qiymat:", i)
	}

	// --- 2. TESKARI YO'NALISH ---
	fmt.Println("\n=== 10 dan 1 gacha (teskari) ===")
	for i := 10; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- 3. QADAM BILAN (step) ---
	fmt.Println("\n=== Juft sonlar 2 dan 20 gacha ===")
	for i := 2; i <= 20; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// --- 4. YIG'INDI HISOBLASH ---
	fmt.Println("\n=== 1 dan 100 gacha sonlar yig'indisi ===")
	yigindi := 0
	for i := 1; i <= 100; i++ {
		yigindi += i
	}
	fmt.Println("Yig'indi:", yigindi) // 5050

	// --- 5. FAKTORIAL ---
	fmt.Println("\n=== 6! (faktorial) ===")
	n := 6
	faktorial := 1
	for i := 1; i <= n; i++ {
		faktorial *= i
	}
	fmt.Printf("%d! = %d\n", n, faktorial) // 720

	// --- 6. KO'PAYTUVCHILAR JADVALI (multiplication table) ---
	fmt.Println("\n=== 7 ning ko'paytirish jadvali ===")
	son := 7
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d × %d = %d\n", son, i, son*i)
	}

	// --- 7. NESTED FOR (ichma-ich) — yulduzcha figura ---
	fmt.Println("\n=== Yulduzcha uchburchak ===")
	for satr := 1; satr <= 5; satr++ {
		for ustun := 1; ustun <= satr; ustun++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	// *
	// * *
	// * * *
	// * * * *
	// * * * * *

	// --- 8. KO'PAYTIRISH JADVALI (to'liq, 1 dan 9 gacha) ---
	fmt.Println("\n=== Pifagor jadvali ===")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}

	// --- 9. AMALIY: O'rtacha hisobi ---
	fmt.Println("\n=== O'rtacha hisobi (10 ta talaba) ===")
	yigindiBal := 0
	talabaSoni := 10
	// Hozircha hard-code, slice'ni 7-darsda o'rganamiz
	ballar := [10]int{85, 92, 78, 65, 88, 91, 72, 84, 79, 95}
	for i := 0; i < talabaSoni; i++ {
		yigindiBal += ballar[i]
	}
	ortacha := float64(yigindiBal) / float64(talabaSoni)
	fmt.Printf("O'rtacha bal: %.2f\n", ortacha)

	// --- 10. INDEKS BILAN: maximum topish ---
	fmt.Println("\n=== Maksimum bal ===")
	maxBal := ballar[0]
	maxIndex := 0
	for i := 1; i < talabaSoni; i++ {
		if ballar[i] > maxBal {
			maxBal = ballar[i]
			maxIndex = i
		}
	}
	fmt.Printf("Eng yuqori bal: %d (talaba #%d)\n", maxBal, maxIndex+1)
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + Aniq sondagi takrorlash (10 marta, N marta).
// + Counter bilan ishlash (index orqali element olish).
// + Akkumulator (yig'indi, ko'paytma, min/max).
//
// QACHON BOSHQA SHAKL AFZAL:
// - Slice/map ustida iteratsiya → `range` (03-range mavzusida).
// - Cheksiz tinglash (server, scanner) → infinite for { }.
// - Ehtiyot bo'ling: i++ EMAS, ++i ham EMAS. Go'da `++` faqat
//   alohida operator, ifoda emas. `j = i++` — XATO.
// ==================================================
