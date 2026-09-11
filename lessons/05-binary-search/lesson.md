# 05 — Binary Search

## THEORY

Qog'oz lug'atdan biror so'zni qidirayotganingizni tasavvur qiling. Uni sahifama-sahifa, boshidan oxirigacha varaqlab chiqmaysiz — o'rtasidan ochasiz, kerakli so'z alifbo bo'yicha oldindami keyindami ekanligini ko'rasiz, va faqat o'sha **yarmini** davom ettirasiz. Har safar qidiruv maydonini **yarmiga** qisqartirasiz, toki so'zni topguningizcha. **Binary search (ikkilik qidiruv)** — aynan shunday ishlaydigan, juda tez qidiruv algoritmi.

**Shart: ro'yxat oldindan saralangan bo'lishi kerak.** Binary search faqat **saralangan** (o'sish tartibida) ro'yxatlarda ishlaydi — chunki u "o'rtadagi qiymat qidirilayotgandan kattami-kichikmi" degan taqqoslashga tayanadi, va bu faqat tartiblangan ma'lumotda mantiqiy bo'ladi.

**Algoritm qadamlari:**
1. Qidiruv oralig'ining chap (`chap`) va o'ng (`ong`) chegaralarini belgilang (dastlab — butun ro'yxat).
2. O'rta indeksni toping: `orta := (chap + ong) / 2`.
3. Agar `sonlar[orta] == maqsad` bo'lsa — topildi, `orta`ni qaytaring.
4. Agar `sonlar[orta] < maqsad` bo'lsa — maqsad **o'ng yarmida**, `chap = orta + 1` qiling.
5. Agar `sonlar[orta] > maqsad` bo'lsa — maqsad **chap yarmida**, `ong = orta - 1` qiling.
6. `chap > ong` bo'lib qolsa (qidiruv oralig'i "tugagan") — maqsad ro'yxatda yo'q, `-1` qaytaring.

```go
func binarQidiruv(sonlar []int, maqsad int) int {
	chap, ong := 0, len(sonlar)-1
	for chap <= ong {
		orta := (chap + ong) / 2
		if sonlar[orta] == maqsad {
			return orta
		} else if sonlar[orta] < maqsad {
			chap = orta + 1
		} else {
			ong = orta - 1
		}
	}
	return -1
}
```

**Nega bu tezroq.** Oddiy, boshidan-oxirigacha qidiruv (`for i, s := range sonlar { if s == maqsad { ... } }`) — har bir elementni birma-bir tekshiradi, ya'ni `n` elementli ro'yxat uchun eng yomon holatda `n` marta tekshirish kerak bo'lishi mumkin ("bu — `O(n)`" deyiladi). Binary search esa har qadamda qidiruv maydonini **yarmiga** qisqartiradi — 1000 elementli ro'yxatda ham atigi ~10 qadamda javob topiladi (`O(log n)`), chunki 1000 → 500 → 250 → 125 → ... → 1. Bu farq katta ma'lumotlar bilan ishlaganda **juda** sezilarli bo'ladi.

**Eslatma.** Go standart kutubxonasida ham tayyor `sort.SearchInts` funksiyasi bor, u xuddi shu algoritmni ichida bajaradi — lekin bu darsda algoritmning o'zini **qo'lda** yozib, uning qanday ishlashini chuqur tushunib olamiz.

## EXAMPLE

```go
package main

import "fmt"

func binarQidiruv(sonlar []int, maqsad int) int {
	chap, ong := 0, len(sonlar)-1
	for chap <= ong {
		orta := (chap + ong) / 2
		if sonlar[orta] == maqsad {
			return orta
		} else if sonlar[orta] < maqsad {
			chap = orta + 1
		} else {
			ong = orta - 1
		}
	}
	return -1
}

func main() {
	sonlar := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println(binarQidiruv(sonlar, 7))  // 3
	fmt.Println(binarQidiruv(sonlar, 4))  // -1 — topilmadi
}
```

Natija:

```
3
-1
```

## TASK

`binarQidiruv(sonlar []int, maqsad int) int` funksiyasi berilgan (`sonlar` har doim o'sish tartibida saralangan holda beriladi). Uni ikkilik qidiruv algoritmi yordamida shunday to'ldiringki, u `maqsad`ning `sonlar` ichidagi indeksini qaytarsin, agar topilmasa `-1` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `chap, ong := 0, len(sonlar)-1` bilan boshlang, `for chap <= ong { orta := (chap + ong) / 2 ...}` tsiklini yozing.
2. `sonlar[orta] == maqsad` bo'lsa `orta`ni qaytaring; kichik bo'lsa `chap = orta + 1`; katta bo'lsa `ong = orta - 1`; tsikldan tashqarida `return -1`.
