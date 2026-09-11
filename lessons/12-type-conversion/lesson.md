# 12 — Type Conversion

## THEORY

Turli mamlakatlarning pulini tasavvur qiling — dollar bilan so'mni to'g'ridan-to'g'ri qo'shib bo'lmaydi, avval birini boshqasiga **ayirboshlash** kerak (aniq kursga ko'ra). Go'da turlar ham xuddi shunday: ko'plab tillardan farqli o'laroq, Go turlarni **avtomatik** aylantirmaydi. Masalan, `int` bilan `float64`ni to'g'ridan-to'g'ri qo'shib bo'lmaydi:

```go
var a int = 5
var b float64 = 2.5
// a + b // XATO: compile bo'lmaydi!
```

Buning uchun turni **aniq** (explicit) aylantirish kerak — xuddi ayirboshlash punktida pulni almashtirgandek:

```go
natija := float64(a) + b   // 7.5
```

Sintaksis: `Tur(qiymat)`. Eng ko'p ishlatiladigan aylantirishlar:

```go
float64(butunSon)   // int → float64
int(kasrSon)         // float64 → int (kasr qismi tashlanadi, yaxlitlanmaydi!)
```

**Sonlar turlari orasida ham aylantirish kerak bo'ladi.** Bu qoida faqat `int`/`float64` orasida emas, `int` oilasining o'zi ichida ham amal qiladi — masalan `int32` bilan `int64`ni ham to'g'ridan-to'g'ri qo'shib bo'lmaydi, avval biri ikkinchisiga aylantirilishi kerak:

```go
var kichik int32 = 100
var katta int64 = 500
// kichik + katta // XATO
natija := int64(kichik) + katta // TO'G'RI
```

**Diqqat: "torroq" turga aylantirishda ma'lumot yo'qolishi mumkin.** Masalan, katta qiymatli `int64`ni kichikroq `int8`ga aylantirsangiz (`int8` faqat -128 dan 127 gacha sig'diradi), qiymat "kesilib" chiqishi mumkin — bu haqiqiy dasturlarda uchraydigan xatolarning sabablaridan biri, shuning uchun turlarni aylantirganda qanday turga o'tayotganingizga har doim e'tibor bering.

**Son bilan matn (string) orasidagi aylantirish — `strconv` paketi.** `float64(...)` yoki `int(...)` kabi to'g'ridan-to'g'ri aylantirish faqat sonlar orasida ishlaydi; sonni matnga yoki matnni songa aylantirish uchun alohida `strconv` (string conversion) paketi kerak:

```go
import "strconv"

matn := strconv.Itoa(42)           // int → string: "42"
son, err := strconv.Atoi("42")      // string → int: 42, va xato
```

`strconv.Atoi` ikkita qiymat qaytaradi — natija va xato ("Multiple Return Values" darsida bu naqshni ko'rgan edingiz). Xatoni e'tiborsiz qoldirib bo'lmaydi (yoki hech bo'lmasa, ataylab `_` bilan e'tiborsiz qoldirilganini ko'rsatish kerak), chunki berilgan matn haqiqiy son bo'lmasligi mumkin (masalan `"abc"` — bunda `err` `nil` bo'lmaydi, ya'ni xato haqiqatan yuz bergan bo'ladi).

Kasr sonlar bilan ishlash uchun mos funksiyalar ham bor:

```go
son, err := strconv.ParseFloat("3.14", 64)  // string → float64
matn := strconv.FormatFloat(3.14, 'f', 2, 64) // float64 → string, 2 xonali kasr bilan
```

**Rune va byte orasidagi aylantirish.** Strings va For Loop darslaridagi `rune` (bitta Unicode belgi) va `byte` turlarini ham son sifatida aylantirish mumkin — masalan, bitta harfning "raqamli" (kod) qiymatini olish yoki, aksincha, raqamdan harf hosil qilish uchun:

```go
harf := 'A'          // bu aslida rune (int32) turida, qiymati 65
son := int(harf)      // 65
qayta := rune(66)     // 'B'
fmt.Println(string(qayta)) // "B"
```

## EXAMPLE

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	yosh := 25
	matn := strconv.Itoa(yosh)
	fmt.Println("Yoshim: " + matn)

	son, _ := strconv.Atoi("100")
	fmt.Println(son + 1)

	var a int32 = 10
	var b int64 = 20
	fmt.Println(int64(a) + b) // 30 — avval turlarni tenglashtirdik

	kasr, _ := strconv.ParseFloat("9.5", 64)
	fmt.Println(kasr + 0.5) // 10
}
```

Natija:

```
Yoshim: 25
101
30
10
```

## TASK

`son` nomli `int` o'zgaruvchi (`42`) berilgan. Uni `strconv.Itoa` yordamida stringga aylantirib, `"Son: "` so'ziga qo'shib (`+` orqali) ekranga chiqaring (natija: `Son: 42`).

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `strconv.Itoa(son)` — `int`ni `string`ga aylantiradi.
2. `fmt.Println("Son: " + strconv.Itoa(son))` — ikkita stringni `+` bilan birlashtirish mumkin.
