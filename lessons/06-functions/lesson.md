# 06 — Functions

## THEORY

Funksiyani **retsept** yoki **avtomat (vending machine)** deb tasavvur qiling. Avtomatga pul solasiz va tugmani bosasiz (bu — kirish, ya'ni parametrlar), avtomat esa sizga mos mahsulotni chiqarib beradi (bu — natija, ya'ni qaytariladigan qiymat). Retseptni bir marta yozib qo'ysangiz, uni xohlagancha marta, har xil ingredientlar bilan qayta-qayta ishlatishingiz mumkin — funksiyaning ham asosiy afzalligi shu: kodni bir joyda yozib, uni butun dastur bo'ylab qayta-qayta chaqirasiz.

**Funksiya qanday e'lon qilinadi:**

```go
func nomi(parametr1 tur1, parametr2 tur2) qaytishTuri {
	// kod
	return natija
}
```

- `func` — "bu funksiya" degan kalit so'z.
- `nomi` — funksiyaning nomi, uni shu nom orqali chaqirasiz.
- `parametr1 tur1, parametr2 tur2` — funksiyaga qanday "ingredientlar" (kirish qiymatlari) kerakligi.
- `qaytishTuri` — funksiya nima qaytarishini bildiradi.
- `return natija` — funksiya ishini tugatib, natijani chaqirgan joyga qaytaradi.

**Bir xil turdagi parametrlarni qisqartirish.** Agar bir nechta parametr ketma-ket bir xil turda bo'lsa, turni faqat oxirgisida bir marta yozish kifoya:

```go
func kop(a, b int) int {
	return a * b
}
```

Bu `func kop(a int, b int) int` bilan aynan bir xil ishlaydi — shunchaki qisqaroq yozilgan.

**Funksiya har doim aniq nechta qiymat qabul qilishi kerak.** Ko'p skript tillaridan farqli, Go'da "ixtiyoriy parametr" yoki "standart qiymat" tushunchasi yo'q — agar funksiya ikkita parametr kutayotgan bo'lsa, uni chaqirganda albatta ikkalasini ham berish kerak, aks holda kod compile bo'lmaydi. (Buning o'rniga ishlatiladigan bir usul bor — **variadik parametrlar**, quyida ko'ramiz.)

**Chaqirish tartibi muhim emas.** Go'da funksiyalarni fayl ichida qayerda yozganingiz muhim emas — hatto funksiyani o'zidan oldin e'lon qilingan boshqa funksiya ichida chaqirsangiz ham, yoki keyinroq yozilgan funksiyani oldinroq chaqirsangiz ham, hammasi ishlayveradi, chunki Go butun faylni birinchi o'qib chiqib, keyin ishga tushiradi.

**Variadik (o'zgaruvchan sonli) parametrlar.** Ba'zan funksiyaga nechta qiymat berilishini oldindan bilmaysiz — masalan, ikkita son yig'indisimi, uchtami, o'ntami. Buning uchun parametr turi oldiga uch nuqta (`...`) qo'yiladi:

```go
func yigindi(sonlar ...int) int {
	natija := 0
	for _, son := range sonlar {
		natija += son
	}
	return natija
}
```

Bunday funksiyani xohlagan sondagi argument bilan chaqirish mumkin:

```go
yigindi(1, 2)        // 3
yigindi(1, 2, 3, 4)  // 10
yigindi()             // 0 — hech narsa bermasa ham bo'ladi
```

Funksiya ichida `sonlar` — oddiy `[]int` slice sifatida ishlaydi, uni `range` bilan aylanib chiqish mumkin (buni "For Loop" darsida ko'rgan edingiz).

**Har bir funksiya nomi paket ichida yagona bo'lishi kerak.** Go'da bir xil nomli, lekin turli parametrli ikkita funksiya yozib bo'lmaydi (ba'zi tillarda bunga "overloading" deyiladi — Go'da bu yo'q). Har bir funksiyaning nomi butun paket ichida bir marta ishlatiladi.

**Katta va kichik harf bilan boshlanish — muhim farq.** Agar funksiya nomi **katta** harf bilan boshlansa (masalan `Kop`), u boshqa paketlardan ham chaqirilishi mumkin bo'ladi — bunga "exported" (eksport qilingan) deyiladi. Kichik harf bilan boshlansa (masalan `kop`), u faqat shu paket ichida ishlatiladi. Bu kurs davomida biz deyarli har doim kichik harfdan foydalanamiz, chunki hammasi bitta `main` paketi ichida yoziladi.

## EXAMPLE

```go
package main

import "fmt"

func kop(a, b int) int {
	return a * b
}

func yigindi(sonlar ...int) int {
	natija := 0
	for _, son := range sonlar {
		natija += son
	}
	return natija
}

func main() {
	fmt.Println(kop(3, 4))
	fmt.Println(yigindi(1, 2, 3, 4, 5))
}
```

Natija:

```
12
15
```

## TASK

To'rtburchakning yuzini hisoblovchi `maydon(uzunlik, kenglik int) int` funksiyasini yozing — u `uzunlik * kenglik` ni qaytarishi kerak.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Funksiya tanasi bitta qatordan iborat bo'lishi mumkin: `return uzunlik * kenglik`.
2. Ikkala parametr ham `int` bo'lgani uchun, ularni `uzunlik, kenglik int` deb birga yozish mumkin — har biriga alohida `int` yozish shart emas.
