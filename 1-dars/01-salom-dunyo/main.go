// =============================================================================
// MAVZU 1: GO TILIGA KIRISH VA BIRINCHI DASTUR
// =============================================================================
//
// REAL HAYOTDAN MISOL:
//   Tasavvur qiling, siz Uber, Netflix, Twitch, Docker, Kubernetes, yoki
//   Cloudflare'da ishlamoqchisiz. Bu kompaniyalarning hammasi backend
//   tizimlarini katta qismini Go'da yozadi. Nega? Chunki Go:
//     - JUDA TEZ kompilyatsiya qilinadi (bir necha soniya)
//     - Native machine code'ga aylanadi (Python yoki JS kabi interpret emas)
//     - Concurrency (parallel ishlash) tilning O'ZIDA bor (goroutine'lar)
//     - Sintaksisi sodda — 1 oyda asoslarini o'rganish mumkin
//     - Kompilyatsiya natijasi BITTA executable fayl — deploy oson
//
// GO QAERDA ISHLATILADI?
//   - Backend API'lar (REST, gRPC)
//   - Microservice'lar
//   - DevOps tools (Docker, Kubernetes, Terraform — hammasi Go!)
//   - CLI dasturlar
//   - Network servislari (proxy, load balancer)
//   - Cloud infrastruktura
//
// NAZARIYA — BIRINCHI DASTURNING ANATOMIYASI:
//
//   1. `package main`
//      Har bir Go fayl biror paketga tegishli bo'lishi shart.
//      `main` paketi — bu dasturning kirish nuqtasi (entry point).
//      Agar dastur ishga tushadigan executable bo'lishi kerak bo'lsa,
//      `package main` bo'lishi shart.
//
//   2. `import "fmt"`
//      Boshqa paketlardan kod ishlatish. `fmt` — "format" qisqartmasi,
//      standart kutubxonadagi I/O (input/output) paketi.
//
//   3. `func main()`
//      Bu funksiya dastur boshlanganda AVTOMATIK chaqiriladi.
//      Har bir `package main` da AYNAN BITTA `main()` funksiya bo'lishi shart.
//
// ISHGA TUSHIRISH:
//   Bu papkada terminal ochib:
//     go run main.go     -> bir martalik ishga tushirish
//     go build main.go   -> executable yaratadi (main yoki main.exe)
//     ./main             -> shu executable'ni ishga tushiradi
// =============================================================================

package main

// fmt — formatlangan I/O paketi
import "fmt"

func main() {
	// Println — yangi qatorga chiqaradi
	fmt.Println("Salom, Dunyo!")
	fmt.Println("Salom, Go!")

	// Print — yangi qatorsiz chiqaradi
	fmt.Print("Bu bitta")
	fmt.Print(" qatorda yoziladi.\n") // \n — yangi qator belgisi

	// Println avtomatik probel qo'yib bir nechta argumentni chiqaradi
	fmt.Println("Mening ismim:", "Ali", "Yoshim:", 25)
}

// =============================================================================
// AMALIY VAZIFA:
//   1. Yuqoridagi kodni o'zgartiring va o'z ismingizni chiqaring.
//   2. 3 ta sevimli kompaniyangizni Println orqali chiqaring.
//   3. `fmt.Println()` ichida hech narsa yozmasangiz nima bo'ladi? Sinab ko'ring.
// =============================================================================
