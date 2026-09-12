# 07 — Middleware

## THEORY

Kontsertga kirishdan oldin, hamma tomoshabin avval **bilet tekshiruvidan** o'tadi — bu tekshiruv konsertning o'zi emas, lekin **har bir** tomoshabin uchun, kontsert zaliga kirishdan oldin, bir xilda bajariladigan qo'shimcha qadam. **Middleware** — veb-dasturlashda aynan shu vazifani bajaradi: bu — asosiy handler ishga tushishidan **oldin** (yoki keyin) bajariladigan, ko'plab handler'lar uchun **umumiy** bo'lgan kod.

**Middleware — bu handler'ni "o'rab oluvchi" funksiya.** "Higher-Order Functions" darsida ko'rgan g'oyani eslang: funksiya boshqa funksiyani argument sifatida qabul qilib, **yangi funksiya** qaytarishi mumkin edi. Middleware aynan shunday ishlaydi:

```go
func logMiddleware(keyingi http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ishladi", "true") // "oldin" bajariladigan qism
		keyingi.ServeHTTP(w, r)              // asosiy handler'ni chaqirish
	})
}
```

`http.HandlerFunc(...)` — oddiy funksiyani `http.Handler` interfeysiga aylantiruvchi maxsus tur ("Custom Types" va "Interfaces" darslarini eslang — bu, aslida, `func(w, r)` imzosiga `ServeHTTP` methodini "qo'shib qo'yadigan" nozik usul).

**Middleware'ni qo'llash:**

```go
asosiyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Salom!")
})

oraladi := logMiddleware(asosiyHandler) // endi "o'ralgan" handler
```

Endi `oraladi.ServeHTTP(w, r)` chaqirilganda, avval `logMiddleware` ichidagi kod (sarlavha o'rnatish) ishlaydi, **so'ng** `keyingi.ServeHTTP(w, r)` orqali asl handler chaqiriladi.

**Nega bu foydali — kodni takrorlamaslik.** Agar loglash, autentifikatsiya tekshiruvi, yoki CORS sarlavhalarini **har bir** handler ichida qo'lda qo'shib chiqsangiz, bu takrorlanadigan, xato qilish oson kod bo'lardi. Middleware orqali bu mantiqni **bir marta** yozib, istalgan sondagi handler'ga, ularning o'zini o'zgartirmasdan "o'rab" qo'llash mumkin — bu naqsh, aslida, "Decorator" (bezovchi) deb ataladigan, ko'plab dasturlash tillarida uchraydigan umumiy naqshning Go'dagi tabiiy ko'rinishi.

**Bir nechta middleware'ni zanjirlash.** Middleware'lar bir-birining ustiga "qatlam-qatlam" qo'llanishi mumkin:

```go
yakuniy := logMiddleware(autentifikatsiyaMiddleware(asosiyHandler))
```

Bu yerda so'rov avval `logMiddleware`, keyin `autentifikatsiyaMiddleware`, so'ng asosiy handler orqali "o'tadi" — xuddi bir necha qatlamli tekshiruv postidan ketma-ket o'tgandek.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func logMiddleware(keyingi http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ishladi", "true")
		keyingi.ServeHTTP(w, r)
	})
}

func main() {
	asosiy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Salom!")
	})

	oraladi := logMiddleware(asosiy)

	req := httptest.NewRequest("GET", "/salom", nil)
	rec := httptest.NewRecorder()
	oraladi.ServeHTTP(rec, req)

	fmt.Println(rec.Header().Get("X-Ishladi"))
	fmt.Println(rec.Body.String())
}
```

Natija:

```
true
Salom!

```

## TASK

`logMiddleware(keyingi http.Handler) http.Handler` funksiyasi berilgan. Uni shunday to'ldiringki, u qaytargan yangi handler:

1. Javobga `X-Ishladi: true` sarlavhasini o'rnatsin.
2. So'ngra `keyingi.ServeHTTP(w, r)` orqali **asl** (o'ralgan) handler'ni chaqirsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { ... })` — yangi handler shu ko'rinishda qaytariladi.
2. Ichida: `w.Header().Set("X-Ishladi", "true")` avval, keyin `keyingi.ServeHTTP(w, r)`.
