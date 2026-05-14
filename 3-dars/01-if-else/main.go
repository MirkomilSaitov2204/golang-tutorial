// ==================================================
// MAVZU: if / else if / else
// ==================================================
// Nazariya:
//   Shartli operator — bu dasturning "qaror qabul qilish" mexanizmi.
//   Agar shart TRUE bo'lsa, kod blokini bajaradi; bo'lmasa — boshqa yo'l.
//
//   Sintaksis:
//     if shart {
//         // bajariladi (shart true bo'lsa)
//     } else if boshqaShart {
//         // bajariladi (birinchi false, lekin bu true bo'lsa)
//     } else {
//         // qolgan barcha holatlar
//     }
//
//   Go'ning O'ZIGA XOS xususiyatlari:
//     1. Shart atrofida QAVS KERAK EMAS:  if x > 0  (NA  `if (x > 0)`).
//     2. Figurali qavs SHART:  `{` AYNAN shu qatorda turishi kerak.
//     3. Tek qatorli if BO'LMAYDI:  `if x>0 doSomething()` — XATO.
//     4. `else` boshqa qatorga tushmaydi: `} else {` shaklida.
//     5. Shart faqat BOOL bo'lishi kerak. `if 1 { }` — XATO (C/Python emas).
//
// Real misol:
//   Bank ATM: foydalanuvchi pul yechmoqchi. Tekshirishlar:
//   - PIN to'g'rimi?
//   - Balansda yetarli pulmi?
//   - Kunlik limit oshmaganmi?
//   - Bankomat ichida shu summa nominalida pul bormi?
//
// Tartib muhim:
//   Eng aniq (specific) shartdan eng umumiyga qarab yozing.
//   Aks holda — keyingilari hech qachon bajarilmaydi.

package main

import "fmt"

func main() {
	// --- 1. SODDA if ---
	yosh := 19
	if yosh >= 18 {
		fmt.Println("Voyaga yetgan.")
	}

	// --- 2. if / else ---
	balans := 75000
	if balans >= 100000 {
		fmt.Println("VIP mijoz.")
	} else {
		fmt.Println("Oddiy mijoz.")
	}

	// --- 3. if / else if / else (kategoriya) ---
	bal := 85
	if bal >= 90 {
		fmt.Println("Baho: A'lo")
	} else if bal >= 75 {
		fmt.Println("Baho: Yaxshi")
	} else if bal >= 60 {
		fmt.Println("Baho: Qoniqarli")
	} else {
		fmt.Println("Baho: Yetarli emas")
	}

	// --- 4. MURAKKAB SHARTLAR (mantiqiy operatorlar) ---
	vip := true
	hisob := 600000
	yoshMijoz := 30

	// VIP bo'lsin VA balansi yetarli VA yoshi 25+ bo'lsin → kredit beriladi
	if vip && hisob >= 500000 && yoshMijoz >= 25 {
		fmt.Println("Kreditga loyiqsiz!")
	} else {
		fmt.Println("Kredit shartlariga to'g'ri kelmaysiz.")
	}

	// --- 5. ATM REAL MISOLI: pul yechish ---
	fmt.Println("\n=== ATM pul yechish simulyatsiyasi ===")

	currentBalans := 500000
	yechmoqchi := 300000
	kunlikLimit := 1000000
	bugunYechilgan := 800000

	if yechmoqchi <= 0 {
		fmt.Println("Xato: summa musbat bo'lishi kerak.")
	} else if yechmoqchi > currentBalans {
		fmt.Println("Xato: balansda yetarli pul yo'q.")
	} else if bugunYechilgan+yechmoqchi > kunlikLimit {
		fmt.Println("Xato: kunlik limit oshib ketadi.")
	} else if yechmoqchi%10000 != 0 {
		// Bankomatda 10ming so'mli kupyuralar
		fmt.Println("Xato: faqat 10 000 so'mga karrali summa.")
	} else {
		yangiBalans := currentBalans - yechmoqchi
		fmt.Printf("Muvaffaqiyat. Yangi balans: %d so'm\n", yangiBalans)
	}

	// --- 6. NESTED if (ichma-ich) — KAMROQ ISHLATING ---
	fmt.Println("\n=== Nested if (saqlanish) ===")
	tarmoqUlanganmi := true
	tizimIshlamoqda := true
	autorizatsiya := false

	if tarmoqUlanganmi {
		if tizimIshlamoqda {
			if autorizatsiya {
				fmt.Println("Hammasi yaxshi — operatsiyaga ruxsat.")
			} else {
				fmt.Println("Avval login qiling.")
			}
		} else {
			fmt.Println("Tizim ta'mirda.")
		}
	} else {
		fmt.Println("Internet yo'q.")
	}

	// YAXSHIROQ — "early return" yoki && bilan birlashtirish:
	//   if !tarmoqUlanganmi { ... return }
	//   if !tizimIshlamoqda { ... return }
	//   if !autorizatsiya { ... return }
	//   // asosiy logika
	// Bu — "guard clauses" pattern. Kodni o'qishni yengillashtiradi.

	// --- 7. STRING SHARTLARI ---
	rol := "admin"
	if rol == "admin" {
		fmt.Println("\nBarcha huquqlar mavjud.")
	} else if rol == "moderator" {
		fmt.Println("\nQisman huquqlar.")
	} else {
		fmt.Println("\nFaqat o'qish.")
	}
}

// ==================================================
// QACHON QO'LLANADI / QO'LLANMAYDI:
// ==================================================
// + Bir nechta turli mantiqiy shartni tekshirish (range, mantiqiy AND/OR).
// + Validatsiya (form, API request).
// + Business rule'lar (kim, qachon, qaysi sharoitda).
//
// QACHON SWITCH AFZAL:
// - 4+ ta diskret qiymatni (rol, status, mavsumi) tekshirsangiz → switch.
// - if-else-if zanjiri 5+ qator → switch o'qilishi yaxshi.
//
// ANTI-PATTERN'LAR:
// - Chuqur nested if (3+ daraja) — guard clause'larga aylantiring.
// - Bir xil kodni har bir tarmoqda takrorlash — refactor qiling.
// ==================================================
