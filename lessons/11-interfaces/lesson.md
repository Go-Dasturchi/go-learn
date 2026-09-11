# 29 — Interfaces

## THEORY

Ish e'loni "haydovchi kerak, mashina yura olishi shart" deb yozilganini tasavvur qiling. Bu e'longa kim javob berishi mumkin? Yengil mashina, yuk mashinasi, hatto avtobus — **qaysi turdagi transport ekanligidan qat'iy nazar**, agar u "yura olsa", ish e'loniga mos keladi. Go'da **interface** aynan shu — bu **"nima qila olishi kerak"** haqidagi shartnoma, "qanday tur bo'lishi kerak" haqida emas.

**E'lon qilish — faqat method imzolari:**

```go
type Shakl interface {
	Yuza() float64
}
```

Bu — "`Shakl` deb hisoblanishi uchun, biror turning `Yuza() float64` methodi bo'lishi kifoya" degani.

**Eng muhim, Go'ga xos xususiyat: implicit (yashirin) implementatsiya.** Ko'p boshqa tillarda (masalan Java), "men bu interfeysni amalga oshiryapman" deb **ochiq aytish** kerak bo'ladi (`implements` kalit so'zi orqali). Go'da bunday emas — agar biror struct'da interfeys talab qilgan methodlar **borligi**ning o'zi kifoya, hech qanday qo'shimcha e'lon shart emas:

```go
type Kvadrat struct {
	Tomon float64
}

func (k Kvadrat) Yuza() float64 {
	return k.Tomon * k.Tomon
}

type Tortburchak struct {
	Uzunlik, Kenglik float64
}

func (t Tortburchak) Yuza() float64 {
	return t.Uzunlik * t.Kenglik
}
```

`Kvadrat` va `Tortburchak` — ikkalasi ham hech qayerda "men Shakl'man" demagan, lekin ikkalasida ham `Yuza() float64` methodi bo'lgani uchun, ikkalasi ham **avtomatik ravishda** `Shakl` interfeysiga mos keladi.

**Nima uchun bu foydali — polimorfizm.** Endi siz `Kvadrat` yoki `Tortburchak`ning aynan qaysi biri ekanligini bilmasdan turib, ularning hammasini **bir xil kod** bilan ishlata olasiz — faqat ularda umumiy `Yuza()` methodi borligini bilish kifoya:

```go
func chopEt(s Shakl) {
	fmt.Println("Yuza:", s.Yuza())
}

chopEt(Kvadrat{Tomon: 4})               // Yuza: 16
chopEt(Tortburchak{Uzunlik: 3, Kenglik: 5}) // Yuza: 15
```

Yoki hatto turli xil turlarni **bitta slice**ga yig'ib, ularning barchasi ustida bir xil amal bajarish mumkin:

```go
shakllar := []Shakl{
	Kvadrat{Tomon: 2},
	Tortburchak{Uzunlik: 3, Kenglik: 4},
}
for _, s := range shakllar {
	fmt.Println(s.Yuza())
}
```

Bu — struct'lar orasidagi umumiylikni ("hammasi yuza hisoblay oladi") kodning o'zida ifodalash imkonini beradi, har bir tur uchun alohida-alohida shart yozishga hojat qoldirmaydi.

## EXAMPLE

```go
package main

import "fmt"

type Shakl interface {
	Yuza() float64
}

type Kvadrat struct {
	Tomon float64
}

func (k Kvadrat) Yuza() float64 {
	return k.Tomon * k.Tomon
}

type Tortburchak struct {
	Uzunlik, Kenglik float64
}

func (t Tortburchak) Yuza() float64 {
	return t.Uzunlik * t.Kenglik
}

func main() {
	shakllar := []Shakl{
		Kvadrat{Tomon: 2},
		Tortburchak{Uzunlik: 3, Kenglik: 4},
	}
	for _, s := range shakllar {
		fmt.Println(s.Yuza())
	}
}
```

Natija:

```
4
12
```

## TASK

`Shakl` interfeysi (`Yuza() float64` methodi bilan), `Kvadrat` va `Tortburchak` struct'lari (ikkalasi ham `Yuza()` methodiga ega) berilgan. `jamiYuza(shakllar []Shakl) float64` funksiyasini shunday to'ldiringki, u ro'yxatdagi barcha shakllarning yuzalari **yig'indisini** qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `for _, s := range shakllar { ... }` — har bir elementning aniq turini bilish shart emas, chunki hammasi `Shakl` interfeysiga mos keladi.
2. Har bir aylanishda `natija += s.Yuza()` qiling, tsikldan keyin `return natija`.
