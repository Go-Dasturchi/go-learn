# 03 — Buffered Channels

## THEORY

"Channels" darsida ko'rgan oddiy kanal — pochtachi va qabul qiluvchi ikkalasi ham bir vaqtda "hozir" bo'lishini talab qiladigan qo'lma-qo'l topshirishga o'xshaydi. Endi pochta qutisini tasavvur qiling — unda bir nechta (masalan, 5 ta) xat sig'adigan bo'sh joy bor: pochtachi xatni tashlab ketishi mumkin, hatto uy egasi hozir bo'lmasa ham — faqat quti **to'lib qolmasa** bo'ldi. **Buferlangan kanal (buffered channel)** — aynan shunday ishlaydi.

**Yaratish — ikkinchi argument sifatida sig'im (capacity):**

```go
ch := make(chan int, 5) // 5 ta joy sig'adigan bufer bilan
```

**Asosiy farq: yuborish darhol "hozir kimdir qabul qiladimi" deb kutmaydi.** Buferlangan kanalga yuborish, agar buferda **bo'sh joy bo'lsa**, darhol tugaydi (bloklanmaydi) — hech kim hozircha uni "olib ketmagan" bo'lsa ham:

```go
ch := make(chan int, 3)
ch <- 1 // darhol tugaydi — buferda joy bor
ch <- 2 // darhol tugaydi
ch <- 3 // darhol tugaydi — bufer endi to'liq (3/3)
// ch <- 4 // BU YERDA BLOKLANARDI — bufer to'lgan, joy yo'q
```

Faqat bufer **to'lib qolgandan keyingi** yuborish bloklanadi — chunki endi kimdir kelib, joy bo'shatmaguncha, yangi xat "sig'maydi".

**`len` va `cap` — buferning joriy holatini bilish.** Kanal ustida ham "Slices" darsida ko'rgan `len`/`cap` ishlaydi: `len(ch)` — hozir kanalda **kutib turgan** (hali olinmagan) qiymatlar soni, `cap(ch)` — kanalning umumiy sig'imi:

```go
fmt.Println(len(ch), cap(ch)) // 3 3 — to'liq to'lgan
```

**Nega buferlangan kanal foydali.** U goroutine'larni bir-biriga **qattiq bog'lamasdan** ishlashga imkon beradi — "ishlab chiqaruvchi" goroutine "iste'molchi" hali tayyor bo'lmasa ham, ma'lum miqdorda oldinga ketib, ishlashda davom etishi mumkin (buferda joy bor ekan). Bu ayniqsa ishlab chiqarish tezligi va iste'mol qilish tezligi bir-biriga to'liq mos kelmaydigan holatlarda foydali — masalan, "Worker Pools" darsida ko'radigan naqshda.

**Diqqat: buferlangan kanal ham cheksiz emas.** Bufer — faqat vaqtinchalik "amortizatsiya" beradi, muammoni butunlay yo'q qilmaydi. Agar bufer sig'imidan **ko'p** narsa yig'ilib qolsa (iste'molchi hech qachon kelmasa), yuborish baribir **abadiy bloklanadi** — xuddi oddiy kanaldagidek.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(len(ch), cap(ch)) // 3 3

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(len(ch), cap(ch)) // 1 3
}
```

Natija:

```
3 3
1
2
1 3
```

## TASK

`bufferGaJoylash(qiymatlar []int) []int` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `len(qiymatlar)` sig'imli buferlangan kanal yarating.
2. Barcha `qiymatlar`ni shu kanalga yuboring (sig'im yetarli bo'lgani uchun bloklanmaydi).
3. Kanalni yoping, so'ng undan barcha qiymatlarni `range` bilan o'qib, `[]int` slice qilib qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ch := make(chan int, len(qiymatlar))` — sig'im aynan elementlar soniga teng bo'lgani uchun, barcha yuborishlar bir xil goroutine'da, bloklanmasdan ketma-ket bajarilishi mumkin.
2. `for _, q := range qiymatlar { ch <- q }`, keyin `close(ch)`, so'ng `natija := []int{}` va `for q := range ch { natija = append(natija, q) }`, oxirida `return natija`.
