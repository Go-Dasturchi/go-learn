# 01 — HTTP Handlers

## THEORY

Restoran ofitsiantini tasavvur qiling: mijoz buyurtma yozilgan qog'ozni beradi (bu — **so'rov, request**), ofitsiant oshxonaga borib, tayyor taomni olib kelib, mijozga topshiradi (bu — **javob, response**). Veb-server ham aynan shunday ishlaydi: brauzer (yoki boshqa dastur) **so'rov (request)** yuboradi, server esa **javob (response)** qaytaradi. Bu — **HTTP** protokolining asosiy g'oyasi, va Go'da buni yozish uchun maxsus tashqi kutubxona kerak emas — standart `net/http` paketi yetarli.

**Handler — so'rovga javob beruvchi funksiya.** Go'da har bir "ofitsiant" — bu **handler** deb ataladigan funksiya, va u har doim bir xil imzoga ega:

```go
func salomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Salom, Go!")
}
```

- **`r *http.Request`** — kelgan so'rov haqidagi barcha ma'lumot: qaysi manzilga (`r.URL`), qaysi usul bilan (`r.Method` — GET, POST va h.k.), qanday sarlavhalar (`r.Header`) bilan kelgani.
- **`w http.ResponseWriter`** — javobni "yozib chiqadigan" joy. `fmt.Fprintln(w, ...)` — "Strings"/"String Formatting" darslarida ko'rgan `fmt` funksiyalariga o'xshaydi, faqat bu safar natija ekranga emas, balki **`w` orqali so'rov yuborgan tomonga** yuboriladi.

**Haqiqiy serverni ishga tushirish (faqat ma'lumot uchun).** Amalda, handler'ni haqiqiy tarmoq orqali ishlatish uchun:

```go
http.HandleFunc("/salom", salomHandler)
http.ListenAndServe(":8080", nil) // serverni 8080-portda ishga tushiradi
```

Bu — dasturni **abadiy ishlab turadigan** qilib qo'yadi (so'rovlarni kutib), shuning uchun uni avtomatik testlarda to'g'ridan-to'g'ri ishlatib bo'lmaydi.

**Handler'ni haqiqiy server ochmasdan sinash — `httptest` paketi.** Go standart kutubxonasi handler'larni **haqiqiy tarmoq orqali server ochmasdan** sinash uchun maxsus vosita beradi:

```go
import "net/http/httptest"

req := httptest.NewRequest("GET", "/salom", nil) // "soxta" so'rov yaratish
rec := httptest.NewRecorder()                     // javobni "yozib oladigan" soxta yozuvchi

salomHandler(rec, req) // handler'ni to'g'ridan-to'g'ri chaqiramiz

fmt.Println(rec.Body.String()) // "Salom, Go!\n"
```

`httptest.NewRequest` — haqiqiy tarmoqsiz, xotirada "soxta" so'rov yasaydi; `httptest.NewRecorder` — `http.ResponseWriter`ning "soxta" versiyasi, handler yozgan hamma narsani xotirada saqlab qoladi, keyin `rec.Body.String()` orqali tekshirish mumkin. Bu darsda va keyingi barcha veb-darslarda handler'larni aynan shu tarzda, real server ochmasdan sinaymiz.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func salomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Salom, Go!")
}

func main() {
	req := httptest.NewRequest("GET", "/salom", nil)
	rec := httptest.NewRecorder()

	salomHandler(rec, req)

	fmt.Println(rec.Body.String())
}
```

Natija:

```
Salom, Go!

```

## TASK

`salomHandler(w http.ResponseWriter, r *http.Request)` funksiyasi berilgan. Uni shunday to'ldiringki, u javob sifatida aynan `"Salom, Go!"` matnini (qatorning oxirida yangi qator belgisi bilan) yozsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `fmt.Fprintln(w, "Salom, Go!")` — `w` ham `io.Writer` sifatida ishlaydi, xuddi `fmt.Println` konsolga yozgandek, bu safar javobga yozadi.
2. `Fprintln` (`Fprint`dan farqli) qatorning oxiriga avtomatik yangi qator belgisi (`\n`) qo'shadi — shuning uchun natija oxirida bo'sh qator ko'rinadi.
