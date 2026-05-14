// ==================================================
// MAVZU: switch operatori
// ==================================================
// Nazariya:
//   `switch` — bu if-else-if zanjirining QULAY shakli.
//   Bir o'zgaruvchini bir necha qiymatga taqqoslab, mos kelganini bajaradi.
//
//   Go switch'ning O'ZIGA XOS jihatlari:
//     1. AUTOMATIC break — har bir `case`dan keyin O'Z-O'ZIDAN to'xtaydi.
//        (C/Java'da har bir case'dan keyin `break` yozish kerak.)
//     2. `fallthrough` — agar SO'NGRA keyingi case'ni HAM bajarmoqchi
//        bo'lsangiz, aniq yozasiz (kamdan-kam ishlatiladi).
//     3. Multi-case — bir case ichida bir nechta qiymat:
//        `case 1, 2, 3:`
//     4. `switch` o'zgaruvchisiz ham ishlaydi — if-else-if zamonaviy shakli:
//        `switch { case x>0: ...; case x<0: ... }`
//     5. Type switch — interface tipini tekshirish (12-darsda batafsil).
//
// Real misol:
//   Buyurtma statusi: "yangi", "to'lov_kutilmoqda", "tayyorlanmoqda",
//   "yetkazilmoqda", "bekor_qilindi", "yetkazildi" — har biri uchun
//   alohida UI rangi va xabari.

package main

import "fmt"

func main() {
	// --- 1. ODDIY switch ---
	hafta := 3
	switch hafta {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6, 7: // MULTI-CASE — dam olish kunlari
		fmt.Println("Dam olish kuni")
	default:
		fmt.Println("Noto'g'ri kun raqami")
	}

	// --- 2. STRING ustida switch ---
	buyurtmaStatus := "tayyorlanmoqda"
	switch buyurtmaStatus {
	case "yangi":
		fmt.Println("To'lovni amalga oshiring.")
	case "to'lov_kutilmoqda":
		fmt.Println("To'lov tekshirilmoqda...")
	case "tayyorlanmoqda":
		fmt.Println("Buyurtmangiz tayyorlanmoqda.")
	case "yetkazilmoqda":
		fmt.Println("Kuryer yo'lda.")
	case "yetkazildi":
		fmt.Println("Sotib olganingiz uchun rahmat!")
	case "bekor_qilindi":
		fmt.Println("Buyurtma bekor qilindi.")
	default:
		fmt.Println("Noma'lum status:", buyurtmaStatus)
	}

	// --- 3. switch O'ZGARUVCHISIZ (if-else-if shakli) ---
	// Agar har bir case'da murakkab shart bo'lsa — bu shakl ishlatiladi.
	bal := 78
	switch {
	case bal >= 90:
		fmt.Println("Baho: A'lo (A)")
	case bal >= 75:
		fmt.Println("Baho: Yaxshi (B)")
	case bal >= 60:
		fmt.Println("Baho: Qoniqarli (C)")
	default:
		fmt.Println("Baho: Yetarli emas (F)")
	}

	// --- 4. fallthrough — keyingi case'ga "tushib" ketish ---
	// Kamdan-kam kerak bo'ladi. Asosan "include lower levels" mantiqida.
	huquq := "admin"
	fmt.Println("\nHuquqlar:")
	switch huquq {
	case "admin":
		fmt.Println("- delete (o'chirish)")
		fallthrough
	case "editor":
		fmt.Println("- edit (tahrirlash)")
		fallthrough
	case "viewer":
		fmt.Println("- read (o'qish)")
	}
	// "admin" bo'lsa: 3 ta qator chiqadi.
	// "editor" bo'lsa: 2 ta qator (edit + read).
	// "viewer" bo'lsa: 1 ta qator (read).

	// --- 5. ATM REAL MISOLI: menu tanlash ---
	fmt.Println("\n=== ATM menyu simulyatsiyasi ===")
	tanlov := 2 // foydalanuvchi 2-bandni tanladi

	switch tanlov {
	case 1:
		fmt.Println("→ Balansni ko'rsatish")
	case 2:
		fmt.Println("→ Pul yechish")
	case 3:
		fmt.Println("→ Pul qo'yish")
	case 4:
		fmt.Println("→ Pul o'tkazma")
	case 5:
		fmt.Println("→ Tarixni ko'rsatish")
	case 0:
		fmt.Println("→ Chiqish")
	default:
		fmt.Println("→ Noto'g'ri tanlov, qayta urinib ko'ring.")
	}

	// --- 6. KO'P SHARTLI switch — soliq stavkasi ---
	fmt.Println("\n=== Soliq stavkasi ===")
	oylik := 8_500_000.0 // 8.5 mln
	var stavka float64

	switch {
	case oylik <= 3_000_000:
		stavka = 0.07 // 7%
	case oylik <= 7_000_000:
		stavka = 0.12 // 12%
	case oylik <= 15_000_000:
		stavka = 0.18 // 18%
	default:
		stavka = 0.23 // 23%
	}
	soliq := oylik * stavka
	fmt.Printf("Oylik: %.0f, stavka: %.0f%%, soliq: %.0f\n", oylik, stavka*100, soliq)

	// --- 7. switch ICHIDA ETKAZMA (initialization) ---
	// switch x := f(); x  →  x faqat shu blokda mavjud
	switch kun := 15; {
	case kun < 10:
		fmt.Println("Oy boshi")
	case kun < 20:
		fmt.Println("Oy o'rtasi")
	default:
		fmt.Println("Oy oxiri")
	}
	// `kun` o'zgaruvchisi switch tashqarisida MAVJUD EMAS.
}

// ==================================================
// QACHON QO'LLANADI / QO'LLANMAYDI:
// ==================================================
// + 3+ ta diskret qiymatni tekshirish (status, rol, menu).
// + if-else-if zanjirini soddalashtirish.
// + Type switch (interface{} bilan ishlash, 12-darsda).
//
// QACHON IF YAXSHIROQ:
// - Faqat 1-2 ta shart bo'lsa.
// - Shart juda murakkab va o'z taglarini talab qilsa.
//
// ANTI-PATTERN'LAR:
// - Har bir case'da kod takrorlanishi → ortiq abstraksiya kerak.
// - fallthrough ortiqcha — agar barcha case'da ishlatilsa, oddiy `if` afzal.
// ==================================================
