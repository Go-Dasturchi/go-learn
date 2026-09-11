# 04 — Closures

## THEORY

Sayohatga chiqayotgan odamni tasavvur qiling — u o'zi bilan bir sumka olib boradi, unda kerakli narsalari (pasport, kalitlar, pul) bor. Qayerga borsa ham, sumkasi va uning ichidagi narsalar u bilan birga yuradi. Go'da **closure** — bu aynan shunday: **o'zi yaratilgan joydagi o'zgaruvchilarni "sumkasiga solib", o'zi bilan birga olib yuradigan funksiya**.

**Oddiy misol.** Bir funksiya boshqa bir funksiyani (ichki, anonim funksiyani) qaytarganda, va bu ichki funksiya tashqi funksiyaning o'zgaruvchilaridan foydalansa — bu closure hosil bo'ladi:

```go
func hisoblagichYarat() func() int {
	son := 0
	return func() int {
		son++
		return son
	}
}
```

`hisoblagichYarat()` chaqirilganda, u `son` nomli mahalliy o'zgaruvchi yaratadi va uni ishlatadigan **kichik funksiyani** qaytaradi. Odatda, "Scope" darsida ko'rganingizdek, funksiya tugagach uning mahalliy o'zgaruvchilari "yo'qolishi" kerak edi. Lekin closure bunga yo'l qo'ymaydi — u `son`ni o'zi bilan birga "sumkasida" olib ketadi:

```go
hisobla := hisoblagichYarat()
fmt.Println(hisobla()) // 1
fmt.Println(hisobla()) // 2
fmt.Println(hisobla()) // 3
```

Har safar `hisobla()` chaqirilganda, u xuddi shu `son`ning **davomini** ko'radi — chunki bu bitta, umumiy "sumka".

**Muhim: har bir chaqiruv o'zining alohida "sumkasini" oladi.** Agar `hisoblagichYarat()`ni yana bir marta chaqirsangiz, bu — **butunlay yangi**, mustaqil `son` bilan yangi closure yaratadi:

```go
hisobla1 := hisoblagichYarat()
hisobla2 := hisoblagichYarat()

fmt.Println(hisobla1()) // 1
fmt.Println(hisobla1()) // 2
fmt.Println(hisobla2()) // 1 — mustaqil, o'zining sumkasi bilan boshlaydi
```

`hisobla1` va `hisobla2` — ikkalasi ham bir xil "retsept" (funksiya tanasi) bo'yicha yaratilgan, lekin har birining o'z, mustaqil `son`i bor — ular bir-biriga hech qanday ta'sir qilmaydi.

**Nega bu foydali.** Closure'lar holatni (state) "yashirincha" saqlash kerak bo'lgan joylarda juda qulay — masalan, hisoblagichlar, keshlash, yoki bir marta sozlanib, keyin qayta-qayta ishlatiladigan funksiyalar yaratishda. Ular keyingi darsda ko'radigan **higher-order functions** (funksiyalarni parametr yoki natija sifatida ishlatish) bilan birga ayniqsa kuchli vosita hisoblanadi.

## EXAMPLE

```go
package main

import "fmt"

func hisoblagichYarat() func() int {
	son := 0
	return func() int {
		son++
		return son
	}
}

func main() {
	hisobla1 := hisoblagichYarat()
	fmt.Println(hisobla1())
	fmt.Println(hisobla1())
	fmt.Println(hisobla1())

	hisobla2 := hisoblagichYarat()
	fmt.Println(hisobla2()) // mustaqil — yana 1 dan boshlaydi
}
```

Natija:

```
1
2
3
1
```

## TASK

`hisoblagichYarat() func() int` funksiyasi berilgan. Uni shunday to'ldiringki, u chaqirilganda, har chaqirilganda 1 dan boshlab ketma-ket oshib boradigan (1, 2, 3, ...) sonni qaytaradigan **closure** yaratib bersin. Har bir yangi `hisoblagichYarat()` chaqiruvi mustaqil, o'zining alohida hisoblagichi bilan boshlanishi kerak.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Funksiya boshida `son := 0` deb e'lon qiling, keyin `return func() int { ... }` bilan ichki funksiyani qaytaring.
2. Ichki funksiya ichida `son++` qiling, keyin `return son` — tashqi `son` o'zgaruvchisi closure orqali "eslab qolinadi".
