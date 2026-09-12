# 06 — Middleware

## THEORY

"Web Development" bo'limidagi "Middleware" darsida, har bir so'rovga sarlavha qo'shuvchi oddiy middleware'ni ko'rgan edingiz. Middleware'ning yana bir, juda keng tarqalgan qo'llanilishi bor: **kirishni cheklash (access control)** — ya'ni, so'rov asosiy handler'ga yetib borishidan oldin, "bu so'rovni davom ettirsam bo'ladimi?" deb tekshirish.

**Autentifikatsiya tekshiruvchi middleware:**

```go
func authMiddleware(keyingi http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized) // 401 — ruxsat yo'q
			return // MUHIM: keyingi.ServeHTTP chaqirilmaydi!
		}
		keyingi.ServeHTTP(w, r)
	})
}
```

**Eng muhim nozik joy: `return` — asosiy handler'ni "to'xtatish".** "Web Development"dagi log middleware har doim `keyingi.ServeHTTP(w, r)`ni chaqirardi — chunki uning vazifasi faqat kuzatish edi, hech kimni "to'xtatmasdi". Bu yerda esa, agar shart bajarilmasa (token yo'q), **`keyingi.ServeHTTP` umuman chaqirilmaydi** — middleware so'rovni shu yerning o'zida "ushlab qoladi", asosiy handler hech qachon ishga tushmaydi. Bu — kirishni cheklashning butun mohiyati: **noto'g'ri** so'rovlar himoyalangan kodga **hech qachon yetib bormasligi** kerak.

**Nega bu naqsh — "Authentication" va "JWT" darslarining poydevori.** Keyingi darslarda ko'radigan, haqiqiy tokenni tekshirish (masalan, JWT imzosini tasdiqlash) — aynan shu middleware **ichiga** joylashtiriladi. Middleware'ning o'zi — "qanday tekshirish" mexanizmidan mustaqil "qobiq" (shell): u har doim bir xil ishlaydi (tekshir, muvaffaqiyatsiz bo'lsa to'xtat, aks holda davom ettir), faqat tekshirish mantig'ining o'zi murakkablashadi.

**Bir nechta himoyalangan handler'ga bitta middleware qo'llash:**

```go
mux := http.NewServeMux()
himoyalanganHandler := authMiddleware(http.HandlerFunc(maxfiyMalumot))
mux.Handle("/maxfiy", himoyalanganHandler)
```

`mux.Handle` (`mux.HandleFunc`dan farqli) — allaqachon `http.Handler` bo'lgan qiymatni to'g'ridan-to'g'ri ro'yxatdan o'tkazadi, bu — middleware bilan "o'ralgan" handler'larni ro'yxatdan o'tkazishning aynan shu usuli.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func authMiddleware(keyingi http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		keyingi.ServeHTTP(w, r)
	})
}

func main() {
	asosiy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Maxfiy ma'lumot")
	})
	himoyalangan := authMiddleware(asosiy)

	req1 := httptest.NewRequest("GET", "/maxfiy", nil)
	req1.Header.Set("Authorization", "token-123")
	rec1 := httptest.NewRecorder()
	himoyalangan.ServeHTTP(rec1, req1)
	fmt.Println(rec1.Code)

	req2 := httptest.NewRequest("GET", "/maxfiy", nil)
	rec2 := httptest.NewRecorder()
	himoyalangan.ServeHTTP(rec2, req2)
	fmt.Println(rec2.Code)
}
```

Natija:

```
200
401
```

## TASK

`authMiddleware(keyingi http.Handler) http.Handler` funksiyasi berilgan. Uni shunday to'ldiringki, u qaytargan yangi handler:

1. `Authorization` sarlavhasini tekshirsin.
2. Agar bo'sh bo'lsa, `http.StatusUnauthorized` (401) qaytarib, **asosiy handler'ni chaqirmasdan** to'xtasin.
3. Aks holda, `keyingi.ServeHTTP(w, r)` orqali asosiy handler'ni chaqirsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `token := r.Header.Get("Authorization")`, `if token == "" { w.WriteHeader(http.StatusUnauthorized); return }` — `return`ni unutmang!
2. Shartdan keyin: `keyingi.ServeHTTP(w, r)`.
