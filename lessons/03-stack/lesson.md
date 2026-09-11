# 03 — Stack

## THEORY

Oshxonada bir-birining ustiga terib qo'yilgan tarelkalar uyumini tasavvur qiling — yangi tarelkani faqat **ustiga** qo'yasiz, va olganda ham faqat **eng ustidagisini** olasiz. Pastdagi tarelkaga to'g'ridan-to'g'ri qo'l tekkizib bo'lmaydi. **Stack (stek)** — aynan shunday ishlaydigan tuzilma: **LIFO** ("Last In, First Out" — oxirgi kirgan birinchi chiqadi).

**Ikkita asosiy amal:**
- **Push** — stekning ustiga yangi element qo'shish.
- **Pop** — stekning eng ustidagi elementni olib tashlash va uni qaytarish.

Go'da stekni oddiy **slice** yordamida amalga oshirish mumkin — slice'ning **oxiri** stekning "usti" bo'ladi:

```go
type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x) // "Slice Append" darsini eslang
}
```

`Push` — "Methods" darsida ko'rgan pointer receiver'dan foydalanadi, chunki u `s.items`ni **chindan ham o'zgartirishi** kerak.

**`Pop` — oxirgi elementni olib tashlash.** Bu, "Slices" va "Multiple Return Values" darslarida ko'rgan ikkita naqshni birlashtiradi: oxirgi elementga `len(s.items)-1` indeksi orqali murojaat qilish, va natija bilan birga "muvaffaqiyatli bo'ldimi" degan `bool`ni ham qaytarish (chunki bo'sh stekdan `Pop` qilib bo'lmaydi):

```go
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // stek bo'sh — olib chiqadigan narsa yo'q
	}
	oxirgi := len(s.items) - 1
	qiymat := s.items[oxirgi]
	s.items = s.items[:oxirgi] // oxirgi elementni "kesib tashlaymiz"
	return qiymat, true
}
```

`s.items[:oxirgi]` — "Slices" darsida ko'rgan slicing sintaksisi: boshidan `oxirgi`-indeksgacha (kirmaydi), ya'ni oxirgi elementni tashlab, qolganlarini qoldiradi.

**Nega stek foydali.** Stek — dasturlashda juda ko'p joyda "yashirincha" ishlatiladi: masalan, funksiya chaqiruvlari o'zi ham stek tamoyilida ishlaydi ("Recursion" darsida ko'rgan zanjirni eslang — eng oxirgi chaqirilgan funksiya birinchi tugaydi). Qavslarning to'g'ri joylashganini tekshirish, "orqaga qaytish" (undo) funksiyasi, va chuqurlik bo'yicha qidiruv (DFS, keyingi darslarda ko'rasiz) — bularning barchasi stek g'oyasiga tayanadi.

## EXAMPLE

```go
package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	oxirgi := len(s.items) - 1
	qiymat := s.items[oxirgi]
	s.items = s.items[:oxirgi]
	return qiymat, true
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	qiymat, ok := s.Pop()
	fmt.Println(qiymat, ok) // 3 true — eng oxirgi qo'shilgan birinchi chiqadi

	qiymat, ok = s.Pop()
	fmt.Println(qiymat, ok) // 2 true
}
```

Natija:

```
3 true
2 true
```

## TASK

`Stack` struct'i (`items []int` maydoni bilan, `Push` methodi allaqachon yozilgan) berilgan. `Pop() (int, bool)` methodini shunday to'ldiringki, u stekning **eng ustidagi** (oxirgi qo'shilgan) elementini olib tashlab, uni qaytarsin. Agar stek bo'sh bo'lsa, `0, false` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval bo'shlikni tekshiring: `if len(s.items) == 0 { return 0, false }`.
2. Oxirgi elementni oling va slice'dan kesib tashlang: `oxirgi := len(s.items) - 1`, `qiymat := s.items[oxirgi]`, `s.items = s.items[:oxirgi]`, keyin `return qiymat, true`.
