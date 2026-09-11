# 04 — Queue

## THEORY

Do'kon kassasidagi navbatni tasavvur qiling — birinchi bo'lib turgan odam birinchi bo'lib xizmat ko'radi, navbatga yangi kelganlar esa **oxiriga** turishadi. **Queue (navbat)** — aynan shunday ishlaydigan tuzilma: **FIFO** ("First In, First Out" — birinchi kirgan birinchi chiqadi). Bu — "Stack" darsida ko'rgan LIFO'ning aynan **teskarisi**.

**Ikkita asosiy amal:**
- **Enqueue** — navbatning **oxiriga** yangi element qo'shish.
- **Dequeue** — navbatning **boshidagi** elementni olib tashlash va qaytarish.

```go
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(x int) {
	q.items = append(q.items, x) // stekdagidek, oxiriga qo'shiladi
}
```

**`Dequeue` — birinchi (boshidagi) elementni olib tashlash.** Bu yerda "Stack"dan farq shu: stekda **oxirgi** element olinardi, navbatda esa **birinchi**:

```go
func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	qiymat := q.items[0]
	q.items = q.items[1:] // birinchi elementni "kesib tashlaymiz"
	return qiymat, true
}
```

`q.items[1:]` — "Slices" darsida ko'rgan slicing sintaksisi: 1-indeksdan (ikkinchi elementdan) oxirigacha, ya'ni birinchi elementni tashlab, qolganlarini qoldiradi.

**Stack va Queue — bir xil amallar, teskari tartib.** E'tibor bering: `Enqueue` va `Push` aslida bir xil ishlaydi (oxiriga qo'shish) — farq faqat **olib chiqishda**: stek oxiridan oladi (LIFO), navbat boshidan oladi (FIFO). Ma'lumotlarni saqlashning tashqi ko'rinishi (slice) bir xil, lekin **qoidasi** (qaysi uchidan olish) boshqacha — bu, "Interfaces" darsida ko'rgan g'oyaga o'xshab, bir xil tashqi shakl ostida turli xatti-harakatlarni ifodalash mumkinligini yana bir bor ko'rsatadi.

**Nega navbat foydali.** Navbat — "adolatli" tartibda ishlov berish kerak bo'lgan har qanday joyda ishlatiladi: masalan, chop etish navbati (birinchi yuborilgan hujjat birinchi chop etiladi), vazifalarni ketma-ket bajarish tizimlari, va **kenglik bo'yicha qidiruv** (BFS, keyingi darslarda graflar bilan ishlashda ko'rasiz) — u aynan navbat yordamida amalga oshiriladi.

## EXAMPLE

```go
package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(x int) {
	q.items = append(q.items, x)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	qiymat := q.items[0]
	q.items = q.items[1:]
	return qiymat, true
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	qiymat, ok := q.Dequeue()
	fmt.Println(qiymat, ok) // 1 true — birinchi qo'shilgan birinchi chiqadi

	qiymat, ok = q.Dequeue()
	fmt.Println(qiymat, ok) // 2 true
}
```

Natija:

```
1 true
2 true
```

## TASK

`Queue` struct'i (`items []int` maydoni bilan, `Enqueue` methodi allaqachon yozilgan) berilgan. `Dequeue() (int, bool)` methodini shunday to'ldiringki, u navbatning **birinchi** (eng oldin qo'shilgan) elementini olib tashlab, uni qaytarsin. Agar navbat bo'sh bo'lsa, `0, false` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval bo'shlikni tekshiring: `if len(q.items) == 0 { return 0, false }`.
2. Birinchi elementni oling va slice'dan kesib tashlang: `qiymat := q.items[0]`, `q.items = q.items[1:]`, keyin `return qiymat, true`.
