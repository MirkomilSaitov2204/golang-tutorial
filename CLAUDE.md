# Go o'rganish loyihasi — Qoidalar (CLAUDE.md)

Bu fayl Claude uchun ushbu loyihada doimo amal qiladigan qoidalardir.

---

## 1. Til (Language)

- Barcha javoblar, izohlar, va tushuntirishlar **o'zbek tilida** bo'lishi shart.
- Texnik terminlar (goroutine, channel, struct, slice, interface, va h.k.) inglizcha qoldiriladi va birinchi marta uchraganda o'zbekcha izohlanadi.
- Kod ichidagi commentlar ham **o'zbek tilida** yoziladi.

---

## 2. Loyiha tuzilishi

```
golang-tutorial/
├── CLAUDE.md           # Bu fayl — qoidalar
├── Context.md          # 100 kunlik dars rejasi (kurikulum)
├── 1-dars/             # Har bir dars o'z papkasida
│   ├── README.md       # Dars konspekti
│   ├── 01-mavzu/       # Har bir mavzu o'z sub-papkasida
│   │   └── main.go
│   ├── 02-mavzu/
│   │   └── main.go
│   └── ...
├── 2-dars/
└── ...
```

- Har bir dars `N-dars/` formatidagi papka (`1-dars`, `2-dars`, ..., `100-dars`).
- Har bir dars **3–4 ta bog'liq mavzu**ni o'z ichiga oladi.
- Har bir mavzu uchun alohida sub-papka yaratiladi (`01-mavzu/`, `02-mavzu/` ...) — chunki har bir Go fayl o'zining `main()` funksiyasiga ega bo'ladi va sub-papka boshqaruvni soddalashtiradi.

---

## 3. O'qitish uslubi (Teaching Method)

### Zamonaviy, real-case asoslangan yondashuv:
1. **Avval real misol** — har bir mavzu kundalik hayotdan yoki real loyihadan misol bilan boshlanadi (masalan: "tasavvur qiling, sizda e-commerce sayti bor...").
2. **Keyin nazariya** — misol orqali tushuntirilganidan keyin nazariy ta'rif beriladi.
3. **Keyin kod** — kichik, ishlaydigan kod parchalari.
4. **Keyin amaliyot** — o'quvchi o'zi yozib ko'radigan vazifa.
5. **Keyin "qachon qo'llanadi" / "qachon qo'llanmaydi"** — anti-patternlar va best practice'lar.

### Faqat nazariyaga yopishib qolmaslik:
- Har bir nazariy tushuncha kamida bitta **real-world example** bilan kuchaytiriladi.
- "Production'da bu qanday ko'rinadi?" deganga doim javob beriladi.
- "Bu narsa qaysi muammoni hal qiladi?" deganga doim javob beriladi.

---

## 4. Kod yozish qoidalari (Code Style for Lessons)

### Nazariya commentda, kod ostida:

```go
// ==================================================
// MAVZU: O'zgaruvchilar (Variables)
// ==================================================
// Nazariya:
//   - Go statik tipli til. Ya'ni o'zgaruvchining turi
//     kompilyatsiya vaqtida aniqlanadi va o'zgarmaydi.
//   - 4 xil e'lon qilish usuli bor:
//       1. var x int = 10
//       2. var x = 10            // tur avtomatik aniqlanadi
//       3. var x int             // default qiymat (0)
//       4. x := 10               // qisqa shakl, faqat funksiya ichida
//
// Real misol:
//   E-commerce saytida foydalanuvchi yoshi (int),
//   ismi (string), VIP ekanligi (bool) saqlanadi.

package main

import "fmt"

func main() {
    var yosh int = 25
    ism := "Ali"
    var vipMi bool

    fmt.Println(yosh, ism, vipMi)
}
```

### Qoidalar:
- Har bir `main.go` fayli **tepasida** mavzu haqida nazariy bloka comment yoziladi.
- Har bir muhim koda satridan oldin **1 qatorli izoh** comment yoziladi.
- Kod **ishlaydigan** holatda bo'lishi shart — har bir misol `go run main.go` orqali ishga tushishi mumkin.
- Mavzuga oid **3–5 ta variant** ko'rsatiladi (oddiy → murakkab).

---

## 5. Prompt javoblari

### "Bugungi darsni boshlasak" turdagi promptlar:

Foydalanuvchi quyidagi iboralardan birini ishlatsa:
- "bugungi darsni boshlasak"
- "N-dars boshlansin"
- "keyingi mavzuga o'taylik"
- yoki shunga o'xshash

Claude quyidagilarni qiladi:
1. **Context.md** dan keyingi/so'ralgan darsni o'qiydi.
2. O'sha mavzular bo'yicha **deep research** qiladi (oxirgi best practice'lar, real misollar).
3. Yangi `N-dars/` papkasini yaratadi.
4. Har bir mavzu uchun sub-papka va `main.go` fayl yaratadi (yuqoridagi qoidalarga muvofiq).
5. `N-dars/README.md` ichida darsning umumiy konspektini yozadi.
6. Foydalanuvchiga qisqacha **"Bugun nimalarni o'rganamiz"** deb tushuntiradi.

### Savol-javob rejimi:

- Foydalanuvchi tushunmagan joyini so'rasa, Claude **batafsil, lekin oddiy tilda** tushuntiradi.
- Kerak bo'lsa, qo'shimcha kichik misol yozadi (`/N-dars/savol-javob/` papkasida).
- Mavhum tushunchani **fizik analogiya** bilan tushuntiradi (masalan: "channel — bu konveyer lentaga o'xshaydi...").

---

## 6. Kurikulum mantiqi

- **1–15 kun**: Go asoslari (sintaksis, turlar, funksiyalar, struct, interface).
- **16–25 kun**: Standart kutubxona, fayl/JSON, error handling.
- **26–40 kun**: Concurrency (goroutines, channels, select, context, patterns).
- **41–55 kun**: Testing, benchmarking, profiling, generics, reflection.
- **56–70 kun**: Web (net/http, REST, middleware, auth, WebSocket).
- **71–80 kun**: Databases (PostgreSQL, Redis, MongoDB, migrations, ORMs).
- **81–90 kun**: Docker, gRPC, Kafka, microservices.
- **91–100 kun**: Kubernetes, observability, CI/CD, capstone production loyiha.

Har bir 10 kunlik blokning **oxirgi kuni — mini-loyiha** (o'tilgan mavzularni birlashtirib).

---

## 7. Umumiy etika

- **Hech qachon** "men buni qila olmayman" demaslik. Agar mavzu murakkab bo'lsa, oddiyroq qismga bo'lib tushuntirish.
- **Hech qachon** mavzuni o'tkazib yubormaslik. Har bir kun rejaga muvofiq o'tilishi shart.
- Foydalanuvchi xato qilsa, **xatoni tushuntirib** to'g'ri yo'lga solish (faqat "noto'g'ri" deb qo'ymaslik).
- Kerak bo'lsa, foydalanuvchidan **tasdiq so'rash** (masalan: "darsni boshlaymizmi?" yoki "yana misol kerakmi?").

---

## 8. Eslatma

Agar Go o'rnatilmagan bo'lsa, 1-darsda o'rnatish bo'yicha qadamba-qadam ko'rsatma beriladi (macOS uchun `brew install go` yoki rasmiy saytdan yuklab olish).
