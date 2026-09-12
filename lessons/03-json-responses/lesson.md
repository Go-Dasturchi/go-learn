# 03 — JSON Responses

## THEORY

Zamonaviy veb-ilovalar (mobil ilova, frontend saytlar) serverdan ko'pincha oddiy matn emas, balki **JSON** kutadi — "JSON" darsida ko'rgan universal ma'lumot formatini eslang. Endi shu bilimni HTTP handler'lar bilan birlashtiramiz.

**Struct'ni to'g'ridan-to'g'ri javobga yozish — `json.NewEncoder(w).Encode(...)`:**

```go
type Odam struct {
	Ism  string `json:"ism"`
	Yosh int    `json:"yosh"`
}

func odamHandler(w http.ResponseWriter, r *http.Request) {
	odam := Odam{Ism: "Ali", Yosh: 25}
	json.NewEncoder(w).Encode(odam)
}
```

`json.NewEncoder(w)` — "JSON" darsida ko'rgan `json.Marshal`ga o'xshaydi, lekin natijani qaytarish o'rniga, uni to'g'ridan-to'g'ri `w`ga (ya'ni, javobga) **yozib yuboradi** — bu ayniqsa katta ma'lumotlar uchun samaraliroq, chunki oraliq `[]byte`ni xotirada alohida saqlashning hojati yo'q.

**`Content-Type` sarlavhasini o'rnatish — mijozga "bu JSON" deb aytish.** Standart holatda, `w`ga yozilgan javob "oddiy matn" deb hisoblanadi. Mijoz tomonga bu aslida JSON ekanligini bildirish uchun maxsus **sarlavha (header)** o'rnatish odat tusiga kirgan:

```go
w.Header().Set("Content-Type", "application/json")
```

**Muhim tartib qoidasi: sarlavhani yozishdan oldin o'rnatish kerak.** HTTP'da sarlavhalar javob **tanasidan oldin** yuboriladi — shuning uchun `w.Header().Set(...)` har doim `Encode(...)` (yoki boshqa yozish amali)dan **oldin** chaqirilishi kerak:

```go
func odamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // AVVAL
	json.NewEncoder(w).Encode(odam)                      // KEYIN
}
```

Agar tartibni teskari qilsangiz, sarlavha e'tiborsiz qoldiriladi — chunki Go birinchi yozish amalida (bu holda `Encode`) sarlavhalarni "muzlatib", ularni avtomatik yuborib yuboradi (standart `200 OK` status bilan birga).

**Diqqat: `Encode` oxiriga yangi qator qo'shadi.** `json.Encoder.Encode` — har doim yozgan JSON'ining oxiriga bitta `\n` (yangi qator) belgisini qo'shadi — bu, ayniqsa test yozganda, aniq mos kelishi kerak bo'lgan kichik, lekin muhim tafsilot.

## EXAMPLE

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Odam struct {
	Ism  string `json:"ism"`
	Yosh int    `json:"yosh"`
}

func odamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	odam := Odam{Ism: "Ali", Yosh: 25}
	json.NewEncoder(w).Encode(odam)
}

func main() {
	req := httptest.NewRequest("GET", "/odam", nil)
	rec := httptest.NewRecorder()

	odamHandler(rec, req)

	fmt.Println(rec.Header().Get("Content-Type"))
	fmt.Println(rec.Body.String())
}
```

Natija:

```
application/json
{"ism":"Ali","yosh":25}

```

## TASK

`Odam` struct'i (`Ism string` — tegi `json:"ism"`, `Yosh int` — tegi `json:"yosh"`) va `odamHandler(w http.ResponseWriter, r *http.Request)` funksiyasi berilgan. Funksiyani shunday to'ldiringki:

1. `Content-Type` sarlavhasini `"application/json"` deb o'rnatsin.
2. `Odam{Ism: "Ali", Yosh: 25}` qiymatini JSON ko'rinishida javobga yozsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `w.Header().Set("Content-Type", "application/json")` — bu qatorni **avval** yozing, `Encode`dan oldin.
2. `json.NewEncoder(w).Encode(Odam{Ism: "Ali", Yosh: 25})` — natijani `w`ga to'g'ridan-to'g'ri yozadi.
