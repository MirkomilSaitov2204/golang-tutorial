# 100 Kunlik Go O'rganish Rejasi (Context.md)

Bu reja **0 dan production darajaga** olib boradi. Har bir kun 3–4 ta bog'liq mavzuni o'z ichiga oladi. Kerakli amaliyot vaqti — kuniga **1.5–3 soat**.

---

## I BOSQICH — Go Asoslari (1–15 kun)

### 1-dars — Go bilan birinchi tanishuv
- Go nima, nima uchun yaratilgan, qaerda ishlatiladi
- Go o'rnatish va `GOPATH`, `GOROOT`, `go.mod` tushunchalari
- "Salom Dunyo" dasturi va `package`, `import`, `func main()`
- O'zgaruvchilar (`var`, `:=`, `const`) va asosiy ma'lumot turlari (`int`, `float64`, `string`, `bool`)

### 2-dars — Asosiy sintaksis
- Arifmetik, mantiqiy va taqqoslash operatorlari
- Tur konvertatsiyasi (type conversion)
- Input/Output: `fmt.Print`, `Println`, `Printf`, `Scan`, `Scanln`
- `strconv` paketi bilan satr ↔ son konvertatsiyasi

### 3-dars — Shartli operatorlar
- `if`, `else if`, `else`
- `switch` (fallthrough, multi-case, type switch)
- `if` ichidagi qisqa e'lon (`if x := f(); x > 0`)
- Real misol: kalkulyator va foydalanuvchi yoshini tekshirish

### 4-dars — Sikllar
- `for` (3 ta shakli: classic, while-like, infinite)
- `range` bilan iteratsiya
- `break`, `continue`, label'lar (yorliqlar)
- Real misol: Fibonacci, FizzBuzz, prime number tekshirish

### 5-dars — Funksiyalar
- E'lon va chaqirish, parametrlar, return qiymat
- Multiple return (`return x, y, err`)
- Named return values
- Variadic parametrlar (`func(nums ...int)`)

### 6-dars — Funksiyalarning chuqurroq mavzulari
- Closure (yopilmalar) — counter misoli
- Recursion (rekursiya) — factorial, tree traversal
- Funksiya — birinchi sinf qiymat (first-class citizen)
- `defer` va uning stack tartibi

### 7-dars — Massivlar va Slice'lar
- Array (fixed size) — kamdan-kam ishlatiladi
- Slice — dinamik massiv, `len`, `cap`, `make`
- `append`, `copy`, slice'ning ichki tuzilishi (ptr/len/cap)
- Real misol: log buffer va to'p (stack) implementatsiyasi

### 8-dars — Map va String'lar
- `map[K]V` — yaratish, qo'shish, o'chirish, mavjudligini tekshirish
- Map iteratsiyasi tartibsizligi
- String — `byte` va `rune` farqi, UTF-8
- `strings` paketi: `Split`, `Join`, `Contains`, `Replace`, `ToLower`

### 9-dars — Pointer'lar
- Pointer nima va xotira modeli
- `&` (address-of) va `*` (dereference)
- Funksiyaga pointer berish (value vs reference semantics)
- `new` vs `make` farqi

### 10-dars — Struct'lar
- Struct e'loni va inicializatsiya (3 xil usuli)
- Embedded struct (kompozitsiya)
- Anonymous struct
- Real misol: `User`, `Order`, `Product` modellari

### 11-dars — Metodlar
- Value receiver vs Pointer receiver
- Qachon qaysi birini tanlash
- Methods on built-in types (custom types orqali)
- Real misol: `BankAccount` struct va uning metodlari

### 12-dars — Interface'lar (1-qism)
- Interface nima va duck typing
- Implicit implementation
- `Stringer`, `Error` interface'lari
- Real misol: bir nechta `Notifier` (Email, SMS, Telegram)

### 13-dars — Interface'lar (2-qism)
- Type assertion (`v.(T)`)
- Type switch (`switch v := x.(type)`)
- Empty interface (`interface{}`) va `any`
- Interface composition (`io.ReadWriter = io.Reader + io.Writer`)

### 14-dars — Error handling
- `error` interface'i
- `errors.New`, `fmt.Errorf`, `%w` bilan wrap
- `errors.Is`, `errors.As`
- Custom error types (sentinel errors, typed errors)

### 15-dars — Panic, Recover, va Best Practice'lar
- `panic` qachon ishlatiladi (kamdan-kam)
- `recover` orqali panic'ni ushlash
- `defer` + `recover` pattern
- **Mini-loyiha**: Komand qatori TODO ilovasi (CRUD, fayl ichiga saqlash)

---

## II BOSQICH — Standart kutubxona va Modullar (16–25 kun)

### 16-dars — Go modullari
- `go mod init`, `go mod tidy`, `go.sum`
- Semantic versioning, `replace` direktivasi
- Vendor directory, private modullar
- Real misol: o'z mini kutubxonangizni yaratish

### 17-dars — Paketlar va vidimost
- Eksport qoidalari (katta harf = public)
- Paket strukturasi (cmd/, internal/, pkg/)
- `init()` funksiyasi
- Circular import muammosi va yechimi

### 18-dars — Fayl bilan ishlash
- `os.Open`, `os.Create`, `os.ReadFile`, `os.WriteFile`
- `bufio.Scanner`, `bufio.Writer`
- `io.Reader`, `io.Writer` interface'lari (eng muhim!)
- Real misol: log fayl yozish va o'qish

### 19-dars — JSON
- `encoding/json` — `Marshal`, `Unmarshal`
- Struct tag'lar (`json:"name,omitempty"`)
- Custom `MarshalJSON`/`UnmarshalJSON`
- Real misol: REST API javobini parse qilish

### 20-dars — Boshqa formatlar
- YAML (`gopkg.in/yaml.v3`)
- TOML (`github.com/BurntSushi/toml`)
- CSV (`encoding/csv`)
- Real misol: konfiguratsiya faylini turli formatlarda o'qish

### 21-dars — `time` va `strconv`
- `time.Now`, `time.Parse`, `time.Format` (Go'ning g'aroyib layout sistemasi)
- Timezone, Duration, Ticker, Timer
- `strconv.Itoa`, `Atoi`, `ParseFloat`, `FormatInt`
- Real misol: smena hisoblagich va shartnoma muddati

### 22-dars — `regexp`, `sort`, `math`
- Regular expression asoslari
- `sort.Slice`, `sort.SliceStable`, custom interface
- `math`, `math/rand` (yangi `math/rand/v2`)
- Real misol: email validator va leaderboard sortlash

### 23-dars — Komand qatori (CLI) ilovalar
- `os.Args`, `flag` paketi
- `cobra` kutubxonasi bilan tanishuv
- Exit code va signal handling
- Real misol: o'z `git`-ga o'xshash CLI

### 24-dars — Loglash asoslari
- `log` paketi (eski)
- `log/slog` (yangi, structured logging — Go 1.21+)
- Log darajalari va JSON output
- Real misol: ilovaning kuzatuv tizimi

### 25-dars — Environment va konfiguratsiya
- `os.Getenv`, `os.Setenv`
- `.env` fayllar (`godotenv`)
- 12-Factor App tamoyillari
- **Mini-loyiha**: konfiguratsiyali markdown-to-HTML konvertor

---

## III BOSQICH — Concurrency (26–40 kun)

### 26-dars — Goroutine'larga kirish
- Goroutine nima, `go` keyword
- OS thread vs goroutine
- `runtime.GOMAXPROCS`, scheduler asoslari
- Real misol: parallel HTTP so'rovlar

### 27-dars — Channel'lar (1-qism)
- Unbuffered channel — sinxronizatsiya
- `ch <- v` va `<-ch` operatorlari
- Channel'ning yopilishi (`close`)
- Real misol: producer-consumer

### 28-dars — Channel'lar (2-qism)
- Buffered channel
- Directional channel (`chan<-`, `<-chan`)
- `for range` channel ustida
- `select` operatori asoslari

### 29-dars — `select` va timeout pattern'lar
- `select` multiple channel
- `default` case
- `time.After` va timeout
- Real misol: HTTP request bilan timeout

### 30-dars — Worker pool pattern
- Job queue va worker'lar
- Graceful shutdown
- `sync.WaitGroup`
- Real misol: rasm konvertor (1000 rasm — 10 worker)

### 31-dars — Mutex va Race condition
- Race condition nima va qanday yuzaga keladi
- `sync.Mutex`, `sync.RWMutex`
- Go race detector (`go run -race`)
- Real misol: bank hisob nomidagi (concurrent transfer)

### 32-dars — `sync` paketining boshqa qismlari
- `sync.Once` — singleton pattern
- `sync.Pool` — obyektlarni qayta ishlatish
- `sync.Map` — concurrent map (qachon ishlatish kerak)
- `atomic` paketi

### 33-dars — `context` paketi
- Context nima va nima uchun kerak
- `context.Background`, `WithCancel`, `WithTimeout`, `WithValue`
- Context'ni HTTP handler'da ishlatish
- Real misol: long-running query'ni bekor qilish

### 34-dars — Pipeline pattern
- Stage'lar va channel'lar zanjiri
- Cancellation propagation
- Backpressure
- Real misol: log fayl → parse → filter → save zanjiri

### 35-dars — Fan-in va Fan-out
- Fan-out: ko'p worker bir sourceni qayta ishlaydi
- Fan-in: ko'p channelni bittaga birlashtirish
- Real misol: web scraper (parallel sahifa yuklash)

### 36-dars — Concurrency anti-pattern'lar
- Goroutine leak
- Channel deadlock
- "Go faster" trap (har joyda go ishlatish)
- Race detection va testing

### 37-dars — `errgroup` va `singleflight`
- `golang.org/x/sync/errgroup` — birinchi xato bilan to'xtash
- `singleflight` — cache stampede oldini olish
- Real misol: bir nechta API'dan parallel olish

### 38-dars — Channel pattern'lar darajalari
- Or-channel pattern
- Tee channel
- Bridge channel
- Rate limiter

### 39-dars — Concurrency primitiv'lar tahlil
- "Share memory by communicating" falsafasi
- Qachon mutex, qachon channel ishlatish
- `runtime.Gosched`, `runtime.Goexit`

### 40-dars — **Mini-loyiha**: Concurrent web crawler
- URL'larni parallel ko'rib chiqish
- Visited URL'larni cache'lash
- Cancellation va graceful shutdown
- Rate limiting

---

## IV BOSQICH — Testing, Generics va Advanced (41–55 kun)

### 41-dars — Testing asoslari
- `testing` paketi, `go test`
- Naming conventions (`TestXxx`)
- `t.Errorf`, `t.Fatalf`, `t.Run` (subtests)
- Real misol: kalkulyator funksiyalarini test qilish

### 42-dars — Table-driven testing
- Test case'lar slice'i
- Subtest'lar bilan `t.Run`
- Helper funksiyalar (`t.Helper`)
- Setup va teardown

### 43-dars — Mocking va Test doubles
- Interface orqali mocking
- `gomock`, `mockery`
- Stub, fake, mock farqi
- Dependency injection asoslari

### 44-dars — Benchmark va Profiling
- `go test -bench`
- `b.N`, `b.ResetTimer`, `b.ReportAllocs`
- `pprof` (CPU, memory, goroutine)
- Real misol: ikki algoritm taqqoslash

### 45-dars — Code coverage va Fuzz testing
- `go test -cover -coverprofile`
- `go tool cover -html`
- Fuzz testing (Go 1.18+)
- Property-based testing

### 46-dars — Integration testing
- `httptest` paketi
- Testcontainers (Postgres, Redis test environment)
- Test fixtures
- CI'da test qanday ishlaydi

### 47-dars — Generics (Go 1.18+)
- Type parameters: `[T any]`
- Constraint'lar (`comparable`, `~int`, custom)
- Generic funksiyalar va tip'lar
- Real misol: generic stack, queue, set

### 48-dars — Generics — chuqurroq
- Type sets, union constraints
- Generic interface'lar
- Qachon generics ishlatish KERAK EMAS
- Standard `slices`, `maps`, `cmp` paketlari

### 49-dars — Reflection
- `reflect` paketi
- `reflect.TypeOf`, `reflect.ValueOf`
- Struct tag'larni o'qish
- Real misol: ORM/JSON marshaller minimal versiyasi

### 50-dars — Code generation
- `go generate`
- `stringer`, `mockgen` misoli
- AST manipulyatsiyasi asoslari
- Qachon generation, qachon reflection

### 51-dars — `unsafe`, `cgo`, `syscall`
- `unsafe.Pointer` — qachon kerak (kamdan-kam)
- `cgo` orqali C kutubxonasini chaqirish
- Plain `syscall` misoli
- Eslatma: bu vositalardan ehtiyot bo'lish kerak

### 52-dars — Linter'lar va kod sifati
- `gofmt`, `goimports`
- `go vet`, `staticcheck`, `golangci-lint`
- Pre-commit hooks
- Effective Go va Code Review Comments

### 53-dars — Loyiha tuzilishi (Project Layout)
- `cmd/`, `internal/`, `pkg/`, `api/` papkalari
- Clean Architecture vs Hexagonal vs Onion
- Domain-driven design (DDD) asoslari
- Real misol: o'rta o'lchamdagi loyihaning rejasi

### 54-dars — Dependency injection
- Manual DI
- `wire` (compile-time DI)
- `fx` (runtime DI)
- Trade-off'lar

### 55-dars — **Mini-loyiha**: TDD bilan REST API
- Test'larni avval yozish
- Refactoring + green tests
- 80%+ coverage
- CI integratsiyasi

---

## V BOSQICH — Web Development (56–70 kun)

### 56-dars — `net/http` asoslari
- `http.HandleFunc`, `http.ListenAndServe`
- `Request`, `ResponseWriter` strukturalari
- Path va query parametrlar
- Real misol: oddiy "echo" server

### 57-dars — Routing va router'lar
- Standart `http.ServeMux` (Go 1.22+ — chiroyli routing)
- `chi`, `gin`, `echo`, `fiber` taqqoslash
- Path parameters, route groups
- Real misol: REST CRUD endpoint'lar

### 58-dars — Middleware pattern
- Middleware nima va qanday ishlaydi
- Request ID, logging, recovery middleware
- Middleware chain
- Real misol: o'z auth middleware'ingiz

### 59-dars — REST API dizayni
- HTTP metodlar va status code'lar
- Resource naming, versioning
- Pagination, filtering, sorting
- HATEOAS asoslari (ixtiyoriy)

### 60-dars — JSON, validation, error response
- `json.NewDecoder`/`Encoder`
- `go-playground/validator` bilan validatsiya
- Standardlashtirilgan error response
- Real misol: signup endpoint validatsiya bilan

### 61-dars — Authentication
- Session-based auth (cookie + Redis)
- Stateless JWT (header, payload, signature)
- Bcrypt bilan parol xeshlash
- Real misol: login + refresh token

### 62-dars — Authorization
- RBAC (role-based) vs ABAC (attribute-based)
- Casbin kutubxonasi
- Middleware orqali rolni tekshirish
- Real misol: admin/user huquqlari

### 63-dars — WebSocket
- `gorilla/websocket` yoki `nhooyr/websocket`
- Ping/pong, connection lifecycle
- Broadcast pattern
- Real misol: chat servisi

### 64-dars — Server-Sent Events (SSE)
- HTTP'da long-lived connection
- WebSocket'dan farqi
- Real misol: realtime notification

### 65-dars — Template'lar
- `html/template` (XSS-safe)
- Layout va component'lar
- HTMX bilan integratsiya
- Real misol: SSR sahifa

### 66-dars — Static fayllar va Upload
- `http.FileServer`
- Multipart form, file upload
- Image resize, S3'ga yuklash
- Real misol: avatar upload

### 67-dars — CORS, CSRF, Security headers
- CORS preflight, qoidalari
- CSRF token
- Helmet-like security headers (HSTS, CSP)
- Rate limiting (`golang.org/x/time/rate`)

### 68-dars — `slog`, structured logging, OpenTelemetry intro
- `log/slog` chuqurroq
- Trace ID, request ID propagation
- OpenTelemetry SDK ulashi
- Log → Loki/Elasticsearch

### 69-dars — Konfiguratsiya: `viper`, env, flags
- Viper bilan multi-source config
- 12-Factor compliance
- Secret management asoslari (Vault, AWS SM)

### 70-dars — **Mini-loyiha**: To'liq REST API
- Auth + RBAC + pagination + validation + logging
- Swagger/OpenAPI
- Test'lar bilan
- Docker'siz versiya (Docker keyingi bosqichda)

---

## VI BOSQICH — Ma'lumotlar bazasi (71–80 kun)

### 71-dars — `database/sql` asoslari
- Driver'lar, connection
- `db.Query`, `QueryRow`, `Exec`
- `Rows.Scan`, `sql.Null*` turlari
- Real misol: oddiy SQLite CRUD

### 72-dars — PostgreSQL bilan `pgx`
- `pgx/v5` driver, connection pool
- `pgxpool`
- PostgreSQL specific turlar (JSONB, array, UUID)
- Real misol: foydalanuvchilar jadvali

### 73-dars — Tranzaksiyalar va Connection management
- `BeginTx`, `Commit`, `Rollback`
- Isolation level'lar
- Connection pool sozlamalari
- Deadlock va retry pattern

### 74-dars — Migration'lar
- `golang-migrate`, `goose`, `atlas`
- Forward va rollback migration'lar
- Versioning strategiyalari
- Real misol: jadval yaratish, kolonna qo'shish

### 75-dars — ORM lar va Query Builder'lar
- GORM (full ORM) — pros/cons
- Ent (graph ORM, Facebook)
- SQLC (SQL → Go code) — nega ko'pchilik buni tanlaydi
- Squirrel (query builder)

### 76-dars — SQLC bilan chuqurroq
- `sqlc.yaml`, query yozish
- Type-safe code generation
- Migratsiyalar bilan ulash
- Real misol: e-commerce schema

### 77-dars — Redis bilan ishlash
- `go-redis/v9`
- Key-value, list, hash, set, sorted set
- Pub/sub
- Real misol: cache layer va session store

### 78-dars — MongoDB
- `mongo-go-driver`
- BSON, kolleksiyalar, query
- Index'lar va aggregation pipeline
- Real misol: log saqlash

### 79-dars — Database pattern'lar
- Repository pattern
- Unit of Work
- N+1 query muammosi va yechimi
- Read replica, sharding asoslari

### 80-dars — **Mini-loyiha**: DB-backed API
- 70-darsdagi loyihani PostgreSQL'ga ko'chirish
- Redis cache qatlamini qo'shish
- Migration'lar
- Repository pattern

---

## VII BOSQICH — Docker, gRPC, Kafka (81–90 kun)

### 81-dars — Docker asoslari
- Container vs VM
- Image, container, registry
- Asosiy Docker komandalari
- Volume, network

### 82-dars — Go uchun Dockerfile
- Multi-stage build (scratch image bilan)
- Image size optimizatsiyasi
- Build cache strategiyasi
- Distroless image'lar

### 83-dars — Docker Compose
- `docker-compose.yml` strukturasi
- Service'lar orasidagi network
- Volume va environment
- Real misol: Go app + Postgres + Redis stack

### 84-dars — gRPC kirish
- Protocol Buffers tili
- `.proto` fayl yozish
- `protoc` va `protoc-gen-go-grpc`
- gRPC vs REST

### 85-dars — gRPC server va client
- Unary RPC
- Generated stub'lar
- Error handling (status code'lar)
- Real misol: User servisi

### 86-dars — gRPC streaming
- Server streaming
- Client streaming
- Bidirectional streaming
- Real misol: chat servisi gRPC orqali

### 87-dars — gRPC pattern'lar
- Interceptor'lar (middleware ekvivalenti)
- Auth, logging, retry interceptor
- gRPC + TLS
- gRPC-Gateway (gRPC → REST proxy)

### 88-dars — Kafka kirish
- Event-driven architecture
- Kafka asoslari: broker, topic, partition, consumer group
- Producer va consumer
- Real misol: order events

### 89-dars — Kafka Go bilan
- `franz-go` (eng tezligi) yoki `segmentio/kafka-go`
- Consumer group, offset management
- Idempotent producer
- Real misol: e-commerce order pipeline

### 90-dars — **Mini-loyiha**: Microservices
- User-service (REST)
- Order-service (gRPC)
- Notification-service (Kafka consumer)
- Hammasi Docker Compose'da

---

## VIII BOSQICH — Kubernetes va Production (91–100 kun)

### 91-dars — Kubernetes asoslari
- Pod, Deployment, Service, Namespace
- `kubectl` asosiy komandalari
- Lokal cluster: `kind`, `minikube`
- Manifest fayllar (YAML)

### 92-dars — Go servisini K8s'ga deploy
- Image push (Docker Hub yoki GHCR)
- Deployment, Service, Ingress
- Replica'lar va rolling update
- Real misol: 90-darsdagi loyihani deploy qilish

### 93-dars — ConfigMap va Secret
- Environment'ga config inject
- Secret management (Sealed Secrets, External Secrets)
- Resource limits, requests
- Liveness va readiness probe

### 94-dars — Helm va deployment patternlar
- Helm chart yaratish
- Values va template'lar
- Blue-green, canary deployment'lar
- ArgoCD (GitOps) asoslari

### 95-dars — Observability — 1
- Prometheus metrics: counter, gauge, histogram
- `prometheus/client_golang`
- Grafana dashboard
- Real misol: HTTP latency monitoring

### 96-dars — Observability — 2
- Distributed tracing: OpenTelemetry + Jaeger
- Trace propagation cross-services
- Log aggregation (Loki/ELK)
- SLI, SLO, SLA

### 97-dars — CI/CD
- GitHub Actions workflow
- Test → build → push image → deploy
- Branch strategy (GitFlow vs Trunk-based)
- Security scanning (trivy, gosec)

### 98-dars — Security best practices
- OWASP Top 10 Go uchun
- Static analysis: `gosec`, `govulncheck`
- Dependency scanning
- Secrets in code anti-pattern

### 99-dars — Performance optimization
- Profiling production'da (`pprof` endpoint)
- Allocation tracking, escape analysis
- Database query optimization
- Caching strategiyalari (multi-layer)

### 100-dars — **CAPSTONE LOYIHA** — Production-ready microservices system
- E-commerce yoki ijtimoiy tarmoq
- 3+ microservice (REST + gRPC + Kafka)
- PostgreSQL + Redis + S3 storage
- K8s'ga deploy
- Prometheus + Grafana + Jaeger
- CI/CD pipeline
- Loyihani GitHub'ga upload
