# 30 — Custom Types

## THEORY

Ikkita shisha idishni tasavvur qiling — ikkalasi ham xuddi shu materialdan (shisha) yasalgan, lekin bittasi "sut shishasi", ikkinchisi "yog' shishasi" deb yorliqlangan. Ular ichidan bir xil ko'rinsa ham, ularni **aralashtirib qo'ymaslik** uchun aniq yorliq muhim. Go'da `type` kalit so'zi orqali xuddi shunday — mavjud bir turga (masalan `int` yoki `float64`) asoslangan, lekin **o'z nomiga ega, alohida** yangi tur yaratish mumkin.

Buni "Structs" darsida allaqachon ko'rgansiz (`type Odam struct {...}`), lekin `type` faqat struct'lar uchun emas — u **istalgan** mavjud tur asosida yangi nom yaratish uchun ishlatiladi:

```go
type Harorat float64
type Ismlar []string
type Som int
```

**Nega bu foydali — tasodifiy aralashtirishning oldini oladi.** Tasavvur qiling, dasturingizda pul bilan ishlaydigan qism bor, va u so'm bilan dollarni ikkalasini ham oddiy `int` sifatida saqlaydi. Bunday holda, tasodifan so'm miqdorini dollar deb hisoblab yuborish oson xato bo'ladi — chunki ikkalasi ham Go uchun "shunchaki int". Agar ularni alohida turlarga ajratsangiz:

```go
type Som int
type Dollar int
```

Endi Go bu ikkalasini **butunlay boshqa-boshqa turlar** deb hisoblaydi, garchi ikkalasi ham asosida `int` bo'lsa ham — ularni to'g'ridan-to'g'ri qo'shib yoki solishtirib bo'lmaydi, avval "Type Conversion" darsida ko'rgan usul bilan aniq aylantirish kerak bo'ladi. Bu — kompilyator sizga "bu yerda so'm bilan dollarni aralashtiryapsiz, ehtiyot bo'ling" deb signal berishining bir usuli.

**Custom type'ga method qo'shish — bu ularning eng kuchli tomoni.** "Methods" darsida struct'larga method qo'shishni ko'rgan edik. Xuddi shu narsani oddiy turlarga asoslangan custom type'larga ham qilish mumkin:

```go
type Harorat float64

func (h Harorat) Fahrengeytga() float64 {
	return float64(h)*9/5 + 32
}

selsiy := Harorat(100)
fmt.Println(selsiy.Fahrengeytga()) // 212
```

Diqqat qiling: oddiy `float64`ning o'zida `Fahrengeytga()` degan method yo'q — bu faqat **sizning** `Harorat` turingizga tegishli. Bu — Go'da mavjud turlarga o'zingizning ma'nodor xatti-harakatingizni "qo'shib qo'yish" usuli, hech qanday murakkab meros (inheritance) mexanizmisiz.

**Muhim: custom type asl turining barcha xususiyatlarini saqlab qoladi.** `Harorat` — asosida `float64` bo'lgani uchun, unda hali ham oddiy arifmetik amallar (`+`, `-`, `*` va h.k.) ishlaydi, faqat endi siz unga o'zingizning nomingiz va methodlaringizni ham qo'sha olasiz:

```go
issiq := Harorat(30) + Harorat(5) // 35 — oddiy qo'shish ishlaydi
```

## EXAMPLE

```go
package main

import "fmt"

type Harorat float64

func (h Harorat) Fahrengeytga() float64 {
	return float64(h)*9/5 + 32
}

func main() {
	suvQaynash := Harorat(100)
	fmt.Println(suvQaynash.Fahrengeytga()) // 212

	xonaHarorati := Harorat(22)
	fmt.Println(xonaHarorati.Fahrengeytga()) // 71.6
}
```

Natija:

```
212
71.6
```

## TASK

`Harorat` custom turi (`float64` asosida) berilgan. `Fahrengeytga() float64` methodini shunday to'ldiringki, u Selsiy darajasini (`h`) Farengeytga aylantirib qaytarsin. Formula: `Farengeyt = Selsiy * 9/5 + 32`.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `h` — bu `Harorat` turida, lekin uni oddiy son sifatida ishlatish uchun `float64(h)` orqali aylantiring.
2. `return float64(h)*9/5 + 32` — formula to'g'ridan-to'g'ri shu.
