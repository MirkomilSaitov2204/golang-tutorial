// ==================================================
// VAZIFA 5: Son topish o'yini (Number Guessing Game)
// ==================================================
// Maqsad: Klassik 1-100 oraligida tasodifiy son tanlanadi.
// Foydalanuvchi taxminni kiritadi, dastur "ko'p" yoki "kam" deydi.
//
// QADAMLAR:
//   1. math/rand bilan 1-100 tasodifiy son.
//   2. Max 7 urinish.
//   3. Topgan bo'lsa — nechta urinishni chiqaring.
//   4. Yakunda — "Yana o'ynaysizmi?".
//
// MAVZULAR: for cheksiz, break, continue, math/rand, mantiqiy.

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const maxUrinish = 7

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Son topish o'yini ===")
	fmt.Println("Men 1-100 oralig'ida son o'yladim. Toping!")
	fmt.Printf("Max urinishlar: %d\n\n", maxUrinish)

	// Tashqi sikl — qayta o'ynash uchun
	for {
		// Yangi tasodifiy son (Go 1.22+ math/rand/v2)
		togri := rand.IntN(100) + 1 // 1-100
		urinish := 0
		topdi := false

		// Ichki sikl — urinishlar
		for urinish < maxUrinish {
			urinish++
			fmt.Printf("Urinish %d/%d. Sizning taxmin: ", urinish, maxUrinish)
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			javob, xato := strconv.Atoi(input)
			if xato != nil {
				fmt.Println("Xato: faqat son.")
				urinish-- // bu urinish hisobga olinmaydi
				continue
			}

			if javob < 1 || javob > 100 {
				fmt.Println("Diapazon: 1-100.")
				urinish--
				continue
			}

			if javob == togri {
				fmt.Printf("\nTOPDINGIZ! %d-urinishda. Javob: %d\n", urinish, togri)
				topdi = true
				break
			} else if javob < togri {
				fmt.Println("→ Ko'proq. Yuqori son ayting.")
			} else {
				fmt.Println("→ Kam. Pastroq son ayting.")
			}

			// Maslahat (hint) — 5-urinishdan keyin
			if urinish == 5 {
				// "Bir oz yaqin" yoki "Juda uzoq" maslahatini berish
				farq := javob - togri
				if farq < 0 {
					farq = -farq
				}
				if farq <= 5 {
					fmt.Println("HINT: juda yaqindasiz!")
				} else if farq >= 30 {
					fmt.Println("HINT: juda uzoqdasiz.")
				}
			}
		}

		if !topdi {
			fmt.Printf("\nMag'lubiyat! To'g'ri javob: %d edi.\n", togri)
		}

		// Yana o'ynash?
		fmt.Print("\nYana o'ynaysizmi? (ha/yo'q): ")
		scanner.Scan()
		davom := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if davom != "ha" && davom != "yes" {
			fmt.Println("\nXayr! Yana keling.")
			break // tashqi sikldan chiqamiz
		}
		fmt.Println()
	}
}
