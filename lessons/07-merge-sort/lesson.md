# 07 — Merge Sort

## THEORY

Ikki kutubxonachi katta bir uyum kitobni muallif ismi bo'yicha saralashi kerak deb tasavvur qiling. Ular uyumni ikkiga bo'lib olishadi, har biri o'z yarmini (yana ikkiga bo'lib, va yana...) alohida-alohida saralaydi, va oxirida ikkita **allaqachon saralangan** kichik uyumni birlashtiradi — bu birlashtirish oson, chunki ikkala uyumning "eng tepasi" (eng kichik nusxasi) taqqoslanadi va navbat bilan olinadi. Bu — **"bo'lib tashla va yeng" (divide and conquer)** strategiyasi, va **merge sort** aynan shu tamoyilga asoslangan.

**Ikki bosqich:**
1. **Bo'lish (divide)** — ro'yxatni ikkiga bo'lib, har birini **rekursiv** ravishda (o'zini o'zi chaqirib, "Recursion" darsini eslang) saralang, toki bitta elementli (allaqachon "saralangan") bo'lakchalarga yetguncha.
2. **Birlashtirish (merge)** — ikkita saralangan bo'lakchani, ularning eng kichik elementlarini navbat bilan solishtirib, bitta katta saralangan ro'yxatga birlashtiring.

```go
func mergeTartiblash(sonlar []int) []int {
	if len(sonlar) <= 1 {
		return sonlar // bazaviy holat — 0 yoki 1 elementli ro'yxat allaqachon "saralangan"
	}
	orta := len(sonlar) / 2
	chap := mergeTartiblash(sonlar[:orta])
	ong := mergeTartiblash(sonlar[orta:])
	return merge(chap, ong)
}

func merge(chap, ong []int) []int {
	natija := []int{}
	i, j := 0, 0
	for i < len(chap) && j < len(ong) {
		if chap[i] <= ong[j] {
			natija = append(natija, chap[i])
			i++
		} else {
			natija = append(natija, ong[j])
			j++
		}
	}
	natija = append(natija, chap[i:]...)
	natija = append(natija, ong[j:]...)
	return natija
}
```

`merge` funksiyasi — ikkita **allaqachon saralangan** ro'yxatni oladi, va ularning boshidan solishtirib, kichigini navbat bilan natijaga qo'shib boradi. Bitta ro'yxat tugagach, ikkinchisining **qolgan qismi** (u allaqachon saralangan, hammasi navbatdagilardan katta) to'g'ridan-to'g'ri qo'shiladi (`chap[i:]...` — "Slice Append" darsida ko'rgan `...` yoyish belgisi).

**Nega merge sort ancha tezroq.** "Bubble Sort" darsida ko'rgan `O(n²)` bilan solishtiring: merge sort har safar ro'yxatni **yarmiga** bo'lgani uchun (xuddi "Binary Search"dagidek), umumiy murakkabligi `O(n log n)` bo'ladi — bu katta `n` uchun `O(n²)`dan sezilarli darajada tezroq. Masalan, 1 million elementli ro'yxat uchun bubble sort taxminan trillion amal talab qilsa, merge sort atigi ~20 million amal talab qiladi.

## EXAMPLE

```go
package main

import "fmt"

func merge(chap, ong []int) []int {
	natija := []int{}
	i, j := 0, 0
	for i < len(chap) && j < len(ong) {
		if chap[i] <= ong[j] {
			natija = append(natija, chap[i])
			i++
		} else {
			natija = append(natija, ong[j])
			j++
		}
	}
	natija = append(natija, chap[i:]...)
	natija = append(natija, ong[j:]...)
	return natija
}

func mergeTartiblash(sonlar []int) []int {
	if len(sonlar) <= 1 {
		return sonlar
	}
	orta := len(sonlar) / 2
	chap := mergeTartiblash(sonlar[:orta])
	ong := mergeTartiblash(sonlar[orta:])
	return merge(chap, ong)
}

func main() {
	fmt.Println(mergeTartiblash([]int{5, 2, 9, 1, 7, 3}))
}
```

Natija:

```
[1 2 3 5 7 9]
```

## TASK

`merge(chap, ong []int) []int` funksiyasi (ikkita saralangan slice'ni birlashtiradi) va `mergeTartiblash(sonlar []int) []int` funksiyasi berilgan. `mergeTartiblash`ni **rekursiya** yordamida shunday to'ldiringki, u `sonlar`ni ikkiga bo'lib, har birini rekursiv saralab, so'ng `merge` yordamida birlashtirib qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bazaviy holat: `if len(sonlar) <= 1 { return sonlar }` — 0 yoki 1 elementli ro'yxat allaqachon saralangan.
2. `orta := len(sonlar) / 2`, keyin `chap := mergeTartiblash(sonlar[:orta])`, `ong := mergeTartiblash(sonlar[orta:])`, va oxirida `return merge(chap, ong)`.
