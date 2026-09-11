# 22 — Slice Indexing

## THEORY

Slice — bu tartiblangan ro'yxat, va uning har bir "joyi" **indeks** deb ataladi, kutubxonadagi javon raflariga o'xshab, har bir rafning o'z raqami bor. Go'da indekslash **0** dan boshlanadi — birinchi element `s[0]`, ikkinchisi `s[1]`, va hokazo:

```go
mevalar := []string{"olma", "banan", "uzum"}
fmt.Println(mevalar[0]) // "olma"
fmt.Println(mevalar[2]) // "uzum"
```

**Oxirgi elementga murojaat.** Slice'ning oxirgi elementi har doim `s[len(s)-1]` indeksida turadi (`len(s)` — uzunlik, lekin indekslar 0dan boshlangani uchun oxirgi indeks uzunlikdan bitta kam):

```go
oxirgi := mevalar[len(mevalar)-1] // "uzum"
```

**Muhim ogohlantirish: chegaradan chiqib ketish (`index out of range`).** Agar mavjud bo'lmagan indeksga murojaat qilsangiz (masalan, 3 elementli slice'ning `s[5]`siga, yoki hatto `s[3]`siga — chunki oxirgi to'g'ri indeks `2`), dastur **darhol ishlashdan to'xtaydi** (bunga "panic" deyiladi):

```go
mevalar := []string{"olma", "banan", "uzum"}
fmt.Println(mevalar[5]) // PANIC: index out of range
```

Go'da manfiy indeks (`mevalar[-1]`, "oxiridan birinchisi" ma'nosida, ba'zi tillarda ishlaydigan) ham **umuman yo'q** — bunday yozish compile xatosi beradi. Oxirgi elementga murojaat qilish uchun har doim `s[len(s)-1]` formulasidan foydalaniladi.

**Indeks orqali yangilash.** Elementni o'qish (`s[i]`) qanday ishlasa, unga yangi qiymat yozish ham xuddi shunday ishlaydi:

```go
mevalar[0] = "shaftoli"
fmt.Println(mevalar) // [shaftoli banan uzum]
```

**Muhim, ko'p adashiladigan nozik joy: `range`dagi qiymat — nusxa, asl elementning o'zi emas!** "For Loop" darsida `range` bilan slice ustida aylanishni ko'rgan edik. Lekin bir muhim narsani eslab qolish kerak: `for _, v := range s` deganda, `v` — bu **elementning nusxasi**, asl slice ichidagi joyning o'zi emas. Shuning uchun `v`ni o'zgartirish asl slice'ga **hech qanday ta'sir qilmaydi**:

```go
sonlar := []int{1, 2, 3}
for _, v := range sonlar {
	v = v * 100 // faqat nusxani o'zgartiryapmiz
}
fmt.Println(sonlar) // [1 2 3] — o'zgarmagan!
```

Agar chindan ham slice ichidagi elementlarni o'zgartirmoqchi bo'lsangiz, **indeks orqali** yozish kerak:

```go
sonlar := []int{1, 2, 3}
for i := range sonlar {
	sonlar[i] = sonlar[i] * 100 // indeks orqali — asl slice o'zgaradi
}
fmt.Println(sonlar) // [100 200 300]
```

Bu — "elementlarni joyida o'zgartirish" kerak bo'lgan har qanday vazifada (masalan, hammasini ikki barobar oshirish, hammasiga bir son qo'shish) yodda tutish kerak bo'lgan eng muhim qoidalardan biri.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	sonlar := []int{10, 20, 30, 40}

	fmt.Println(sonlar[0])            // 10
	fmt.Println(sonlar[len(sonlar)-1]) // 40 — oxirgi element

	for _, v := range sonlar {
		v = v + 1000 // nusxani o'zgartiryapmiz, asl slice o'zgarmaydi
	}
	fmt.Println(sonlar) // [10 20 30 40] — o'zgarmagan

	for i := range sonlar {
		sonlar[i] = sonlar[i] + 1000 // indeks orqali — asl slice o'zgaradi
	}
	fmt.Println(sonlar) // [1010 1020 1030 1040]
}
```

Natija:

```
10
40
[10 20 30 40]
[1010 1020 1030 1040]
```

## TASK

`ikkilantir(sonlar []int) []int` funksiyasi berilgan. Uni shunday to'ldiringki, u **indeks orqali** har bir elementni ikki baravar oshirsin va shu (o'zgartirilgan) slice'ning o'zini qaytarsin (masalan, `ikkilantir([]int{1, 2, 3})` → `[2 4 6]`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `range` qiymati (`v`) bu yerda ishlamaydi, chunki u faqat nusxa — indeks kerak: `for i := range sonlar { ... }`.
2. Har bir aylanishda `sonlar[i] = sonlar[i] * 2` qiling, tsikldan keyin `return sonlar`.
