# 08 — JSON API

## THEORY

"Web Development" bo'limida JSON javob yozish va JSON so'rov tanasini o'qishni ko'rgan edingiz. Endi, ko'plab haqiqiy JSON API'larda uchraydigan bir naqshni ko'rib chiqamiz: natijani **oddiy massiv** sifatida emas, balki qo'shimcha ma'lumot (masalan, umumiy son) bilan birga, **"konvert" (envelope)** ichiga o'rab qaytarish.

**Nega faqat massiv qaytarish yetarli emas.** Agar API to'g'ridan-to'g'ri `["Ali", "Vali", "Guli"]` qaytarsa, mijoz "jami nechta element bor" yoki "yana ko'proq sahifa bormi" kabi savollarga javob topa olmaydi (ayniqsa "Pagination" darsida ko'radigan holatlarda). Buning o'rniga, ko'plab API'lar natijani quyidagicha "o'raydi":

```json
{
	"data": ["Ali", "Vali", "Guli"],
	"count": 3
}
```

**Bu struct'lar orqali qanday amalga oshiriladi:**

```go
type RoyxatJavobi struct {
	Data  []string `json:"data"`
	Count int      `json:"count"`
}

func royxatHandler(w http.ResponseWriter, r *http.Request) {
	royxat := []string{"Ali", "Vali", "Guli"}
	javob := RoyxatJavobi{
		Data:  royxat,
		Count: len(royxat),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(javob)
}
```

**Nega bu naqsh foydali — "kelajakka moslashuvchanlik".** Agar hozir siz to'g'ridan-to'g'ri massivni qaytarsangiz, va keyinroq API'ga `count` yoki `sahifa` kabi yangi maydon qo'shmoqchi bo'lsangiz — bu **eski mijozlarni sindiradi** (chunki ular "javob — bu massiv" deb kutayotgan bo'ladi, endi esa u struct bo'lib qoladi). Boshidanoq "konvert" ichida qaytarish, kelajakda **yangi maydon qo'shishni** hech kimni sindirmasdan amalga oshirish imkonini beradi — bu, "Modules" darsida ko'rgan semantik versiyalash g'oyasiga o'xshab, "orqaga mos kelish"ni saqlashning yana bir usuli.

**Struct ichida struct — "Nested Structs" darsining amaliy qo'llanilishi.** Katta API'larda konvert ko'pincha yanada murakkab bo'ladi — masalan, xato va muvaffaqiyat holatlari uchun alohida maydonlar, yoki "Pagination" darsida ko'radigan sahifalash ma'lumotlari uchun ichma-ich struct. Bugungi oddiy `RoyxatJavobi` — shu murakkabroq naqshlarning boshlang'ich nuqtasi.

## EXAMPLE

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type RoyxatJavobi struct {
	Data  []string `json:"data"`
	Count int      `json:"count"`
}

func royxatHandler(w http.ResponseWriter, r *http.Request) {
	royxat := []string{"Ali", "Vali", "Guli"}
	javob := RoyxatJavobi{Data: royxat, Count: len(royxat)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(javob)
}

func main() {
	req := httptest.NewRequest("GET", "/royxat", nil)
	rec := httptest.NewRecorder()
	royxatHandler(rec, req)
	fmt.Println(rec.Body.String())
}
```

Natija:

```
{"data":["Ali","Vali","Guli"],"count":3}

```

## TASK

`RoyxatJavobi` struct'i (`Data []string` — tegi `json:"data"`, `Count int` — tegi `json:"count"`) va `royxatHandler` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u `royxat := []string{"Ali", "Vali", "Guli"}` ro'yxatini `RoyxatJavobi{Data: royxat, Count: len(royxat)}` ko'rinishida, `Content-Type: application/json` sarlavhasi bilan JSON qilib javobga yozsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `w.Header().Set("Content-Type", "application/json")` — bu qatorni **avval** yozing.
2. `json.NewEncoder(w).Encode(RoyxatJavobi{Data: royxat, Count: len(royxat)})`.
