## Amaliy ish: 1-4 darslar bo'yicha 5 ta mustaqil loyiha

> **Maqsad**: ATM loyihasidan keyin, ko'nikmangizni mustahkamlash uchun bu 5 ta vazifani **o'zingiz** yoza olishingiz kerak. Har biri o'rtacha 1-3 soat vaqt oladi.

Har bir vazifa shu papkada (`amaliy-ish-01-04/`) **alohida sub-papka**da bajariladi. Skeleton fayllar (`vazifa-N/main.go`) tayyorlangan — siz `// TODO` qatorlarini to'ldirasiz.

---

## VAZIFA 1: Restoran Schyot (Bill Calculator)

**Real ssenariy**: Restoran kassasida ishlayapsiz. Mijoz buyurtmasini hisoblash kerak.

**Talab**:
- 4 ta turli taom narxi (`const` orqali e'lon qilinsin)
- Foydalanuvchi har bir taomdan nechta olganini kiritadi
- Hisobotda chiqsin:
  - Har bir taom subtotal
  - Umumiy summa
  - **Servis to'lovi** (10%)
  - **QQS** (15%)
  - **YAKUNIY** (umumiy + servis + QQS)
- Agar yakuniy summa 500 000+ bo'lsa, **5% chegirma** beriladi (loyalty)
- Tushlik vaqti (12:00-15:00) — qo'shimcha 10% chegirma

**Texnologiyalar**: `if/else`, `switch`, `fmt.Scan`, arifmetik, `bufio.Scanner`.

📁 [vazifa-1-restoran-schyot/main.go](./vazifa-1-restoran-schyot/main.go)

---

## VAZIFA 2: Yonilg'i Quyish Stansiyasi

**Real ssenariy**: Avtotalovida yonilg'i quyish stansiyasi terminali.

**Talab**:
- 3 xil yonilg'i: AI-80 (8500 so'm), AI-92 (9500), AI-95 (10500)
- Mijoz tanlaydi va litr miqdorini kiritadi
- Yoki **so'm bo'yicha**: necha so'mlik quymoqchisiz?
- Hisoblang: litr/jami so'm/qoldiq
- **Loyalty card** soraydi (true/false):
  - Bor: 3% chegirma
  - Yo'q: oddiy narx
- **Yakun**: chek (kvitansiya) shaklida ko'rsatish

**Texnologiyalar**: `switch`, `strconv.ParseFloat`, mantiqiy operatorlar.

📁 [vazifa-2-yonilgi-stansiyasi/main.go](./vazifa-2-yonilgi-stansiyasi/main.go)

---

## VAZIFA 3: Smart Quiz (Aqlli viktorina)

**Real ssenariy**: Maktabga uchun matematika viktorinasi.

**Talab**:
- 5 ta savol (matematika misollari): "5 + 7 = ?", "12 * 8 = ?" ...
- Foydalanuvchi javob beradi
- Har to'g'ri javob — **10 ball**
- Agar 3 ta ketma-ket to'g'ri bersa — **bonus** (+5 ball ham)
- Agar 2 ta ketma-ket xato bersa — **darslarni qayta o'qishni** maslahat
- Yakuniy ball:
  - 50: A'lo (`Genius!`)
  - 40-49: Yaxshi
  - 25-39: O'rta
  - <25: Qoniqarli emas

**Texnologiyalar**: `for`, `if-else-if`, akkumulyator, ketma-ketlik tekshirish.

📁 [vazifa-3-smart-quiz/main.go](./vazifa-3-smart-quiz/main.go)

---

## VAZIFA 4: Smart Parking (Avtoturargoh)

**Real ssenariy**: Aeroport avtoturargohida narxni hisoblash.

**Talab**:
- Foydalanuvchi: kirish soati va chiqish soati (24-soatlik formatda, masalan 14)
- Sutkalar va soatlar farqi avtomatik hisoblansin (oddiy holat — bir kun ichida)
- Narxlar:
  - Birinchi 1 soat: **bepul**
  - 2-4 soat: har soati 5 000 so'm
  - 5-8 soat: har soati 4 000 so'm
  - 8+: butun kuni 30 000 so'm (sutkalik)
- **Kechki vaqt** (20:00-08:00): har soati 50% qimmat
- **Hafta oxiri** (input: true/false): +20% qo'shimcha
- Yakuniy narxni chiqaring

**Texnologiyalar**: bir nechta `if-else-if` zanjiri, mantiqiy operatorlar.

📁 [vazifa-4-smart-parking/main.go](./vazifa-4-smart-parking/main.go)

---

## VAZIFA 5: Number Guessing Game (Sonni topish o'yini)

**Real ssenariy**: Klassik 100 dan kichik son topish o'yini.

**Talab**:
- Dastur 1-100 oraligida tasodifiy son tanlaydi (`math/rand`)
- Foydalanuvchi son kiritadi
- Dastur "ko'p" yoki "kam" deydi
- **Max 7 ta urinish** (binary search optimum 7)
- Har urinish — hisoblanadi
- Topgan bo'lsa: nechta urinishda ekanini chiqaradi
- Topa olmasa: to'g'ri javobni ko'rsatadi
- Yakunda: "Yana o'ynaysizmi? (ha/yo'q)" — agar `ha` bo'lsa, qayta boshlanadi

**Texnologiyalar**: `for` cheksiz, `break`, `continue`, `math/rand`, mantiqiy.

📁 [vazifa-5-son-topish-oyini/main.go](./vazifa-5-son-topish-oyini/main.go)

---

## Qanday yechishni boshlash kerak?

1. **Skeleton fayllarni o'qing** — har bir vazifa-N/main.go da `// TODO:` qatorlari bor.
2. **Tushuncha qiling** — vazifani o'zingizning so'zlaringiz bilan tushuntirib oling.
3. **Pseudo-code yozing** — kodga qadar, qog'ozda algoritmni rejalashtiring.
4. **Bosqichma-bosqich yozing** — birinchi input olish, keyin logika, oxirida output.
5. **Test qiling** — chekka holatlarni (0, manfiy, juda katta) ham tekshiring.
6. **Agar tiqilib qolsangiz** — `loyiha-01-04-bank-atm/main.go` ga qarang, o'xshash pattern bor.

---

## Baholash mezonlari

| Mezon | Ball |
|-------|------|
| Dastur kompilyatsiyadan o'tadi (`go build`) | 20% |
| Asosiy logika to'g'ri ishlaydi | 40% |
| Chekka holatlar (xato input, 0, manfiy) qoplangan | 20% |
| Kod toza (comment'lar, mantiqiy nomlash) | 10% |
| Foydalanuvchiga qulay UX (xabarlar, formatlash) | 10% |

---

## Yordam

Tiqilib qolsangiz, savol bering:
- "Vazifa N da X nima qilish kerak?"
- "Bu xatoni qanday tuzataman?"
- "Bu pattern qaerda ishlatilgan edi?"

Men sizga **kod yozib bermayman** — yo'l-yo'riq beraman. O'zingiz yozsangizgina o'rganasiz!
