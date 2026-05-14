// ==================================================
// LOYIHA: AQLLI ATM SIMULYATORI (Smart ATM)
// ==================================================
// Bu loyiha 1-4 darslarning XAMMASINI birlashtiradi:
//
//   1-DARS:
//     - O'zgaruvchilar (var, :=, const)
//     - Asosiy turlar (int, float64, string, bool)
//     - I/O (fmt.Print, Scanln)
//
//   2-DARS:
//     - Arifmetik + taqqoslash + mantiqiy operatorlar
//     - Tur konvertatsiyasi (float64 ↔ int)
//     - strconv.Atoi, strconv.ParseFloat
//     - bufio.Scanner — to'liq qatorni o'qish
//
//   3-DARS:
//     - if / else if / else (balans, limit, PIN tekshirish)
//     - switch (menu, operatsiya turi)
//     - if x := f(); x != nil { ... } — qisqa e'lon (error handling)
//
//   4-DARS:
//     - for klassik (tarix indeksi)
//     - for while-like (PIN urinishlar)
//     - for infinite (asosiy menu loop)
//     - break, continue
//     - range (tarix ustida iteratsiya)
//
// LOYIHA FUNKSIONALI:
//   1) PIN orqali kirish (max 3 urinish, blok)
//   2) Balansni ko'rish
//   3) Pul yechish (10 000 so'mga karrali, kunlik limit)
//   4) Pul qo'yish (depozit)
//   5) Boshqa kartaga o'tkazma (komissiya 1%, o'z bankka 0%)
//   6) Operatsiyalar tarixini ko'rish (oxirgi 5 ta)
//   7) Hisobotni ko'rish (umumiy kiritma, chiqim, soni)
//   8) Xavfsizlik: PIN o'zgartirish
//   0) Chiqish
//
// BIZNES QOIDALARI:
//   - PIN — 4 raqamli
//   - Boshlang'ich balans: 1 500 000 so'm
//   - Kunlik yechish limiti: 5 000 000 so'm
//   - Bir martalik yechish max: 2 000 000 so'm
//   - Pul yechish kupyura: 10 000 so'mga karrali
//   - Komissiya boshqa bankka: 1% (min 5 000 so'm)
//   - VIP shartlari: balans > 10 000 000 → komissiya yo'q
//   - 3 marta xato PIN → karta blok
//
// ESLATMA: Slice, map, struct hali o'rgatilmaganda, shuning uchun
// tarix uchun ALOHIDA o'zgaruvchilar ishlatamiz (oxirgi 5 ta).
// Bu loyiha 7-darsdan keyin slice bilan QISQARTIRILADI.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ----- KONFIGURATSIYA (const'lar) -----
const (
	togriPIN_default    = "1234"
	kunlikYechishLimit  = 5_000_000.0
	birMartalikLimit    = 2_000_000.0
	kupyuraKarrasi      = 10_000.0
	maxPINUrinish       = 3
	komissiyaFoiz       = 0.01    // 1%
	komissiyaMin        = 5_000.0 // min komissiya summa
	vipBalansChegarasi  = 10_000_000.0
	boshlangichBalans   = 1_500_000.0
	birlashtiruvchiChiziq = "=========================================="
)

func main() {
	// ----- DASTLABKI HOLAT -----
	scanner := bufio.NewScanner(os.Stdin)

	balans := boshlangichBalans
	bugunYechilgan := 0.0
	pin := togriPIN_default
	bloklandi := false

	// Statistika
	umumiyKiritma := 0.0
	umumiyChiqim := 0.0
	operatsiyaSoni := 0

	// Tarix — oxirgi 5 ta operatsiya (slice hali yo'q!)
	var tarix1, tarix2, tarix3, tarix4, tarix5 string

	// ----- 1) AUTORIZATSIYA — PIN tekshirish (while-loop) -----
	fmt.Println(birlashtiruvchiChiziq)
	fmt.Println("    AQLLI BANK — ATM SIMULYATORI v1.0")
	fmt.Println(birlashtiruvchiChiziq)

	urinish := 0
	for urinish < maxPINUrinish {
		fmt.Printf("\nPIN kodingizni kiriting (urinish %d/%d): ", urinish+1, maxPINUrinish)
		scanner.Scan()
		kiritilganPIN := strings.TrimSpace(scanner.Text())

		// Validatsiya: 4 raqam
		if len(kiritilganPIN) != 4 {
			fmt.Println("Xato: PIN aynan 4 ta raqamdan iborat bo'lishi kerak.")
			urinish++
			continue
		}

		// Faqat raqam ekanligini tekshirish (strconv.Atoi)
		if _, err := strconv.Atoi(kiritilganPIN); err != nil {
			fmt.Println("Xato: PIN faqat raqamlardan iborat bo'lishi kerak.")
			urinish++
			continue
		}

		if kiritilganPIN == pin {
			fmt.Println("\nMuvaffaqiyatli kirish! Xush kelibsiz.")
			break
		}

		urinish++
		qoldi := maxPINUrinish - urinish
		fmt.Printf("Noto'g'ri PIN. Qolgan urinishlar: %d\n", qoldi)
	}

	if urinish >= maxPINUrinish {
		bloklandi = true
		fmt.Println("\n" + birlashtiruvchiChiziq)
		fmt.Println("KARTA BLOKLANDI!")
		fmt.Println("3 marta xato PIN kiritildi. Bankga murojaat qiling.")
		fmt.Println(birlashtiruvchiChiziq)
		return
	}
	_ = bloklandi

	// ----- 2) ASOSIY MENU (cheksiz sikl) -----
	for {
		fmt.Println("\n" + birlashtiruvchiChiziq)
		fmt.Println("                 ASOSIY MENU")
		fmt.Println(birlashtiruvchiChiziq)
		fmt.Println("  1) Balansni ko'rish")
		fmt.Println("  2) Pul yechish")
		fmt.Println("  3) Pul qo'yish (depozit)")
		fmt.Println("  4) Boshqa kartaga o'tkazma")
		fmt.Println("  5) Operatsiyalar tarixi (oxirgi 5 ta)")
		fmt.Println("  6) Hisobotni ko'rish")
		fmt.Println("  7) PIN kodni o'zgartirish")
		fmt.Println("  0) Chiqish")
		fmt.Print("\nTanlovingiz: ")
		scanner.Scan()
		tanlov := strings.TrimSpace(scanner.Text())

		// Switch — har bir tanlov uchun alohida logika
		switch tanlov {

		// ============== 1) BALANS ==============
		case "1":
			fmt.Println("\n>>> BALANS <<<")
			fmt.Printf("Sizning balans: %.2f so'm\n", balans)
			if balans >= vipBalansChegarasi {
				fmt.Println("Status: VIP mijoz (komissiya yo'q)")
			} else {
				qoldi := vipBalansChegarasi - balans
				fmt.Printf("VIP gacha qoldi: %.2f so'm\n", qoldi)
			}

		// ============== 2) PUL YECHISH ==============
		case "2":
			fmt.Println("\n>>> PUL YECHISH <<<")
			fmt.Printf("Joriy balans: %.2f so'm\n", balans)
			fmt.Printf("Bugun yechilgan: %.2f so'm (limit: %.0f)\n", bugunYechilgan, kunlikYechishLimit)
			fmt.Print("Qancha yechmoqchisiz? ")
			scanner.Scan()
			summa, xato := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

			// Bir necha bosqichli validatsiya — if-else-if zanjiri
			if xato != nil {
				fmt.Println("Xato: faqat son kiriting.")
			} else if summa <= 0 {
				fmt.Println("Xato: summa musbat bo'lishi kerak.")
			} else if summa > birMartalikLimit {
				fmt.Printf("Xato: bir martalik max %.0f so'm.\n", birMartalikLimit)
			} else if summa > balans {
				fmt.Printf("Xato: balansda yetarli pul yo'q (sizda %.2f).\n", balans)
			} else if bugunYechilgan+summa > kunlikYechishLimit {
				qoldiLimit := kunlikYechishLimit - bugunYechilgan
				fmt.Printf("Xato: kunlik limit oshib ketadi. Bugun yechish mumkin: %.0f\n", qoldiLimit)
			} else if int(summa)%int(kupyuraKarrasi) != 0 {
				fmt.Printf("Xato: faqat %.0f so'mga karrali summalar.\n", kupyuraKarrasi)
			} else {
				// Hammasi OK — pul beriladi
				balans -= summa
				bugunYechilgan += summa
				umumiyChiqim += summa
				operatsiyaSoni++
				yangiYozuv := fmt.Sprintf("YECHILDI: -%.0f so'm", summa)
				tarix1, tarix2, tarix3, tarix4, tarix5 = tarixQoshish(yangiYozuv, tarix1, tarix2, tarix3, tarix4)
				fmt.Printf("Muvaffaqiyat! Yechildi: %.0f so'm. Yangi balans: %.2f\n", summa, balans)
			}

		// ============== 3) PUL QO'YISH (DEPOZIT) ==============
		case "3":
			fmt.Println("\n>>> PUL QO'YISH <<<")
			fmt.Print("Qancha qo'ymoqchisiz? ")
			scanner.Scan()
			summa, xato := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

			if xato != nil {
				fmt.Println("Xato: faqat son kiriting.")
			} else if summa <= 0 {
				fmt.Println("Xato: summa musbat bo'lishi kerak.")
			} else if summa > 50_000_000 {
				// AML compliance — 50 mln dan ko'p bo'lsa, qo'shimcha tekshirish kerak
				fmt.Println("Xato: 50 mln dan ko'p — bank ofisiga murojaat qiling.")
			} else {
				balans += summa
				umumiyKiritma += summa
				operatsiyaSoni++
				yangiYozuv := fmt.Sprintf("QO'YILDI: +%.0f so'm", summa)
				tarix1, tarix2, tarix3, tarix4, tarix5 = tarixQoshish(yangiYozuv, tarix1, tarix2, tarix3, tarix4)
				fmt.Printf("Muvaffaqiyat! Qo'yildi: %.0f. Yangi balans: %.2f\n", summa, balans)
			}

		// ============== 4) O'TKAZMA ==============
		case "4":
			fmt.Println("\n>>> O'TKAZMA <<<")
			fmt.Print("Qabul qiluvchi karta raqami (16 raqam): ")
			scanner.Scan()
			karta := strings.TrimSpace(scanner.Text())

			// Karta uzunligi va raqam ekanligi tekshiruvi
			kartaXatoMi := false
			if len(karta) != 16 {
				fmt.Println("Xato: karta raqami 16 ta raqamdan iborat bo'lishi kerak.")
				kartaXatoMi = true
			} else if _, err := strconv.Atoi(karta); err != nil {
				fmt.Println("Xato: karta raqami faqat raqamlardan iborat bo'lishi kerak.")
				kartaXatoMi = true
			}

			if kartaXatoMi {
				break // switch case'dan chiqamiz (asosiy for davom etadi)
			}

			fmt.Print("Qaysi bankka? (own/other): ")
			scanner.Scan()
			bank := strings.ToLower(strings.TrimSpace(scanner.Text()))

			fmt.Print("Summa: ")
			scanner.Scan()
			summa, xato := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

			if xato != nil || summa <= 0 {
				fmt.Println("Xato: noto'g'ri summa.")
				break
			}

			// Komissiya hisoblash — switch va if birgalikda
			var komissiya float64
			bankXato := false
			switch bank {
			case "own":
				komissiya = 0
			case "other":
				if balans >= vipBalansChegarasi {
					// VIP — komissiya yo'q
					komissiya = 0
					fmt.Println("VIP imtiyozi: komissiya yo'q!")
				} else {
					komissiya = summa * komissiyaFoiz
					// Min komissiya cheklash
					if komissiya < komissiyaMin {
						komissiya = komissiyaMin
					}
				}
			default:
				fmt.Println("Xato: bank 'own' yoki 'other' bo'lishi kerak.")
				bankXato = true
			}

			if bankXato {
				// Asosiy menu sikliga qaytamiz
				continue
			}

			umumiy := summa + komissiya
			if umumiy > balans {
				fmt.Printf("Xato: balansda yetarli emas. Kerak: %.2f, bor: %.2f\n", umumiy, balans)
			} else {
				balans -= umumiy
				umumiyChiqim += umumiy
				operatsiyaSoni++
				oxirgi4 := karta[12:] // karta oxirgi 4 raqami
				yangiYozuv := fmt.Sprintf("O'TKAZMA: -%.0f (***%s, komissiya: %.0f)", summa, oxirgi4, komissiya)
				tarix1, tarix2, tarix3, tarix4, tarix5 = tarixQoshish(yangiYozuv, tarix1, tarix2, tarix3, tarix4)
				fmt.Printf("Muvaffaqiyat! O'tkazildi: %.0f, komissiya: %.0f. Yangi balans: %.2f\n",
					summa, komissiya, balans)
			}

		// ============== 5) TARIX ==============
		case "5":
			fmt.Println("\n>>> OXIRGI 5 TA OPERATSIYA <<<")
			// for-range tushunchasini ko'rsatish uchun array ishlataylik
			tarixArr := [5]string{tarix1, tarix2, tarix3, tarix4, tarix5}
			bushTarix := true
			for i, t := range tarixArr {
				if t == "" {
					continue // bo'sh slot — o'tkazib yuborish
				}
				bushTarix = false
				fmt.Printf("  %d) %s\n", i+1, t)
			}
			if bushTarix {
				fmt.Println("  (tarix bo'sh)")
			}

		// ============== 6) HISOBOT ==============
		case "6":
			fmt.Println("\n>>> SESSIYA HISOBOTI <<<")
			fmt.Println(birlashtiruvchiChiziq)
			fmt.Printf("Joriy balans:        %.2f so'm\n", balans)
			fmt.Printf("Umumiy kiritma:      %.2f so'm\n", umumiyKiritma)
			fmt.Printf("Umumiy chiqim:       %.2f so'm\n", umumiyChiqim)
			fmt.Printf("Operatsiyalar soni:  %d ta\n", operatsiyaSoni)
			fmt.Printf("Bugun yechilgan:     %.2f so'm\n", bugunYechilgan)
			qoldiLimit := kunlikYechishLimit - bugunYechilgan
			fmt.Printf("Kunlik limitda qoldi: %.2f so'm\n", qoldiLimit)

			// Sof o'zgarish (kiritma - chiqim)
			sof := umumiyKiritma - umumiyChiqim
			if sof > 0 {
				fmt.Printf("Sof o'zgarish:       +%.2f so'm (ko'paydi)\n", sof)
			} else if sof < 0 {
				fmt.Printf("Sof o'zgarish:       %.2f so'm (kamaydi)\n", sof)
			} else {
				fmt.Println("Sof o'zgarish:       0 (o'zgarmadi)")
			}
			fmt.Println(birlashtiruvchiChiziq)

		// ============== 7) PIN O'ZGARTIRISH ==============
		case "7":
			fmt.Println("\n>>> PIN O'ZGARTIRISH <<<")
			fmt.Print("Joriy PIN: ")
			scanner.Scan()
			joriy := strings.TrimSpace(scanner.Text())
			if joriy != pin {
				fmt.Println("Xato: joriy PIN noto'g'ri.")
				break
			}
			fmt.Print("Yangi PIN (4 raqam): ")
			scanner.Scan()
			yangi := strings.TrimSpace(scanner.Text())

			// Yangi PIN validatsiyasi
			if len(yangi) != 4 {
				fmt.Println("Xato: PIN aynan 4 ta raqam.")
				break
			}
			if _, err := strconv.Atoi(yangi); err != nil {
				fmt.Println("Xato: PIN faqat raqam.")
				break
			}
			if yangi == pin {
				fmt.Println("Xato: yangi PIN eski PIN'dan farq qilishi kerak.")
				break
			}
			// Oddiy PIN'lardan saqlanish (simple check)
			if yangi == "0000" || yangi == "1111" || yangi == "1234" || yangi == "9999" {
				fmt.Println("Xato: bunday PIN juda oson, boshqasini tanlang.")
				break
			}
			fmt.Print("Tasdiqlash uchun yana kiriting: ")
			scanner.Scan()
			tasdiq := strings.TrimSpace(scanner.Text())
			if tasdiq != yangi {
				fmt.Println("Xato: tasdiq mos kelmadi.")
				break
			}
			pin = yangi
			fmt.Println("PIN muvaffaqiyatli o'zgartirildi!")

		// ============== 0) CHIQISH ==============
		case "0":
			fmt.Println("\n" + birlashtiruvchiChiziq)
			fmt.Println("Bankimizdan foydalanganingiz uchun rahmat!")
			fmt.Printf("Yakuniy balans: %.2f so'm\n", balans)
			fmt.Printf("Operatsiyalar: %d ta\n", operatsiyaSoni)
			fmt.Println(birlashtiruvchiChiziq)
			return

		default:
			fmt.Println("Noto'g'ri tanlov. 0-7 oralig'idagi raqam kiriting.")
		}
	}
}

// ==================================================
// YORDAMCHI: tarixga yangi yozuv qo'shish
// ==================================================
// Funksiya 5-darsda batafsil. Hozircha — oddiy "queue" pattern.
// Eski 5-element olib tashlanadi, qolganlari siljiydi, yangisi 1-o'ringa.
//
// Boshqacha aytganda: tarix1 = yangi, tarix2 = eski tarix1, ...
func tarixQoshish(yangi, t1, t2, t3, t4 string) (string, string, string, string, string) {
	return yangi, t1, t2, t3, t4
}

// ==================================================
// LOYIHA HAQIDA UMUMIY:
// ==================================================
// + 1-darsdan o'rganganlar: const, var, :=, fmt.Print, fmt.Scan
// + 2-darsdan: %, /, &&, ||, !, strconv.Atoi/ParseFloat, bufio.Scanner
// + 3-darsdan: if/else, switch, qisqa e'lon (if err := ...; ...)
// + 4-darsdan: for klassik, for-while, for-infinite, break, continue, range
//
// KEYINGI DARS'LARDA YAXSHILANADIGAN JOYLAR:
// - 5-dars (funksiya): har bir case'ni alohida funksiyaga ajratish
// - 7-dars (slice): tarix slice ([]string) bilan, len cheklov
// - 10-dars (struct): `type Account struct { Balans float64; ...}`
// - 14-dars (error): xatolar to'liq qaytarish
// - 18-dars (fayl): balansni faylga saqlash (JSON)
// ==================================================
