# 23 — Slice Copy

## THEORY

Kutubxonadagi kitob katalogi kartochkasini tasavvur qiling — kartochkaning o'zi kitob emas, balki kitob **qayerda turganini ko'rsatadigan yorliq**. Agar kimdir sizning kartochkangizni qo'lda ko'chirib olsa (nusxa yozib olsa), ikkala kartochka ham **bir xil javon** ni ko'rsatib turadi — ular ikkita alohida kitob emas! Go'da slice ham xuddi shunday ishlaydi: slice — bu asl ma'lumotning o'zi emas, balki unga ishora qiluvchi **"yorliq" (header)**.

**Muhim, ko'p adashiladigan haqiqat: `b := a` slice'ni "chinakam" nusxalamaydi.** Array bilan solishtiring ("Arrays" darsida ko'rgan edingiz) — array'ni tenglashtirish uning **to'liq mustaqil nusxasini** yaratadi. Slice'da esa bunday emas: `b := a` faqat ikkinchi "yorliq" yaratadi, ikkalasi ham **bitta xuddi shu orqa fondagi ma'lumotni** ko'rsatadi:

```go
a := []int{1, 2, 3}
b := a          // b — a bilan BIR XIL orqa fon ma'lumotiga ishora qiladi
b[0] = 99
fmt.Println(a)  // [99 2 3] — a HAM o'zgardi!
fmt.Println(b)  // [99 2 3]
```

Bu — array'dan farqli, ba'zida kutilmagan xatolarga sabab bo'lishi mumkin bo'lgan muhim xususiyat: agar bir slice'ni boshqa o'zgaruvchiga "shunchaki tenglashtirsangiz", ular aslida bir xil ma'lumotni birga "ishlatishadi" — birontasini o'zgartirsangiz, ikkalasi ham o'zgaradi.

**Chinakam, mustaqil nusxa olish — `copy` funksiyasi.** Agar sizga haqiqatan ham mustaqil, bir-biriga bog'liq bo'lmagan ikkita slice kerak bo'lsa (xuddi kitobning o'zini fotokopiya qilib, alohida javonga qo'yganday), maxsus `copy` funksiyasidan foydalanish kerak:

```go
a := []int{1, 2, 3}
b := make([]int, len(a)) // avval kerakli hajmda YANGI joy ajratamiz
copy(b, a)                // a ichidagi ma'lumotni b ga fizik ko'chiramiz

b[0] = 99
fmt.Println(a) // [1 2 3] — a o'zgarmadi!
fmt.Println(b) // [99 2 3] — faqat b o'zgardi
```

**Muhim qoida: `copy` faqat ikkalasining eng kichigi qadar ko'chiradi.** `copy(dst, src)` funksiyasi `dst` va `src`dan qaysi biri qisqaroq bo'lsa, o'sha uzunlik qadar element ko'chiradi va nechta element ko'chirilganini qaytaradi. Shu sababli, `b`ni oldindan `make([]int, len(a))` bilan **to'g'ri hajmda** tayyorlab qo'yish muhim — aks holda ba'zi elementlar ko'chirilmay qolishi mumkin:

```go
a := []int{1, 2, 3, 4, 5}
b := make([]int, 3) // faqat 3 ta joy!
n := copy(b, a)
fmt.Println(b, n) // [1 2 3] 3 — faqat birinchi 3 tasi ko'chdi
```

## EXAMPLE

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a // faqat "yorliq" nusxasi — bir xil ma'lumotga ishora qiladi
	b[0] = 99
	fmt.Println(a) // [99 2 3] — a ham o'zgardi
	fmt.Println(b) // [99 2 3]

	x := []int{1, 2, 3}
	y := make([]int, len(x))
	copy(y, x) // haqiqiy, mustaqil nusxa
	y[0] = 99
	fmt.Println(x) // [1 2 3] — x o'zgarmadi
	fmt.Println(y) // [99 2 3]
}
```

Natija:

```
[99 2 3]
[99 2 3]
[1 2 3]
[99 2 3]
```

## TASK

`nusxala(sonlar []int) []int` funksiyasi berilgan. Uni shunday to'ldiringki, u `sonlar`ning **mustaqil, haqiqiy nusxasini** qaytarsin — ya'ni qaytarilgan slice'ni o'zgartirish asl `sonlar`ga hech qanday ta'sir qilmasin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval to'g'ri hajmda yangi slice yarating: `nusxa := make([]int, len(sonlar))`.
2. `copy(nusxa, sonlar)` bilan ma'lumotni ko'chiring, keyin `return nusxa` qiling — `return sonlar` emas, chunki u hali ham asl slice'ning o'zi bo'lardi.
