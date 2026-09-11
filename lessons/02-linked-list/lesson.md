# 02 — Linked List

## THEORY

Xazina qidiruv o'yinini tasavvur qiling: birinchi qog'ozda "keyingi maslahat u yerda" deb yozilgan, u yerga borsangiz, yana bir qog'oz — "keyingisi bu yerda" — va hokazo, toki oxirgi qog'ozgacha (unda "bu — oxiri" deb yozilgan). Har bir qog'oz o'zidan keyingisi **qayerdaligini** biladi, lekin hammasi bir joyda, ketma-ket turgan emas — ular xonaning turli burchaklarida bo'lishi mumkin. **Linked list (bog'langan ro'yxat)** — aynan shunday tuzilma: har bir element (**node**, tugun) o'z qiymatini va **keyingi tugunga ishora qiluvchi pointer**ni saqlaydi.

**Slice bilan solishtiring.** "Slices" darsida ko'rgan slice — bu xotirada **ketma-ket** joylashgan elementlar to'plami; istalgan elementga indeks orqali darhol murojaat qilish mumkin. Linked list esa — xotirada **tarqoq** joylashgan tugunlar, ularni faqat **boshidan boshlab, zanjir bo'ylab yurib** o'tish mumkin — 5-elementga to'g'ridan-to'g'ri "sakrab" o'tib bo'lmaydi, avval 1, 2, 3, 4-tugunlar orqali o'tish kerak.

**Tugun (Node) qanday e'lon qilinadi:**

```go
type Node struct {
	Value int
	Next  *Node // keyingi tugunga ishora — oxirgi tugunda bu nil bo'ladi
}
```

`Next *Node` — "Pointers" darsida ko'rgan pointer turi, lekin bu safar **o'zining turiga** ishora qilyapti (`Node` o'z ichida `*Node`ni saqlayapti) — bu Go'da butunlay normal, chunki pointer — bu shunchaki manzil, va manzilning o'zi qancha joy egallashi oldindan ma'lum (struct'ning o'zi qancha joy egallashidan qat'iy nazar).

**Qo'lda ro'yxat yasash.** Tugunlarni birma-bir yaratib, ularni bir-biriga ulash mumkin:

```go
uchinchi := &Node{Value: 3}
ikkinchi := &Node{Value: 2, Next: uchinchi}
birinchi := &Node{Value: 1, Next: ikkinchi}
// birinchi -> ikkinchi -> uchinchi -> nil
```

**Ro'yxat bo'ylab yurish.** Boshidan boshlab, `Next` orqali oxirigacha (`nil`ga yetguncha) yurish — linked list bilan ishlashning eng asosiy naqshi, va u ko'pincha `for` tsikli orqali (klassik, "For Loop" darsida ko'rgan sanoqli emas, **shart-only** ko'rinishida) amalga oshiriladi:

```go
joriy := birinchi
for joriy != nil {
	fmt.Println(joriy.Value)
	joriy = joriy.Next
}
```

**Nega linked list foydali (va nega kam ishlatiladi).** Linked list'ning afzalligi — boshiga yoki o'rtasiga yangi element qo'shish (agar tugunni bilsangiz) juda tez, chunki faqat bir nechta pointer'ni qayta yo'naltirish kifoya, butun ro'yxatni "surish" shart emas (slice'da elementni o'rtaga qo'shish esa qolgan hamma narsani surishga to'g'ri keladi). Lekin amalda, Go'da slice'lar ancha ko'proq ishlatiladi, chunki ular tezroq murojaat qilinadi (indeks orqali to'g'ridan-to'g'ri) va xotirani samaraliroq ishlatadi. Linked list — ko'proq **tushunish uchun muhim, poydevor** hisoblanadigan tuzilma: u pointer'lar, rekursiya va ma'lumotlar tuzilmalarining "qo'lda qurilishi" haqida chuqur tasavvur beradi.

## EXAMPLE

```go
package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	uchinchi := &Node{Value: 3}
	ikkinchi := &Node{Value: 2, Next: uchinchi}
	birinchi := &Node{Value: 1, Next: ikkinchi}

	joriy := birinchi
	for joriy != nil {
		fmt.Println(joriy.Value)
		joriy = joriy.Next
	}
}
```

Natija:

```
1
2
3
```

## TASK

`Node` struct'i (`Value int`, `Next *Node`) va `royxatGaAylantir(bosh *Node) []int` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u `bosh`dan boshlab, ro'yxat bo'ylab yurib, barcha qiymatlarni tartib bilan o'z ichiga olgan `[]int` slice qaytarsin. Agar `bosh` `nil` bo'lsa, bo'sh slice qaytarilishi kerak.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'sh slice bilan boshlang: `natija := []int{}`, keyin `joriy := bosh` deb belgilang.
2. `for joriy != nil { natija = append(natija, joriy.Value); joriy = joriy.Next }` — tsikldan keyin `return natija`.
