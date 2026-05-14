// ==================================================
// MAVZU: Shartli operatorlar — Real loyihaviy misollar
// ==================================================
// Bu faylda biz 3-darsda o'rganganlarni REAL hayotiy
// muammolarga qo'llaymiz:
//
//   1. Kalkulyator (4 amal + 0 ga bo'lish himoyasi + xato)
//   2. Yosh kategoriyasi (8 ta toifa)
//   3. BMI hisoblovchi (sog'liq indeksi)
//   4. Telegram XizmatPaketi tanlash
//   5. Avtobus chiptasi narxi (yosh, masofa, mavsumi)
//
// Bularning hammasida if-else, switch, qisqa e'lon, &&/||
// va strconv ishlatiladi.

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

	fmt.Println("===========================================")
	fmt.Println("  Real misollar to'plami — qaysi birini")
	fmt.Println("  ishga tushirmoqchisiz?")
	fmt.Println("===========================================")
	fmt.Println("  1) Kalkulyator")
	fmt.Println("  2) Yosh kategoriyasi")
	fmt.Println("  3) BMI hisoblovchi")
	fmt.Println("  4) Mobil tarif tanlash")
	fmt.Println("  5) Avtobus chiptasi narxi")
	fmt.Print("Tanlovingiz: ")

	scanner.Scan()
	tanlovStr := strings.TrimSpace(scanner.Text())
	tanlov, xato := strconv.Atoi(tanlovStr)
	if xato != nil {
		fmt.Println("Faqat 1-5 oralig'idagi son.")
		return
	}

	switch tanlov {
	case 1:
		kalkulyator(scanner)
	case 2:
		yoshKategoriyasi(scanner)
	case 3:
		bmiHisobi(scanner)
	case 4:
		mobilTarif(scanner)
	case 5:
		avtobusChipta(scanner)
	default:
		fmt.Println("Bunday tanlov yo'q.")
	}
}

// ---------- 1. KALKULYATOR ----------
func kalkulyator(s *bufio.Scanner) {
	fmt.Println("\n--- Kalkulyator ---")
	fmt.Print("1-son: ")
	s.Scan()
	a, xato1 := strconv.ParseFloat(strings.TrimSpace(s.Text()), 64)

	fmt.Print("Operator (+ - * / %): ")
	s.Scan()
	op := strings.TrimSpace(s.Text())

	fmt.Print("2-son: ")
	s.Scan()
	b, xato2 := strconv.ParseFloat(strings.TrimSpace(s.Text()), 64)

	if xato1 != nil || xato2 != nil {
		fmt.Println("Xato: son emas kiritildi.")
		return
	}

	var natija float64
	switch op {
	case "+":
		natija = a + b
	case "-":
		natija = a - b
	case "*":
		natija = a * b
	case "/":
		if b == 0 {
			fmt.Println("Xato: 0 ga bo'lib bo'lmaydi!")
			return
		}
		natija = a / b
	case "%":
		// modul faqat int uchun — float ni int ga aylantirib qilamiz
		if b == 0 {
			fmt.Println("Xato: 0 ga modul bo'lmaydi.")
			return
		}
		natija = float64(int(a) % int(b))
	default:
		fmt.Println("Bunday operator yo'q:", op)
		return
	}
	fmt.Printf("Natija: %.4f\n", natija)
}

// ---------- 2. YOSH KATEGORIYASI ----------
func yoshKategoriyasi(s *bufio.Scanner) {
	fmt.Println("\n--- Yosh kategoriyasi ---")
	fmt.Print("Yoshingizni kiriting: ")
	s.Scan()
	yosh, xato := strconv.Atoi(strings.TrimSpace(s.Text()))
	if xato != nil || yosh < 0 || yosh > 150 {
		fmt.Println("Yosh 0-150 oralig'ida bo'lishi kerak.")
		return
	}

	// switch o'zgaruvchisiz — range tekshirish
	switch {
	case yosh < 1:
		fmt.Println("Chaqaloq (Newborn)")
	case yosh < 3:
		fmt.Println("Yosh bola (Toddler)")
	case yosh < 6:
		fmt.Println("Maktabgacha yosh")
	case yosh < 13:
		fmt.Println("Maktab yoshi")
	case yosh < 18:
		fmt.Println("O'smir")
	case yosh < 30:
		fmt.Println("Yoshlar")
	case yosh < 60:
		fmt.Println("Katta yoshli")
	default:
		fmt.Println("Keksa yosh — hurmatga sazovor")
	}
}

// ---------- 3. BMI HISOBLOVCHI ----------
func bmiHisobi(s *bufio.Scanner) {
	fmt.Println("\n--- BMI (Body Mass Index) ---")
	// Formula: BMI = vazn / (bo'y * bo'y),  bo'y metr'da

	fmt.Print("Vazningiz (kg): ")
	s.Scan()
	vazn, x1 := strconv.ParseFloat(strings.TrimSpace(s.Text()), 64)

	fmt.Print("Bo'yingiz (sm): ")
	s.Scan()
	boySm, x2 := strconv.ParseFloat(strings.TrimSpace(s.Text()), 64)

	if x1 != nil || x2 != nil || vazn <= 0 || boySm <= 0 {
		fmt.Println("Noto'g'ri qiymat.")
		return
	}

	boyM := boySm / 100         // sm → m
	bmi := vazn / (boyM * boyM) // formula

	fmt.Printf("Sizning BMI: %.2f\n", bmi)

	// Toifalash — WHO standartiga muvofiq
	if bmi < 18.5 {
		fmt.Println("Toifa: Vazn yetishmaydi")
	} else if bmi < 25.0 {
		fmt.Println("Toifa: Normal vazn — ajoyib!")
	} else if bmi < 30.0 {
		fmt.Println("Toifa: Ortiqcha vazn")
	} else if bmi < 35.0 {
		fmt.Println("Toifa: 1-darajali semizlik")
	} else if bmi < 40.0 {
		fmt.Println("Toifa: 2-darajali semizlik")
	} else {
		fmt.Println("Toifa: 3-darajali semizlik — shifokorga murojaat qiling")
	}
}

// ---------- 4. MOBIL TARIF TANLASH ----------
func mobilTarif(s *bufio.Scanner) {
	fmt.Println("\n--- Mobil tarif maslahatchisi ---")

	fmt.Print("Oylik internet sarfingiz (GB): ")
	s.Scan()
	gb, _ := strconv.Atoi(strings.TrimSpace(s.Text()))

	fmt.Print("Oylik suhbat (daqiqada): ")
	s.Scan()
	daq, _ := strconv.Atoi(strings.TrimSpace(s.Text()))

	fmt.Print("Boshqa operatorga ko'p qo'ng'iroq qilasizmi? (true/false): ")
	s.Scan()
	tashqi, _ := strconv.ParseBool(strings.TrimSpace(s.Text()))

	// Murakkab shart — mantiqiy operatorlar bilan
	switch {
	case gb < 5 && daq < 100:
		fmt.Println("Tavsiya: LITE — 30 ming/oy")
	case gb < 20 && daq < 500:
		fmt.Println("Tavsiya: STANDARD — 60 ming/oy")
	case gb < 50 && daq < 1500:
		fmt.Println("Tavsiya: PREMIUM — 120 ming/oy")
	case gb >= 50 || daq >= 1500:
		if tashqi {
			fmt.Println("Tavsiya: BUSINESS+ (tashqi pakeyt bilan) — 250 ming/oy")
		} else {
			fmt.Println("Tavsiya: BUSINESS — 200 ming/oy")
		}
	default:
		fmt.Println("Tavsiya: STANDARD — 60 ming/oy")
	}
}

// ---------- 5. AVTOBUS CHIPTASI NARXI ----------
func avtobusChipta(s *bufio.Scanner) {
	fmt.Println("\n--- Avtobus chiptasi narxi ---")

	fmt.Print("Yoshingiz: ")
	s.Scan()
	yosh, _ := strconv.Atoi(strings.TrimSpace(s.Text()))

	fmt.Print("Masofa (km): ")
	s.Scan()
	km, _ := strconv.Atoi(strings.TrimSpace(s.Text()))

	fmt.Print("Mavsum (yoz/qish/bahor/kuz): ")
	s.Scan()
	mavsum := strings.ToLower(strings.TrimSpace(s.Text()))

	fmt.Print("Talabamisiz? (true/false): ")
	s.Scan()
	talaba, _ := strconv.ParseBool(strings.TrimSpace(s.Text()))

	// Asosiy narx — har 1 km uchun 300 so'm
	narx := km * 300

	// Yosh chegirmasi
	switch {
	case yosh < 7:
		fmt.Println("→ 7 yoshgacha BEPUL")
		narx = 0
	case yosh < 16:
		narx = narx * 50 / 100 // 50% chegirma
		fmt.Println("→ Bola chegirmasi (50%)")
	case yosh >= 65:
		narx = narx * 30 / 100 // 70% chegirma
		fmt.Println("→ Nafaqaxor chegirmasi (70%)")
	}

	// Talaba qo'shimcha chegirma — agar avval to'liq narxda bo'lsa
	if talaba && yosh >= 16 && yosh < 65 {
		narx = narx * 70 / 100 // 30% chegirma
		fmt.Println("→ Talaba chegirmasi (30%)")
	}

	// Mavsum qo'shimchasi — qishda yo'l qiyin, +20%
	if mavsum == "qish" {
		narx = narx * 120 / 100
		fmt.Println("→ Qish qo'shimchasi (+20%)")
	}

	// Uzoq masofa qo'shimchasi — 500+ km
	if km > 500 {
		narx += 10000 // express xizmat
		fmt.Println("→ Uzoq masofa qo'shimchasi (+10 000 so'm)")
	}

	fmt.Printf("\nYAKUNIY NARX: %d so'm\n", narx)
}

// ==================================================
// AMALIY VAZIFA:
// ==================================================
// 1. Yangi toifa qo'shing: "Soliq stavkasini hisoblovchi".
//    Oylik daromad asosida soliqning 3 turli stavkasini hisoblang.
// 2. BMI hisoblovchiga JINSGA qarab chegara qo'shing
//    (erkak/ayol uchun farq qiladi).
// 3. Avtobus chiptasi narxiga FRIDA (Juma) chegirmasini qo'shing — 10%.
// ==================================================
