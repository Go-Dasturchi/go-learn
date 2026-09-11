# 03 — Types

## THEORY

Tur (type) — bu tortmaning **shakli**: unga qanday narsa solish mumkinligini belgilaydi. Suyuqlik uchun idishga non solib bo'lmaydi, non uchun tarelkaga suv quysangiz oqib ketadi — har bir narsaning o'z idishi bor. Go'da ham xuddi shunday: har bir o'zgaruvchining turi bor, va bu tur dastur ishlash davomida **o'zgarmaydi** (buni "statik tipdagi til" deyiladi). Bir marta "matn" uchun yaratilgan tortmaga keyin son solib bo'lmaydi.

**Eng ko'p ishlatiladigan asosiy turlar:**

- `string` — matn: `"Ali"`, `"Toshkent shahri"`
- `int` — butun son: `25`, `-7`, `0`
- `float64` — kasr (nuqtali) son: `36.6`, `-2.5`
- `bool` — mantiqiy qiymat: `true` yoki `false`

`:=` ishlatilganda Go turini qiymatning **ko'rinishidan** o'zi aniqlab oladi:

```go
ism := "Ali"        // qo'shtirnoq ichida — string
yosh := 25           // nuqtasiz butun son — int
boyi := 1.75         // nuqtali son — float64
talaba := true       // true/false — bool
```

**Sonlarning boshqa turlari ham bor.** `int` — eng ko'p ishlatiladigan butun son turi, lekin kerak bo'lganda aniq o'lchamdagi turlar ham mavjud: `int8`, `int16`, `int32`, `int64` (ishorali, ya'ni manfiy bo'lishi mumkin) va `uint8`, `uint16`, `uint32`, `uint64` (faqat musbat, "unsigned"). Kundalik dasturlashda deyarli har doim oddiy `int` yetarli — boshqalari faylning aniq formatlari, tarmoq protokollari kabi maxsus holatlarda kerak bo'ladi.

Ikkita qiziq alias (taxallus) turi bor:
- `byte` — aslida `uint8`ning boshqa nomi, bitta baytni ifodalaydi (0 dan 255 gacha).
- `rune` — aslida `int32`ning boshqa nomi, bitta Unicode belgisini (masalan, bitta harf yoki emoji) ifodalaydi. Bu — Strings darsida yana uchraydi.

**Turni tekshirish.** `fmt.Printf` bilan `%T` belgisidan foydalanib, o'zgaruvchining aynan qaysi turdaligini ekranga chiqarib ko'rish mumkin:

```go
son := 42
fmt.Printf("%T\n", son) // int
```

**Muhim qoida: turlar avtomatik aralashmaydi.** Boshqa ba'zi tillardan farqli o'laroq, Go `int` bilan `float64`ni, yoki `int32` bilan `int64`ni to'g'ridan-to'g'ri birga qo'shishga ruxsat bermaydi — bu haqda "Type Conversion" darsida batafsil gaplashamiz. Hozircha shuni bilib qo'yish kifoya: **bir turdagi qiymatlarni boshqa turdagisi bilan to'g'ridan-to'g'ri aralashtirib bo'lmaydi**, buning uchun avval ataylab aylantirish kerak bo'ladi.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	ism := "Malika"
	yosh := 22
	boyi := 1.68
	talaba := true

	fmt.Println(ism, yosh, boyi, talaba)

	fmt.Printf("%T\n", ism)    // string
	fmt.Printf("%T\n", yosh)   // int
	fmt.Printf("%T\n", boyi)   // float64
	fmt.Printf("%T\n", talaba) // bool
}
```

Natija:

```
Malika 22 1.68 true
string
int
float64
bool
```

## TASK

1. `temperature` nomli `float64` o'zgaruvchi yarating, qiymati `36.6` bo'lsin.
2. `isRaining` nomli `bool` o'zgaruvchi yarating, qiymati `false` bo'lsin.
3. Ikkalasini `fmt.Println` orqali, ikkita alohida qatorda ekranga chiqaring (avval `temperature`, keyin `isRaining`).

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Kasr sonlar uchun nuqta ishlatiladi, vergul emas: `36.6`, `36,6` emas.
2. `bool` turidagi qiymat qo'shtirnoqsiz yoziladi: `false`, `"false"` emas.
