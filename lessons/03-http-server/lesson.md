# 03 — HTTP Server

## THEORY

"HTTP Handlers" darsida `httptest` orqali handler'larni haqiqiy server ochmasdan sinashni ko'rgan edik. Endi, haqiqiy dasturda server **chindan** qanday ishga tushirilishini ko'ramiz.

**Eng oddiy usul — `http.ListenAndServe`:**

```go
http.HandleFunc("/salom", salomHandler)
http.ListenAndServe(":8080", nil) // 8080-portda tinglashni boshlaydi, abadiy ishlaydi
```

Bu — tez, sodda, lekin **moslashuvchanligi cheklangan**: vaqt limitlari (timeout), maxsus sozlamalarni bera olmaysiz.

**Professional usul — `http.Server` struct'i.** Haqiqiy loyihalarda, server odatda **struct** sifatida, aniq sozlamalar bilan quriladi ("Structs" darsini eslang):

```go
server := &http.Server{
	Addr:         ":8080",
	Handler:      mux,
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}

server.ListenAndServe() // endi shu struct'ning o'z sozlamalari bilan ishga tushadi
```

- **`Addr`** — qaysi portda tinglash kerakligi.
- **`Handler`** — barcha so'rovlarni qabul qiluvchi asosiy handler (odatda "Routing" darsida ko'rgan `*http.ServeMux`).
- **`ReadTimeout` / `WriteTimeout`** — "Time and Duration" darsida ko'rgan `time.Duration` turi — agar mijoz juda sekin so'rov yuborsa yoki javobni juda sekin qabul qilsa, server cheksiz kutib qolmasligi uchun chegaralar.

**Nega struct orqali qurish afzalroq.** `http.ListenAndServe(":8080", nil)` — bu qisqa yo'l, lekin **standart** (juda uzoq yoki cheksiz) timeout'lar bilan ishlaydi — bu, ishlab chiqarish muhitida xavfli bo'lishi mumkin (masalan, sekin yoki yomon niyatli mijozlar serverning resurslarini "band qilib" qo'yishi mumkin). `http.Server` struct'ini o'zingiz qurish, har bir sozlamani **aniq va nazorat ostida** belgilash imkonini beradi.

**Testda serverni ishga tushirish — nega `ListenAndServe`ni to'g'ridan-to'g'ri chaqirmaymiz.** "HTTP Client" va oldingi darslarda ko'rgan `httptest.NewServer` — aslida, ichkarida xuddi shu `http.Server` mexanizmidan foydalanadi, lekin **tasodifiy, bo'sh portda**, va uni avtomatik boshqaradi (`Close()` chaqirilganda to'xtaydi). Shu sababli, bu kursda serverni sinash uchun har doim `httptest.NewServer` ishlatiladi — `ListenAndServe`ning o'zi **abadiy bloklanadi**, shuning uchun uni avtomatik testda to'g'ridan-to'g'ri chaqirib bo'lmaydi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func serverTuzish(handler http.Handler, manzil string) *http.Server {
	return &http.Server{
		Addr:         manzil,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func main() {
	mux := http.NewServeMux()
	server := serverTuzish(mux, ":8080")

	fmt.Println(server.Addr)
	fmt.Println(server.ReadTimeout)
}
```

Natija:

```
:8080
5s
```

## TASK

`serverTuzish(handler http.Handler, manzil string) *http.Server` funksiyasi berilgan. Uni shunday to'ldiringki, u `Addr: manzil`, `Handler: handler`, `ReadTimeout: 5*time.Second`, `WriteTimeout: 10*time.Second` sozlamalari bilan `*http.Server` qaytarsin (serverni ishga tushirmasdan, faqat sozlangan struct'ni qaytaring).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `&http.Server{Addr: manzil, Handler: handler, ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second}` — struct literalini to'g'ridan-to'g'ri qaytaring.
2. `time` paketini import qilishni unutmang — `time.Second` shu paketda.
