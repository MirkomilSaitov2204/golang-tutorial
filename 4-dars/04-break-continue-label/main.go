// ==================================================
// MAVZU: break, continue va labellar (yorliqlar)
// ==================================================
// Nazariya:
//   Siklni boshqarish uchun 3 ta operator:
//
//   - break    → ENG YAQIN siklni TO'XTATADI va undan chiqib ketadi.
//   - continue → joriy iteratsiyani O'TKAZIB YUBORADI, keyingisiga o'tadi.
//   - return   → butun funksiyani to'xtatadi (funksiyani 5-darsda chuqurroq).
//
//   LABEL (yorliq) — agar NESTED (ichma-ich) sikllarda
//   ENG TASHQI sikldan chiqish kerak bo'lsa:
//
//     tashqi:
//     for i := 0; i < 5; i++ {
//         for j := 0; j < 5; j++ {
//             if shart {
//                 break tashqi   // ikkala sikldan ham chiqadi!
//             }
//         }
//     }
//
//   Esda saqlab oling:
//     - break i++ — i++ NA, faqat label.
//     - Label nomi bilan bir qatorda `for` boshlanmaydi — yangi qatorda.
//     - Label kamdan-kam kerak — ko'p hollarda funksiyaga ajratish yaxshiroq.
//
// Real misol:
//   E-commerce search: 1000 ta mahsulot ichidan birinchi mos kelganini topish.
//   Topilgach — siklni darhol to'xtatish (vaqt tejash).

package main

import "fmt"

func main() {
	// --- 1. break: birinchi ko'paytuvchini topish ---
	fmt.Println("=== break — 42 ning birinchi ko'paytuvchisini topish ===")
	son := 42
	var birinchi int
	for i := 2; i <= son; i++ {
		if son%i == 0 {
			birinchi = i
			break // topdik — to'xtaymiz
		}
	}
	fmt.Println("Birinchi ko'paytuvchi:", birinchi) // 2

	// --- 2. continue: faqat juftlarni chop etish ---
	fmt.Println("\n=== continue — 1-10 dagi toq sonlarni o'tkazib yuborish ===")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // toq son — keyingi iteratsiyaga
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	// 2 4 6 8 10

	// --- 3. break va continue birga ---
	fmt.Println("\n=== Birinchi 5 ta tub son ===")
	topildi := 0
	n := 2
	for {
		// Tub son tekshirish
		tubMi := true
		for j := 2; j*j <= n; j++ {
			if n%j == 0 {
				tubMi = false
				break // ichki sikldan chiqamiz
			}
		}

		if !tubMi {
			n++
			continue
		}

		fmt.Print(n, " ")
		topildi++
		if topildi == 5 {
			break // 5 tasini topdik — tashqi sikldan chiqamiz
		}
		n++
	}
	fmt.Println()
	// 2 3 5 7 11

	// --- 4. LABEL (YORLIQ) — nested sikldan chiqish ---
	fmt.Println("\n=== Label — 2D matritsada birinchi 0 ni topish ===")
	matritsa := [3][3]int{
		{1, 2, 3},
		{4, 0, 6}, // 0 bu yerda!
		{7, 8, 9},
	}
	xRow, xCol := -1, -1

tashqi: // bu yorliq — yangi qatorda yozilishi kerak
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if matritsa[i][j] == 0 {
				xRow, xCol = i, j
				break tashqi // har ikki sikldan chiqamiz
			}
		}
	}
	fmt.Printf("0 topildi: satr %d, ustun %d\n", xRow, xCol)

	// --- 5. continue label ---
	fmt.Println("\n=== continue label — diagonalni o'tkazib yuborish ===")
qator:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue qator // butun satrni emas, faqat shu (i,j) ni o'tkazamiz
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()
	// (0,1) (0,2) (1,0) (1,2) (2,0) (2,1)
	// XATO TUSHUNCHA: continue label — TASHQI sikl iteratsiyasini boshlaydi,
	// ya'ni ichki sikl to'xtaydi va tashqi siklning keyingi iteratsiyasiga o'tadi.

	// --- 6. break TASHQARIDA — XATO ---
	// break faqat sikl ICHIDA ishlaydi. Aks holda kompilyatsiya xatosi.
	// `break` agar `switch` ichida bo'lsa — switch'dan chiqadi (sikldan emas)!
	// Buni eslab qoling:
	fmt.Println("\n=== switch ichidagi break ===")
	for i := 0; i < 5; i++ {
		switch i {
		case 3:
			fmt.Println("3 ga yetdik, switch'dan chiqamiz")
			break // FAQAT switch'dan chiqadi, FOR davom etadi!
		}
		fmt.Println("i =", i)
	}
	// Agar haqiqatan ham SIKLDAN chiqmoqchi bo'lsangiz — label kerak:
	fmt.Println("\n=== switch ichidan SIKL'dan chiqish — label kerak ===")
siklBoshi:
	for i := 0; i < 5; i++ {
		switch i {
		case 3:
			fmt.Println("3 ga yetdik, butun sikldan chiqamiz")
			break siklBoshi // label bilan — sikldan chiqadi!
		}
		fmt.Println("i =", i)
	}

	// --- 7. AMALIY: do'kondan birinchi tugagan mahsulotni topish ---
	fmt.Println("\n=== Mahsulotlar inventarizatsiyasi ===")
	mahsulotlar := []string{"Olma", "Banan", "Anor", "Behi", "Anjir"}
	zaxiralar := []int{50, 0, 30, 15, 25} // banan tugagan

	for i, m := range mahsulotlar {
		if zaxiralar[i] == 0 {
			fmt.Printf("DIQQAT: %s tugagan! (indeks %d)\n", m, i)
			break // birinchi topilganida to'xtaymiz
		}
		fmt.Printf("%s: %d ta\n", m, zaxiralar[i])
	}
}

// ==================================================
// QACHON QO'LLANADI:
// ==================================================
// + break — natija topilganda erta to'xtash (vaqt tejash).
// + continue — ba'zi iteratsiyalarni "yo'q" qilib o'tkazish (filter).
// + label — chuqur nested sikllardan chiqish (kamdan-kam).
//
// ANTI-PATTERN'LAR:
// - Label'ni har joyda ishlatish — kod o'qib bo'lmaydigan bo'ladi.
//   YAXSHIROQ: nested sikllarni alohida funksiyaga ajrating + return.
// - switch ICHIDA `break` orqali sikldan chiqmoqchi bo'lish — bu ISHLAMAYDI.
//   Label ishlating yoki sikl shartini o'zgartiring.
// ==================================================
