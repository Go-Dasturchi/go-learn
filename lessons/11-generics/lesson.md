# 11 — Generics

## THEORY

Oddiy kalitni tasavvur qiling — u faqat bitta o'lchamdagi murvatga mos keladi. Endi moslashuvchan (o'lchamini o'zgartirsa bo'ladigan) kalitni tasavvur qiling — u har xil o'lchamdagi murvatlarga mos kelaveradi, alohida-alohida kalit sotib olishning hojati yo'q. Go'da **generics** (umumlashtirilgan turlar) aynan shunday ishlaydi: bitta funksiyani **istalgan mos turdagi** qiymatlar bilan ishlaydigan qilib yozish imkonini beradi.

**Muammo — generics'siz.** Ikkita `int`dan kattasini topadigan funksiya yozsangiz:

```go
func KattaInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

Endi xuddi shu narsani `float64` yoki `string` uchun ham kerak bo'lsa — mantiq **bir xil**, faqat tur boshqacha. Generics'siz, har bir tur uchun alohida funksiya yozishga to'g'ri keladi.

**Yechim — turlangan parametr (type parameter).** Generics funksiya nomidan keyin, kvadrat qavs ichida, "bu funksiya qanday tur bilan ishlashi mumkinligi"ni bildiradi:

```go
type Ordered interface {
	~int | ~float64 | ~string
}

func Katta[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

- `[T Ordered]` — "bu funksiya `T` nomli, `Ordered` cheklovi (constraint)ga mos keladigan **istalgan** tur bilan ishlaydi" degani.
- `Ordered interface { ~int | ~float64 | ~string }` — bu maxsus interfeys, "constraint" deb ataladi: u oddiy method to'plamini emas, balki **qaysi turlar ruxsat etilganini** belgilaydi (`|` — "yoki" degani, ya'ni `int`, `float64` yoki `string` bo'lishi mumkin).

Endi bitta `Katta` funksiyasini istalgan mos tur bilan chaqirish mumkin:

```go
fmt.Println(Katta(3, 7))         // 7 — int bilan
fmt.Println(Katta(3.5, 2.1))     // 3.5 — float64 bilan
fmt.Println(Katta("olma", "nok")) // olma — string bilan (alifbo tartibida "solishtirilgan")
```

Go **avtomatik ravishda** argumentlarning turiga qarab `T` nima ekanligini aniqlaydi — alohida `Katta[int](3, 7)` deb yozishning hojati yo'q (garchi xohlasangiz shunday ham yozish mumkin).

**`~` belgisi nimani anglatadi.** `~int` — "aynan `int`ning o'zi, yoki uning asosida yaratilgan istalgan custom type" degani ("Custom Types" darsida ko'rgan `type Son int` kabi turlarni ham qamrab oladi). Agar shunchaki `int` deb yozilsa (tildasiz), faqat aynan `int`ning o'ziga ruxsat berilardi, undan yasalgan custom type'larga emas.

**Nega bu foydali.** Generics — kod takrorlanishini kamaytiradi: bir marta yozilgan, sinovdan o'tgan mantiq, ko'plab turlar bilan xavfsiz ishlaydi. Bu ayniqsa "Slices" darsida ko'rgan umumiy vositalar (masalan, "ro'yxatdan eng kattasini topish", "ikkita ro'yxatni birlashtirish") uchun juda foydali — ular deyarli har doim tur nomiga bog'liq bo'lmagan, umumiy mantiqqa ega bo'ladi.

## EXAMPLE

```go
package main

import "fmt"

type Ordered interface {
	~int | ~float64 | ~string
}

func Katta[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Katta(3, 7))
	fmt.Println(Katta(3.5, 2.1))
	fmt.Println(Katta("olma", "nok"))
}
```

Natija:

```
7
3.5
olma
```

## TASK

`Ordered` cheklovi (`~int | ~float64 | ~string`, allaqachon yozilgan) va `Katta[T Ordered](a, b T) T` generic funksiyasi berilgan. Funksiyani shunday to'ldiringki, u `a` va `b`dan **kattarog'ini** qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu — oddiy `if`/`else` bilan solishtirish, faqat `a` va `b` turi `int`, `float64` yoki `string` bo'lishi mumkin: `if a > b { return a }`.
2. `return b` — agar `a > b` yolg'on bo'lsa (ya'ni `b` katta yoki teng bo'lsa).
