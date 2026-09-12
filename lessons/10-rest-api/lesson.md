# 10 — Building a REST API

## THEORY

Bu darsda butun bo'lim davomida o'rgangan hamma narsani — handler'lar, JSON, status kodlar, va endi **so'rov usuli (method)** bo'yicha shoxlanish — birlashtirib, kichik bir **REST API** quramiz. "REST" — bu, resurslar (masalan, "kitoblar", "foydalanuvchilar") ustida **standart HTTP usullari** orqali amal bajarish g'oyasiga asoslangan uslub: `GET` — o'qish, `POST` — yaratish, `PUT`/`PATCH` — yangilash, `DELETE` — o'chirish.

**So'rov usulini tekshirish — `r.Method`.** Bitta manzil (masalan, `/kitoblar`) turli usullar orqali **turlicha** ma'no anglatishi mumkin — `GET /kitoblar` (barcha kitoblarni ko'rsat) va `POST /kitoblar` (yangi kitob qo'sh) ikkalasi ham bir xil manzilga, lekin boshqa-boshqa maqsadda yuboriladi. Buni "Switch" darsida ko'rgan naqsh bilan ajratish mumkin:

```go
func (d *KitoblarDukoni) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// barcha kitoblarni JSON qilib qaytarish
	case http.MethodPost:
		// yangi kitobni qo'shish
	default:
		w.WriteHeader(http.StatusMethodNotAllowed) // 405 — bu usul qo'llab-quvvatlanmaydi
	}
}
```

**Holatni saqlash — struct + method.** Bu handler **holatga ega** (u qaysi kitoblar borligini "eslab qolishi" kerak) — shuning uchun oddiy funksiya emas, balki "Methods" darsida ko'rgan **pointer receiver'li method** sifatida yoziladi:

```go
type Kitob struct {
	Nomi string `json:"nomi"`
}

type KitoblarDukoni struct {
	Kitoblar []Kitob
}
```

`KitoblarDukoni` — o'zining `Kitoblar` slice'ini saqlaydi, va `Handler` methodi shu slice bilan ishlaydi (`d.Kitoblar`ni o'qiydi va o'zgartiradi) — bu, aslida, "Data Structures" bo'limida ko'rgan struct+method naqshining, HTTP dunyosidagi tabiiy davomi.

**Yangi resurs yaratish — `201 Created`.** Yangi narsa muvaffaqiyatli yaratilganda, oddiy `200 OK` o'rniga, REST konvensiyasi bo'yicha **`201 Created`** status qaytariladi — bu mijozga "so'rovingiz nafaqat qabul qilindi, balki natijasida **yangi narsa paydo bo'ldi**" deb aniq bildiradi:

```go
w.WriteHeader(http.StatusCreated) // 201
json.NewEncoder(w).Encode(yangiKitob) // yaratilgan narsani qaytarib yuborish odat
```

**Butun oqim — POST so'rovni qanday qayta ishlash kerak:** 1) so'rov tanasidagi JSON'ni struct'ga o'qish ("Reading JSON Request Bodies" darsini eslang), 2) uni ichki ro'yxatga qo'shish (`append`), 3) `201 Created` status va yaratilgan narsaning o'zini JSON qilib javob berish.

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

type Kitob struct {
	Nomi string `json:"nomi"`
}

type KitoblarDukoni struct {
	Kitoblar []Kitob
}

func (d *KitoblarDukoni) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(d.Kitoblar)
	case http.MethodPost:
		var yangi Kitob
		if err := json.NewDecoder(r.Body).Decode(&yangi); err != nil {
			http.Error(w, "Noto'g'ri JSON", http.StatusBadRequest)
			return
		}
		d.Kitoblar = append(d.Kitoblar, yangi)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(yangi)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	dukon := &KitoblarDukoni{}

	postReq := httptest.NewRequest("POST", "/kitoblar", strings.NewReader(`{"nomi":"O'tkan kunlar"}`))
	postRec := httptest.NewRecorder()
	dukon.Handler(postRec, postReq)
	fmt.Println(postRec.Code, postRec.Body.String())

	getReq := httptest.NewRequest("GET", "/kitoblar", nil)
	getRec := httptest.NewRecorder()
	dukon.Handler(getRec, getReq)
	fmt.Println(getRec.Code, getRec.Body.String())
}
```

Natija:

```
201 {"nomi":"O'tkan kunlar"}

200 [{"nomi":"O'tkan kunlar"}]

```

## TASK

`Kitob`, `KitoblarDukoni` struct'lari va `Handler` methodi berilgan (`GET` va `default` holatlar allaqachon yozilgan). `POST` holatini shunday to'ldiringki:

1. So'rov tanasidagi JSON'ni `Kitob` struct'iga o'qing. Agar xatolik bo'lsa, `http.StatusBadRequest` (400) bilan `http.Error(w, "Noto'g'ri JSON", http.StatusBadRequest)` qaytaring.
2. Yangi kitobni `d.Kitoblar`ga qo'shing.
3. `Content-Type`ni `"application/json"` qilib, `http.StatusCreated` (201) status bilan, yaratilgan kitobni JSON qilib javob bering.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `var yangi Kitob`, `if err := json.NewDecoder(r.Body).Decode(&yangi); err != nil { http.Error(w, "Noto'g'ri JSON", http.StatusBadRequest); return }`.
2. `d.Kitoblar = append(d.Kitoblar, yangi)`, keyin `w.Header().Set("Content-Type", "application/json")`, `w.WriteHeader(http.StatusCreated)`, `json.NewEncoder(w).Encode(yangi)` — status va sarlavha har doim `Encode`dan **oldin**.
