# 02 — Query Parameters

## THEORY

`https://sayt.uz/qidiruv?matn=salom&til=uz` kabi manzilni ko'rgan bo'lsangiz kerak — `?`dan keyingi qism (`matn=salom&til=uz`) **query parametrlar** deb ataladi: bu — URL orqali serverga qo'shimcha ma'lumot (masalan, qidiruv so'zi, sahifa raqami, filtr) uzatishning standart usuli. Buni pochta konvertiga yozilgan "qo'shimcha ko'rsatma"ga o'xshatish mumkin — asosiy manzil bir xil, lekin konvertga qo'shimcha yozuv qo'shib, xat qanday ishlov berilishini aniqlashtirasiz.

**O'qish — `r.URL.Query()`:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	qiymatlar := r.URL.Query()      // url.Values turi — map[string][]string'ga o'xshaydi
	ism := qiymatlar.Get("ism")      // birinchi "ism" qiymatini oladi, yo'q bo'lsa "" qaytaradi
	fmt.Fprintf(w, "Salom, %s!", ism)
}
```

`Get("ism")` — "Maps" darsida ko'rgan xatti-harakatga o'xshaydi: agar `ism` parametri URL'da umuman bo'lmasa, xato bermaydi, shunchaki **bo'sh string** (`""`) qaytaradi.

**Standart qiymat bilan ishlash.** Ko'pincha parametr berilmasa, "standart" qiymatdan foydalanish kerak bo'ladi — bu, "Boolean" va "If/Else" darslarida ko'rgan oddiy shart tekshiruvi bilan hal qilinadi:

```go
ism := qiymatlar.Get("ism")
if ism == "" {
	ism = "mehmon"
}
```

**So'rovni to'g'ridan-to'g'ri sinash uchun query parametrlarni qanday qo'shish.** `httptest.NewRequest`ga to'liq URL (query qismi bilan birga) berish kifoya:

```go
req := httptest.NewRequest("GET", "/salom?ism=Ali", nil)
```

Handler ichida `r.URL.Query().Get("ism")` chaqirilganda, u avtomatik ravishda `"Ali"`ni qaytaradi — xuddi haqiqiy brauzerdan kelgan so'rovdagidek.

**Bir nechta qiymat va bir nechta parametr.** URL'da bir nechta turli parametr bo'lishi mumkin (`?ism=Ali&yosh=25`), va hatto bitta parametr **bir necha marta** takrorlanishi mumkin (`?rang=qizil&rang=ko'k`) — bunday holatda `Get()` faqat **birinchisini** qaytaradi, hammasini olish uchun `qiymatlar["rang"]` orqali to'liq `[]string`ga murojaat qilinadi. Kundalik ishlarning aksariyatida esa, oddiy `Get()` yetarli bo'ladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func salomlashHandler(w http.ResponseWriter, r *http.Request) {
	ism := r.URL.Query().Get("ism")
	if ism == "" {
		ism = "mehmon"
	}
	fmt.Fprintf(w, "Salom, %s!", ism)
}

func main() {
	req1 := httptest.NewRequest("GET", "/salom?ism=Ali", nil)
	rec1 := httptest.NewRecorder()
	salomlashHandler(rec1, req1)
	fmt.Println(rec1.Body.String())

	req2 := httptest.NewRequest("GET", "/salom", nil)
	rec2 := httptest.NewRecorder()
	salomlashHandler(rec2, req2)
	fmt.Println(rec2.Body.String())
}
```

Natija:

```
Salom, Ali!
Salom, mehmon!
```

## TASK

`salomlashHandler(w http.ResponseWriter, r *http.Request)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `r.URL.Query().Get("ism")` orqali `ism` query parametrini o'qisin.
2. Agar `ism` bo'sh bo'lsa, `"mehmon"` deb qabul qilsin.
3. Javob sifatida `"Salom, <ism>!"` (masalan, `"Salom, Ali!"`) matnini yozsin (yangi qatorsiz, `Fprintf` bilan).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ism := r.URL.Query().Get("ism")`, keyin `if ism == "" { ism = "mehmon" }`.
2. `fmt.Fprintf(w, "Salom, %s!", ism)` — `Fprintln` emas, `Fprintf` ishlatiladi, chunki bu safar yangi qator kerak emas.
