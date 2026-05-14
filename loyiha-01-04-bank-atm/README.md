## LOYIHA: Aqlli ATM Simulyatori (Smart Bank ATM)

> **Maqsad**: 1-4 darslarda o'rganganlarning HAMMASINI bitta to'liq, ishlaydigan, real-tilda loyihada birlashtirish.

---

### Loyiha haqida

Bu — terminal-asoslangan **bank ATM** simulyatori. Foydalanuvchi:
- PIN orqali tizimga kiradi
- Balansni ko'radi, pul yechadi, qo'yadi
- Boshqa kartaga o'tkazma qiladi (komissiya bilan)
- Operatsiyalar tarixini ko'radi
- PIN'ni o'zgartiradi

Hamma asosiy bank ATM funksiyasi — Go'da, faqat 1-4 darslarda o'rganilgan vositalar bilan.

---

### Biznes qoidalari (real bankka o'xshash)

| Qoida | Qiymat |
|-------|--------|
| Boshlang'ich balans | 1 500 000 so'm |
| Boshlang'ich PIN | 1234 |
| Kunlik yechish limiti | 5 000 000 so'm |
| Bir martalik yechish max | 2 000 000 so'm |
| Kupyura nominali | 10 000 so'mga karrali |
| Komissiya boshqa bankka | 1% (min 5 000 so'm) |
| VIP chegarasi | 10 000 000+ balans → komissiya yo'q |
| Max PIN urinish | 3 marta, keyin BLOK |
| Pul qo'yish (AML) | 50 mln dan kam |

---

### Loyiha qaysi tushunchalarni qoplaydi?

#### 1-dars
- ✅ `const`, `var`, `:=`
- ✅ `int`, `float64`, `string`, `bool`
- ✅ `fmt.Print`, `fmt.Printf`, `fmt.Println`

#### 2-dars
- ✅ Arifmetik operatorlar (`+`, `-`, `*`, `/`, `%`)
- ✅ Taqqoslash (`>=`, `<=`, `==`, `!=`)
- ✅ Mantiqiy (`&&`, `||`, `!`)
- ✅ Tur konvertatsiyasi (`float64 → int` modul uchun)
- ✅ `strconv.Atoi`, `strconv.ParseFloat`
- ✅ `bufio.Scanner` (to'liq qatorni o'qish)
- ✅ `strings.TrimSpace`, `strings.ToLower`

#### 3-dars
- ✅ `if`, `else if`, `else` — 5-bosqichli validatsiya
- ✅ `switch` — menu va bank turi
- ✅ Qisqa e'lon: `if _, err := strconv.Atoi(s); err != nil`

#### 4-dars
- ✅ `for` while-shakl — PIN urinish hisoblagich
- ✅ `for {}` cheksiz — asosiy menu sikli
- ✅ `for-range` — tarix ustida
- ✅ `break`, `continue`

---

### Loyihani ishga tushirish

```bash
cd loyiha-01-04-bank-atm
go run main.go
```

#### Demo session

```text
==========================================
    AQLLI BANK — ATM SIMULYATORI v1.0
==========================================

PIN kodingizni kiriting (urinish 1/3): 1234

Muvaffaqiyatli kirish! Xush kelibsiz.

==========================================
                 ASOSIY MENU
==========================================
  1) Balansni ko'rish
  2) Pul yechish
  3) Pul qo'yish (depozit)
  4) Boshqa kartaga o'tkazma
  5) Operatsiyalar tarixi (oxirgi 5 ta)
  6) Hisobotni ko'rish
  7) PIN kodni o'zgartirish
  0) Chiqish

Tanlovingiz: 2

>>> PUL YECHISH <<<
Joriy balans: 1500000.00 so'm
Bugun yechilgan: 0.00 so'm (limit: 5000000)
Qancha yechmoqchisiz? 250000
Muvaffaqiyat! Yechildi: 250000 so'm. Yangi balans: 1250000.00
```

---

### Test qilish ssenariylari

Quyidagi holatlarni o'zingiz tekshirib ko'ring:

#### Xavfsizlik
- [ ] 3 marta xato PIN → karta bloklanadimi?
- [ ] PIN faqat raqamlardan iboratligini tekshiruvchi
- [ ] PIN aynan 4 ta belgi ekanligi

#### Pul yechish
- [ ] Balansdan ko'p yechmoqchi bo'lganda
- [ ] Manfiy summa
- [ ] 0
- [ ] 12 345 (kupyuraga karrali emas)
- [ ] Bir martalik 2 mln dan ko'p
- [ ] Kunlik 5 mln dan ko'p (bir nechta operatsiyada)

#### O'tkazma
- [ ] Karta 16 raqam emas
- [ ] Karta harf bilan
- [ ] VIP statusda — komissiyasiz
- [ ] Oddiy — 1% komissiya
- [ ] Min komissiya (5000) ishlashi (kichik summa)
- [ ] Bank turi noto'g'ri (`other` yoki `own` emas)

#### PIN o'zgartirish
- [ ] Joriy PIN noto'g'ri
- [ ] Yangi PIN eski bilan bir xil
- [ ] Yangi PIN — kuchsiz (`0000`, `1234`)
- [ ] Tasdiqlash mos kelmadi

---

### Loyiha tuzilishi

```
loyiha-01-04-bank-atm/
├── README.md       # bu fayl
└── main.go         # ~350 qator — to'liq ATM simulyatori
```

> **Eslatma**: Kelajak darslarda (5-dars: funksiya, 7-dars: slice, 10-dars: struct) bu loyihani qayta yozish va **toza arxitektura**ga keltirish vazifasi beriladi.

---

### Loyihani kengaytirish g'oyalari

Agar darsdan ortib qolgan vaqtingiz bo'lsa:

1. **Foiz hisoblash** — har oy oxirida balansga 0.5% qo'shilsin.
2. **Mavjudlik tekshiruv** — pul qo'yishda yuqori limit (1 mlrd) qo'ying.
3. **Foydalanuvchilar bo'yicha** — 3 ta turli kartani (3 ta turli PIN) qo'llang. Switch orqali kim kirayotganini tanlash.
4. **Bilingual** — o'zbek/rus tilini tanlash imkoniyati.
5. **Mini hisob-kitob** — har bir o'tkazma uchun **kurs konvertatsiyasi** (USD/UZS).

---

**Keyingi qadam**: [Amaliy ish vazifalari](../amaliy-ish-01-04/README.md) — yana 5 ta loyiha vazifasi sizni kutmoqda!
