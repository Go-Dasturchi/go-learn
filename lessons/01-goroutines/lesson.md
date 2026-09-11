# 01 — Goroutines

## THEORY

Bir o'zingiz uy ishlarini qilish o'rniga, har bir ishni (kir yuvish, ovqat pishirish, tozalash) alohida yordamchiga topshirib, hammasini **bir vaqtda, parallel** qilib bitirganingizni tasavvur qiling. Go'da bunday "yordamchi" — **goroutine** deb ataladi: bu — juda yengil (oddiy operatsion tizim oqimidan ancha arzon) parallel bajariladigan vazifa.

**Goroutine qanday ishga tushiriladi — `go` kalit so'zi.** Istalgan funksiya chaqiruvi oldiga `go` deb yozsangiz, u **alohida, parallel** oqimda ishga tushadi, va joriy kod **kutib turmasdan** darhol davom etadi:

```go
func salomlash() {
	fmt.Println("Salom, goroutine'dan!")
}

func main() {
	go salomlash() // parallel ishga tushadi
	fmt.Println("main davom etyapti")
}
```

**Muhim, xavfli nozik joy: `main()` hech kimni kutmaydi.** Yuqoridagi misolda, `salomlash()`ning chiqishi **ko'rinmasligi ham mumkin** — chunki `main()` funksiyasi tugashi bilan, dasturning **o'zi butunlay** to'xtaydi, hatto boshqa goroutine'lar hali ishini tugatmagan bo'lsa ham. Bu — ko'p yangi dasturchilarni chalg'itadigan birinchi narsa: goroutine ishga tushirilgandan keyin, uni **kutish** uchun maxsus vosita kerak bo'ladi (buni "WaitGroup" darsida ko'ramiz).

**Natijani qanday "qaytarib olish" — channel bilan tanishuv.** Goroutine oddiy funksiya kabi `return` orqali natija qaytara olmaydi (chunki uni chaqirgan kod allaqachon davom etib ketgan bo'lishi mumkin). Natijani xavfsiz "qaytarib olish" uchun **channel** (kanal) ishlatiladi — bu, aslida, goroutine'lar orasida ma'lumot yuborish/qabul qilish uchun mo'ljallangan "quvur" (keyingi darsda batafsil ko'ramiz):

```go
ch := make(chan int) // int qiymatlar uchun kanal yaratish
go func() {
	ch <- 42 // kanalga qiymat yuborish
}()
natija := <-ch // kanaldan qiymatni kutib olish
fmt.Println(natija) // 42
```

`<-` belgisi ikki xil ishlatiladi: `ch <- 42` — kanalga yuborish, `<-ch` — kanaldan olish. `natija := <-ch` qatori — goroutine qiymat yuborguncha **kutadi** (bloklanadi), shu orqali natijani xavfsiz "qaytarib olish" muammosi hal qilinadi.

**Bir nechta goroutine'ni parallel ishlatish va natijalarni indeks bo'yicha yig'ish.** Agar bir nechta qiymatni parallel qayta ishlab, natijalarni **to'g'ri tartibda** yig'ib olmoqchi bo'lsangiz, har bir natijaga "bu qaysi elementga tegishli" degan indeksni ham qo'shib yuborish qulay usul:

```go
type Natija struct {
	Indeks int
	Qiymat int
}

func kvadratlarParallel(sonlar []int) []int {
	ch := make(chan Natija)
	for i, son := range sonlar {
		go func(indeks, s int) {
			ch <- Natija{Indeks: indeks, Qiymat: s * s}
		}(i, son)
	}

	natijalar := make([]int, len(sonlar))
	for range sonlar {
		n := <-ch
		natijalar[n.Indeks] = n.Qiymat
	}
	return natijalar
}
```

Diqqat qiling: `go func(indeks, s int) { ... }(i, son)` — `i` va `son`ni **funksiya argumenti sifatida** uzatyapmiz, tashqi `for` tsiklining o'zgaruvchilariga to'g'ridan-to'g'ri emas. Bu — eski Go versiyalarida (1.22'dan oldin) tsikl o'zgaruvchisi barcha goroutine'lar orasida **umumiy** bo'lganligi sababli paydo bo'lgan, mashhur xatoning oldini oluvchi, xavfsiz odat — hozirgi Go'da bu muammo tuzatilgan bo'lsa ham, argument sifatida aniq uzatish har doim eng aniq va tushunarli yo'l bo'lib qoladi.

Natijalarni `natijalar[n.Indeks] = n.Qiymat` orqali **indeksga qarab** joylashtirish — goroutine'lar qaysi tartibda tugashidan qat'iy nazar (bu tartib **kafolatlanmagan**!), yakuniy natija har doim **to'g'ri tartibda** chiqishini ta'minlaydi.

## EXAMPLE

```go
package main

import "fmt"

type Natija struct {
	Indeks int
	Qiymat int
}

func kvadratlarParallel(sonlar []int) []int {
	ch := make(chan Natija)
	for i, son := range sonlar {
		go func(indeks, s int) {
			ch <- Natija{Indeks: indeks, Qiymat: s * s}
		}(i, son)
	}

	natijalar := make([]int, len(sonlar))
	for range sonlar {
		n := <-ch
		natijalar[n.Indeks] = n.Qiymat
	}
	return natijalar
}

func main() {
	fmt.Println(kvadratlarParallel([]int{1, 2, 3, 4, 5}))
}
```

Natija:

```
[1 4 9 16 25]
```

## TASK

`Natija` struct'i (`Indeks int`, `Qiymat int`) va `kvadratlarParallel(sonlar []int) []int` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u har bir son uchun **alohida goroutine** ishga tushirib, uning kvadratini hisoblasin, natijani `Natija{Indeks, Qiymat}` ko'rinishida kanalga yuborsin, va asosiy funksiya barcha natijalarni yig'ib, **to'g'ri tartibda** (kirish tartibida) `[]int` qilib qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ch := make(chan Natija)` bilan kanal yarating, keyin `for i, son := range sonlar { go func(indeks, s int) { ch <- Natija{Indeks: indeks, Qiymat: s * s} }(i, son) }` bilan har bir son uchun goroutine ishga tushiring.
2. Natijalarni yig'ish uchun: `natijalar := make([]int, len(sonlar))`, so'ng `for range sonlar { n := <-ch; natijalar[n.Indeks] = n.Qiymat }`, oxirida `return natijalar`.
