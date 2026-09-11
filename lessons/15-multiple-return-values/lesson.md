# 15 — Multiple Return Values

## THEORY

Bozordagi elektron tarozini tasavvur qiling — mahsulotni tarozining ustiga qo'yganingizda, u sizga bir vaqtda ikkita ma'lumot beradi: **og'irlik** va **narx**. Ikkalasi ham bitta o'lchashning natijasi, va ularni alohida-alohida so'rashning hojati yo'q. Go funksiyasi ham xuddi shunday ishlay oladi — ko'p tillardan farqli o'laroq, Go funksiyasi **bir nechta qiymatni bir vaqtda** qaytarishi mumkin, bu esa alohida struktura yoki obyekt yaratishga hojat qoldirmaydi:

```go
func bolish(a, b int) (int, int) {
	bolinma := a / b
	qoldiq := a % b
	return bolinma, qoldiq
}
```

Chaqirishda ikkala qiymat ham qabul qilinadi:

```go
b, q := bolish(17, 5)
fmt.Println(b, q)  // 3 2
```

**Bu naqsh Go'da xatoliklarni qaytarish uchun juda keng qo'llaniladi.** Go'da alohida "try/catch" mexanizmi yo'q — buning o'rniga, deyarli har bir standart kutubxona funksiyasi natija bilan birga **xato** ham qaytaradi, va bu xatoni chaqiruvchi darhol tekshiradi:

```go
son, err := strconv.Atoi("abc")
if err != nil {
	fmt.Println("Xato:", err)
}
```

Bu yerda `err` — "hech qanday xato yo'q" degan maxsus `nil` qiymatida bo'lishi ham, yoki nimadir noto'g'ri ketganda xato haqidagi ma'lumotni o'zida saqlashi ham mumkin. Odat bo'yicha, agar funksiya xato qaytarishi mumkin bo'lsa, **xato har doim oxirgi qaytariladigan qiymat** bo'ladi — bu Go kodini o'qigan har bir kishi darhol tanib oladigan naqsh.

**Kerak bo'lmagan qiymatni e'tiborsiz qoldirish — `_`.** Agar biror qaytarilgan qiymat kerak bo'lmasa, uni pastki chiziqcha `_` bilan almashtirish mumkin — bu Go'ga "bilaman, bu qiymat qaytadi, lekin men uni ataylab ishlatmayman" deyishning rasmiy usuli:

```go
bolinma, _ := bolish(17, 5)  // qoldiq bizga kerak emas
```

Buni umuman yozmasdan qoldirib bo'lmaydi (`bolinma := bolish(17, 5)` — XATO beradi, chunki funksiya ikkita qiymat qaytaradi, biттasini emas) — Go har doim barcha qaytariladigan qiymatlarga aniq nom yoki `_` talab qiladi.

**Qiymatlar pozitsiyasi bo'yicha aniqlanadi, nomi bo'yicha emas.** `func bolish(a, b int) (int, int)` imzosida ikkala qaytariladigan qiymat ham shunchaki `int` deb yozilgan, ularning "ma'nosi" (bo'linmami yoki qoldiqmi) faqat funksiya ichidagi `return` qatoridagi **tartib** orqali belgilanadi. Chaqiruvchi tomon `b, q := bolish(...)` deb yozganda, birinchi qaytariladigan qiymat `b`ga, ikkinchisi `q`ga tushadi — nomlarning o'zi hech qanday bog'liqlik yaratmaydi, faqat **tartib** muhim. (Buni yanada aniqroq qilish uchun qaytariladigan qiymatlarga ham nom berish mumkin — buni keyingi darsda ko'ramiz.)

**"Bor-yo'qligini" tekshirish naqshi.** Ko'p qiymat qaytarish, ayniqsa keyingi darslarda (masalan, lug'atlardan qidirish) juda foydali bo'ladigan yana bir naqshni ham ta'minlaydi: funksiya nafaqat natijani, balki "bu natija haqiqatan topildimi" degan ikkinchi `bool` qiymatni ham qaytarishi mumkin — bu odatda "comma ok" naqshi deb ataladi, va uni keyingi mavzularda amalda ko'p uchratasiz.

## EXAMPLE

```go
package main

import "fmt"

func minMax(royxat []int) (int, int) {
	kichik, katta := royxat[0], royxat[0]
	for _, son := range royxat {
		if son < kichik {
			kichik = son
		}
		if son > katta {
			katta = son
		}
	}
	return kichik, katta
}

func bolish(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	kichik, katta := minMax([]int{5, 2, 9, 1, 7})
	fmt.Println(kichik, katta)

	bolinma, qoldiq := bolish(17, 5)
	fmt.Println(bolinma, qoldiq)

	// qoldiq kerak bo'lmasa — _ bilan e'tiborsiz qoldiramiz
	faqatBolinma, _ := bolish(20, 4)
	fmt.Println(faqatBolinma)
}
```

Natija:

```
1 9
3 2
5
```

## TASK

`bolishVaQoldiq(a, b int) (int, int)` funksiyasi berilgan. Uni shunday to'ldiringki, u ikkita qiymatni **bir vaqtda** qaytarsin:

1. birinchi — `a / b` (bo'linma)
2. ikkinchi — `a % b` (qoldiq)

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Funksiya imzosida ikkita qaytish turi qavs ichida ko'rsatiladi: `func bolishVaQoldiq(a, b int) (int, int)`.
2. `return a / b, a % b` — bitta `return` qatorida ikkita qiymatni vergul bilan ajratib qaytaring.
