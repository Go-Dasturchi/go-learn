# 06 — Indexes

## THEORY

Katta lug'atda biror so'zni topish uchun, uni sahifama-sahifa qidirmaysiz — chunki lug'at **alifbo tartibida**, va siz "M" harfiga tegishli sahifaga to'g'ridan-to'g'ri o'tasiz. Agar lug'at tartibsiz bo'lganida, so'zni topish uchun **hamma sahifani** ko'rib chiqishga to'g'ri kelardi. **Index (indeks)** — ma'lumotlar bazasida aynan shu vazifani bajaradi: u ma'lum ustun bo'yicha, ma'lumotni **tezda topish** uchun qo'shimcha "yo'l ko'rsatgich" tuzilmasini yaratadi.

**Indeks yaratish — `CREATE INDEX`:**

```sql
CREATE INDEX idx_foydalanuvchilar_ism ON foydalanuvchilar(ism);
```

Endi `WHERE ism = 'Ali'` kabi so'rovlar, jadvaldagi **har bir qatorni** tekshirish o'rniga, indeks orqali kerakli qatorlarni **deyarli darhol** topa oladi.

**Nega bu — "Binary Search" darsida ko'rgan g'oyaning davomi.** Indeks (odatda) daraxt-shaklidagi tuzilma ("Binary Search Tree" darsini eslang) sifatida saqlanadi — bu, oddiy, tartiblanmagan ro'yxatda **chiziqli qidiruv** (`O(n)`, har bir elementni tekshirish) o'rniga, **logarifmik qidiruv** (`O(log n)`, "Binary Search" darsida ko'rgan tezlik) imkonini beradi.

**Farqni ko'rish — `EXPLAIN ANALYZE`.** PostgreSQL'da so'rov haqiqatan qanday bajarilishini ko'rish mumkin:

```sql
EXPLAIN ANALYZE SELECT * FROM foydalanuvchilar WHERE ism = 'Ali';
```

Indeks bo'lmasa, natijada `Seq Scan` (ketma-ket, chiziqli qidiruv) ko'rinadi; indeks qo'shilgandan keyin — `Index Scan` (indeks orqali qidiruv), va katta jadvallarda bu farq **sezilarli** bo'ladi.

**Nega har bir ustunga indeks qo'shmaslik kerak.** Indeks — bepul emas: u qo'shimcha xotira egallaydi, va **har safar** yangi qator qo'shilganda yoki yangilanganda, indeksning o'zi ham yangilanishi kerak — bu, `INSERT`/`UPDATE` amallarini **sekinlashtiradi**. Shuning uchun indekslar odatda faqat **tez-tez qidiriladigan** ustunlarga (masalan, `WHERE`, `JOIN`, `ORDER BY`da ko'p ishlatiladigan ustunlar) qo'shiladi — bu, tezlik va xotira/yozish tezligi orasidagi **savdolashuv (trade-off)**.

**Go'da indeks g'oyasini sinash — chiziqli qidiruv va xarita orqali qidiruvni solishtirish:**

```go
func chiziqliQidiruv(royxat []string, maqsad string) int {
	for i, s := range royxat {
		if s == maqsad {
			return i
		}
	}
	return -1
}

func indeksYarat(royxat []string) map[string]int {
	indeks := make(map[string]int)
	for i, s := range royxat {
		indeks[s] = i
	}
	return indeks
}
```

`indeksYarat` — "Hash Table" darsida ko'rgan g'oyani qo'llab, `royxat`dagi har bir qiymatdan uning **o'rniga** to'g'ridan-to'g'ri o'tish imkonini beradi (`indeks["Ali"]` — deyarli bir zumda), butun ro'yxatni qayta-qayta aylanib chiqishga hojat qolmaydi — aynan SQL indeksining, Go'dagi soddalashtirilgan ko'rinishi.

## EXAMPLE

```go
package main

import "fmt"

func indeksYarat(royxat []string) map[string]int {
	indeks := make(map[string]int)
	for i, s := range royxat {
		indeks[s] = i
	}
	return indeks
}

func main() {
	royxat := []string{"Ali", "Vali", "Guli"}
	indeks := indeksYarat(royxat)

	fmt.Println(indeks["Vali"]) // 1 — darhol, ro'yxatni aylanmasdan
}
```

Natija:

```
1
```

## TASK

`indeksYarat(royxat []string) map[string]int` funksiyasi berilgan. Uni shunday to'ldiringki, u `royxat`dagi har bir qiymatdan uning **indeksiga** (o'rniga) `map[qiymat]indeks` ko'rinishida xarita qurib qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'sh map bilan boshlang: `indeks := make(map[string]int)`, keyin `for i, s := range royxat { ... }`.
2. Har bir aylanishda `indeks[s] = i` qiling, tsikldan keyin `return indeks`.
