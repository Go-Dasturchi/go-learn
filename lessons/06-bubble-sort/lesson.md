# 06 — Bubble Sort

## THEORY

Bir qator odamni bo'yi bo'yicha tizish kerak bo'lsa, va siz faqat **yonma-yon turgan ikkitasini** solishtirib, kerak bo'lsa joyini almashtira olsangiz, buni qanday qilasiz? Chapdan boshlab, har bir juftlikni solishtirib, agar noto'g'ri tartibda bo'lsa — joylarini almashtirib, o'ngga qarab yurasiz. Bir marta oxirigacha yurganingizda, **eng baland** odam albatta oxiriga "suzib chiqadi". Shuni bir necha marta takrorlaysiz, toki hammasi joyida bo'lguncha. Aynan shu — **bubble sort (pufakcha saralash)** algoritmi, nomini elementlarning asta-sekin o'z joyiga "suzib chiqishi"dan olgan.

**Algoritm:**

```go
func pufakchaSaralash(sonlar []int) []int {
	n := len(sonlar)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if sonlar[j] > sonlar[j+1] {
				sonlar[j], sonlar[j+1] = sonlar[j+1], sonlar[j]
			}
		}
	}
	return sonlar
}
```

- **Ichki tsikl** (`j`) — yonma-yon turgan elementlarni solishtirib, kerak bo'lsa almashtirib, chapdan o'ngga yuradi.
- **Tashqi tsikl** (`i`) — ichki tsiklni bir necha marta takrorlaydi. Har bir to'liq aylanishdan keyin, eng katta topilmagan element o'z joyiga (oxiriga) "cho'kadi" — shuning uchun `n-1-i` bilan ichki tsiklning oxirini har safar bittaga qisqartiramiz, chunki oxirgi `i` ta element allaqachon to'g'ri joyida turibdi.

**`sonlar[j], sonlar[j+1] = sonlar[j+1], sonlar[j]` — Go'ning qulay almashtirish sintaksisi.** Ko'p tillarda ikkita qiymatni almashtirish uchun vaqtinchalik o'zgaruvchi kerak bo'ladi (`vaqtincha := a; a = b; b = vaqtincha`). Go'da esa buni bitta qatorda, "Multiple Return Values" darsida ko'rgan ko'p qiymatli tayinlash orqali to'g'ridan-to'g'ri qilish mumkin.

**Nega bubble sort sekin, lekin tushunish uchun foydali.** Bubble sort — eng oddiy, lekin eng **sekin** saralash algoritmlaridan biri: har doim ikkita ichma-ich tsikl kerak bo'lgani uchun, `n` elementli ro'yxat uchun taxminan `n × n` marta solishtirish talab qiladi (`O(n²)` — "Nested Loops" darsida ko'rgan tushunchani eslang). Katta ro'yxatlar uchun bu sezilarli darajada sekin. Shunga qaramay, uni o'rganish muhim — chunki u saralash algoritmlarining eng sodda, tushunarli namunasi, va keyingi darslarda ko'radigan **ancha tezroq** algoritmlar (Merge Sort, Quick Sort) bilan solishtirish uchun yaxshi boshlang'ich nuqta bo'ladi.

## EXAMPLE

```go
package main

import "fmt"

func pufakchaSaralash(sonlar []int) []int {
	n := len(sonlar)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if sonlar[j] > sonlar[j+1] {
				sonlar[j], sonlar[j+1] = sonlar[j+1], sonlar[j]
			}
		}
	}
	return sonlar
}

func main() {
	fmt.Println(pufakchaSaralash([]int{5, 2, 9, 1, 7}))
}
```

Natija:

```
[1 2 5 7 9]
```

## TASK

`pufakchaSaralash(sonlar []int) []int` funksiyasi berilgan. Uni **bubble sort** algoritmi yordamida (Go'ning tayyor `sort` paketidan foydalanmasdan, algoritmning o'zini yozib) shunday to'ldiringki, u `sonlar`ni o'sish tartibida saralab, shu (saralangan) slice'ning o'zini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikkita ichma-ich `for` tsikli kerak: tashqisi `i := 0; i < n; i++`, ichkisi `j := 0; j < n-1-i; j++`.
2. Ichki tsiklda `if sonlar[j] > sonlar[j+1] { sonlar[j], sonlar[j+1] = sonlar[j+1], sonlar[j] }` — Go'ning bitta qatorlik almashtirish sintaksisidan foydalaning.
