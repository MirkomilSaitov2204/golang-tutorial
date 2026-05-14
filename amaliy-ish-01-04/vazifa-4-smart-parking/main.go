// ==================================================
// VAZIFA 4: Smart Parking (Avtoturargoh)
// ==================================================
// Maqsad: Aeroport avtoturargohida narxni hisoblash.
//
// QADAMLAR:
//   1. Foydalanuvchi kirish soati va chiqish soatini kiritadi.
//   2. Farqi hisoblanadi (bir kun ichida).
//   3. Bosqichli narx tariflari qo'llanadi.
//   4. Kechki vaqt (20-08) +50%.
//   5. Hafta oxiri +20%.
//
// MAVZULAR: if-else-if zanjiri, mantiqiy operatorlar, arifmetik.

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

	fmt.Println("=== Smart Parking (Aeroport) ===")
	fmt.Println("Tarif:")
	fmt.Println("  1 soat:   bepul")
	fmt.Println("  2-4:      5 000 so'm/soat")
	fmt.Println("  5-8:      4 000 so'm/soat")
	fmt.Println("  8+:       30 000 so'm (sutkalik)")
	fmt.Println()

	kirish := soatSora(scanner, "Kirish soati (0-23): ")
	chiqish := soatSora(scanner, "Chiqish soati (0-23): ")

	// Soatlar farqi (bir kun ichida deb hisoblaymiz)
	soatlar := chiqish - kirish
	if soatlar < 0 {
		// Tunda chiqish (masalan, kirgan 22, chiqdi 03) — keyingi kun
		soatlar = (24 - kirish) + chiqish
	}
	if soatlar == 0 {
		fmt.Println("Siz hech narsa to'lamaysiz. Yaxshi yo'l!")
		return
	}

	fmt.Printf("Avtoturargohda turish: %d soat\n", soatlar)

	fmt.Print("Hafta oxiri (shanba/yakshanba)? (ha/yo'q): ")
	scanner.Scan()
	haftaOxiri := strings.ToLower(strings.TrimSpace(scanner.Text())) == "ha"

	// --- Asosiy narx hisoblash ---
	var narx float64

	switch {
	case soatlar <= 1:
		narx = 0
	case soatlar <= 4:
		// 1 soat bepul, qolgan soatlar 5000 so'mdan
		narx = float64(soatlar-1) * 5_000
	case soatlar <= 8:
		// 1 soat bepul, 2-4 soatlar — 3 ta * 5000 = 15000
		// 5-8 soatlar — (soatlar-4) * 4000
		narx = 3*5_000 + float64(soatlar-4)*4_000
	default:
		narx = 30_000 // sutkalik
	}

	// --- Kechki vaqt qo'shimchasi ---
	// 20:00 - 08:00 orasi (kirish soati shu oraliqda bo'lsa)
	if kirish >= 20 || kirish < 8 {
		// Bir kunlik bo'lmasa, kechki +50%
		if soatlar <= 8 {
			narx = narx * 1.5
			fmt.Println("→ Kechki vaqt qo'shimchasi (+50%)")
		}
	}

	// --- Hafta oxiri ---
	if haftaOxiri {
		narx = narx * 1.2
		fmt.Println("→ Hafta oxiri (+20%)")
	}

	fmt.Println()
	fmt.Println("==========================")
	fmt.Printf("Yakuniy narx: %.0f so'm\n", narx)
	fmt.Println("==========================")
}

func soatSora(s *bufio.Scanner, savol string) int {
	for {
		fmt.Print(savol)
		s.Scan()
		v, err := strconv.Atoi(strings.TrimSpace(s.Text()))
		if err != nil || v < 0 || v > 23 {
			fmt.Println("Xato: 0-23 oralig'idagi son.")
			continue
		}
		return v
	}
}
