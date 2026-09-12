# GO ZERO TO HERO

Terminal ichida ishlaydigan, Vim orqali kod yozib o'rganiladigan interaktiv Go kursi.

```
lesson tanlash → theory → task → Vim → test → pass/fail → hint → progress → next lesson
```

## Project structure

```
go-learn/
├── assets.go                 # lessons/ papkasini binary ichiga joylaydi (go:embed)
├── cmd/go-learn/main.go      # entry point
├── internal/
│   ├── cli/                  # menyu, dars oqimi, subcommandlar
│   ├── lesson/                # lesson.md parser va manifest loader
│   ├── runner/                # workspace fayllari va Vim integratsiyasi
│   ├── checker/               # kodni compile/test qilish (izolyatsiyalangan temp modul)
│   └── progress/              # ~/.go-learn/progress.json bilan ishlash
├── lessons/
│   ├── manifest.json          # darslar tartibi va kategoriyasi
│   └── 01-hello/ ...          # har bir dars: lesson.md, starter/, solution/, tests/
├── go.mod
├── Makefile
└── README.md
```

## How to install

Avval repo'ni klon qiling, so'ng ichidan o'rnating (`./cmd/go-learn` — nisbiy yo'l, shuning uchun repo papkasi ichida turib ishga tushirish kerak):

```bash
git clone https://github.com/Go-Dasturchi/go-learn.git
cd go-learn
go install ./cmd/go-learn
```

Bu buyruq `go-learn` binary'sini `$(go env GOPATH)/bin` (odatda `~/go/bin`) papkasiga qo'yadi — o'sha papka `$PATH`da bo'lsa, `go-learn` buyrug'i istalgan joydan ishlaydi.

### macOS / Linux / Windows

Yuqoridagi uchta buyruq (`git clone`, `cd`, `go install`) macOS, Linux va Windows'da bir xil ishlaydi — Git va Go ikkalasi ham cross-platform. Dastur Windows'da ham ishlaydi, lekin **eng silliq tajriba uchun avval WSL (Windows Subsystem for Linux) o'rnatib, keyin shu qadamlarni WSL terminali ichida bajarish tavsiya etiladi** — chunki dastur ichida terminal kengligini aniqlash uchun Unix-uslubidagi `stty` buyrug'i ishlatiladi, va Vim odatda Windows'da emas, Linux/macOS'da tayyor o'rnatilgan bo'ladi. WSL ichida bular hammasi muammosiz ishlaydi.

Yoki klon qilingan papka ichida lokal build:

```bash
make build
```

## How to run

```bash
go-learn          # asosiy menyu
go-learn start    # xuddi shu menyu
go-learn list     # barcha darslar ro'yxati (bajarilganlari ✓ bilan)
go-learn progress # progress dashboard
go-learn reset    # progress'ni tozalash
go-learn lesson 01-hello   # ma'lum bir darsni to'g'ridan-to'g'ri ochish
go-learn check    # joriy workspace papkasidagi main.go'ni tekshiradi
go-learn next     # birinchi tugallanmagan darsni to'g'ridan-to'g'ri ochadi
go-learn help     # yordam
```

`$VISUAL` yoki `$EDITOR` o'rnatilgan bo'lsa, o'sha editor ishlatiladi; aks holda standart `vim` ochiladi.

### Vim'dan chiqmasdan tekshirish va keyingi darsga o'tish

Dars ichida Vim ikkita oynada ochiladi (dars matni + kod), va odatda tekshirish uchun `:wqa` bilan Vim'dan butunlay chiqish kerak. Agar buni xohlamasangiz — Ctrl+T bilan terminal split oching (u avtomatik joriy dars papkasida ochiladi) va shunchaki:

```bash
go-learn check
```

Muvaffaqiyatli o'tsa, keyingi dars nomi ko'rsatiladi. Terminaldan chiqmasdan o'sha darsga o'tish uchun:

```bash
go-learn next
```

deb yozing — bu joriy papka nomidan (`~/.go-learn/workspace/<dars-id>`) darsni aniqlab, `main.go`ni sinovdan o'tkazadi va natijani darhol ko'rsatadi, Vim'dan chiqmasdan. Muvaffaqiyatli o'tsa, dars ham avtomatik "bajarilgan" deb belgilanadi.

## How to start the first lesson

```bash
go-learn
```

so'ng `1` (Fundamentals) → `1` (01-hello) ni tanlang, `ENTER` bosing — Vim ochiladi (agar `$EDITOR`/`$VISUAL` vim yoki nvim bo'lsa, ikkita ustunda: chapda dars matni, o'ngda kod). Topshiriqni yeching va **`:wqa`** (yoki `:xa`) bilan chiqing — ikkita oyna ochiq bo'lgani uchun oddiy `:wq`/`:q` faqat bitta oynani yopadi, ikkinchisi ochiq qolib ketadi va kod tekshiruvga umuman yetib bormaydi. Kod avtomatik tekshiriladi.

## How progress works

Har bir darsning kodi `~/.go-learn/workspace/<lesson-id>/main.go` faylida saqlanadi — darsni qayta ochsangiz, oldingi yozgan kodingiz saqlanib qoladi (qayta yozib tashlanmaydi).

Tugallangan darslar ro'yxati `~/.go-learn/progress.json` faylida saqlanadi:

```json
{
  "completed": ["01-hello", "02-variables"]
}
```

Dastur qayta ishga tushirilganda ham progress saqlanib qoladi.

## How to add a new lesson

1. `lessons/<id>/` papkasini yarating (masalan `lessons/07-constants/`).
2. Ichiga `lesson.md` yozing — quyidagi formatga rioya qiling:

   ```markdown
   # <id> — <Title>

   ## THEORY
   ...

   ## EXAMPLE
   ```go
   ...
   ```

   ## TASK
   ...

   ## HINTS
   1. kichik hint
   2. kuchliroq hint
   ```

3. `starter/main.go.tmpl` — TODO bilan boshlang'ich kod.
4. `solution/main.go.tmpl` — to'liq yechim (foydalanuvchiga hech qachon ko'rsatilmaydi).
5. `tests/main_test.go.tmpl` — `package main` ichida, starter/solution fayllaridagi funksiyalarni sinovdan o'tkazadigan testlar.

   > Fayllar ataylab `.tmpl` kengaytmasi bilan saqlanadi (`.go` emas) — shunda ular loyihaning o'z `go build`/`go vet`/`go test`iga aralashmaydi (ular hali TODO bilan to'liqmas bo'lgani uchun compile bo'lmasligi mumkin), lekin `go:embed` ularni baribir binary ichiga oddiy fayl sifatida joylaydi.

6. `lessons/manifest.json`ga yangi yozuv qo'shing: `{"id": "07-constants", "title": "Constants", "category": "Fundamentals"}`.
7. `make check` orqali build/vet/test'ni tekshiring.

## Tests result

```
$ make check
go fmt ./...
go vet ./...
go test ./...
ok  	go-learn/internal/checker
ok  	go-learn/internal/lesson
ok  	go-learn/internal/progress
```

## Loyiha holati

Jami **538 ta dars** tayyor va test qilingan, o'n bitta bosqichga bo'lingan holda (menyuda ko'rsatiladigan tartib bo'yicha). Kurs dasturi asl texnik topshiriqdagi (`go_zero_to_hero_claude_code_prompt.md`) LEVEL 1–9 ro'yxatiga (spesifikatsiyadagi barcha darajalar) to'g'ridan-to'g'ri mos qilib tuzilgan:

- **0. Problems** — 410 ta amaliy masala: 205 tasi Abramyan to'plamidan (Begin/Integer/Array/... bo'yicha bo'lingan), 205 tasi LeetCode uslubidagi masalalar (Easy/Medium/Hard bo'yicha bo'lingan). Har birida ikkita haqiqiy, masalaga xos hint bor (generik emas).
- **1. Fundamentals** — 18 ta dars: Hello World → Variables → Types → If/Else → For Loop → Functions → Constants → Strings → Numbers → Boolean → Operators → Type Conversion → Switch → Nested Loops → Multiple Return Values → Named Return Values → Defer → Scope.
- **2. Data Structures** — 12 ta dars: Arrays → Slices → Slice Append → Slice Indexing → Slice Copy → Maps → Structs → Nested Structs → Pointers → Methods → Interfaces → Custom Types.
- **3. Go Core** — 19 ta dars: Errors → Custom Errors → Panic and Recover → Closures → Higher-Order Functions → String Formatting → Sorting → Time and Duration → JSON → File I/O → Generics → Regular Expressions → Packages → Modules → Exported / Unexported → Error Handling → Encoding / Decoding → Command Line Arguments → Environment Variables.
- **4. Data Structures & Algorithms** — 12 ta dars: Recursion → Linked List → Stack → Queue → Binary Search → Bubble Sort → Merge Sort → Quick Sort → Binary Tree → Binary Search Tree → Hash Table → Graphs and BFS (klassik struktura/algoritmlar Go'da noldan yozib chiqiladi, `sort`/`container` paketlaridan foydalanmasdan).
- **5. Concurrency** — 13 ta dars: Goroutines → Channels → Buffered Channels → Select → WaitGroup → Mutex → Worker Pools → Context → sync.Once and Atomic → Race Conditions → Unbuffered Channels → Producer / Consumer → Concurrent HTTP Requests. Har bir mashq real vaqtga bog'liq bo'lmagan, deterministik natija beradigan qilib loyihalangan (indeksli natijalar, oldindan bekor qilingan context, bitta yuboruvchili kanallar) — bu esa `go test -race` ostida ham barqaror ishlashini ta'minlaydi.
- **6. Web Development** — 10 ta dars: HTTP Handlers → Query Parameters → JSON Responses → Reading JSON Request Bodies → HTTP Status Codes → Routing and Path Parameters → Middleware → Form Data → HTTP Client → Building a REST API (faqat standart `net/http` va `net/http/httptest` paketlari bilan, tashqi framework'siz; barcha mashqlar haqiqiy tarmoq/port ochmasdan, `httptest.NewRecorder`/`httptest.NewServer` orqali sinaladi).
- **7. Testing** — 6 ta dars: Unit Testing → Table Driven Tests → Test Helpers → Benchmarks → Fuzz Testing → Mocking Concepts.
- **8. HTTP & Backend** — 16 ta dars: HTTP Basics → HTTP Client → HTTP Server → net/http → Routing → Middleware → REST API → JSON API → HTTP Status Codes → Authentication → JWT → Password Hashing → Pagination → Filtering → Validation → Error Responses (JWT imzosi `crypto/hmac`/`crypto/sha256` bilan, parol xeshi `crypto/sha256` bilan — ikkalasi ham faqat standart kutubxona, real loyihada esa maxsus kutubxonalar — masalan JWT uchun tasdiqlangan kutubxona, parol uchun bcrypt/argon2 — ishlatilishi lozimligi darsda alohida ta'kidlangan).
- **9. Database** — 9 ta dars: SQL Basics → PostgreSQL → Database Connection → CRUD → Transactions → Indexes → Joins → Connection Pool → Repository Pattern. Bu bosqich boshqalardan farqli — har bir dars ikki qismli: (1) **haqiqiy terminal/psql bo'yicha to'liq amaliy qo'llanma** (PostgreSQL'ni brew/apt orqali o'rnatish va ishga tushirish, `psql` orqali ma'lumotlar bazasi/foydalanuvchi/jadval yaratish, barcha asosiy SQL so'rovlari — `SELECT`/`INSERT`/`UPDATE`/`DELETE`/`GROUP BY`/agregat funksiyalar/`JOIN`/`CREATE INDEX`/`EXPLAIN ANALYZE` — o'quvchi buni **o'z terminalida** bajaradi, checker buni tekshirmaydi) va (2) **standart Go bilan yechiladigan, avtomatik tekshiriladigan mashq**, tegishli tushunchani (masalan, transaction — oldindan tekshirib keyin o'zgartirish, connection pool — bufferlangan kanal, repository pattern — interfeys orqali) real bazasiz, sof Go orqali mustahkamlaydi. `database/sql` + `lib/pq` orqali Go'dan PostgreSQL'ga qanday ulanish, so'rov yuborish va natijani `Scan` qilish — barchasi nazariy qismda to'liq, kod misollari bilan tushuntirilgan.
- **10. Production Go** — 13 ta dars: Clean Architecture → Dependency Injection → Configuration → Logging → Graceful Shutdown → Docker → Docker Compose → Redis → Caching → gRPC → Kafka → Microservices → Observability. "Database" bosqichi bilan bir xil ikki qismli tuzilma: tashqi vosita talab qiladigan mavzularda (Docker, Docker Compose, Redis, gRPC, Kafka) — **haqiqiy terminal qo'llanmasi** (o'rnatish, ishga tushirish, real buyruqlar — `docker build`/`docker run`, `docker compose up`, `redis-cli`, `protoc`, Kafka konsol producer/consumer) berilgan, checker buni tekshirmaydi; har bir darsda esa, tegishli tushunchani (interfeys orqali qatlamlash, bog'liqlikni in'ektsiya qilish, TTL'li kesh, FIFO xabar navbati, tarmoq xatoligini o'rash) sof Go orqali mustahkamlaydigan, avtomatik tekshiriladigan mashq bor.

To'liq oqim (menyu → theory → task → Vim → test → hint → progress → next lesson) barcha 538 ta dars uchun ishlaydi. Yangi dars qo'shish yuqoridagi "How to add a new lesson" bo'yicha bosqichma-bosqich amalga oshiriladi.
