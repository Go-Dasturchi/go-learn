# 07 — REST API

## THEORY

"Web Development" bo'limidagi "Building a REST API" darsida `GET` (ro'yxatni ko'rish) va `POST` (yangi narsa qo'shish)ni ko'rgan edingiz. To'liq **CRUD** (Create, Read, Update, Delete) uchun, oxirgi muhim qism yetishmayapti: **`DELETE`** — narsani o'chirish.

**`DELETE` so'rovini qayta ishlash — odatda manzil orqali "qaysi narsa" ko'rsatiladi.** Ko'pincha o'chiriladigan narsaning ID'si (yoki indeksi) URL manzilining o'zida beriladi (masalan, `/kitoblar/2`) — "Web Development" bo'limidagi "Routing" darsida ko'rgan `strings.TrimPrefix` naqshi bilan buni ajratib olamiz:

```go
func (d *KitoblarDukoni) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// ... (avvalgi darsdagidek) ...
	case http.MethodPost:
		// ... (avvalgi darsdagidek) ...
	case http.MethodDelete:
		indeksMatni := strings.TrimPrefix(r.URL.Path, "/kitoblar/")
		indeks, err := strconv.Atoi(indeksMatni)
		if err != nil || indeks < 0 || indeks >= len(d.Kitoblar) {
			w.WriteHeader(http.StatusNotFound) // 404 — bunday indeks yo'q
			return
		}
		d.Kitoblar = append(d.Kitoblar[:indeks], d.Kitoblar[indeks+1:]...)
		w.WriteHeader(http.StatusNoContent) // 204 — muvaffaqiyatli, qaytariladigan tana yo'q
	}
}
```

**`d.Kitoblar[:indeks]` + `d.Kitoblar[indeks+1:]` — slice'dan elementni o'chirish naqshi.** "Slices" darsida ko'rgan slicing sintaksisini eslang: bu naqsh — `indeks`gacha bo'lgan qismni, `indeks`dan keyingi qism bilan **birlashtirib**, aynan `indeks`dagi elementni "tashlab ketadi". Bu — Go'da array/slice'dan bitta elementni o'chirishning eng ko'p qo'llaniladigan usuli (chunki Go'da tayyor "o'chirish" methodi yo'q).

**`204 No Content` — muvaffaqiyatli, lekin qaytariladigan narsa yo'q.** "HTTP Handlers" va "HTTP Status Codes" darslarida ko'rgan `200`/`201`dan farqli, `204` — "amal muvaffaqiyatli bajarildi, lekin sizga qaytarish uchun hech qanday ma'lumot yo'q" degani. O'chirish amali uchun bu — juda mos: narsa o'chirildi, uni qayta ko'rsatishning ma'nosi yo'q.

**Nega ID/indeks orqali manzil qurish — REST'ning "resurs" g'oyasi.** "Building a REST API" darsida ko'rgan REST uslubida, har bir **resurs** (bu holda, har bir kitob) o'zining **noyob manzili**ga ega bo'lishi kerak deb hisoblanadi: `/kitoblar` — barcha kitoblar to'plami, `/kitoblar/2` — aynan 2-indeksdagi (yoki ID'li) bitta kitob. Bu naqsh, API'ni izchil va bashorat qilinadigan qiladi: har qanday resurs uchun, uning manzilini bilib, unga mos HTTP usulini qo'llash kifoya.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

type KitoblarDukoni struct {
	Kitoblar []string
}

func (d *KitoblarDukoni) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		indeksMatni := strings.TrimPrefix(r.URL.Path, "/kitoblar/")
		indeks, err := strconv.Atoi(indeksMatni)
		if err != nil || indeks < 0 || indeks >= len(d.Kitoblar) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		d.Kitoblar = append(d.Kitoblar[:indeks], d.Kitoblar[indeks+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	dukon := &KitoblarDukoni{Kitoblar: []string{"Sarob", "O'tkan kunlar"}}

	req := httptest.NewRequest("DELETE", "/kitoblar/0", nil)
	rec := httptest.NewRecorder()
	dukon.Handler(rec, req)

	fmt.Println(rec.Code)
	fmt.Println(dukon.Kitoblar)
}
```

Natija:

```
204
[O'tkan kunlar]
```

## TASK

`KitoblarDukoni` va `Handler` methodi berilgan. `Handler`ning `DELETE` holatini shunday to'ldiringki:

1. `r.URL.Path`dan `/kitoblar/` prefiksini kesib, qolgan qismni songa (indeksga) aylantiring.
2. Agar aylantirish xato bersa, yoki indeks chegaradan tashqarida bo'lsa, `http.StatusNotFound` (404) qaytarib to'xtang.
3. Aks holda, `d.Kitoblar`dan shu indeksdagi elementni o'chirib, `http.StatusNoContent` (204) qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `indeksMatni := strings.TrimPrefix(r.URL.Path, "/kitoblar/")`, `indeks, err := strconv.Atoi(indeksMatni)`, keyin `if err != nil || indeks < 0 || indeks >= len(d.Kitoblar) { w.WriteHeader(http.StatusNotFound); return }`.
2. `d.Kitoblar = append(d.Kitoblar[:indeks], d.Kitoblar[indeks+1:]...)`, so'ng `w.WriteHeader(http.StatusNoContent)`.
