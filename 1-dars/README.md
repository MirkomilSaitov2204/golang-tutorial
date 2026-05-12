# 1-DARS — Go bilan birinchi tanishuv

## Bugun nimalarni o'rganamiz?

Bugun Go tilida birinchi qadamlarimizni qo'yamiz. Tasavvur qiling, siz yangi ishga keldingiz va kompaniyangiz Go'da yangi backend yozmoqchi. Birinchi kuni sizdan **kalkulyator dasturi** yozish so'raldi. Buni qila olishingiz uchun avval:
1. Go'ni o'rnatishni,
2. Birinchi dasturni ishga tushirishni,
3. O'zgaruvchilar bilan ishlashni,
4. Ma'lumotlarni ekranga chiqarish va foydalanuvchidan olishni

bilishimiz kerak. Mana 4 ta mavzu — aynan shular.

---

## Mavzular ro'yxati

### [01-salom-dunyo/](./01-salom-dunyo/main.go)
- Go nima, qaerda ishlatiladi
- Go'ni o'rnatish (macOS uchun)
- Birinchi dasturning anatomiyasi: `package`, `import`, `func main()`
- `go run`, `go build` farqi

### [02-ozgaruvchilar/](./02-ozgaruvchilar/main.go)
- `var` va `:=` (qisqa shakl)
- `const` (o'zgarmas)
- Zero values (default qiymatlar)
- Multiple assignment

### [03-malumot-turlari/](./03-malumot-turlari/main.go)
- Asosiy turlar: `int`, `float64`, `string`, `bool`
- Tur konvertatsiyasi
- `rune` va `byte`
- `len()` va string xususiyatlari

### [04-input-output/](./04-input-output/main.go)
- `fmt.Print`, `Println`, `Printf` farqlari
- Verb'lar: `%d`, `%s`, `%f`, `%v`, `%T`, `%+v`
- `fmt.Scan`, `Scanln` orqali input olish
- **Yakuniy misol**: Mini-kalkulyator

---

## Go o'rnatish — qisqa ko'rsatma (macOS)

Terminalda quyidagini bajaring:

```bash
# Variant 1: Homebrew orqali (tavsiya etiladi)
brew install go

# Tekshirish
go version
# Natija: go version go1.23.x darwin/arm64

# Variant 2: Rasmiy saytdan
# https://go.dev/dl/ → macOS .pkg yuklab oling, o'rnating
```

Agar `brew` o'rnatilmagan bo'lsa:
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

---

## Har bir mavzuni qanday ishga tushirish?

Har bir mavzu o'z papkasida `main.go` fayl. Uni ishga tushirish uchun:

```bash
cd 1-dars/01-salom-dunyo
go run main.go
```

Yoki bir martagina kompilyatsiya qilib, executable yaratish:

```bash
go build -o salom main.go
./salom
```

---

## Bugungi dars uchun amaliy vazifa

1. Barcha 4 ta mavzudagi `main.go` fayllarni o'qing va ishga tushiring.
2. `02-ozgaruvchilar` ichidagi misollarni o'zgartirib ko'ring — masalan, o'z ismingiz va yoshingizni qo'shing.
3. `04-input-output` oxiridagi mini-kalkulyatorni kengaytirib, **bo'lish** va **modul** operatsiyalarini qo'shing.

Tushunmagan joyingiz bo'lsa — **savol bering**! Men "savol-javob" rejimida tushuntiraman.

---

**Keyingi dars:** 2-dars — Asosiy sintaksis (operatorlar, type conversion, `strconv`).
