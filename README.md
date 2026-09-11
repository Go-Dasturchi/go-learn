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

```bash
go install ./cmd/go-learn
```

Yoki lokal build:

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

Jami **440 ta dars** tayyor va test qilingan, uchta bosqichga bo'lingan holda (menyuda ko'rsatiladigan tartib bo'yicha):

- **0. Problems** — 410 ta amaliy masala: 205 tasi Abramyan to'plamidan (Begin/Integer/Array/... bo'yicha bo'lingan), 205 tasi LeetCode uslubidagi masalalar (Easy/Medium/Hard bo'yicha bo'lingan).
- **1. Fundamentals** — 18 ta dars: Hello World → Variables → Types → If/Else → For Loop → Functions → Constants → Strings → Numbers → Boolean → Operators → Type Conversion → Switch → Nested Loops → Multiple Return Values → Named Return Values → Defer → Scope.
- **2. Data Structures** — 12 ta dars: Arrays → Slices → Slice Append → Slice Indexing → Slice Copy → Maps → Structs → Nested Structs → Pointers → Methods → Interfaces → Custom Types.

To'liq oqim (menyu → theory → task → Vim → test → hint → progress → next lesson) barcha 440 ta dars uchun ishlaydi. Yangi dars qo'shish yuqoridagi "How to add a new lesson" bo'yicha bosqichma-bosqich amalga oshiriladi.
