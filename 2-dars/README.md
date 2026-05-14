# 2-DARS — Asosiy sintaksis

## Bugun nimalarni o'rganamiz?

1-darsda biz o'zgaruvchilar bilan tanishdik. Endi ular ustida **amallar bajarishni** o'rganamiz. Tasavvur qiling, e-commerce saytida:
- Mahsulot narxiga **20% chegirma** berish kerak (`*` va `-` operatorlari)
- Foydalanuvchi yoshini tekshirish: `yosh >= 18` (taqqoslash)
- VIP **VA** balansi 100 ming so'mdan ko'p (`&&` mantiqiy operator)
- Foydalanuvchidan kiritilgan **"25"** stringini sonta aylantirish (`strconv`)

Bularning hammasi — bugun!

---

## Mavzular ro'yxati

### [01-operatorlar/](./01-operatorlar/main.go)
- Arifmetik: `+`, `-`, `*`, `/`, `%`
- Taqqoslash: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Mantiqiy: `&&`, `||`, `!`
- Bitwise (qisqacha): `&`, `|`, `^`, `<<`, `>>`
- Operatorlar prioriteti

### [02-tur-konvertatsiya/](./02-tur-konvertatsiya/main.go)
- Sonli turlar orasida konvertatsiya
- Aniqlik (precision) yo'qotish — diqqat!
- `byte` va `rune` konvertatsiyasi
- String va son orasidagi farq (string ↔ int — `strconv` orqali)

### [03-strconv/](./03-strconv/main.go)
- `strconv.Itoa`, `strconv.Atoi`
- `strconv.ParseFloat`, `FormatFloat`
- `strconv.ParseBool`, `FormatBool`
- Xatolarni qaytaradigan funksiyalar — birinchi tanishuv

### [04-input-techniques/](./04-input-techniques/main.go)
- `fmt.Scan`, `fmt.Scanln`, `fmt.Scanf` farqlari
- `bufio.Scanner` — eng ishonchli usul
- Bir nechta qiymatni bir qatorda olish
- Foydalanuvchi xatosini ushlash (validation)

---

## Amaliy vazifa
Har bir mavzu oxirida vazifa bor. Eng muhimi — har bir misolni o'zgartirib ko'rib, **"nima bo'larkin?"** deb tajriba qiling.

**Keyingi dars:** 3-dars — Shartli operatorlar (`if/else`, `switch`).
