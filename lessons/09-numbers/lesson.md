# 09 — Numbers

## THEORY

Kundalik hayotda ikki xil son bilan ishlaymiz: **butun sanoq** (nechta olma bor — 1, 2, 3...) va **o'lchov** (necha kilogramm — 1.5, 2.75...). Pulni tangalarda sanaysiz (butun), lekin suyuqlikni o'lchagichda o'lchaysiz (aniq emas, taxminiy, kasrli). Go'da ham xuddi shunday ikki katta guruh bor:

- `int` — butun son (kompyuter arxitekturasiga qarab 32 yoki 64 bit): `5`, `-3`, `0`
- `int8`, `int16`, `int32`, `int64` — aniq o'lchamli butun sonlar (maxsus holatlar uchun)
- `float64` — kasr son (standart tanlov): `3.14`, `-0.5`
- `float32` — kamroq aniqlikdagi kasr son

Bu turlar bir-biriga **avtomatik aylanmaydi** (bu haqda "Type Conversion" darsida batafsil gaplashamiz) — `int` bilan `float64`ni to'g'ridan-to'g'ri qo'shib bo'lmaydi.

**Odatiy arifmetik amallar:** `+` (qo'shish), `-` (ayirish), `*` (ko'paytirish), `/` (bo'lish), `%` (qoldiq).

**Eng muhim nozik joy: butun sonni butun songa bo'lish.** Ikkita `int`ni bo'lganda, natija ham `int` bo'ladi va **kasr qismi tashlab yuboriladi** (yaxlitlanmaydi, shunchaki kesib tashlanadi):

```go
fmt.Println(7 / 2)   // 3, 3.5 emas!
fmt.Println(7 % 2)   // 1 (qoldiq)
```

Buni pul almashtirish kassasiga o'xshatish mumkin: agar sizda 7 ta 1000-so'mlik bo'lsa va 2 tadan bir "to'plam" qilmoqchi bo'lsangiz, 3 ta to'liq to'plam chiqadi, va 1 ta ortiqcha (qoldiq) qoladi — yarim to'plam degan narsa yo'q.

Kasr natija olish uchun kamida bitta tomon `float64` bo'lishi kerak:

```go
fmt.Println(7.0 / 2)  // 3.5
fmt.Println(7 / 2.0)  // 3.5 — ikkalasidan bittasi kasr bo'lsa yetarli
```

**Manfiy sonlar bilan qoldiq.** `%` operatori manfiy sonlar bilan ishlaganda, natijaning ishorasi **bo'linuvchi**ning ishorasiga mos keladi:

```go
fmt.Println(-7 % 2)  // -1
fmt.Println(7 % -2)  // 1
```

**Kasr sonlarni `==` bilan solishtirish — xavfli.** Kompyuter kasr sonlarni xotirada aniq emas, taxminiy saqlaydi (xuddi o'lchov asbobida har doim ozgina xatolik bo'lgani kabi). Shu sababli, ikkita hisoblangan `float64` qiymatni `==` bilan solishtirish kutilmagan natija berishi mumkin:

```go
fmt.Println(0.1 + 0.2 == 0.3) // false! garchi qog'ozda rost bo'lsa ham
```

Buning o'rniga, ikkala son bir-biriga **yetarlicha yaqinmi** deb tekshirish to'g'riroq (masalan, `math.Abs` yordamida farqini hisoblab, juda kichik bir sondan kamligini tekshirish).

**`math` paketi — qo'shimcha matematik funksiyalar:**

```go
math.Sqrt(16)     // 4  — kvadrat ildiz
math.Abs(-5)      // 5  — moduli (manfiylikni olib tashlaydi)
math.Max(3, 7)    // 7  — kattasi
math.Min(3, 7)    // 3  — kichikroq
math.Pow(2, 10)   // 1024 — 2 ning 10-darajasi
math.Floor(4.7)   // 4  — pastga yaxlitlash
math.Ceil(4.2)    // 5  — yuqoriga yaxlitlash
math.Round(4.5)   // 5  — odatdagi yaxlitlash qoidasi bo'yicha
```

Diqqat: `math` paketidagi funksiyalar deyarli hammasi `float64` bilan ishlaydi va `float64` qaytaradi — `int` bilan to'g'ridan-to'g'ri ishlatmoqchi bo'lsangiz, avval `float64(...)` orqali aylantirish kerak bo'ladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(17 / 5)     // 3 — butun bo'lish
	fmt.Println(17 % 5)     // 2 — qoldiq
	fmt.Println(17.0 / 5)   // 3.4 — kasr bo'lish

	fmt.Println(math.Sqrt(64))   // 8
	fmt.Println(math.Max(3, 9))  // 9
	fmt.Println(math.Abs(-12))   // 12
}
```

Natija:

```
3
2
3.4
8
9
12
```

## TASK

`qoldiq(a, b int) int` funksiyasi berilgan. Uni shunday to'ldiringki, u `a` ni `b` ga bo'lgandagi **qoldiqni** qaytarsin (masalan, `qoldiq(17, 5)` → `2`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. qoldiq operatori — `%` belgisi, bo'lish operatori `/` emas.
2. `return a % b` — funksiya tanasi shu bitta qatordan iborat bo'lishi mumkin.
