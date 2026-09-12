# 04 — Reading JSON Request Bodies

## THEORY

"JSON Responses" darsida serverdan **chiqadigan** JSON'ni ko'rdik. Endi teskarisiga qaraymiz: mijoz serverga JSON **yuborganda** (odatda `POST` so'rovi bilan), server uni qanday "o'qib", tushunadi?

**So'rov tanasini JSON'dan struct'ga aylantirish — `json.NewDecoder(r.Body).Decode(&v)`:**

```go
type Sonlar struct {
	A int `json:"a"`
	B int `json:"b"`
}

func yigindiHandler(w http.ResponseWriter, r *http.Request) {
	var kirish Sonlar
	json.NewDecoder(r.Body).Decode(&kirish)

	natija := kirish.A + kirish.B
	fmt.Fprintf(w, "%d", natija)
}
```

`r.Body` — kelgan so'rovning "tanasi" (ya'ni, yuborilgan xom ma'lumot), `io.Reader` turida (fayl yoki matnni o'qishga o'xshaydi — "File I/O" darsini eslang). `json.NewDecoder(r.Body).Decode(&kirish)` — bu ma'lumotni o'qib, JSON'ni `kirish` struct'iga "yozib" beradi. Diqqat qiling: `&kirish` — pointer beryapmiz ("Pointers" darsini eslang), chunki `Decode` natijani **qaytarish** o'rniga, siz bergan struct'ning **o'zini, manzili orqali** to'ldiradi — "JSON" darsida `json.Unmarshal` bilan ko'rgan xuddi shu naqsh.

**So'rovni JSON tanasi bilan sinash.** `httptest.NewRequest`ning uchinchi argumenti — so'rov tanasi (`io.Reader` turida), uni `strings.NewReader(...)` orqali berish mumkin:

```go
tanaJSON := `{"a": 5, "b": 3}`
req := httptest.NewRequest("POST", "/yigindi", strings.NewReader(tanaJSON))
```

**Xatolikni tekshirish — noto'g'ri JSON kelishi mumkin.** `Decode` ham "Errors" darsida ko'rgan naqsh bo'yicha xatolik qaytaradi — agar kelgan matn haqiqiy JSON bo'lmasa (masalan, buzilgan yoki bo'sh bo'lsa):

```go
if err := json.NewDecoder(r.Body).Decode(&kirish); err != nil {
	http.Error(w, "Noto'g'ri JSON", http.StatusBadRequest)
	return
}
```

`http.Error` — xatolik xabarini **va** mos status kodini (`http.StatusBadRequest` — 400) bir vaqtda yozib beruvchi qulay yordamchi funksiya ("HTTP Status Codes" darsida bu haqda batafsil gaplashamiz).

## EXAMPLE

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Sonlar struct {
	A int `json:"a"`
	B int `json:"b"`
}

func yigindiHandler(w http.ResponseWriter, r *http.Request) {
	var kirish Sonlar
	json.NewDecoder(r.Body).Decode(&kirish)
	fmt.Fprintf(w, "%d", kirish.A+kirish.B)
}

func main() {
	tana := `{"a": 5, "b": 3}`
	req := httptest.NewRequest("POST", "/yigindi", strings.NewReader(tana))
	rec := httptest.NewRecorder()

	yigindiHandler(rec, req)

	fmt.Println(rec.Body.String())
}
```

Natija:

```
8
```

## TASK

`Sonlar` struct'i (`A int` — tegi `json:"a"`, `B int` — tegi `json:"b"`) va `yigindiHandler(w http.ResponseWriter, r *http.Request)` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u so'rov tanasidagi JSON'ni `Sonlar` struct'iga o'qib, `A` va `B`ning yig'indisini (oddiy son sifatida, JSON emas) javobga yozsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `var kirish Sonlar`, keyin `json.NewDecoder(r.Body).Decode(&kirish)` — `&kirish` orqali pointer bering, `Decode` natijani shu manzilga yozadi.
2. `fmt.Fprintf(w, "%d", kirish.A+kirish.B)` — yig'indini javobga yozing.
