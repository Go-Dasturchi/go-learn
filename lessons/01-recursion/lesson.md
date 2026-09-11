# 01 — Recursion

## THEORY

Matryoshka (bir-birining ichiga solingan rus qo'g'irchoqlari)ni tasavvur qiling — har birini ochsangiz, ichida xuddi shu qo'g'irchoqning **kichikroq nusxasi** chiqadi, toki eng kichigiga (ichida boshqa qo'g'irchoq bo'lmagan) yetguningizcha. **Rekursiya** — funksiyaning **o'zini o'zi chaqirishi**, har safar muammoning kichikroq versiyasi bilan, toki eng oddiy, to'g'ridan-to'g'ri javob beriladigan holatga (bunga **bazaviy holat**, "base case" deyiladi) yetguncha.

**Klassik misol — faktorial.** `n!` (n faktorial) — `1 * 2 * 3 * ... * n` ga teng. Buni rekursiv ta'riflash mumkin: `n! = n * (n-1)!`, va `0! = 1` (bazaviy holat):

```go
func faktorial(n int) int {
	if n == 0 {
		return 1 // bazaviy holat — bu yerda rekursiya to'xtaydi
	}
	return n * faktorial(n-1) // rekursiv chaqiruv — kichikroq muammo bilan
}
```

`faktorial(4)` chaqirilganda, aslida quyidagicha "zanjir" hosil bo'ladi:

```
faktorial(4) = 4 * faktorial(3)
             = 4 * (3 * faktorial(2))
             = 4 * (3 * (2 * faktorial(1)))
             = 4 * (3 * (2 * (1 * faktorial(0))))
             = 4 * (3 * (2 * (1 * 1)))
             = 24
```

**Har bir rekursiv funksiyada ikkita narsa bo'lishi shart:**
1. **Bazaviy holat (base case)** — rekursiya qachon to'xtashini bildiruvchi shart. Bu bo'lmasa, funksiya **cheksiz** o'zini chaqiraveradi, toki dastur xotira yetishmovchiligidan qulab tushmaguncha ("stack overflow").
2. **Rekursiv qadam** — muammoni **kichikroq** versiyasiga qisqartirib, o'zini shu kichikroq versiya bilan chaqirish. Muhim: har safar muammo albatta bazaviy holatga **yaqinlashishi** kerak (masalan, `n` kamayib borishi), aks holda rekursiya hech qachon to'xtamaydi.

**Nega rekursiya foydali.** Ba'zi masalalar rekursiv tarzda **tabiiy ravishda** oddiyroq ifodalanadi — ayniqsa daraxt tuzilmalari (keyingi darslarda ko'rasiz), yoki "katta muammoni bir xil, kichikroq muammolarga bo'lish" g'oyasiga asoslangan masalalar (masalan, "Merge Sort" va "Quick Sort" darslarida ko'rasiz). Har qanday rekursiv yechimni, nazariy jihatdan, oddiy `for` tsikli bilan ham yozish mumkin — lekin ba'zan rekursiv yozuv ancha tushunarli va qisqa chiqadi.

**Ehtiyot chorasi.** Rekursiya chiroyli, lekin ehtiyotsiz ishlatilsa xavfli — bazaviy holatni unutib qo'yish yoki muammoni to'g'ri kichraytirmaslik, dasturni "cheksiz chuqurlik"ka olib borib, qulab tushishiga sabab bo'ladi. Har doim: "bu albatta to'xtaydimi?" deb o'zingizga savol bering.

## EXAMPLE

```go
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func fibonachi(n int) int {
	if n <= 1 {
		return n
	}
	return fibonachi(n-1) + fibonachi(n-2)
}

func main() {
	fmt.Println(faktorial(5))  // 120
	fmt.Println(fibonachi(7))  // 13
}
```

Natija:

```
120
13
```

## TASK

`faktorial(n int) int` funksiyasi berilgan. Uni **rekursiya** yordamida shunday to'ldiringki, u `n!` (n faktorial) ni qaytarsin. `0!` — `1` ga teng deb qabul qilinadi (bazaviy holat).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bazaviy holatni unutmang: `if n == 0 { return 1 }`.
2. Aks holda: `return n * faktorial(n-1)` — funksiya o'zini kichikroq (`n-1`) qiymat bilan chaqiradi.
