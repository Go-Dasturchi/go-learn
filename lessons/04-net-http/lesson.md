# 04 — net/http

## THEORY

Shu paytgacha `net/http` paketining ko'plab qismlarini alohida-alohida ko'rgan edingiz — handler'lar, klient, server, status kodlar. Bu dars, paketning yana bir muhim qismini — **so'rovni qo'lda, batafsil qurish**ni ko'rsatadi, chunki `http.Get`/`http.Post` har doim yetarli emas (masalan, maxsus sarlavha qo'shish kerak bo'lsa).

**`http.NewRequest` — so'rovni to'liq nazorat ostida qurish:**

```go
sorov, err := http.NewRequest("GET", url, nil)
if err != nil {
	return err
}
sorov.Header.Set("X-Mijoz-Nomi", "go-learn")
```

(Diqqat: Go o'zgaruvchi nomlarida `'` belgisi ishlatilmaydi, shuning uchun kodda "so'rov" so'zini `sorov` deb yozamiz — apostrofsiz.)

`http.NewRequest(usul, url, tana)` — "HTTP Basics" darsida ko'rgan usul (`GET`, `POST` va h.k.), manzil, va so'rov tanasini (agar kerak bo'lmasa `nil`) qabul qiladi, va **hali yuborilmagan** `*http.Request` qaytaradi — bu, uni yuborishdan oldin **sozlash** imkonini beradi.

**`http.Client` — so'rovni yuborish.** `http.Get`/`http.Post` — aslida, orqa fonda standart `http.Client`dan foydalanadi. Qo'lda qurilgan so'rovni yuborish uchun, klientni ochiq chaqirish kerak:

```go
klient := &http.Client{}
javob, err := klient.Do(sorov)
```

`klient.Do(sorov)` — sizning maxsus sozlangan so'rovingizni yuboradi va oddiy `http.Get` kabi `*http.Response` qaytaradi.

**Nega bu naqsh kerak — maxsus sarlavhalar.** Real API'lar ko'pincha maxsus sarlavhalar talab qiladi — masalan, "Authentication" darsida ko'radigan `Authorization` sarlavhasi, yoki qaysi tilda javob kerakligini bildiruvchi `Accept-Language`. `http.Get`/`http.Post` bunday qo'shimcha sarlavha qo'shish imkonini bermaydi — faqat `http.NewRequest` + `Header.Set` + `Client.Do` naqshi orqali bu mumkin bo'ladi.

**Serverda kelgan sarlavhani o'qish.** Mos ravishda, server tomonida ham kelgan so'rovning sarlavhasini o'qish mumkin ("HTTP Handlers" darsidagi `r *http.Request`ni eslang):

```go
qiymat := r.Header.Get("X-Mijoz-Nomi")
```

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func sarlavhaBilanSorov(url, sarlavhaQiymati string) (int, error) {
	sorov, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	sorov.Header.Set("X-Mijoz-Nomi", sarlavhaQiymati)

	klient := &http.Client{}
	javob, err := klient.Do(sorov)
	if err != nil {
		return 0, err
	}
	defer javob.Body.Close()
	return javob.StatusCode, nil
}

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Kelgan sarlavha:", r.Header.Get("X-Mijoz-Nomi"))
		w.WriteHeader(http.StatusOK)
	}))
	defer testServer.Close()

	status, _ := sarlavhaBilanSorov(testServer.URL, "go-learn")
	fmt.Println(status)
}
```

Natija:

```
Kelgan sarlavha: go-learn
200
```

## TASK

`sarlavhaBilanSorov(url, sarlavhaQiymati string) (int, error)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `http.NewRequest("GET", url, nil)` orqali so'rov quring.
2. `X-Mijoz-Nomi` sarlavhasini `sarlavhaQiymati`ga o'rnating.
3. `http.Client{}` orqali so'rovni yuboring, javobning `StatusCode`sini qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `sorov, err := http.NewRequest("GET", url, nil)`, xatolikni tekshiring, keyin `sorov.Header.Set("X-Mijoz-Nomi", sarlavhaQiymati)`.
2. `klient := &http.Client{}`, `javob, err := klient.Do(sorov)`, xatolikni tekshiring, `defer javob.Body.Close()`, so'ng `return javob.StatusCode, nil`.
