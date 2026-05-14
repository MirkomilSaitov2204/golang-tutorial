// ==================================================
// VAZIFA 2: Yonilg'i Quyish Stansiyasi
// ==================================================
// Maqsad: Avtotalovida yonilg'i tanlash va miqdorni hisoblash.
//
// QADAMLAR:
//   1. Foydalanuvchi yonilg'i turini tanlaydi (1, 2, 3).
//   2. Tanlaydi: a) litr bo'yicha, b) so'm bo'yicha.
//   3. Loyalty card mi (true/false).
//   4. Chek (kvitansiya) shaklida natija chiqsin.
//
// MAVZULAR: switch, strconv.ParseFloat, mantiqiy operatorlar.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	narx80 = 8_500.0  // AI-80
	narx92 = 9_500.0  // AI-92
	narx95 = 10_500.0 // AI-95

	loyaltyFoiz = 0.03 // 3% chegirma
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Yonilg'i quyish stansiyasi ===")
	fmt.Println()
	fmt.Println("Yonilg'i turini tanlang:")
	fmt.Printf("  1) AI-80 — %.0f so'm/litr\n", narx80)
	fmt.Printf("  2) AI-92 — %.0f so'm/litr\n", narx92)
	fmt.Printf("  3) AI-95 — %.0f so'm/litr\n", narx95)
	fmt.Print("Tanlovingiz (1-3): ")
	scanner.Scan()
	tur, xato := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if xato != nil || tur < 1 || tur > 3 {
		fmt.Println("Noto'g'ri tanlov.")
		return
	}

	// Switch — narxni tanlash
	var narxLitr float64
	var nom string
	switch tur {
	case 1:
		narxLitr, nom = narx80, "AI-80"
	case 2:
		narxLitr, nom = narx92, "AI-92"
	case 3:
		narxLitr, nom = narx95, "AI-95"
	}

	fmt.Println("\nQanday hisoblaymiz?")
	fmt.Println("  1) Litr bo'yicha")
	fmt.Println("  2) So'm bo'yicha")
	fmt.Print("Tanlov: ")
	scanner.Scan()
	usul, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	var litr, summa float64
	switch usul {
	case 1:
		fmt.Print("Necha litr? ")
		scanner.Scan()
		l, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil || l <= 0 {
			fmt.Println("Noto'g'ri qiymat.")
			return
		}
		litr = l
		summa = litr * narxLitr
	case 2:
		fmt.Print("Necha so'm? ")
		scanner.Scan()
		s, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil || s <= 0 {
			fmt.Println("Noto'g'ri qiymat.")
			return
		}
		summa = s
		litr = summa / narxLitr
	default:
		fmt.Println("Bunday usul yo'q.")
		return
	}

	// Loyalty
	fmt.Print("Loyalty card bormi? (ha/yo'q): ")
	scanner.Scan()
	loyalty := strings.ToLower(strings.TrimSpace(scanner.Text())) == "ha"

	chegirma := 0.0
	if loyalty {
		chegirma = summa * loyaltyFoiz
	}

	yakuniy := summa - chegirma

	// Min summa cheklash (foydalanuvchi 5 so'mga so'rasa)
	if litr < 0.1 {
		fmt.Println("\nMinimum 0.1 litr bo'lishi kerak.")
		return
	}

	// --- Chek formatida chiqarish ---
	fmt.Println()
	fmt.Println("====================================")
	fmt.Println("       NEFT-GAZ MARKAZI")
	fmt.Println("====================================")
	fmt.Printf("Yonilg'i:        %s\n", nom)
	fmt.Printf("Litr:            %.2f L\n", litr)
	fmt.Printf("Litr narxi:      %.0f so'm\n", narxLitr)
	fmt.Println("------------------------------------")
	fmt.Printf("Subtotal:        %.0f so'm\n", summa)
	if loyalty {
		fmt.Printf("Loyalty (-3%%):   -%.0f so'm\n", chegirma)
	}
	fmt.Println("====================================")
	fmt.Printf("JAMI:            %.0f so'm\n", yakuniy)
	fmt.Println("====================================")
	fmt.Println("Yo'lda omad tilaymiz!")
}
