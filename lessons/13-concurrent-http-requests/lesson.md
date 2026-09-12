# 13 — Concurrent HTTP Requests

## THEORY

Tasavvur qiling, sizga 5 ta do'stingizga qo'ng'iroq qilib, ulardan xabar olishingiz kerak. Ularga **birma-bir**, birinchisi javob berib bo'lgach ikkinchisiga qo'ng'iroq qilsangiz, bu besh baravar ko'proq vaqt oladi. Aksincha, agar sizda besh telefon bo'lib, **hammasiga bir vaqtda** qo'ng'iroq qilsangiz — umumiy vaqt, eng sekin javob bergan do'stingizning vaqtigacha qisqaradi. Bu dars — "Goroutines", "Channels" va "HTTP Client" darslarida ko'rgan bilimlarni birlashtirib, aynan shu g'oyani amalga oshiradi: bir nechta HTTP so'rovini **ketma-ket emas, parallel** yuborish.

**Muammo — ketma-ket so'rovlar sekin.** "HTTP Client" darsida ko'rgan `matnniOl` funksiyasini eslang. Agar bir nechta URL manzilidan ma'lumot olish kerak bo'lsa, oddiy `for` tsikli bilan **birin-ketin** so'rov yuborish — har bir so'rovning javobini navbat bilan kutishga majbur qiladi:

```go
// SEKIN — har bir so'rov oldingisi tugaguncha kutadi
for _, url := range urllar {
	natija, _ := matnniOl(url)
	natijalar = append(natijalar, natija)
}
```

**Yechim — har bir so'rov uchun alohida goroutine.** "Goroutines" darsida ko'rgan **indeksli natija** naqshini eslang — bu yerda aynan shu naqsh, endi HTTP so'rovlari uchun qo'llaniladi:

```go
type Natija struct {
	Indeks int
	Matn   string
}

func barchaSahifalarniOl(urllar []string) []string {
	ch := make(chan Natija, len(urllar))

	for i, url := range urllar {
		go func(indeks int, u string) {
			matn, _ := matnniOl(u)
			ch <- Natija{Indeks: indeks, Matn: matn}
		}(i, url)
	}

	natijalar := make([]string, len(urllar))
	for range urllar {
		n := <-ch
		natijalar[n.Indeks] = n.Matn
	}
	return natijalar
}
```

**Nega tartib saqlanadi, garchi so'rovlar tasodifiy tartibda tugasa ham.** So'rovlarning **qaysi biri qachon** tugashi kafolatlanmagan — birinchi URL oxirgi bo'lib javob qaytarishi ham mumkin. Lekin har bir natija o'zining **asl indeksi** bilan birga kanalga yuborilgani uchun, ularni `natijalar[n.Indeks]` orqali **to'g'ri joyga** qo'yish — yakuniy natija har doim **kiritilgan tartibda** chiqishini kafolatlaydi, garchi ular parallel va tasodifiy tartibda hisoblangan bo'lsa ham.

**Nega bu haqiqiy dasturlarda muhim.** Tashqi xizmatlarga (ob-havo API, valyuta kursi, boshqa mikroservislar) ko'p so'rov yuboradigan dasturlar uchun, ularni parallel yuborish — umumiy javob vaqtini sezilarli darajada qisqartiradi. Bu — Go'ning goroutine'lari nima uchun "yengil" deb atalishining, va ular I/O (kiritish/chiqarish, masalan tarmoq so'rovlari) bilan ishlashda ayniqsa qulay bo'lishining amaliy sababi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func matnniOl(url string) (string, error) {
	javob, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer javob.Body.Close()
	baytlar, err := io.ReadAll(javob.Body)
	return string(baytlar), err
}

type Natija struct {
	Indeks int
	Matn   string
}

func barchaSahifalarniOl(urllar []string) []string {
	ch := make(chan Natija, len(urllar))
	for i, url := range urllar {
		go func(indeks int, u string) {
			matn, _ := matnniOl(u)
			ch <- Natija{Indeks: indeks, Matn: matn}
		}(i, url)
	}

	natijalar := make([]string, len(urllar))
	for range urllar {
		n := <-ch
		natijalar[n.Indeks] = n.Matn
	}
	return natijalar
}

func main() {
	s1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "birinchi")
	}))
	defer s1.Close()
	s2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ikkinchi")
	}))
	defer s2.Close()

	fmt.Println(barchaSahifalarniOl([]string{s1.URL, s2.URL}))
}
```

Natija:

```
[birinchi ikkinchi]
```

## TASK

`matnniOl` funksiyasi va `Natija` struct'i berilgan. `barchaSahifalarniOl(urllar []string) []string` funksiyasini shunday to'ldiringki, u har bir URL uchun alohida goroutine ishga tushirib (parallel ravishda `matnniOl` chaqirib), natijalarni **kiritilgan tartibda** (indeks bo'yicha) `[]string` qilib qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ch := make(chan Natija, len(urllar))`, keyin `for i, url := range urllar { go func(indeks int, u string) { matn, _ := matnniOl(u); ch <- Natija{Indeks: indeks, Matn: matn} }(i, url) }`.
2. `natijalar := make([]string, len(urllar))`, `for range urllar { n := <-ch; natijalar[n.Indeks] = n.Matn }`, oxirida `return natijalar`.
