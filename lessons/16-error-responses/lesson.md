# 16 — Error Responses

## THEORY

Bu bo'limning yakuniy darsi — "HTTP Status Codes", "JSON API" va "Validation" darslarida ko'rgan g'oyalarni birlashtirib, API'ning **xatolik javoblarini izchil** qilib qurishni ko'rsatadi.

**Muammo — turli xato matnlari, turli shakllarda.** Agar bir handler xatolikni oddiy matn (`http.Error`) bilan, boshqasi esa boshqa shakldagi JSON bilan qaytarsa, API'ni ishlatuvchi mijoz (masalan, mobil ilova) har bir holat uchun **alohida** kod yozishga majbur bo'ladi. Yechim — **barcha** xato javoblari uchun **bitta, izchil JSON shakl**ni belgilash:

```go
type XatoJavobi struct {
	Xato string `json:"xato"`
	Kod  int    `json:"kod"`
}
```

**Yordamchi funksiya — xato javobini yozish ("Test Helpers" darsidagi yordamchi funksiya g'oyasini eslang):**

```go
func xatoYoz(w http.ResponseWriter, xabar string, kod int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(kod)
	json.NewEncoder(w).Encode(XatoJavobi{Xato: xabar, Kod: kod})
}
```

Diqqat qiling: `w.Header().Set(...)` va `w.WriteHeader(...)` — "JSON Responses" va "HTTP Status Codes" darslarida ko'rgan qat'iy tartibda: sarlavha va status kod, javob tanasini yozishdan **oldin** o'rnatiladi.

**Handler'larda qanday ishlatiladi — butun bo'lim davomida o'rgangan hamma narsa birlashadi:**

```go
func kitobHandler(w http.ResponseWriter, r *http.Request) {
	var kirish Kitob
	if err := json.NewDecoder(r.Body).Decode(&kirish); err != nil {
		xatoYoz(w, "noto'g'ri JSON", http.StatusBadRequest)
		return
	}
	if err := kitobniValidatsiyaQil(kirish); err != nil {
		xatoYoz(w, err.Error(), http.StatusBadRequest)
		return
	}
	// ... hammasi joyida, davom etamiz ...
}
```

**Nega bu — butun bo'limning "yopilishi".** Endi har bir handler, muvaffaqiyatsizlik yuz berganda, **bir xil, bashorat qilinadigan** JSON shaklida javob beradi — mijoz doim `{"xato": "...", "kod": ...}` ko'rinishidagi javobni kutishi mumkin, handler qaysi bo'limda ekanligidan qat'iy nazar. Bu — "REST API" darsida ko'rgan izchillik g'oyasining, xatolik holatlari uchun ham qo'llanilishi: yaxshi API — nafaqat muvaffaqiyat holatida, balki **xatolik holatida ham** izchil va bashorat qilinadigan bo'lishi kerak.

## EXAMPLE

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type XatoJavobi struct {
	Xato string `json:"xato"`
	Kod  int    `json:"kod"`
}

func xatoYoz(w http.ResponseWriter, xabar string, kod int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(kod)
	json.NewEncoder(w).Encode(XatoJavobi{Xato: xabar, Kod: kod})
}

func main() {
	rec := httptest.NewRecorder()
	xatoYoz(rec, "ism bo'sh bo'lishi mumkin emas", http.StatusBadRequest)

	fmt.Println(rec.Code)
	fmt.Println(rec.Body.String())
}
```

Natija:

```
400
{"xato":"ism bo'sh bo'lishi mumkin emas","kod":400}

```

## TASK

`XatoJavobi` struct'i (`Xato string` — tegi `json:"xato"`, `Kod int` — tegi `json:"kod"`) berilgan. `xatoYoz(w http.ResponseWriter, xabar string, kod int)` funksiyasini shunday to'ldiringki, u:

1. `Content-Type`ni `"application/json"` qilib o'rnatsin.
2. `kod` status bilan javob boshini yozsin.
3. `XatoJavobi{Xato: xabar, Kod: kod}`ni JSON qilib javobga yozsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `w.Header().Set("Content-Type", "application/json")` va `w.WriteHeader(kod)` — ikkalasi ham `Encode`dan **oldin**.
2. `json.NewEncoder(w).Encode(XatoJavobi{Xato: xabar, Kod: kod})`.
