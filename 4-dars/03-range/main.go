// ==================================================
// MAVZU: range — kolleksiya ustida iteratsiya
// ==================================================
// Nazariya:
//   `range` — kolleksiya (slice, array, map, string, channel)
//   ustida QULAY iteratsiya qilish uchun.
//
//   Sintaksis (slice/array uchun):
//     for indeks, qiymat := range kolleksiya {
//         // ...
//     }
//
//   Variantlar:
//     1) for i, v := range a   — ikkala
//     2) for i := range a      — faqat indeks
//     3) for _, v := range a   — faqat qiymat (indeks kerak emas)
//     4) for range a           — faqat takrorlash (n marta), Go 1.22+
//
//   Turi bo'yicha range nima qaytaradi:
//     - SLICE/ARRAY:  i (int), v (element)
//     - MAP:          k (kalit), v (qiymat) — tartibi RANDOM
//     - STRING:       i (bayt indeksi), r (rune — Unicode kod)
//     - CHANNEL:      v (qiymat) — channel yopilguncha
//
// Real misol:
//   E-commerce: barcha buyurtmalarni ko'rib chiqib,
//   "yangi" statusdagilarning umumiy summasini hisoblash.

package main

import "fmt"

func main() {
	// --- 1. SLICE USTIDA range ---
	mevalar := []string{"Olma", "Banan", "Anor", "Anjir", "Behi"}

	fmt.Println("=== Indeks va qiymat ===")
	for i, m := range mevalar {
		fmt.Printf("[%d] %s\n", i, m)
	}

	fmt.Println("\n=== Faqat qiymat ===")
	for _, m := range mevalar {
		fmt.Println("Meva:", m)
	}

	fmt.Println("\n=== Faqat indeks ===")
	for i := range mevalar {
		fmt.Println("Indeks:", i)
	}

	// --- 2. SONLAR slice ustida ---
	narxlar := []int{15000, 22000, 8500, 30000, 12000}
	jami := 0
	for _, narx := range narxlar {
		jami += narx
	}
	fmt.Println("\nUmumiy summa:", jami) // 87500

	// --- 3. STRING ustida range — BAYTLAR EMAS, RUNE'LAR! ---
	fmt.Println("\n=== O'zbekcha matn ustida ===")
	matn := "Salom, dunyo!"
	for i, r := range matn {
		fmt.Printf("Bayt indeksi %d: %c (kod %d)\n", i, r, r)
	}

	// Diqqat — Kirillcha yoki o'zbekcha "ё" UTF-8'da 2 bayt egallaydi.
	// `i` har bayt sayin emas, har RUNE'ning BIRINCHI BAYT'i.
	fmt.Println("\n=== Kirillcha matn ===")
	kirill := "ЗGo"
	for i, r := range kirill {
		fmt.Printf("i=%d, rune=%c (kod %d)\n", i, r, r)
	}
	// i=0, rune=З (2 baytni egallaydi)
	// i=2, rune=G
	// i=3, rune=o

	// --- 4. MAP ustida range — tartib RANDOM! ---
	fmt.Println("\n=== Map ustida ===")
	yoshlar := map[string]int{
		"Ali":   25,
		"Vali":  30,
		"Gani":  19,
		"Hasan": 45,
	}
	for ism, yosh := range yoshlar {
		fmt.Printf("%s: %d yosh\n", ism, yosh)
	}
	// Har gal ishga tushirilganda tartib BOSHQACHA bo'lishi mumkin.
	// Agar tartib kerak bo'lsa — kalitlarni alohida sort qilish kerak (22-dars).

	// --- 5. range FAQAT TAKRORLASH (Go 1.22+) ---
	fmt.Println("\n=== 3 marta salom ===")
	for range 3 {
		fmt.Println("Salom!")
	}
	// Eski versiyada: for i := 0; i < 3; i++ { ... }

	// --- 6. AMALIY: filter va count ---
	fmt.Println("\n=== Filter: 15000+ narxda nechta? ===")
	qimmatlar := 0
	for _, narx := range narxlar {
		if narx >= 15000 {
			qimmatlar++
		}
	}
	fmt.Println("Qimmat mahsulotlar soni:", qimmatlar)

	// --- 7. AMALIY: maksimum + indeksini topish ---
	fmt.Println("\n=== Eng qimmat mahsulot ===")
	maxIdx := 0
	for i, narx := range narxlar {
		if narx > narxlar[maxIdx] {
			maxIdx = i
		}
	}
	fmt.Printf("Eng qimmat: %d so'm (indeks %d)\n", narxlar[maxIdx], maxIdx)

	// --- 8. NESTED range — buyurtmalar ro'yxati ---
	fmt.Println("\n=== Buyurtmalar tarixi ===")
	tarix := []string{
		"2026-01-15: Olma 5kg",
		"2026-01-16: Banan 2kg",
		"2026-01-18: Anjir 1kg",
	}
	for i, b := range tarix {
		fmt.Printf("#%d → %s\n", i+1, b)
	}

	// --- 9. range BILAN PASTGA SANASH ---
	// range to'g'ridan-to'g'ri teskari yo'nalishni QO'LLAB-QUVVATLAMAYDI.
	// Buni qilish uchun klassik for ishlatiladi:
	fmt.Println("\n=== Teskari yo'nalish (klassik for kerak) ===")
	for i := len(mevalar) - 1; i >= 0; i-- {
		fmt.Printf("[%d] %s\n", i, mevalar[i])
	}
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + Slice/array ustida iteratsiya — DEFOLT shakl.
// + Map ustida (tartib muhim bo'lmasa).
// + String ustida — Unicode-safe iteratsiya.
// + Channel ustida (27-darsda batafsil).
//
// QACHON KLASSIK for AFZAL:
// - Teskari yo'nalish (i--).
// - Bir nechta qadam (i += 2).
// - Indekslarni boshqarish (skip, jump).
//
// MUHIM:
// - range MAP'da DOIM tartib boshqacha. Bunga ishonmang!
// - range STRING'da i — RUNE'ning birinchi bayti, RUNE indeksi EMAS.
// - range SLICE'da `v` — NUSXA. Asl elementni o'zgartirish uchun
//   indeks orqali: `a[i] = yangi`, yoki `&a[i]` (10-darsda).
// ==================================================
