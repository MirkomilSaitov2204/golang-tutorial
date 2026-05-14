## 3-DARS — Shartli operatorlar

### Bugun nimalarni o'rganamiz?

Tasavvur qiling, siz ATM bankomat yozyapsiz. U **qaror qabul qilishi** kerak:
- PIN to'g'rimi yoki noto'g'rimi?
- Balansda yetarli pulmi?
- Foydalanuvchi VIP statusda mi?
- Karta bloklanganmi?

Bu qarorlar — `if`, `else`, `switch` operatorlari orqali bajariladi. Bugungi 4 ta mavzu — aynan shu haqida.

---

### Mavzular ro'yxati

#### [01-if-else/](./01-if-else/main.go)
- `if`, `else if`, `else` sintaksisi
- Go ning O'ZIGA XOS xususiyatlari (qavs yo'q, `{` aynan shu qatorda)
- Murakkab shartlar (`&&`, `||`, `!`)
- ATM real misol: pul yechishni 5 bosqichli tekshirish
- Nested if va guard clause pattern

#### [02-switch/main.go](./02-switch/main.go)
- Avtomatik `break` (Go'da o'z-o'zidan to'xtaydi)
- Multi-case: `case 1, 2, 3:`
- `fallthrough` — keyingi case'ga "tushib" ketish
- `switch` o'zgaruvchisiz (if-else-if zamonaviy shakli)
- Real misollar: ATM menu, soliq stavkasi, buyurtma statusi

#### [03-qisqa-elon/](./03-qisqa-elon/main.go)
- `if x := f(); shart` — eng go'zal Go pattern'i
- Error handling'ning asosiy idiomalari: `if err := f(); err != nil`
- Scope toraytirish — bug'larni kamaytirish
- Map'da kalitning mavjudligini tekshirish (comma-ok)

#### [04-real-misollar/](./04-real-misollar/main.go)
**Interaktiv menyu** — 5 ta real misol bilan:
1. Kalkulyator (4 amal + 0 ga bo'lish himoyasi)
2. Yosh kategoriyasi (8 ta toifa)
3. BMI sog'liq indeksi
4. Mobil tarif tanlash maslahatchisi
5. Avtobus chiptasi narxi (yosh, masofa, mavsumi)

---

### Qanday ishga tushiriladi?

```bash
cd 3-dars/01-if-else
go run main.go
```

Yoki barcha 4 tasini ketma-ket:
```bash
for d in 3-dars/*/; do
  echo "=== $d ==="
  (cd "$d" && go run main.go) || break
done
```

---

### Amaliy vazifa (3-dars)

1. `04-real-misollar/main.go` ga **soliq kalkulyatori** qo'shing — oylik daromad asosida.
2. ATM misoliga (`01-if-else`) **PIN'ni kuchsizligini tekshiruvchi** shart qo'shing (`0000`, `1234` ga ruxsat bermang).
3. `02-switch/main.go` da soliq stavkalarini O'zbekiston bo'yicha real qiymatlarga sozlang.

---

**Keyingi dars:** [4-dars — Sikllar (for, range, FizzBuzz, prime)](../4-dars/README.md)
