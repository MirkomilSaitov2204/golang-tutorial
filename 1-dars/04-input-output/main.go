// =============================================================================
// MAVZU 4: INPUT / OUTPUT — fmt PAKETI VA MINI-KALKULYATOR
// =============================================================================
//
// REAL HAYOTDAN MISOL:
//   Sizdan mini-kalkulyator yozish so'ralgan. Foydalanuvchi:
//     1. Birinchi sonni kiritadi
//     2. Operatsiyani tanlaydi (+, -, *, /)
//     3. Ikkinchi sonni kiritadi
//   Dastur natijani chiroyli ko'rinishda chiqaradi.
//   Mana shu — bizning bugungi 1-darsimizning yakuniy mahsuloti!
//
// NAZARIYA — `fmt` PAKETI:
//
//   CHIQARISH FUNKSIYALARI:
//     - fmt.Print     -> qator oxiriga yangi qator QO'YMAYDI
//     - fmt.Println   -> oxiriga "\n" QO'YADI, argumentlar orasiga PROBEL
//     - fmt.Printf    -> formatlangan output, "verb"lar bilan
//
//   FORMATLASH VERB'lari (eng muhimlari):
//     %v   -> qiymatning standart ko'rinishi (har qanday tur uchun ishlaydi)
//     %+v  -> struct uchun maydon nomlari bilan
//     %#v  -> Go-sintaksisidagi ko'rinishi
//     %T   -> qiymatning TURI
//     %d   -> butun son (decimal)
//     %b   -> ikkilik (binary)
//     %o   -> sakkizlik (octal)
//     %x   -> o'n oltilik (hex)
//     %f   -> kasr son (default 6 raqamli)
//     %.2f -> kasr son, 2 ta kasr raqami
//     %e   -> ilmiy notatsiya (1.5e+09)
//     %s   -> string
//     %q   -> qo'shtirnoq ichidagi string
//     %t   -> bool (true/false)
//     %c   -> rune (belgi)
//     %p   -> pointer
//
//   INPUT (kiritish) FUNKSIYALARI:
//     - fmt.Scan      -> probel/yangi qator bilan ajratilgan qiymatlar
//     - fmt.Scanln    -> yangi qatorgacha o'qiydi
//     - fmt.Scanf     -> formatlangan input
//     - bufio.Scanner -> bir qatorni butun olish uchun (yaxshiroq)
//
//   MUHIM: input olish uchun & (manzil) ishlatamiz, chunki funksiya
//   o'zgaruvchining MANZILIga yozadi (keyinroq pointerlar darsida tushuntiramiz).
// =============================================================================

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// -----------------------------------------------------------------
	// 1-QISM: CHIQARISH (Output) MISOLLARI
	// -----------------------------------------------------------------
	fmt.Println("=== Output misollar ===")

	ism := "Aziz"
	yosh := 25
	bo_yi := 1.78
	talabaMisiz := true

	// %v — universal verb
	fmt.Printf("Ism: %v, Yosh: %v, Bo'yi: %v, Talaba: %v\n", ism, yosh, bo_yi, talabaMisiz)

	// Aniq verb'lar bilan
	fmt.Printf("Ism: %s, Yosh: %d, Bo'yi: %.2f m, Talaba: %t\n", ism, yosh, bo_yi, talabaMisiz)

	// Turini ham chiqaramiz
	fmt.Printf("Yosh turi: %T, Bo'yi turi: %T\n", yosh, bo_yi)

	// Raqamlarning turli ko'rinishlari
	son := 255
	fmt.Printf("255: decimal=%d, binary=%b, octal=%o, hex=%x\n", son, son, son, son)

	// Stringning quotalangan ko'rinishi
	matn := "Salom\tdunyo"
	fmt.Printf("Oddiy: %s | Quoted: %q\n", matn, matn)

	// Maydon kengligi va to'ldirish
	fmt.Printf("|%10d|\n", 42)   // o'ngga yotqazish, 10 ta belgi joy
	fmt.Printf("|%-10d|\n", 42)  // chapga yotqazish
	fmt.Printf("|%010d|\n", 42)  // nollar bilan to'ldirish

	// Sprintf — string sifatida qaytaradi (chop etmaydi)
	xabar := fmt.Sprintf("Salom %s, sizning balansingiz %.2f so'm.", ism, 1500.5)
	fmt.Println("Sprintf natijasi:", xabar)

	// -----------------------------------------------------------------
	// 2-QISM: INPUT (kiritish) MISOLLARI
	// -----------------------------------------------------------------
	fmt.Println("\n=== Mini Kalkulyator ===")

	// bufio.NewReader yoki Scanner — eng ishonchli usul
	reader := bufio.NewReader(os.Stdin)

	// Birinchi sonni o'qish
	fmt.Print("Birinchi sonni kiriting: ")
	birinchiQator, _ := reader.ReadString('\n')
	birinchiQator = strings.TrimSpace(birinchiQator) // \n va probellarni olib tashlaymiz

	birinchi, xato1 := strconv.ParseFloat(birinchiQator, 64)
	if xato1 != nil {
		fmt.Println("❌ Xato: birinchi qiymat son emas!", xato1)
		return
	}

	// Operatsiyani o'qish
	fmt.Print("Operatsiya (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	// Ikkinchi sonni o'qish
	fmt.Print("Ikkinchi sonni kiriting: ")
	ikkinchiQator, _ := reader.ReadString('\n')
	ikkinchiQator = strings.TrimSpace(ikkinchiQator)

	ikkinchi, xato2 := strconv.ParseFloat(ikkinchiQator, 64)
	if xato2 != nil {
		fmt.Println("❌ Xato: ikkinchi qiymat son emas!", xato2)
		return
	}

	// Hisoblash (oddiy if/else — keyingi darsda batafsil)
	var natija float64
	var xatoMatni string

	switch op {
	case "+":
		natija = birinchi + ikkinchi
	case "-":
		natija = birinchi - ikkinchi
	case "*":
		natija = birinchi * ikkinchi
	case "/":
		if ikkinchi == 0 {
			xatoMatni = "0 ga bo'lib bo'lmaydi!"
		} else {
			natija = birinchi / ikkinchi
		}
	default:
		xatoMatni = "Bunday operatsiya yo'q: " + op
	}

	// Natijani chiroyli chiqarish
	fmt.Println("\n--- Natija ---")
	if xatoMatni != "" {
		fmt.Println("❌", xatoMatni)
		return
	}
	fmt.Printf("%.2f %s %.2f = %.2f\n", birinchi, op, ikkinchi, natija)
}

// =============================================================================
// AMALIY VAZIFA:
//   1. Kalkulyatorga `%` (modul — qoldiqni olish) operatsiyasini qo'shing.
//   2. Foydalanuvchidan ismni so'rab, "Salom, [ism]!" deb chiqarib bering.
//   3. Bir nechta sonni o'qib, ularning yig'indisi va o'rtachasini chiqaring.
//   4. Printf orqali jadval shaklida natijalarni chiqarib ko'ring (maydon kengligi bilan).
//
// QO'SHIMCHA O'YLAB KO'RING:
//   - Nima uchun input olganda `&` ishlatildi (Scan) yoki ReadString ishlatildi?
//   - Sprintf qachon kerak bo'ladi? (javob: log yozish, error xabarlari, va h.k.)
// =============================================================================
