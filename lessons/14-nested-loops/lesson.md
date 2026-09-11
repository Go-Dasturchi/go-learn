# 14 — Nested Loops

## THEORY

Sinf xonasidagi partalarni tasavvur qiling — bir necha **qator**, har qatorda bir necha **joy**. O'qituvchi har bir qatorni birma-bir aylanib, har qatorda barcha joylarni tekshirib chiqadi, keyingi qatorga o'tadi va u yerdagi barcha joylarni ham tekshiradi. Bu — aynan **ichma-ich tsikl (nested loop)** mantig'i: `for` tsiklini boshqa `for` tsikli ichiga joylashtirish. Ichki tsikl har safar tashqi tsiklning bitta qadami uchun **to'liq** aylanib chiqadi:

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 2; j++ {
		fmt.Println(i, j)
	}
}
```

Bu quyidagi juftliklarni chop etadi: `(1,1) (1,2) (2,1) (2,2) (3,1) (3,2)` — jami `3 × 2 = 6` marta. Har bir "tashqi qator" (`i`) uchun "ichki joy" (`j`) to'liq aylanib chiqadi.

Ichma-ich tsikllar ko'pincha ikki o'lchamli ma'lumot (jadval, matritsa, shaxmat taxtasi) bilan ishlashda ishlatiladi:

```go
for qator := 1; qator <= 3; qator++ {
	for ustun := 1; ustun <= 3; ustun++ {
		fmt.Print(qator*ustun, " ")
	}
	fmt.Println()
}
```

**`break` va `continue` faqat eng yaqin tsiklga ta'sir qiladi.** Bu — ko'p yangi dasturchilarni chalg'itadigan muhim nozik joy: agar ichki tsikl ichida `break` yozsangiz, u faqat **ichki** tsiklni to'xtatadi, tashqi tsikl esa davom etaveradi:

```go
for i := 1; i <= 3; i++ {
	for j := 1; j <= 3; j++ {
		if j == 2 {
			break // faqat ichki tsikldan chiqadi
		}
		fmt.Println(i, j)
	}
}
// chiqadi: (1,1) (2,1) (3,1) — har bir i uchun faqat j=1 chop etiladi,
// keyin ichki tsikl to'xtaydi, lekin tashqi i davom etadi
```

**Ikkala tsikldan birdaniga chiqish kerak bo'lsa — yorliq (label).** Agar tashqi tsikldan ham to'liq chiqmoqchi bo'lsangiz, Go'da tsiklga **nom (label)** berib, `break` shu nomga ishora qilishi mumkin:

```go
tashqi:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				break tashqi // ikkala tsikldan ham birdaniga chiqadi
			}
			fmt.Println(i, j)
		}
	}
```

Bu naqsh ancha kamdan-kam kerak bo'ladi (ko'pincha ichki mantiqni alohida funksiyaga chiqarib, u yerdan oddiy `return` qilish tozaroq yechim bo'ladi), lekin uni bilib qo'yish foydali — chunki yorliqsiz `break` faqat bitta, eng yaqin tsiklni to'xtatadi, deb kutilmagan natijaga duch kelmaslik uchun.

**Muhim ogohlantirish: ichma-ich tsikllar murakkablikni tez oshiradi.** Ikkita `n` martali tsikl birga `n × n` marta aylanadi (buni `O(n²)` deb belgilashadi) — masalan, `n = 1000` bo'lsa, bu million marta aylanish degani. Bu katta ma'lumotlar bilan ishlaganda sezilarli darajada sekinlashishi mumkin, shuning uchun kod yozganda "bu ichma-ich tsikl chindan ham kerakmi" deb o'ylab ko'rish foydali odat.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i, "-", j)
		}
	}

	fmt.Println("---")

	// ichki tsikldagi break faqat ichkisiga ta'sir qiladi
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				break
			}
			fmt.Println(i, j)
		}
	}
}
```

Natija:

```
1 - 1
1 - 2
1 - 3
2 - 1
2 - 2
2 - 3
---
1 1
2 1
3 1
```

## TASK

`yigindi2D(qatorlar, ustunlar int) int` funksiyasi berilgan. Ichma-ich `for` tsikli yordamida uni shunday to'ldiringki, u `qatorlar × ustunlar` o'lchamli "jadval"dagi barcha katakchalar sonini hisoblab qaytarsin (ya'ni oddiygina `qatorlar * ustunlar`ni tsikl orqali hisoblab chiqing — har bir katakcha uchun hisoblagichni bittaga oshiring).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Tashqi tsikl qatorlar bo'yicha, ichki tsikl ustunlar bo'yicha yurishi kerak: `for i := 0; i < qatorlar; i++ { for j := 0; j < ustunlar; j++ { ... } }`.
2. Har bir ichki aylanishda `soni++` qiling, tsikllardan keyin `return soni`.
