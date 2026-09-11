# 07 — Sorting

## THEORY

Kutubxonachi kitoblarni javonga tartib bilan (masalan, muallif ismi bo'yicha alifbo tartibida) terib qo'yganini tasavvur qiling — endi kerakli kitobni topish ancha oson. **Saralash (sorting)** — ma'lumotlarni ma'lum bir tartibda joylashtirish. Go'da buning uchun har safar o'zingiz algoritm yozishning hojati yo'q — standart `sort` paketi tayyor vositalarni beradi.

**Oddiy turlar uchun tayyor funksiyalar:**

```go
sonlar := []int{5, 2, 9, 1, 7}
sort.Ints(sonlar)
fmt.Println(sonlar) // [1 2 5 7 9]

sozlar := []string{"banan", "olma", "uzum"}
sort.Strings(sozlar)
fmt.Println(sozlar) // [banan olma uzum]
```

Diqqat qiling: `sort.Ints`/`sort.Strings` slice'ni **joyida** (in-place) o'zgartiradi — "Slice Indexing" darsida ko'rgan indeks orqali o'zgartirish kabi, yangi slice qaytarmaydi.

**Maxsus tartib — `sort.Slice`.** Agar sizga struct'lar slice'ini biror maydon bo'yicha saralash kerak bo'lsa (masalan, odamlarni yoshi bo'yicha), yoki kamayish tartibida saralash kerak bo'lsa, `sort.Slice` ishlatiladi — unga slice va **"solishtirish qoidasi"**ni bildiruvchi funksiya beriladi:

```go
type Odam struct {
	Ism  string
	Yosh int
}

odamlar := []Odam{
	{Ism: "Vali", Yosh: 30},
	{Ism: "Ali", Yosh: 25},
	{Ism: "Guli", Yosh: 22},
}

sort.Slice(odamlar, func(i, j int) bool {
	return odamlar[i].Yosh < odamlar[j].Yosh
})
```

`func(i, j int) bool` — bu "Higher-Order Functions" darsida ko'rgan naqsh: `sort.Slice`ga funksiya beryapmiz, va u ichida **"i-elementi j-elementidan oldin turishi kerakmi?"** degan savolga javob beramiz. `odamlar[i].Yosh < odamlar[j].Yosh` — "kichikroq yoshdagi oldinroq tursin" degani, ya'ni o'sish tartibida saralash. Agar kamayish tartibida kerak bo'lsa, shunchaki `>` belgisini ishlatish kifoya.

**Nega `sort.Slice` shunchalik moslashuvchan.** `sort.Ints`/`sort.Strings` faqat oddiy o'sish tartibida ishlaydi, lekin `sort.Slice` orqali **istalgan qoida** bilan saralash mumkin — ism uzunligi bo'yicha, bir nechta maydon bo'yicha (avval yoshi, teng bo'lsa ism bo'yicha), yoki hatto teskari tartibda. Bu — Go'ning "algoritmning o'zi bir marta yozilgan, faqat qoidasi o'zgaradi" falsafasining yana bir yorqin namunasi (buni "Interfaces" darsida ham ko'rgan edingiz).

## EXAMPLE

```go
package main

import (
	"fmt"
	"sort"
)

type Odam struct {
	Ism  string
	Yosh int
}

func main() {
	sonlar := []int{5, 2, 9, 1, 7}
	sort.Ints(sonlar)
	fmt.Println(sonlar)

	odamlar := []Odam{
		{Ism: "Vali", Yosh: 30},
		{Ism: "Ali", Yosh: 25},
		{Ism: "Guli", Yosh: 22},
	}
	sort.Slice(odamlar, func(i, j int) bool {
		return odamlar[i].Yosh < odamlar[j].Yosh
	})
	fmt.Println(odamlar)
}
```

Natija:

```
[1 2 5 7 9]
[{Guli 22} {Ali 25} {Vali 30}]
```

## TASK

`Odam` struct'i (`Ism string`, `Yosh int`) va `yoshBoyicha(odamlar []Odam) []Odam` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u berilgan `odamlar` slice'ini **yoshi bo'yicha o'sish tartibida** saralab, shu (saralangan) slice'ning o'zini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `sort` paketini import qilishni unutmang.
2. `sort.Slice(odamlar, func(i, j int) bool { return odamlar[i].Yosh < odamlar[j].Yosh })` keyin `return odamlar`.
