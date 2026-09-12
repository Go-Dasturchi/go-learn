# 11 — Unbuffered Channels

## THEORY

"Channels" darsida buferlanmagan kanalning **bloklovchi** tabiatini ko'rgan edingiz — yuborish, boshqa tomonda kimdir qabul qilishga tayyor turmaguncha, "kutib qoladi". Bu darsda, aynan shu xususiyatning yana bir, juda foydali qo'llanilishini ko'ramiz: ma'lumot yuborish uchun emas, balki **faqat signal berish** uchun.

**Signal sifatida ishlatish — `chan struct{}`.** Ba'zan sizga kanaldan **hech qanday ma'lumot** kerak emas, faqat "tayyor" yoki "tugadi" degan **signal**ning o'zi kerak. Bunday holatlarda, hech narsa egallamaydigan maxsus tur — **bo'sh struct (`struct{}`)** — ishlatiladi:

```go
tugadi := make(chan struct{})

go func() {
	// ... uzoq davom etadigan ish ...
	tugadi <- struct{}{} // "tugadim" signali
}()

<-tugadi // signal kelguncha shu yerda kutamiz
fmt.Println("Ish haqiqatan tugadi, endi davom etsak bo'ladi")
```

`struct{}{}` — "bo'sh struct'ning bo'sh qiymati" degani; u xotirada deyarli hech qanday joy egallamaydi, chunki tashiydigan **ma'lumoti yo'q** — faqat "voqea sodir bo'ldi" degan faktning o'zi muhim.

**Nega bu — "rendezvous" (uchrashuv nuqtasi) deyiladi.** Buferlanmagan kanal orqali signal berish, ikkita goroutine'ning **aniq bir lahzada** "uchrashishini" kafolatlaydi: yuboruvchi goroutine `tugadi <- struct{}{}` qatorida, qabul qiluvchi esa `<-tugadi` qatorida — ikkalasi ham **bir vaqtda**, shu nuqtada "qo'l berishadi". Bu kafolat orqali, siz **ishonch bilan** bilasizki, signal kelgan payt, ishlab chiqaruvchi goroutine ishini **chindan ham** tugatgan bo'ladi — bu, "Race Conditions" darsida ko'rgan muammolarning oldini olishning yana bir usuli.

**`close(ch)` orqali "bir nechtaga birdaniga" signal berish.** Agar bitta emas, balki **bir nechta** goroutine'ga bir vaqtda signal berish kerak bo'lsa, kanalni `<-` orqali emas, balki `close(...)` orqali "yopish" ishlatiladi — "Channels" darsida ko'rganingizdek, yopilgan kanaldan o'qish darhol (standart qiymat bilan) qaytadi, bu esa **barcha** kutayotgan goroutine'larni bir vaqtda "uyg'otadi":

```go
signal := make(chan struct{})
for i := 0; i < 3; i++ {
	go func(id int) {
		<-signal // hammasi shu yerda kutadi
		fmt.Println("Goroutine", id, "signalni oldi")
	}(i)
}
close(signal) // hammasiga BIRDANIGA signal beradi
```

**Amaliy misol — natijani "kutib olish".** Quyidagi naqsh — funksiya ichida bir goroutine natijani hisoblab, uni umumiy o'zgaruvchiga yozib, so'ng signal berishi, tashqi kod esa signalni kutib, shundan keyingina natijani xavfsiz o'qishi:

```go
func hisoblaVaKut(qiymat int) int {
	natija := 0
	tugadi := make(chan struct{})

	go func() {
		natija = qiymat * 2
		close(tugadi)
	}()

	<-tugadi // natija hisoblab bo'linishini kafolatlaydi
	return natija
}
```

## EXAMPLE

```go
package main

import "fmt"

func hisoblaVaKut(qiymat int) int {
	natija := 0
	tugadi := make(chan struct{})

	go func() {
		natija = qiymat * 2
		close(tugadi)
	}()

	<-tugadi
	return natija
}

func main() {
	fmt.Println(hisoblaVaKut(21))
}
```

Natija:

```
42
```

## TASK

`hisoblaVaKut(qiymat int) int` funksiyasi berilgan. Uni shunday to'ldiringki, u buferlanmagan `chan struct{}` signal kanali yordamida: 1) alohida goroutine ishga tushirib, `natija = qiymat * 2` hisoblasin va kanalni yopsin (`close`), 2) tashqi kod signalni kutib olgach, `natija`ni xavfsiz qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `tugadi := make(chan struct{})` bilan signal kanali yarating, `go func() { natija = qiymat * 2; close(tugadi) }()` bilan goroutine ishga tushiring.
2. `<-tugadi` bilan signalni kuting, so'ng `return natija`.
