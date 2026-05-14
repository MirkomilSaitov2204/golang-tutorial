// ==================================================
// VAZIFA 3: Smart Quiz (Matematika viktorinasi)
// ==================================================
// Maqsad: 5 ta matematika savoli berib, foydalanuvchi javobini tekshirish.
//
// QADAMLAR:
//   1. 5 ta savol ketma-ket beriladi.
//   2. Har to'g'ri javob — 10 ball.
//   3. Agar 3 ta ketma-ket TO'G'RI bersa — bonus +5 ball.
//   4. Agar 2 ta ketma-ket XATO bersa — ogohlantirish.
//   5. Yakuniy bahoni chiqaring.
//
// MAVZULAR: for klassik, if-else, akkumulyator, ketma-ketlik tracking.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Savollar va to'g'ri javoblar — bir nechta o'zgaruvchi orqali
	// (slice 7-darsda)
	savol1, javob1 := "5 + 7 = ?", 12
	savol2, javob2 := "12 * 8 = ?", 96
	savol3, javob3 := "144 / 12 = ?", 12
	savol4, javob4 := "23 - 7 + 4 = ?", 20
	savol5, javob5 := "9 * 9 - 1 = ?", 80

	fmt.Println("=== Matematika Quiz ===")
	fmt.Println("5 ta savol. Har to'g'ri javob — 10 ball.")
	fmt.Println("3 ta ketma-ket to'g'ri → bonus +5!")
	fmt.Println()

	ball := 0
	ketmaTogri := 0
	ketmaXato := 0

	// Massiv ishlatish o'rniga har bir savolni alohida ko'rib chiqamiz
	// (lekin sikl ichida har bir indeksni tanlash uchun for-switch ham bo'lardi)
	for i := 1; i <= 5; i++ {
		var savol string
		var togriJavob int
		switch i {
		case 1:
			savol, togriJavob = savol1, javob1
		case 2:
			savol, togriJavob = savol2, javob2
		case 3:
			savol, togriJavob = savol3, javob3
		case 4:
			savol, togriJavob = savol4, javob4
		case 5:
			savol, togriJavob = savol5, javob5
		}

		fmt.Printf("Savol %d: %s ", i, savol)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		javob, xato := strconv.Atoi(input)
		if xato != nil {
			fmt.Println("Xato: faqat son. Savol o'tkazib yuborildi.")
			ketmaXato++
			ketmaTogri = 0
			continue
		}

		if javob == togriJavob {
			ball += 10
			ketmaTogri++
			ketmaXato = 0
			fmt.Printf("To'g'ri! (+10 ball) Joriy: %d\n", ball)

			// Bonus: 3 ta ketma-ket to'g'ri
			if ketmaTogri == 3 {
				ball += 5
				fmt.Println("BONUS! 3 ta ketma-ket to'g'ri — +5 ball")
				ketmaTogri = 0 // reset
			}
		} else {
			ketmaXato++
			ketmaTogri = 0
			fmt.Printf("Xato. To'g'ri javob: %d. Joriy: %d\n", togriJavob, ball)

			// Ogohlantirish: 2 ta ketma-ket xato
			if ketmaXato == 2 {
				fmt.Println("DIQQAT: 2 ta ketma-ket xato. Asoslarni qayta o'qib chiqing!")
				ketmaXato = 0
			}
		}
		fmt.Println()
	}

	// Yakuniy baho
	fmt.Println("====================")
	fmt.Println("    YAKUNIY BAHO    ")
	fmt.Println("====================")
	fmt.Printf("Sizning balingiz: %d / 50 (maksimal: 55 bonus bilan)\n", ball)

	switch {
	case ball >= 50:
		fmt.Println("Daraja: GENIUS! (Yulduz)")
	case ball >= 40:
		fmt.Println("Daraja: Yaxshi! Davom eting.")
	case ball >= 25:
		fmt.Println("Daraja: O'rta. Mashqlar kerak.")
	default:
		fmt.Println("Daraja: Yetarli emas. Asoslarni qayta o'qing.")
	}
}
