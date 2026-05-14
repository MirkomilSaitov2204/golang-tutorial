## 4-DARS — Sikllar (Loops)

### Bugun nimalarni o'rganamiz?

Sikl — bu dasturlashning ENG MUHIM operatorlardan biri. Tasavvur qiling:
- E-commerce sayti **1000 ta** mahsulotni ekranga chiqarishi kerak
- ATM **noto'g'ri PIN kiritilsa qayta so'raydi** (3 marotaba)
- Web server **cheksiz** mijozlarni tinglaydi

Bularning hammasi — `for` sikllarining turli shakllari orqali.

> **Eslatma**: Go'da FAQAT bitta sikl operatori bor — `for`. Lekin u 3-4 xil shaklda ishlatiladi.

---

### Mavzular ro'yxati

#### [01-for-classic/](./01-for-classic/main.go)
- Klassik C-shakl: `for i := 0; i < n; i++`
- Teskari yo'nalish, qadam (step) bilan
- Yig'indi, ko'paytma, faktorial, ko'paytirish jadvali
- Nested for (yulduzcha figura, Pifagor jadvali)
- Maximum / minimum topish algoritmi

#### [02-for-while-infinite/](./02-for-while-infinite/main.go)
- While-shakl: `for shart { ... }`
- Infinite: `for { ... }` + `break`
- Collatz ketma-ketligi (matematik misol)
- ATM PIN urinish (3 marta)
- REPL menyu (cheksiz tinglash)

#### [03-range/](./03-range/main.go)
- `range` slice, string, map, channel ustida
- 4 ta variantni: `i, v`, `_, v`, `i`, `range N`
- String'da UTF-8 / rune indeksi
- Map'da TARTIB RANDOM ekanligi
- Filter, count, max indeksini topish

#### [04-break-continue-label/](./04-break-continue-label/main.go)
- `break` — sikldan chiqish
- `continue` — joriy iteratsiyani o'tkazib yuborish
- LABEL (yorliq) — nested sikldan chiqish
- `switch` ICHIDAGI `break` — TUZOQ! (`for`ga ta'sir qilmaydi)
- Matritsada qiymatni topish

#### [05-fizzbuzz-prime-fibonacci/](./05-fizzbuzz-prime-fibonacci/main.go)
**Algoritmik misollar** (interview da ko'p so'raladi):
1. FizzBuzz klassikasi
2. Tub sonlar (prime) — optimal algoritm (`i*i <= n`)
3. Fibonacci ketma-ketligi (iteratsiya bilan, O(n))
4. Sonni teskari aylantirish (123 → 321)
5. Raqamlar yig'indisi (digital sum — checksum'larda kerak)
6. Palindrom tekshiruv
7. GCD — Yevklid algoritmi

---

### Qanday ishga tushiriladi?

```bash
cd 4-dars/01-for-classic
go run main.go
```

---

### Amaliy vazifa (4-dars)

Har bir mavzu oxirida vazifa berilgan. Eng muhimlari:
1. **FizzBuzz kengaytirish**: 7 ga bo'lingan sonlar uchun "Boom" qo'shing.
2. **Tub sonlar**: birinchi 50 ta tub son orasidagi maksimum farq.
3. **Fibonacci**: 1 mln dan kichik eng katta Fibonacci son.
4. **Luhn algoritmi**: kredit karta validatsiyasi (Google'da o'qing).
5. **Palindrom STRING**: "kabak", "ana" — palindrom; "olma" — emas.

---

### Bog'liq matematik tushunchalar

| Algoritm | Murakkablik | Real ishlatish |
|----------|-------------|----------------|
| FizzBuzz | O(n) | Interview tezkor savol |
| Tub son tekshirish | O(√n) | Cryptografiya (RSA), hash |
| Fibonacci (iter) | O(n) | Dinamik dasturlash misoli |
| Digital sum | O(log n) | Luhn (kredit karta), checksum |
| Palindrom | O(n) | String processing |
| GCD (Euclidean) | O(log min(a,b)) | Kasr soddalashtirish, RSA |

---

**Keyingi dars:** 5-dars — Funksiyalar (parametrlar, return, multiple return, variadic).

> **Loyiha**: 1-4 darslarda o'rganganlaringizni mustahkamlash uchun [Aqlli ATM Simulyatori](../loyiha-01-04-bank-atm/README.md) loyihasini ko'rib chiqing!
