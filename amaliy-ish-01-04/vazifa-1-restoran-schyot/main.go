// ==================================================
// VAZIFA 1: Restoran Schyot (Bill Calculator)
// ==================================================
// Maqsad: Restoran kassasida buyurtma uchun yakuniy summani hisoblovchi
// dastur yozish. Servis to'lovi, QQS, chegirma va tushlik chegirmasi qo'shilishi kerak.
//
// QADAMLAR:
//   1. Foydalanuvchidan har bir taom uchun nechta buyurtma berganini soraysiz.
//   2. Subtotal hisoblang.
//   3. Tushlik vaqti (12:00-15:00)mi tekshiring.
//   4. Chegirma qo'llang (loyalty + tushlik).
//   5. Servis (10%) va QQS (15%) qo'shing.
//   6. Yakuniy summani chiqaring.
//
// MAVZULAR: const, var, fmt.Scan, if-else-if, switch, mantiqiy operatorlar.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Taom narxlari (const — chunki o'zgarmaydigan biznes qoidalar)
const (
	plovNarxi    = 35_000 // Osh
	shashlikNarxi = 18_000 // Shashlik (har bir cho'p)
	lagmonNarxi  = 28_000 // Lag'mon
	choyNarxi    = 5_000  // Choynak

	servisFoiz    = 0.10 // 10% servis
	qqsFoiz       = 0.15 // 15% QQS
	loyaltyChegara = 500_000
	loyaltyFoiz   = 0.05 // 5% chegirma
	tushlikFoiz   = 0.10 // 10% chegirma (12:00-15:00)
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Restoran 'Mehmondo'st' ===")
	fmt.Println("Menyu:")
	fmt.Printf("  1) Osh         — %d so'm\n", plovNarxi)
	fmt.Printf("  2) Shashlik    — %d so'm (1 cho'p)\n", shashlikNarxi)
	fmt.Printf("  3) Lag'mon     — %d so'm\n", lagmonNarxi)
	fmt.Printf("  4) Choynak     — %d so'm\n", choyNarxi)
	fmt.Println()

	// TODO: Har bir taom uchun nechta olganini so'rang.
	// Maslahat: Funksiya yo'q hozircha, har birini alohida `scanner.Scan()` bilan.
	plovSoni := soranSon(scanner, "Osh nechta?")
	shashlikSoni := soranSon(scanner, "Shashlik nechta?")
	lagmonSoni := soranSon(scanner, "Lag'mon nechta?")
	choySoni := soranSon(scanner, "Choynak nechta?")

	// Soat — hozirgi vaqt
	soat := soranSon(scanner, "Hozir nechi soat (0-23)?")

	fmt.Print("Loyalty card bormi? (ha/yo'q): ")
	scanner.Scan()
	loyaltyJavob := strings.ToLower(strings.TrimSpace(scanner.Text()))
	loyaltyBor := loyaltyJavob == "ha" || loyaltyJavob == "yes" || loyaltyJavob == "true"

	// --- Subtotal hisoblash ---
	subtotal := plovSoni*plovNarxi + shashlikSoni*shashlikNarxi +
		lagmonSoni*lagmonNarxi + choySoni*choyNarxi

	if subtotal == 0 {
		fmt.Println("\nHech narsa buyurmadingiz. Xayr!")
		return
	}

	fmt.Println("\n=== HISOB ===")
	if plovSoni > 0 {
		fmt.Printf("  Osh × %d        = %d\n", plovSoni, plovSoni*plovNarxi)
	}
	if shashlikSoni > 0 {
		fmt.Printf("  Shashlik × %d   = %d\n", shashlikSoni, shashlikSoni*shashlikNarxi)
	}
	if lagmonSoni > 0 {
		fmt.Printf("  Lag'mon × %d    = %d\n", lagmonSoni, lagmonSoni*lagmonNarxi)
	}
	if choySoni > 0 {
		fmt.Printf("  Choynak × %d    = %d\n", choySoni, choySoni*choyNarxi)
	}
	fmt.Printf("  ----------------\n")
	fmt.Printf("  Subtotal         = %d so'm\n", subtotal)

	// --- Chegirmalar ---
	// 1) Tushlik chegirmasi (12-15 oralig'i)
	tushlikChegirma := 0.0
	if soat >= 12 && soat < 15 {
		tushlikChegirma = float64(subtotal) * tushlikFoiz
		fmt.Printf("  Tushlik chegirma (-%.0f%%) = -%.0f\n", tushlikFoiz*100, tushlikChegirma)
	}

	// 2) Loyalty chegirmasi (500k+)
	loyaltyChegirma := 0.0
	if loyaltyBor && subtotal >= loyaltyChegara {
		loyaltyChegirma = float64(subtotal) * loyaltyFoiz
		fmt.Printf("  Loyalty chegirma (-%.0f%%) = -%.0f\n", loyaltyFoiz*100, loyaltyChegirma)
	}

	pastSubtotal := float64(subtotal) - tushlikChegirma - loyaltyChegirma

	// --- Servis va QQS ---
	servis := pastSubtotal * servisFoiz
	qqs := pastSubtotal * qqsFoiz
	yakuniy := pastSubtotal + servis + qqs

	fmt.Printf("  Chegirmali        = %.0f\n", pastSubtotal)
	fmt.Printf("  Servis (10%%)     = +%.0f\n", servis)
	fmt.Printf("  QQS (15%%)        = +%.0f\n", qqs)
	fmt.Println("  ================")
	fmt.Printf("  YAKUNIY          = %.0f so'm\n", yakuniy)
	fmt.Println("\nRahmat, yana keling!")
}

// soranSon — yordamchi funksiya. (Funksiya 5-darsda, lekin bu yerda sodda foydalanish)
func soranSon(s *bufio.Scanner, savol string) int {
	for {
		fmt.Print(savol + " ")
		s.Scan()
		v, err := strconv.Atoi(strings.TrimSpace(s.Text()))
		if err != nil || v < 0 {
			fmt.Println("Xato: musbat son kiriting.")
			continue
		}
		return v
	}
}
