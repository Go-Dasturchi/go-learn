# 20 — Slices

## THEORY

Oldingi darsda array — o'lchami qattiq belgilangan, "tuzatib bo'lmaydigan" karton quti ekanligini ko'rdik. Endi tasavvur qiling: xarid ro'yxatini qog'ozga emas, balki telefoningizdagi ilovaga yozayapsiz — kerak bo'lsa yangi mahsulot qo'shasiz, kerak bo'lmaganini o'chirasiz, ro'yxat istalgan uzunlikda bo'lishi mumkin. Go'da bunday **moslashuvchan ro'yxat** — **slice** deb ataladi, va u kundalik dasturlashda arraydan ancha ko'proq ishlatiladi.

**Yaratish usullari.** Eng oddiyi — literal (tayyor qiymatlar bilan):

```go
mevalar := []string{"olma", "banan", "uzum"}
```

Diqqat qiling: array yozuvidan (`[3]string{...}`) farqli, slice yozuvida qavs ichida **hech qanday son yo'q** (`[]string{...}`) — chunki uzunlik qattiq belgilanmagan.

Bo'sh, lekin ma'lum uzunlikdagi slice yaratish uchun `make` ishlatiladi:

```go
sonlar := make([]int, 5) // 5 ta element, hammasi 0 bilan to'ldirilgan
```

**Uzunlik va sig'im.** `len(slice)` — slice ichida hozir nechta element borligini qaytaradi. Slice'ning ichki qurilishida yana **sig'im (capacity)** degan tushuncha ham bor — bu haqda keyingi ("Slice Append") darsda batafsil gaplashamiz, chunki u aynan `append` bilan ishlaganda muhim bo'ladi.

**Slicing sintaksisi — `s[boshlanish:tugash]`.** Slice nomining o'zi ham fe'l sifatida ishlatiladi: mavjud slice (yoki hatto array)dan **bir qismini** kesib olish mumkin, `boshlanish` indeksidan boshlab, `tugash` indeksigacha (tugash **kirmaydi**):

```go
sonlar := []int{10, 20, 30, 40, 50}

fmt.Println(sonlar[1:3])  // [20 30] — indeks 1 dan 3 gacha (3 kirmaydi)
fmt.Println(sonlar[:2])   // [10 20] — boshidan 2-indeksgacha (boshlanish tashlab qo'yilsa, 0 deb olinadi)
fmt.Println(sonlar[3:])   // [40 50] — 3-indeksdan oxirigacha (tugash tashlab qo'yilsa, охиригача deb olinadi)
fmt.Println(sonlar[:])    // [10 20 30 40 50] — hammasi
```

Buni non bo'lagini pichoq bilan kesishga o'xshatish mumkin: `sonlar[1:3]` — 1-chizidan 3-chizigacha bo'lgan bo'lakni oling, degani.

**Nil slice — "hali hech narsa yo'q" holati.** Agar slice'ni faqat e'lon qilib, hech narsa bilan to'ldirmasangiz, u **nil** (bo'sh, hech narsaga ishora qilmaydigan) holatda bo'ladi:

```go
var royxat []int
fmt.Println(royxat == nil) // true
fmt.Println(len(royxat))   // 0 — nil slice ustida len() chaqirish xavfsiz
```

Bu — Go'ning qulay tomonlaridan biri: nil slice ustida `len()` chaqirish yoki `range` bilan aylanish xatolik bermaydi (shunchaki 0 marta aylanadi) — faqat unga to'g'ridan-to'g'ri indeks orqali yozishga (`royxat[0] = 5`) urinish xato beradi, chunki hali hech qanday joy ajratilmagan.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	sonlar := []int{10, 20, 30, 40, 50}

	fmt.Println(sonlar)        // [10 20 30 40 50]
	fmt.Println(len(sonlar))   // 5
	fmt.Println(sonlar[1:3])   // [20 30]
	fmt.Println(sonlar[:2])    // [10 20]
	fmt.Println(sonlar[3:])    // [40 50]

	boshQismi := make([]int, 3)
	fmt.Println(boshQismi) // [0 0 0]

	var royxat []int
	fmt.Println(royxat == nil, len(royxat)) // true 0
}
```

Natija:

```
[10 20 30 40 50]
5
[20 30]
[10 20]
[40 50]
[0 0 0]
true 0
```

## TASK

`oraliq(sonlar []int, boshlanish, tugash int) []int` funksiyasi berilgan. Uni shunday to'ldiringki, u `sonlar` slice'ining `boshlanish` (kiritilgan) dan `tugash` (kiritilmagan) gacha bo'lgan qismini qaytarsin — ya'ni oddiygina Go'ning o'zining slicing sintaksisidan foydalaning.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Funksiya tanasi bitta qatordan iborat bo'lishi mumkin: `return sonlar[boshlanish:tugash]`.
2. `tugash` indeksi natijaga **kirmaydi** — masalan `oraliq([10,20,30,40], 1, 3)` → `[20 30]`, `40` emas.
