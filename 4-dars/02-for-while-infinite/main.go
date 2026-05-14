// ==================================================
// MAVZU: for ning while va infinite shakllari
// ==================================================
// Nazariya:
//   Go'da ALOHIDA `while` operatori YO'Q. Lekin `for` ni
//   while sifatida ishlatish mumkin:
//
//   WHILE-SHAKL:
//     for shart {
//         // tana
//     }
//
//   INFINITE-SHAKL:
//     for {
//         // tana — break bilan to'xtatiladi
//     }
//
//   Eslab qoling:
//     - while-like: shart oldindan ma'lum bo'lmasa (foydalanuvchi inputi).
//     - infinite: serverlar, listener'lar, retry loop'lar.
//
// Real misol:
//   1) ATM PIN: foydalanuvchi 3 marotaba to'g'ri PIN kiritmasa, blok bo'ladi.
//      Bu — while shakli: "while (xato_soni < 3)".
//
//   2) Web server: cheksiz tinglaydi, har bir HTTP so'rovga javob beradi.
//      Bu — infinite for.

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

	// --- 1. WHILE-SHAKL: hisoblagich shart ---
	fmt.Println("=== Sonni 2 ga bo'lib borish (1 dan kichik bo'lguncha) ===")
	x := 100
	for x >= 1 {
		fmt.Print(x, " ")
		x /= 2 // ikkiga bo'lish
	}
	fmt.Println()
	// 100 50 25 12 6 3 1

	// --- 2. WHILE-SHAKL: Collatz sequence (3n+1 problemasi) ---
	fmt.Println("\n=== Collatz: 27 dan 1 ga ===")
	n := 27
	qadam := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		qadam++
	}
	fmt.Println("Qadamlar soni:", qadam) // 111

	// --- 3. WHILE-SHAKL: foydalanuvchi to'g'ri kiritmaguncha ---
	fmt.Println("\n=== To'g'ri yosh kiritish (validatsiyali) ===")
	yosh := -1
	for yosh < 0 || yosh > 150 {
		fmt.Print("Yoshingizni kiriting (0-150): ")
		scanner.Scan()
		matn := strings.TrimSpace(scanner.Text())
		y, err := strconv.Atoi(matn)
		if err != nil {
			fmt.Println("Faqat son. Qayta urinib ko'ring.")
			continue // siklning boshiga qaytadi
		}
		yosh = y
	}
	fmt.Println("Qabul qilindi, yoshingiz:", yosh)

	// --- 4. WHILE-SHAKL: PIN 3 marta urinish ---
	fmt.Println("\n=== ATM: PIN tekshirish (max 3 urinish) ===")
	togriPIN := "1234"
	urinish := 0
	blok := false

	for urinish < 3 {
		fmt.Printf("PIN kiriting (urinish %d/3): ", urinish+1)
		scanner.Scan()
		pin := strings.TrimSpace(scanner.Text())

		if pin == togriPIN {
			fmt.Println("Salom, hisobingizga xush kelibsiz!")
			break // sikldan chiqish — PIN to'g'ri
		}

		urinish++
		fmt.Println("Noto'g'ri PIN.")
	}
	if urinish == 3 {
		blok = true
		fmt.Println("BLOK: 3 marta xato. Karta bloklandi.")
	}
	_ = blok

	// --- 5. INFINITE for + break ---
	fmt.Println("\n=== Menyu (chiqishgacha cheksiz) ===")
	for {
		fmt.Println("\n--- MENYU ---")
		fmt.Println("1. Salomlashish")
		fmt.Println("2. Joriy vaqt (soat) hisoblash")
		fmt.Println("0. Chiqish")
		fmt.Print("Tanlov: ")
		scanner.Scan()
		tanlov := strings.TrimSpace(scanner.Text())

		if tanlov == "0" {
			fmt.Println("Xayr!")
			break // cheksiz sikldan chiqish
		}

		switch tanlov {
		case "1":
			fmt.Println("Salom, foydalanuvchi!")
		case "2":
			fmt.Print("Daqiqada kiriting: ")
			scanner.Scan()
			daq, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			fmt.Printf("Bu = %d soat %d daqiqa\n", daq/60, daq%60)
		default:
			fmt.Println("Noto'g'ri tanlov.")
		}
	}
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + while-shakl: shart soni oldindan noma'lum (input, network kutish).
// + infinite for: serverlar, retry, daemon process'lar.
// + REPL (Read-Eval-Print Loop) — bu fayldagi menyu kabi.
//
// XAVFLAR:
// - INFINITE LOOP — agar `break` yoki shart o'zgarishi YO'Q bo'lsa, dastur osilib qoladi.
//   Doim "sikldan chiqish yo'lini" o'ylab qoldiring (timeout, max iteratsiya).
//
// - O'zgaruvchi yangilanmaslik — for x < 10 { } va x hech qachon o'zgarmasa,
//   bu cheksiz sikl.
//
// MIGRATIYA (boshqa tillardan):
// - C/Java:  while(x < 10) {}    →  Go:  for x < 10 {}
// - Python:  while True:         →  Go:  for {}
// - Python:  do { } while(x)     →  Go:  yo'q — for-break kombinatsiya qiling
// ==================================================
