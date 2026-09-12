# 09 — HTTP Client

## THEORY

Shu paytgacha barcha darslarda biz **server** tomonini — kelgan so'rovlarga javob berishni — ko'rdik. Endi teskarisiga qaraymiz: Go dasturi o'zi boshqa bir serverga **so'rov yuborishi** kerak bo'lsa-chi? Masalan, ob-havo ma'lumotini olish uchun tashqi xizmatga murojaat qilish. Buning uchun `net/http` paketining **klient** qismi ishlatiladi.

**Oddiy `GET` so'rovi — `http.Get`:**

```go
javob, err := http.Get("https://misol.uz/malumot")
if err != nil {
	// tarmoq xatosi (masalan, server javob bermadi)
}
defer javob.Body.Close() // MUHIM: javob tanasini har doim yopish kerak
```

`http.Get` — "Errors" darsida ko'rgan naqsh bo'yicha, natija (`*http.Response`) va xatolikni birga qaytaradi. **`defer javob.Body.Close()`ni unutmaslik juda muhim** — "Defer" darsini eslang: bu tarmoq resurslarining (masalan, ulanishning) to'g'ri bo'shatilishini kafolatlaydi, hatto funksiya qanday tugashidan qat'iy nazar.

**Javob tanasini o'qish.** `javob.Body` — "File I/O" darsida ko'rgan `io.Reader`ning yana bir ko'rinishi. Uni to'liq matn sifatida o'qish uchun `io.ReadAll` ishlatiladi:

```go
baytlar, err := io.ReadAll(javob.Body)
matn := string(baytlar)
```

**O'zingizning kodni haqiqiy internetga bog'lamasdan sinash — `httptest.NewServer`.** Avtomatik testlar haqiqiy, tashqi veb-saytlarga bog'liq bo'lishi **yaramaydi** (internet bo'lmasa, sayt o'chib qolsa, yoki javob o'zgarsa — test buzilib qoladi). Buning o'rniga, Go standart kutubxonasi **vaqtinchalik, lokal test-server** ochish imkonini beradi:

```go
testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test ma'lumoti")
}))
defer testServer.Close() // test tugagach, serverni albatta o'chirish kerak

natija, _ := matnniOl(testServer.URL) // testServer.URL — masalan, "http://127.0.0.1:54321"
```

`httptest.NewServer` — haqiqiy, lekin **vaqtinchalik va faqat shu kompyuterning o'zida** ishlaydigan mini-server ochadi (tasodifiy, bo'sh portda). Bu — sizning HTTP klient kodingizni, hech qanday tashqi tarmoqqa bog'liq bo'lmasdan, to'liq real sharoitda sinash imkonini beradi.

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
	if err != nil {
		return "", err
	}
	return string(baytlar), nil
}

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Salom, mijoz!")
	}))
	defer testServer.Close()

	natija, _ := matnniOl(testServer.URL)
	fmt.Println(natija)
}
```

Natija:

```
Salom, mijoz!

```

## TASK

`matnniOl(url string) (string, error)` funksiyasi berilgan. Uni shunday to'ldiringki, u berilgan `url`ga `GET` so'rovi yuborsin, javob tanasini o'qib, uni `string` sifatida qaytarsin (agar xatolik yuz bersa, uni ham qaytaring).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `javob, err := http.Get(url)`, xatolikni tekshiring, so'ng `defer javob.Body.Close()` qiling.
2. `baytlar, err := io.ReadAll(javob.Body)`, xatolikni tekshiring, so'ng `return string(baytlar), nil` qiling.
