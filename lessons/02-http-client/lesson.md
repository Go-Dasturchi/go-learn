# 02 — HTTP Client

## THEORY

"Web Development" bo'limidagi "HTTP Client" darsida `http.Get` bilan oddiy so'rov yuborishni ko'rgan edingiz. Endi buni bir qadam oldinga olib boramiz: **ma'lumot yuborib** (POST) so'rov yuborish, va **javobning status kodini** tekshirish.

**`http.Post` — JSON tana bilan so'rov yuborish:**

```go
import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Kirish struct {
	Ism string `json:"ism"`
}

kirish := Kirish{Ism: "Ali"}
baytlar, _ := json.Marshal(kirish) // "JSON" darsini eslang

javob, err := http.Post(url, "application/json", bytes.NewReader(baytlar))
```

`http.Post` uchta argument oladi: manzil, `Content-Type` (server yuborilayotgan ma'lumot turini bilishi uchun), va tana (`io.Reader` turida — shuning uchun `[]byte`ni `bytes.NewReader(...)` orqali "o'qiladigan" shaklga aylantiramiz).

**Javobning status kodini tekshirish.** "HTTP Status Codes" darsida ko'rgan kodlarni, endi **klient** tomonidan ham tekshirish mumkin — bu, so'rov "yuborilgani" bilan "muvaffaqiyatli qabul qilingani" bir xil emasligini bildiradi:

```go
javob, err := http.Post(url, "application/json", tana)
if err != nil {
	return err // tarmoq xatosi — so'rov umuman yetib bormadi
}
defer javob.Body.Close()

if javob.StatusCode != http.StatusOK {
	return fmt.Errorf("kutilmagan status: %d", javob.StatusCode)
}
```

`javob.StatusCode` — `int` turidagi, javobning haqiqiy status kodi (masalan, `200`, `404`, `500`). Diqqat qiling: `err != nil` faqat **tarmoq darajasidagi** muammoni bildiradi (server umuman javob bermadi); server javob bergan, lekin "xato" statusi (masalan, `500`) bilan javob bergan bo'lsa, `err` baribir `nil` bo'ladi — shuning uchun status kodni **alohida** tekshirish har doim muhim.

**Nega bu muhim.** Real dasturlarda, tashqi xizmatga so'rov yuborganda, "so'rov yetib bordimi" va "so'rov **muvaffaqiyatli** bajarildimi" — ikkita **alohida** savol. Ikkalasini ham tekshirmasdan qoldirish, xatolarni jimgina "yutib yuborishi" mumkin — bu esa, dasturning nima uchun noto'g'ri ishlayotganini tushunishni ancha qiyinlashtiradi.

## EXAMPLE

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Kirish struct {
	Ism string `json:"ism"`
}

func royxatgaOl(url, ism string) (int, error) {
	baytlar, err := json.Marshal(Kirish{Ism: ism})
	if err != nil {
		return 0, err
	}
	javob, err := http.Post(url, "application/json", bytes.NewReader(baytlar))
	if err != nil {
		return 0, err
	}
	defer javob.Body.Close()
	return javob.StatusCode, nil
}

func main() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer testServer.Close()

	status, _ := royxatgaOl(testServer.URL, "Ali")
	fmt.Println(status)
}
```

Natija:

```
201
```

## TASK

`Kirish` struct'i (`Ism string`, tegi `json:"ism"`) va `royxatgaOl(url, ism string) (int, error)` funksiyasi berilgan. Funksiyani shunday to'ldiringki:

1. `Kirish{Ism: ism}`ni JSON'ga kodlab (`json.Marshal`).
2. `http.Post` orqali `url`ga, `"application/json"` turi bilan yuborsin.
3. Javobning `StatusCode`sini qaytarsin (xatolik bo'lsa, uni ham qaytaring).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `baytlar, err := json.Marshal(Kirish{Ism: ism})`, xatolikni tekshiring, keyin `http.Post(url, "application/json", bytes.NewReader(baytlar))`.
2. `defer javob.Body.Close()` qiling, so'ng `return javob.StatusCode, nil`.
